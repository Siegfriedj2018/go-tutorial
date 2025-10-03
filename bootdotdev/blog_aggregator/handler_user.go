package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("a username is required, please try again")
	}
  
	ctx := context.Background()
	cmdName := cmd.Args[0]
  user, err := s.db.GetUser(ctx, cmdName)
	if err != nil {
		log.Fatalf("user is not found: %v", err)
		return err
	}

	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	log.Println("User has been set to", user.Name)
	return nil
}