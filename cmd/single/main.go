package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	argsWithProg := os.Args
	validate(argsWithProg)

	podName := argsWithProg[1]
	podNamespace := argsWithProg[2]
	routines, _ := strconv.Atoi(argsWithProg[3])
	path := "target"

	increaseUlimit()
	createDirectory(path)

	// run the actual program
	var wg sync.WaitGroup
	for i := 0; i < routines; i++ {
		wg.Add(1)
		routineNumber := i
		go func() {
			defer wg.Done()
			execIntoPod(path, podName, podNamespace, routineNumber)
		}()
	}

	wg.Wait()
}

func execIntoPod(path, name, namespace string, routineNumber int) error {
	cmd := exec.Command("kubectl", "exec", "-i", name, "-n", namespace, "--", "cat", "/dev/urandom")

	output, err := cmd.Output()
	if err != nil {
		log.Printf("An error occurred reading output: %s. Exiting...", err)
		return err
	}
	cmd.Start()

	f, err := os.Create(fmt.Sprintf("%s/%d.txt", path, routineNumber))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(string(output))
	if err != nil {
		log.Fatal(err)
	}

	return nil

	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	log.Fatalf("An error occurred reading stream: %s. Exiting...", err)
	// }
	// cmd.Start()

	// scanner := bufio.NewScanner(stderr)
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	m := scanner.Text()
	// 	fmt.Print(m)
	// 	fmt.Print(" ")
	// }
	// cmd.Wait()
}

func validate(args []string) {
	if len(args) != 4 {
		log.Fatalf("This is not the correct number of arguments! Please, read the manual.")
	}

	_, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatalf("Could not parse number of open connections, please enter a real number as your third argument!")
	}
}

func increaseUlimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Getting Rlimit %s", err)
	}
	log.Print("Old session limit: ")
	log.Println(rLimit)
	rLimit.Max = 999999
	rLimit.Cur = 999999
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Setting Rlimit %s", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Getting Rlimit %s", err)
	}
	log.Print("New session limit: ")
	log.Println(rLimit)
}

func createDirectory(path string) {
	os.RemoveAll("target")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalf("An error occured creating dir: %s", err)
		}
	}
}
