package rank

import (
   "fmt"
   "math"
   "sort"
   "time"

   "github.com/pocketbase/pocketbase/models"
   "github.com/pocketbase/pocketbase/tools/types"
) 
// Result holds a post record along with its computed score and rank.
type Result struct {
   Post  *models.Record
   Score float64
   Rank  int
}

// TimeDecayAlgorithm implements the Algorithm interface using exponential time decay.
type TimeDecayAlgorithm struct {
   HalfLife time.Duration
   NowFunc  func() time.Time
}

// NewTimeDecayAlgorithm creates a new TimeDecayAlgorithm with the given half-life duration.
func NewTimeDecayAlgorithm(halfLife time.Duration) *TimeDecayAlgorithm {
   return &TimeDecayAlgorithm{
       HalfLife: halfLife,
       NowFunc:  time.Now,
   }
}

// Compute calculates scores for the posts using an exponential decay based on creation time.
func (a *TimeDecayAlgorithm) Compute(posts []*models.Record) ([]Result, error) {
   var results []Result

   for _, post := range posts {
       votes, ok := post.Get("votes").(float64)
       if !ok {
           return nil, fmt.Errorf("invalid votes for post %s", post.Id)
       }

       createdAt, ok := post.Get("created").(types.DateTime)
       if !ok {
           return nil, fmt.Errorf("invalid created date for post %s", post.Id)
       }

       createdTime := createdAt.Time()
       timeDiff := a.NowFunc().Sub(createdTime)
       decayFactor := math.Pow(2, -timeDiff.Seconds()/a.HalfLife.Seconds())
       score := votes * decayFactor

       results = append(results, Result{
           Post:  post,
           Score: score,
       })
   }

   // Sort by score descending
   sort.Slice(results, func(i, j int) bool {
       return results[i].Score > results[j].Score
   })

   // Assign ranks
   for i := range results {
       results[i].Rank = i + 1
   }

   return results, nil
}