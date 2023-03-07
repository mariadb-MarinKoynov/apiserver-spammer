# apiserver-spammer

To start the program, you need three arguments. First, the pod name you want to run this to. Second, the namespace of the pod. Third, the number of parallel connection you want to open. Run 
```bash
go run cmd/main.go <POD-NAME> <POD-NAMESPACE> <NUMBER-OF-PARALLEL-OPEN-CONNECTIONS>
```

For example
```bash
go run cmd/main.go biraishkembe-mdb-ms-0 biraishkembe 400
```

The result will be placed in folders in your target directory. Note, rerunning the program will delete everything in your target dir, and recreate it from scrach, so make sure to backup whatever you wish!