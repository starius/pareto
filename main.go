package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/NebulousLabs/fastrand"
)

var (
	nplayers = flag.Int("nplayers", 1000, "Number of players")
	score    = flag.Int("score", 10, "Initial score")
	period   = flag.Int("period", 1000000, "Report period")
)

func main() {
	flag.Parse()
	scores := make([]int, *nplayers)
	for i := range scores {
		scores[i] = *score
	}

	remaining := *nplayers

	printScores := func() {
		counters := make(map[int]int)
		if remaining != *nplayers {
			counters[0] = *nplayers - remaining
		}
		for p := 0; p < remaining; p++ {
			score := scores[p]
			counters[score]++
		}
		type pair struct {
			score, count int
		}
		var pairs []pair
		for score, count := range counters {
			pairs = append(pairs, pair{score, count})
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].score < pairs[j].score
		})
		var parts []string
		for _, p := range pairs {
			parts = append(parts, fmt.Sprintf("#%d=%d", p.score, p.count))
		}
		fmt.Println(strings.Join(parts, " "))
	}

	for i := 0; ; i++ {
		if i%*period == 0 {
			printScores()
		}

		// Choose two random players.
		p1 := fastrand.Intn(remaining)
		p2 := fastrand.Intn(remaining - 1)
		if p2 >= p1 {
			p2++
		}
		var loser, winner int
		if fastrand.Intn(2) == 0 {
			loser = p1
			winner = p2
		} else {
			loser = p2
			winner = p1
		}
		scores[loser]--
		scores[winner]++
		if scores[loser] == 0 {
			remaining--
			scores[loser], scores[remaining] = scores[remaining], scores[loser]
			if remaining == 1 {
				printScores()
				break
			}
		}
	}
}
