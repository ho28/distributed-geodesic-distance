package worker

import (
  //"net/rpc"
  "os"
  //"net"
  "encoding/csv"
  "strconv"
  "runtime"
  "fmt"
  "proj/geo_dis"
  "proj/graph"
)

type GeoDisWorker struct {
  addr string
  g *graph.Graph
  killChans []chan bool
  //mem scheduler begin
  curr_mem float64
  prev_mem float64
  seekingjobs bool
  //mem scheduler end

}

func (self *GeoDisWorker) alive (succ *bool) error {
	*succ=true
	return  nil
}



func NewWorker(addr string) *GeoDisWorker {
  worker := new(GeoDisWorker)
  worker.addr = addr
  return worker
}
/*
func (self *GeoDisWorker) GetAddr() string {
  return self.addr
}
*/
func (self *GeoDisWorker) Init(g *graph.Graph, succ *bool) error {
  self.g = g
  self.killChans = make([]chan bool,0,20)
  *succ = true
  return nil
}

func (self *GeoDisWorker) GeoDis(startNode int, outLocation *string) error {
  killChan := make(chan bool)
  for i := range(self.killChans){
    if self.killChans[i] == nil {
      self.killChans[i] = killChan
    }
  }
  res,err := geo_dis.Geo_Dis(self.g.Vertices, self.g.Edges, startNode, self.g.NV, self.g.NE, killChan)
  if err != nil {
    return err
  }
  //write res to file and return the location written in form addr:filepath
  fname := self.g.Name + "_" + strconv.Itoa(startNode) + ".csv"
  file, err := os.Create(fname)
  if err != nil {
    return err
  }
  writer := csv.NewWriter(file)
  resStr := make([]string,len(res))
  for i,val := range res {
    resStr[i] = strconv.FormatFloat(val,'f',-1,64)
  }
  err = writer.Write(resStr)
  if err != nil {
    return err
  }
  writer.Flush()
  file.Close()
  *outLocation = fname
  res = nil
  runtime.GC()
  return nil
}

func (self *GeoDisWorker) KillTask(task int, succ *bool) error {
  //this is shitty, no idea if it works.
  //TODO: kills random job, doesn't use first (task) argument
  for _,v := range(self.killChans){
    if v != nil {
      v<-true
      runtime.GC()
      return nil
    }
  }
  return fmt.Errorf("no jobs to kill")
}

func (self *GeoDisWorker) Echo(word *string, echo *string) error {
  echo = word
  return nil
}
/*
func main() {
  //os.Args to figure out what address to start service on
  thisWorker := GeoDisWorker{addr:"169.228.66.157:8899"} //TODO
  server := rpc.NewServer()
  server.RegisterName("Worker", &thisWorker)

  l,e := net.Listen("tcp",thisWorker.addr)
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
}*/
