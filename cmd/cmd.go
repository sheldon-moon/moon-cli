package cmd

import (
	"log"
	"os/exec"
)


func execute(name string , arg string)  []byte {
	cmd := exec.Command("echo", "'hello world!'")
	out, err := cmd.Output()
	if err != nil {
		log.Print("execute command error :" , name , arg , err )
		return []byte(err.Error())
	}
	return out
}


