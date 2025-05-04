package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/cron"

	"github.com/we-be/vibecheckr/internal/rank"
	_ "github.com/we-be/vibecheckr/migrations"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	// register migrate CLI commands (create, up, down, etc.)
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	// Register event hooks for likes
	app.OnRecordAfterCreateRequest("likes").Add(func(e *core.RecordCreateEvent) error {
		return updatePostVotes(app, e.Record, 1)
	})

	app.OnRecordAfterDeleteRequest("likes").Add(func(e *core.RecordDeleteEvent) error {
		return updatePostVotes(app, e.Record, -1)
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// Calculate rank every 10 minutes
		scheduler.MustAdd("calculate-rank", "*/10 * * * *", func() {
			if err := rank.CalculateRank(app); err != nil {
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

// updatePostVotes updates the votes count for a post when a like is created or deleted
func updatePostVotes(app *pocketbase.PocketBase, likeRecord *models.Record, delta int) error {
	postId := likeRecord.GetString("post")
	if postId == "" {
		return fmt.Errorf("like record is missing post ID")
	}

	// Fetch the post
	post, err := app.Dao().FindRecordById("posts", postId)
	if err != nil {
		return fmt.Errorf("failed to find post with ID %s: %w", postId, err)
	}

	// Get current votes count
	currentVotes := post.GetInt("votes")

	// Update votes count
	post.Set("votes", currentVotes+delta)

	// Save the updated post
	if err := app.Dao().SaveRecord(post); err != nil {
		return fmt.Errorf("failed to update votes for post %s: %w", postId, err)
	}

	log.Printf("Updated votes for post %s: %d -> %d", postId, currentVotes, currentVotes+delta)
	return nil
}
