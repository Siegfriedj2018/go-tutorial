package main

import (
	"fmt"
	"strings"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	comFunc, ok := c.cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("that is an invalid command, try again")
	}

	err := comFunc(s, cmd)
	return err
}

func (c *commands) Register(name string, f func(*state, command) error) {
	if c.cmds == nil {
		c.cmds = make(map[string]func(*state, command) error)
	}

	if strings.ToLower(name) == "login" {
		c.cmds[name] = f
	}
}
