package main

import (
	"context"
	"fmt"
	"go_tutorial/bootdotdev/blog_aggregator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRss(s *state, _ command) error {
	log.Println("Connecting to rss feed...")

	url := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching: %w", err)
	}

	fmt.Printf("%s\n%s\n%s\n", rss.Channel.Title, rss.Channel.Link, rss.Channel.Description)
	for _, item := range rss.Channel.Item {
		fmt.Printf("  %s\n  %s\n  %s\n  %s\n", item.Title, item.Link, item.PubDate, item.Description)
	}
	return nil
}

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("feed name and a url are required, Usage: addfeed <name> <url>")
	}

	users, err := s.db.GetUser(context.Background(), s.conf.CurrentUser)
	if err != nil {
		return fmt.Errorf("failed to retrieve current user: %w", err)
	}
	
	params := database.CreateFeedParams{
		ID: 				uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name: 			cmd.Args[0],
		Url: 				cmd.Args[1],
		UserID: 		users.ID,
	}
	
	log.Println("Creating feed data...")
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to save feed to database: %w", err)
	}

	fmt.Printf("Feed records:\n - ID: %v\n - Created at: %v\n - Updated at: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt)
	fmt.Printf(" - Name: %v\n - Url: %v\n - User Id: %v\n", feed.Name, feed.Url, feed.UserID)
	return nil
}

func handlerFeeds(s *state, _ command) error {
	allFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds, try again: %w", err)
	}

	for _, fed := range allFeeds {
		fmt.Printf("Feed: %s\nUrl: %s\nUsername: %s\n", fed.Name, fed.Url, fed.Name_2)
	}
	
	return nil
}