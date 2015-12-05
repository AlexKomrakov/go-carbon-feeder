package main

import "os/exec"

type Command struct {
	Name        string
	Arguments []string
}

func (c Command) Run() ([]byte, error) {
	return exec.Command(c.Name, c.Arguments...).Output()
}