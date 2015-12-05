package main

import "os/exec"

type Command struct {
	Command   string
	Arguments []string
}

func (c Command) Run() ([]byte, error) {
	return exec.Command(c.Command, c.Arguments).Output()
}