package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	command := `top -b -n1 | grep "Cpu(s)" | awk '{print $2 + $4}'`
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}