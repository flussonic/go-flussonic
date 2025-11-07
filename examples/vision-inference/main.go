package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flussonic/go-flussonic/config"
	visioninference "github.com/flussonic/go-flussonic/vision-inference"
)

func main() {
	ctx := context.Background()

	// Create Vision Inference client
	cfg, err := config.ParseURL("http://your-cluster-key@vision.example.com:80")
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	client, err := visioninference.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Get Vision configuration
	fmt.Println("=== Example 1: Get Vision Config ===")
	visionConfig, err := client.ConfigGet(ctx)
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	} else if visionConfig != nil {
		fmt.Println("Vision config retrieved successfully")
	}

	// Example 2: Get list of streams
	fmt.Println("\n=== Example 2: List Streams ===")
	streams, err := client.StreamsList(ctx, &visioninference.StreamsListQuery{})
	if err != nil {
		log.Printf("Failed to get streams: %v", err)
	} else if streams != nil {
		fmt.Printf("Found %d streams\n", len(streams.Streams()))
		for i, stream := range streams.Streams() {
			if stream.Name() != "" {
				fmt.Printf("  Stream %d: %s\n", i+1, stream.Name())
			}
			if i >= 4 {
				break
			}
		}
	}

	// Example 3: Get list of episodes with iterator
	fmt.Println("\n=== Example 3: List Episodes with Iterator ===")
	episodesIter := client.EpisodesListIterator(ctx, &visioninference.EpisodesListQuery{
		Select: "episode_id,media,opened_at",
		Limit:  5,
		Sort:   "timestamp",
	})

	episodeCount := 0
	for episode, err := range episodesIter {
		if err != nil {
			log.Printf("Error iterating episodes: %v", err)
			break
		}
		episodeCount++
		episodeID := episode.EpisodeID()
		fmt.Printf("  Episode %d: %d", episodeCount, episodeID)
		if episode.Media() != "" {
			fmt.Printf(" (media: %s)", episode.Media())
		}
		fmt.Println()
		if episodeCount >= 5 {
			break
		}
	}
	fmt.Printf("Total episodes processed: %d\n", episodeCount)

	// Example 4: Get list of counter records with iterator
	fmt.Println("\n=== Example 4: List Counter Records with Iterator ===")
	counterRecordsIter := client.CounterRecordsListIterator(ctx, &visioninference.CounterRecordsListQuery{})

	recordCount := 0
	for record, err := range counterRecordsIter {
		if err != nil {
			log.Printf("Error iterating counter records: %v", err)
			break
		}
		recordCount++
		// VisionCounterRecord has various fields
		_ = record // Use record
		fmt.Printf("  Record %d processed\n", recordCount)
		if recordCount >= 5 {
			break
		}
	}
	fmt.Printf("Total counter records processed: %d\n", recordCount)

	// Example 5: Get Vision metrics
	fmt.Println("\n=== Example 5: Get Vision Metrics ===")
	metrics, err := client.MetricsGet(ctx)
	if err != nil {
		log.Printf("Failed to get metrics: %v", err)
	} else if metrics != nil {
		fmt.Println("Vision metrics retrieved successfully")
	}

	// Example 6: Get worker statistics
	fmt.Println("\n=== Example 6: Get Worker Stats ===")
	stats, err := client.StatsGet(ctx)
	if err != nil {
		log.Printf("Failed to get stats: %v", err)
	} else if stats != nil {
		fmt.Println("Worker stats retrieved successfully")
		// VisionWorkerStats may have various embedded interfaces
	}

	// Example 7: Get episodes with filtering by media
	fmt.Println("\n=== Example 7: List Episodes with Filtering ===")
	episodes, err := client.EpisodesList(ctx, &visioninference.EpisodesListQuery{})
	if err != nil {
		log.Printf("Failed to get episodes: %v", err)
	} else if episodes != nil {
		fmt.Printf("Found %d episodes\n", len(episodes.Episodes()))
		for i, episode := range episodes.Episodes() {
			episodeID := episode.EpisodeID()
			fmt.Printf("  Episode %d: %d", i+1, episodeID)
			if episode.Media() != "" {
				fmt.Printf(" (media: %s)", episode.Media())
			}
			fmt.Println()
		}
		// Check if there are more pages
		if episodes.Next() != nil {
			fmt.Printf("More episodes available (cursor: %s)\n", *episodes.Next())
		}
	}
}
