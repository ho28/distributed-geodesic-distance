package master

import (
    "proj/graph"
    "sync"
    "fmt"
    //"strconv"
    "time"
)

type Master struct {
    workers []*WorkerClient
    lock *sync.Mutex
    g *graph.Graph
    nodeIndex int
    Backs []string
    Dead []bool //array index corresponds to worker num. false is dead true is alive
    Done bool
    Result map[int]string
    TaskList Tasks
}

func (self *Master) LoadGraph () {
    g := new(graph.Graph)
    g.LoadFromFile("scan73", "../src/proj/data/pointmat_scan73.csv", "../src/proj/data/edge4geomat_scan73.csv")

    self.g = g
}

func (self *Master) LoadWorkers () {

    /* Load workers */
    workers := make([]*WorkerClient,len(self.Backs))

    for i,v := range self.Backs {
        workers[i] = &WorkerClient{addr:v}
    }

    self.workers = workers
}

func (self *Master) InitWorker (workerIndex int) {

        e := self.workers[workerIndex].Init(self.g)

        if e != nil {
            fmt.Println("init error:",e.Error())
            self.Dead[workerIndex] = true
        }
}

func (self *Master) Init () {
    self.LoadGraph()
    self.Dead = make([]bool,len(self.Backs))
    self.LoadWorkers()
    for i := range self.Backs {
      self.InitWorker(i)
    }
    self.nodeIndex = 0
    self.lock = &sync.Mutex{}
    self.nodeIndex = 1
    self.lock = &sync.Mutex{}
    self.g.NV=45//just for debug
    jobs := make([]int,0,self.g.NV)
    for i:=1; i<=self.g.NV; i++ {
      jobs = append(jobs,i)
    }
    self.TaskList = StoreTasks(jobs)
    self.Result = make(map[int]string)
    self.Done = false
}

func (self *Master) Work (workerIndex int, errChan chan<- error) {

    worker := self.workers[workerIndex]
    if worker.Validate() == false {
      time.Sleep(time.Second)
      errChan<-fmt.Errorf("worker dead")
      return
    } else {
      self.Dead[workerIndex] = false
    }

    for !self.Dead[workerIndex] {
      self.lock.Lock()

      if i:=self.TaskList.NextTask();i!=-1 {

        self.lock.Unlock()

        var ret string
        resultChan := make(chan error)
	      go func() {
          e := worker.GeoDis(i, &ret)
          resultChan <-e
        }()
        e := <-resultChan
        if e != nil {
          //check if error is "killed", cuz if it is then we killed it on purpose, the worker isn't dead
          fmt.Println("error", e.Error())
          self.Dead[workerIndex] = true//report self as dead
          errChan<-e//do we need this? will this break it?
          return
        }
        fmt.Print("Worker", workerIndex)
        fmt.Print(", output: ", ret)
        fmt.Println("")
        self.lock.Lock()
        self.TaskList.FinishedTask(i)
        self.lock.Unlock()
        self.Result[workerIndex] = ret
      } else {
        self.lock.Unlock()
        self.Done = true
        errChan<-nil
        break
      }
	}
  return
}
      //handle this in main.go: if a worker returns, it died. The master main() can occasionally probe it and if it detects it started up again then create a new worker and add it back to the worker queue and call work. It can accomplish this by, if a call to work() returns, occasionally, every minute or so, call work() again. it won't affect correctness it will just die and return. if the worker comes back up it will connect

