package main

import (
	"flag"
	"fmt"
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
	scores := make([]int, *score+1)
	scores[*score] = *nplayers

	// Returns score.
	randomPlayer := func(removed int) int {
		player := fastrand.Intn(*nplayers - scores[0] - removed)
		s := 0
		for score := 1; score < len(scores); score++ {
			s1 := s + scores[score]
			if s1 > player {
				return score
			}
			s = s1
		}
		panic("broken logic in randomPlayer")
	}

	printScores := func() {
		var parts []string
		for score, x := range scores {
			if x != 0 {
				parts = append(parts, fmt.Sprintf("#%d=%d", score, x))
			}
		}
		fmt.Println(strings.Join(parts, " "))
	}

	for i := 0; ; i++ {
		if i%*period == 0 {
			printScores()
		}

		// Choose two random players.
		s1 := randomPlayer(0)
		scores[s1]--
		s2 := randomPlayer(1)
		scores[s2]--
		if fastrand.Intn(2) == 0 {
			s1--
			s2++
		} else {
			s1++
			s2--
		}
		if s1 == len(scores) || s2 == len(scores) {
			scores = append(scores, 0)
		}
		scores[s1]++
		scores[s2]++

		if scores[0] == *nplayers-1 {
			printScores()
			break
		}
	}
}
