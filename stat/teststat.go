package main

import (
	"log"
	lp "proj/goprocinfo_tmp/linux"
	"fmt"
)


func main(){
stat, err := lp.ReadMaxPID("/proc/sys/kernel/pid_max")
if err !=nil{
	log.Fatal("stat read fail")
}

fmt.Println(stat)
}
