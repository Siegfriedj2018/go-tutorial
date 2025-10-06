package main

import (
	"context"
	"fmt"
	"log"
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