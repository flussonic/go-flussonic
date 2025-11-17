package main

import (
	"context"
	"fmt"
	"log"
	"time"

	central "github.com/flussonic/go-flussonic/central"
	"github.com/flussonic/go-flussonic/central/model"
	"github.com/flussonic/go-flussonic/config"
)

func FetchAndProcessEpisodes() {
	ctx, cancel := context.WithCancel(context.Background())

	// Create Central client
	cfg, err := config.ParseURL("http://api_key@localhost")
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	client, err := central.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	query := central.EpisodesListQuery{
		EpisodeType: "vehicle",
	}

	go client.EpisodesStreaming(ctx, &query, callback)

	// fetching stopped after 1 minute
	<-time.After(1 * time.Minute)
	cancel()
}

func callback(_ context.Context, episode model.Episode) error {
	fmt.Printf("Fetched episode with id: %d\n", episode.EpisodeID())
	return nil
}
