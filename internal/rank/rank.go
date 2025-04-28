package rank

import (
   "fmt"
   "log"
   "time"

   "github.com/pocketbase/dbx"
   "github.com/pocketbase/pocketbase"
)

// CalculateRank fetches posts grouped by events, computes scores, and updates their ranks.
func CalculateRank(app *pocketbase.PocketBase) error {
   log.Println("Calculating rank for posts...")

   events, err := app.Dao().FindRecordsByExpr("events")
   if err != nil {
       return fmt.Errorf("error fetching events: %w", err)
   }

   // instantiate the time decay algorithm with a one-week half-life
   alg := NewTimeDecayAlgorithm(7 * 24 * time.Hour)

   for _, event := range events {
       posts, err := app.Dao().FindRecordsByExpr("posts", dbx.HashExp{"event": event.Id})
       if err != nil {
           log.Printf("Error fetching posts for event %s: %v", event.Id, err)
           continue
       }

       results, err := alg.Compute(posts)
       if err != nil {
           log.Printf("Ranking error for event %s: %v", event.Id, err)
           continue
       }

       for _, r := range results {
           post := r.Post
           log.Printf("Post %s: votes=%.0f, score=%.2f, rank=%d", post.Id, post.Get("votes"), r.Score, r.Rank)
           post.Set("rank", r.Rank)
           if err := app.Dao().SaveRecord(post); err != nil {
               log.Printf("Error updating rank for post %s: %v", post.Id, err)
           }
       }
   }

   log.Println("Rank calculation completed.")
   return nil
}