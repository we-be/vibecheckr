package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/pocketbase/pocketbase/tools/types"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// Calculate rank every hour
		scheduler.MustAdd("calculate-rank", "*/5 * * * *", func() {
			if err := calculateRank(app); err != nil {
				log.Printf("Error calculating rank: %v", err)
			}
		})

		scheduler.Start()

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func calculateRank(app *pocketbase.PocketBase) error {
	log.Println("Calculating rank for posts...")

	// Get all events
	events, err := app.Dao().FindRecordsByExpr("events")
	if err != nil {
		return fmt.Errorf("error fetching events: %w", err)
	}

	for _, event := range events {
		// Get posts for this event
		posts, err := app.Dao().FindRecordsByExpr("posts",
			dbx.HashExp{"event": event.Id},
		)
		if err != nil {
			log.Printf("Error fetching posts for event %s: %v", event.Id, err)
			continue
		}

		for _, post := range posts {
			votes, ok := post.Get("votes").(float64)
			if !ok {
				log.Printf("Error: votes for post %s is not a number", post.Id)
				continue
			}

			createdAt, ok := post.Get("created").(types.DateTime)
			if !ok {
				log.Printf("Error: created date for post %s is invalid", post.Id)
				continue
			}

			// Convert types.DateTime to time.Time
			createdTime := createdAt.Time()

			// Calculate time decay factor (example: 1 week half-life)
			timeDiff := time.Since(createdTime)
			decayFactor := math.Pow(0.5, timeDiff.Hours()/(24*7))

			// Calculate rank
			rank := 1 / int(votes * decayFactor)
			log.Printf("post with votes %f (@%s) is rank %d with %f decay", votes, createdTime, rank, decayFactor)

			// Update post with new rank
			post.Set("rank", rank)
			if err := app.Dao().SaveRecord(post); err != nil {
				log.Printf("Error updating rank for post %s: %v", post.Id, err)
			}
		}
	}

	log.Println("Rank calculation completed.")
	return nil
}
