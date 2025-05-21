package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"slices"
)

type Team struct {
	name        string
	playerNames []string
}

type League struct {
	Teams []string
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) error {
	var matchName string
	if score1 > score2 {
		matchName = team1
	} else if score2 > score1 {
		matchName = team2
	}
	if slices.Contains(l.Teams, matchName) == true {
		l.Wins[matchName] += 1
		return nil
	}
	return errors.New("Team name is not in League")
}

func (l League) Ranking() []string {
	out := make([]string, 0, len(l.Teams))
	for k, v := range l.Wins {
		fmt.Println(k, v)
	}
	return out
}

func main() {
	t1 := Team{"t1Telecom", []string{"bob", "alice"}}
	t2 := Team{"g2", []string{"foo", "bar"}}
	t3 := Team{"fnatic", []string{"feviben", "hilyssang"}}
	l := League{[]string{t1.name, t2.name, t3.name}, map[string]int{}}
	for i := range 20 {
		r := big.NewInt(100)
		err := l.MatchResult(t1.name, rand.Int(i, r), t2.name, 13)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(t1, t2)
	fmt.Println(l)
	fmt.Println(l.Ranking())
}
