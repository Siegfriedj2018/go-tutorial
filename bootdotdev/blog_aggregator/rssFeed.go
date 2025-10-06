package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,		
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to create a get request: %v", err) 
	}
	req.Header.Set("User-Agent", "gator_from_bootdev_user")

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("there was an error sending your request, try again: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading xml doc: %w", err)
	}

	var rssString *RSSFeed
	err = xml.Unmarshal(body, &rssString)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshalling xml: %w", err)
	}

	cleanedTitleChan := html.UnescapeString(rssString.Channel.Title)
	cleanedDesChan := html.UnescapeString(rssString.Channel.Description)
	rssString.Channel.Title = cleanedTitleChan
	rssString.Channel.Description = cleanedDesChan
	for idx, feed := range rssString.Channel.Item {
		cleanedTitleFeed := html.UnescapeString(feed.Title)
		cleanedDesFeed := html.UnescapeString(feed.Description)
		rssString.Channel.Item[idx].Title = cleanedTitleFeed
		rssString.Channel.Item[idx].Description = cleanedDesFeed
	}
	return rssString, nil
}