# apiserver-spammer

To start the program, you need three arguments. First, the pod name you want to run this to. Second, the namespace of the pod. Third, the number of parallel connection you want to open. Run 
```bash
go run cmd/single/main.go <POD-NAME> <POD-NAMESPACE> <NUMBER-OF-PARALLEL-OPEN-CONNECTIONS>
```

For example
```bash
go run cmd/single/main.go biraishkembe-mdb-ms-0 biraishkembe 1000
```

The result will be placed in folders in your target directory. If you see ANY resulting files in the `target` dir, that means that there was an error, and the error will appear in the program's output. Note, rerunning the program will delete everything in your target dir, and recreate it from scrach, so make sure to backup whatever you wish!