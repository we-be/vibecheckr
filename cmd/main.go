package main

import (
   "log"

   "github.com/pocketbase/pocketbase"
   "github.com/pocketbase/pocketbase/core"
   "github.com/pocketbase/pocketbase/tools/cron"
   "github.com/we-be/vibecheckr/internal/rank"
)

func main() {
   app := pocketbase.New()

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