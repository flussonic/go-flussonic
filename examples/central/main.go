package main

import (
	"context"
	"fmt"
	"log"

	central "github.com/flussonic/go-flussonic/central"
	"github.com/flussonic/go-flussonic/central/model"
	"github.com/flussonic/go-flussonic/config"
)

func main() {
	ctx := context.Background()

	// Create Central client
	cfg, err := config.ParseURL("http://api_key@localhost")
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	client, err := central.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Get Central configuration
	fmt.Println("=== Example 1: Get Central Config ===")
	centralConfig, err := client.ConfigGet(ctx, nil)
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	} else if centralConfig != nil && centralConfig.Loglevel() != nil {
		fmt.Printf("Central loglevel: %s\n", *centralConfig.Loglevel())
	}

	fmt.Println("=== Creating a stream ===")

	stream := model.NewCentralStreamConfig()
	fmt.Printf("stream is %+v\n", stream)
	fmt.Printf(" stream is nil? %t\n", stream == nil)

	input := model.NewStreamInput()
	fmt.Println(1)
	input.SetURL("rtsp://localhost/stream")
	fmt.Println(2)
	stream.SetName("some-name")
	fmt.Println(3)
	stream.SetInputs([]model.StreamInput{input})
	fmt.Println(4)

	stream2, err := client.StreamSave(ctx, string(stream.Name()), stream)
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	fmt.Printf("Stream saved: %+v\n", stream2)

	// Example 2: Get list of agents with pagination
	fmt.Println("\n=== Example 2: List Agents with Pagination ===")
	agents, err := client.AgentsList(ctx, &central.AgentsListQuery{
		Select: []string{"id", "name", "status"},
		Limit:  10,
	})
	if err != nil {
		log.Printf("Failed to get agents: %v", err)
	} else if agents != nil {
		fmt.Printf("Found %d agents\n", len(agents.Agents()))
		for i, agent := range agents.Agents() {
			if agent == nil {
				fmt.Println("Agent is nil")
				continue
			}
			if agent.ID() != nil {
				fmt.Printf("  Agent %d: %d", i+1, agent.ID())
				fmt.Println()
			}
			if i >= 4 {
				break
			}
		}
		// Check if there are more pages
		if agents.Next() != nil {
			fmt.Printf("More agents available (cursor: %s)\n", *agents.Next())
		}
	}

	// Example 3: Get single agent
	fmt.Println("\n=== Example 3: Get Single Agent ===")
	agentID := "agent-123"
	agent, err := client.AgentGet(ctx, agentID)
	if err != nil {
		log.Printf("Failed to get agent '%s': %v", agentID, err)
	} else if agent != nil {
		fmt.Printf("Agent '%s' found\n", agentID)
		if agent.ID() != nil {
			fmt.Printf("  ID: %d\n", agent.ID())
		}
	}
}
