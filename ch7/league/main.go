package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	teams := []Team{
		{name: "Bulldogs", players: []string{"James", "Bob"}},
		{name: "Roosters", players: []string{"Billy", "Wayne"}},
		{name: "Eels", players: []string{"Jimmy", "Greg"}},
		{name: "Broncos", players: []string{"George", "Brett"}},
	}
	league := League{
		Teams: teams,
	}

	league.MatchResult("Bulldogs", "Roosters", 24, 12)
	league.MatchResult("Bulldogs", "Eels", 24, 12)
	league.MatchResult("Bulldogs", "Broncos", 12, 24)
	league.MatchResult("Eels", "Broncos", 12, 12)
	league.MatchResult("Eels", "Roosters", 8, 12)
	league.MatchResult("Broncos", "Roosters", 32, 18)

	RankPrinter(league, os.Stdout)
}

func RankPrinter(r Ranker, w io.Writer) {
	x := r.Ranking()
	for i := range len(x) {
		_, err := io.WriteString(w, x[i]+"\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(firstTeam, secondTeam string, firstTeamScore, secondTeamScore int) {
	// Initialise map if it is nil
	if l.Wins == nil {
		l.Wins = map[string]int{}
	}

	switch firstTeamScore >= secondTeamScore {
	case true:
		// We do not need the comma ok idiom here as the value is an int and defaults to 0.
		l.Wins[firstTeam]++
		if firstTeamScore == secondTeamScore {
			l.Wins[secondTeam]++
		}
	default:
		l.Wins[secondTeam]++
	}
}

func (l League) Ranking() []string {
	x := make([]string, len(l.Teams))
	for i := range len(x) {
		x[i] = l.Teams[i].name
	}

	sort.Slice(x, func(i, j int) bool {
		return l.Wins[x[i]] < l.Wins[x[j]]
	})

	return x
}

type Ranker interface {
	Ranking() []string
}
