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

func handlerFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("feed name and a url are required, Usage: addfeed <name> <url>")
	}

	// users, err := s.db.GetUser(context.Background(), s.conf.CurrentUser)
	// if err != nil {
	// 	return fmt.Errorf("failed to retrieve current user: %w", err)
	// }
	
	params := database.CreateFeedParams{
		ID: 				uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name: 			cmd.Args[0],
		Url: 				cmd.Args[1],
		UserID: 		user.ID,
	}
	
	log.Println("Creating feed data...")
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to save feed to database: %w", err)
	}
  
	params2 := database.CreateFeedFollowsParams{
		ID: 				uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	}

	_, err = s.db.CreateFeedFollows(context.Background(), params2)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	} 
	
	fmt.Printf("User: %v\n", user.Name)
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

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please add a url, Usage: follow <url>")
	}
	ctx := context.Background()
	// user, err := s.db.GetUser(ctx, s.conf.CurrentUser)
	// if err != nil {
	// 	return fmt.Errorf("there was an err with current users: %w", err)
	// }
	feed, err := s.db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("err with the url: %w", err)
	}

	params := database.CreateFeedFollowsParams{
		ID: 				uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		UserID: 		user.ID,
		FeedID: 		feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollows(ctx, params)
	if err != nil {
		return fmt.Errorf("creating feed follow has an err: %w", err)
	}

	fmt.Printf("%v is now following %v", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, _ command, user database.User) error {
	// currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUser)
	// if err != nil {
	// 	return fmt.Errorf("error retrieving user: %w", err)
	// }

	userFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error retrieving users feed: %w", err)
	}

	fmt.Printf("User: %v\n", user.Name)
	for _, feed := range userFeeds {
		fmt.Printf(" - Feed: %v\n", feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("there was a error getting feed: %w", err)
	}

	params := database.DeleteFeedFollowByUserParams{
		UserID: 	user.ID,
		FeedID: 	feed.ID,
	}

	err = s.db.DeleteFeedFollowByUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	log.Printf("You (%v) have unfollowed %v\n", user.Name, feed.Name)
	return nil
}