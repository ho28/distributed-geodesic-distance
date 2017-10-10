package main

import (
  "proj/master"
  "time"
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    fmt.Println("Please specify addresses of worker nodes")
    os.Exit(1)
  }
  Master := master.Master{Backs:args}
  start := time.Now()
  Master.Init()
  errChan := make(chan error,1)

  for i,_ := range Master.Backs {
    go Master.Work(i,errChan)
  }
  for !Master.Done {
    ret := <-errChan
    if ret != nil {
      //fmt.Println(ret.Error())
      for i,v := range Master.Dead {
        if v {
          Master.InitWorker(i)
          go Master.Work(i,errChan)
        }
      }
    }
  }
  elapsed := time.Since(start)
  fmt.Printf("execution took %s",elapsed)
}
