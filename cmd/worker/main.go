package main


import (
  "net"
  "net/rpc"
  "fmt"
  "os"
  "runtime"
  "proj/worker"
)


func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  //os.Args to figure out what address to start service on
  argsWithoutProg := os.Args[1:]
  if len(argsWithoutProg) != 1 {
    fmt.Println("Usage: worker [addr:port]")
    os.Exit(1)
  }
  address := argsWithoutProg[0]
  //thisWorker := worker.GeoDisWorker{addr:"169.228.66.157:8899"} //TODO
  thisWorker := worker.NewWorker(address)
  server := rpc.NewServer()
  server.RegisterName("Worker", thisWorker)

  l,e := net.Listen("tcp",address)
  if e != nil {
    fmt.Errorf("net.Listen error")
    fmt.Println("error:",e.Error())
    //handle error
  }

  for{
    conn,e := l.Accept()
    if e != nil {
      fmt.Errorf("listener.Accept error")
      //handle error
    }
    go server.ServeConn(conn)
  }
}
