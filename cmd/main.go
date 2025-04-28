package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/cron"

	"github.com/we-be/vibecheckr/internal/rank"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
			// enable auto creation of migration files when making collection changes in the Dashboard
			// (the isGoRun check is to enable it only during development)
			Automigrate: isGoRun,
		})

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
