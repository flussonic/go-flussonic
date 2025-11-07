package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flussonic/go-flussonic/config"
	"github.com/flussonic/go-flussonic/flussonic"
	"github.com/flussonic/go-flussonic/flussonic/model"
)

func main() {
	ctx := context.Background()

	// Create client using ParseURL (recommended way)
	cfg, err := config.ParseURL("http://your-cluster-key@streamer.example.com:80")
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	client, err := flussonic.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Server readiness check
	fmt.Println("=== Example 1: Readiness Probe ===")
	probe, err := client.ReadinessProbe(ctx)
	if err != nil {
		log.Printf("Readiness probe failed: %v", err)
	} else if probe != nil && probe.ServerVersion() != nil {
		fmt.Printf("Server is ready. Version: %s\n", string(*probe.ServerVersion()))
	}

	// Example 2: Get server configuration
	fmt.Println("\n=== Example 2: Get Server Config ===")
	serverConfig, err := client.ConfigGet(ctx, nil)
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	} else if serverConfig != nil {
		stats := serverConfig.Stats()
		if stats != nil && stats.ServerVersion() != nil {
			fmt.Printf("Server version: %s\n", string(*stats.ServerVersion()))
		}
	}

	streamsListQuery := &flussonic.StreamsListQuery{
		Select: []string{"name", "stats"},
		Limit:  10, // Limit number of items per page
	}

	for stream, err := range client.StreamsListIterator(ctx, streamsListQuery) {
		if err != nil {
			log.Printf("Failed to get streams: %v", err)
		}

		if stream != nil {
			fmt.Printf("Stream: %s\n", stream.Name())
		}
	}

	// Example 3: Get list of streams with pagination
	fmt.Println("\n=== Example 3: List Streams with Pagination ===")
	streams, err := client.StreamsList(ctx, &flussonic.StreamsListQuery{
		Select: []string{"name", "stats"},
		Limit:  10, // Limit number of items per page
	})
	if err != nil {
		log.Printf("Failed to get streams: %v", err)
	} else if streams != nil {
		streamCount := 0
		for _, stream := range streams.Streams() {
			streamCount++
			name := stream.Name()
			if name != "" {
				fmt.Printf("  Stream %d: %s\n", streamCount, name)
			}
			if streamCount >= 5 {
				break
			}
		}
		fmt.Printf("Total streams on page: %d\n", len(streams.Streams()))
		// Check if there are more pages
		if streams.Next() != nil {
			fmt.Printf("More streams available (cursor: %s)\n", *streams.Next())
		}
	}

	// Example 4: Get single stream
	fmt.Println("\n=== Example 4: Get Single Stream ===")
	streamName := "test-stream"
	stream, err := client.StreamGet(ctx, streamName)
	if err != nil {
		log.Printf("Failed to get stream '%s': %v", streamName, err)
	} else if stream != nil {
		fmt.Printf("Stream '%s' found\n", streamName)
		if stream.Name() != "" {
			fmt.Printf("  Name: %s\n", stream.Name())
		}
	}

	// Example 5: Create/update stream
	fmt.Println("\n=== Example 5: Create/Update Stream ===")
	streamConfig := model.NewStreamConfig()
	streamConfig.SetName("example-stream")
	// For example create minimal configuration
	// In real usage need to set all required fields

	savedStream, err := client.StreamSave(ctx, "example-stream", streamConfig)
	if err != nil {
		log.Printf("Failed to save stream: %v", err)
	} else if savedStream != nil {
		fmt.Printf("Stream saved successfully: %s\n", savedStream.Name())
	}

	// Example 6: Get list of DVR archives
	fmt.Println("\n=== Example 6: List DVR Archives ===")
	dvrs, err := client.DvrsList(ctx, &flussonic.DvrsListQuery{
		Sort:  []string{"name"},
		Limit: 10,
	})
	if err != nil {
		log.Printf("Failed to get DVRs: %v", err)
	} else if dvrs != nil {
		fmt.Printf("Found %d DVR archives\n", len(dvrs.Dvrs()))
		for i, dvr := range dvrs.Dvrs() {
			name := dvr.Name()
			if name != "" {
				fmt.Printf("  DVR %d: %s\n", i+1, name)
			}
			if i >= 2 {
				break
			}
		}
	}

	// Example 7: Get list of sessions with manual pagination
	fmt.Println("\n=== Example 7: List Sessions with Manual Pagination ===")
	sessions, err := client.SessionsList(ctx, &flussonic.SessionsListQuery{
		Limit: 5,
	})
	if err != nil {
		log.Printf("Failed to get sessions: %v", err)
	} else if sessions != nil {
		fmt.Printf("Found %d sessions\n", len(sessions.Sessions()))
		for i, session := range sessions.Sessions() {
			sessionID := session.ID()
			if sessionID != nil {
				fmt.Printf("  Session %d: %s\n", i+1, string(*sessionID))
			}
		}
		// Check if there are more pages
		if sessions.Next() != nil {
			fmt.Printf("More sessions available (cursor: %s)\n", *sessions.Next())
		}
	}

	// Example 8: Delete stream
	fmt.Println("\n=== Example 8: Delete Stream ===")
	if err := client.StreamDelete(ctx, "example-stream"); err != nil {
		log.Printf("Failed to delete stream: %v", err)
	} else {
		fmt.Println("Stream deleted successfully")
	}

	// Example 9: Get list of packages with filtering
	fmt.Println("\n=== Example 9: List Packages with Filtering ===")
	packages, err := client.PackagesList(ctx, &flussonic.PackagesListQuery{
		Select: []string{"name", "streams"},
		Limit:  10,
		Sort:   []string{"name"},
	})
	if err != nil {
		log.Printf("Failed to get packages: %v", err)
	} else if packages != nil {
		fmt.Printf("Found %d packages\n", len(packages.Packages()))
		for i, pkg := range packages.Packages() {
			if pkg.Name() != "" {
				fmt.Printf("  Package %d: %s\n", i+1, pkg.Name())
			}
		}
	}
}
