package main

import (
	"context"
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("name is required, usage: 'register <name>'")
	}
	
	nameReg := cmd.Args[0]
	ctx := context.Background()
	params := database.CreateUserParams{
		ID:					uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:				nameReg,
	}
	
	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("there was an error creating the user: %v", err)
	}

	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("there was an error setting the user, %w", err)
	}
	fmt.Printf("User created successfully. User: %v\n", user.Name)
	return nil
}