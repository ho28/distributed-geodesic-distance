package main

import (
	//"encoding/json"
	//"fmt"
	"log"
	"os/exec"
	)

func main() {

	cmd := exec.Command("./worker", "169.228.66.155:8899")
	err := cmd.Start()
	

	cmd2:= exec.Command("./master","169.228.66.155:8899")
	err2:= cmd2.Start()
	if err != nil {	
		log.Fatal(err)
	}
	if err2 != nil{
		log.Fatal(err)
	}

      log.Printf("Waiting for command to finish...")
      err = cmd.Wait()
      log.Printf("Command finished with error: %v", err)


/*
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := cmd.Start(); err != nil {
				log.Fatal(err)
		}	
		var person struct {
			Name string
			Age  int
		}
		if err := json.NewDecoder(stdout).Decode(&person); err != nil {	
			log.Fatal(err)
			}
																			if err := cmd.Wait(); err != nil {
			log.Fatal(err)	
		}
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
*/
}



