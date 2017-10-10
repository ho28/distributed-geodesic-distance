package main

//THIS TEST FILE IS AVAILABLE FOR SIMPLE AD HOC UNIT TESTS

import (
  "fmt"
  "runtime"
  "time"
  "strconv"
  "proj/geo_dis"
  "proj/graph"
)

func main() {
  start := time.Now()
  g := new(graph.Graph)
  g.LoadFromFile("scan74", "data/pointmat_scan74.csv", "data/edge4geomat_scan74.csv")
  killChan := make(chan bool,1)//make buffered otherwise we will block on  sending the kill signal until someone reads from that channel?
  for j:=10000;j<10003;j++{
    res,_ := geo_dis.Geo_Dis(g.Vertices, g.Edges, j, g.NV, g.NE, killChan)
    fmt.Println("after geo_dis, j = ",strconv.Itoa(j))
    if j == 10001 {
      killChan<-true
    }
    fmt.Println("asdf")
    fmt.Println("before printing result, j = ",strconv.Itoa(j))
    for i:=range res {
      fmt.Printf("%f ", res[i])
    }
    res = nil//make res eligible for garbage collection
    runtime.GC()//force garbage collection or we will run out of memory
  }
  close(killChan)
  elapsed := time.Since(start)
  fmt.Printf("execution took %s\n",elapsed)
}

