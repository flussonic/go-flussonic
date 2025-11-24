package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/flussonic/go-flussonic/authorization"
	central "github.com/flussonic/go-flussonic/central"
	"github.com/flussonic/go-flussonic/central/model"
	"github.com/flussonic/go-flussonic/config"
)

func main() {
	ctx := context.Background()

	// Get Central URL from environment or use default
	centralURL := os.Getenv("CENTRAL_URL")
	if centralURL == "" {
		centralURL = "http://localhost:9019"
	}
	apiKey := os.Getenv("CENTRAL_API_KEY")
	if apiKey == "" {
		apiKey = "api_key"
	}

	// Create Central client
	cfg, err := config.ParseURL(centralURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}
	cfg.Auth = authorization.BearerAuth(apiKey)
	cfg.Retry = 3

	client, err := central.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Step 1: Create streamers in central
	fmt.Println("=== Step 1: Creating streamers ===")
	if err := createStreamers(ctx, client); err != nil {
		log.Fatalf("Failed to create streamers: %v", err)
	}
	fmt.Println("Streamers created successfully")

	// Step 2: Read streams.csv and create streams
	fmt.Println("\n=== Step 2: Reading streams.csv and creating streams ===")
	if err := createStreamsFromCSV(ctx, client); err != nil {
		log.Fatalf("Failed to create streams from CSV: %v", err)
	}
	fmt.Println("Streams created successfully")

	// Step 3: Start streaming episodes with type qr_code
	fmt.Println("\n=== Step 3: Starting episode streaming (type: qr_code) ===")
	streamEpisodes(ctx, client)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}

func streamEpisodes(ctx context.Context, client central.Central) {
	query := &central.EpisodesListQuery{
		EpisodeType: "qr_code",
		PollTimeout: 30,
	}

	// Start streaming episodes
	go client.EpisodesStreaming(ctx, query, episodesCallback)
}

func episodesCallback(_ context.Context, episode model.Episode) error {
	// Check if episode type is qr_code
	episodeType := episode.EpisodeType()
	if episodeType == nil || *episodeType != "qr_code" {
		return nil
	}

	// Get payload
	payload := episode.Payload()
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		payloadJSON = fmt.Appendf(nil, "%v", payload)
	}

	// Output episode information
	fmt.Printf("\tEpisode ID: %d, Type: %s, Media: %s, OpenedAt: %d\n", episode.EpisodeID(), *episodeType, episode.Media(), episode.OpenedAt())
	fmt.Printf("\tPayload: %s\n", string(payloadJSON))
	fmt.Println("---")

	return nil
}

func createStreamers(ctx context.Context, client central.Central) error {
	// Create a streamer configuration
	streamerConfig := model.NewStreamerConfig().
		SetHostname("flussonic-1").
		SetRole("streamer").
		SetAPIURL(model.URL("http://flussonic-1:80")).
		SetClusterKey("cluster_key")

	// Save the streamer
	_, err := client.StreamerSave(ctx, string(*streamerConfig.Hostname()), nil, streamerConfig)
	if err != nil {
		return fmt.Errorf("failed to save streamer: %w", err)
	}

	fmt.Printf("Created streamer: %s\n", string(*streamerConfig.Hostname()))

	inferenceConfig := model.NewStreamerConfig().
		SetHostname("demo-inference").
		SetRole("inference").
		SetAPIURL(model.URL("http://demo-inference:8020"))

	_, err = client.StreamerSave(ctx, string(*inferenceConfig.Hostname()), nil, inferenceConfig)
	if err != nil {
		return fmt.Errorf("failed to save inference streamer: %w", err)
	}

	fmt.Printf("Created streamer: %s\n", string(*streamerConfig.Hostname()))

	return nil
}

func createStreamsFromCSV(ctx context.Context, client central.Central) error {
	// Open and read CSV file
	file, err := os.Open("streams.csv")
	if err != nil {
		return fmt.Errorf("failed to open streams.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	// Skip header row if present
	startIdx := 0
	if len(records) > 0 && (records[0][0] == "name" || records[0][0] == "Name") {
		startIdx = 1
	}

	// Create streams from CSV records
	for i := startIdx; i < len(records); i++ {
		record := records[i]
		if len(record) < 2 {
			log.Printf("Skipping invalid record at line %d: insufficient columns", i+1)
			continue
		}

		name := record[0]
		url := record[1]

		if name == "" || url == "" {
			log.Printf("Skipping invalid record at line %d: empty name or URL", i+1)
			continue
		}

		// Create stream configuration
		stream := model.NewCentralStreamConfig()
		stream.SetName(model.MediaName(name))

		// Create input
		input := model.NewStreamInput()
		input.SetURL(model.InputURL(url))
		stream.SetInputs([]model.StreamInput{input})

		stream.SetDvr(model.NewStreamDvrSpec().
			SetReference("watcher").
			SetExpiration(1800).
			SetEpisodesExpiration(3600),
		)

		stream.SetVision(
			model.NewVisionSpec().SetAlg(model.VisionSpecAlgFaces),
		)

		// Save stream
		_, err := client.StreamSave(ctx, string(stream.Name()), stream)
		if err != nil {
			log.Printf("Failed to create stream '%s': %v", name, err)
			continue
		}

		fmt.Printf("Created stream: %s -> %s\n", name, url)
	}

	return nil
}
