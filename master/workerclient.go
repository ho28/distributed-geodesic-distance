package master

import (
    "net/rpc"
    "fmt"
    "proj/graph"
)

// RPC Function Strings
const sWorkerEcho   string = "Worker.Echo"
const sWorkerInit   string = "Worker.Init"
const sWorkerGeoDis string = "Worker.GeoDis"

// Error Strings
const sInitFailed string = "Init Failed"

// Misc Strings
const sTcp string = "tcp"

type WorkerClient struct {
    addr  string
    conn *rpc.Client
    connected bool
}

func (self *WorkerClient) Init(g *graph.Graph) error {

    valid := self.Validate()

    if !self.connected || !valid {

        conn, e := rpc.Dial(sTcp, self.addr)

        if e != nil {
            self.connected = false
            return e
        }

        self.connected = true
        self.conn = conn
    }

    var succ bool

    e := self.conn.Call(sWorkerInit, g, &succ)

    if e != nil {
        self.connected = false
        self.conn.Close()
        return e
    }

    if !succ {
        return fmt.Errorf(sInitFailed)
    }

    return nil
}

func (self *WorkerClient) GeoDis(startNode int, output *string) error {

    valid := self.Validate()

    if !self.connected || !valid {

        conn, e := rpc.Dial(sTcp, self.addr)

        if e != nil {
            self.connected = false
            return e
        }

        self.connected = true
        self.conn = conn
    }

    e := self.conn.Call(sWorkerGeoDis, startNode, output)

    if e != nil {
        self.connected = false
        self.conn.Close()
        return e
    }

    return nil
}

func (self *WorkerClient) Validate() bool {

    var fooVal string

    if self.connected == false {
        return false
    }

    e := self.conn.Call(sWorkerEcho, "foo", &fooVal)

    if e != nil {
       	    fmt.Println("worker could not validate..poss dead")
	    return false
    }

    return true
}
