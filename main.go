package main

import (
	"fmt"
	"log"
)

func main() {
	out, err := Command{`bash`, [2]string{`-c`, `top -b -n1 | grep "Cpu(s)" | awk '{print $2 + $4}'`}}.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}