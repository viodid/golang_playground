package main

import "fmt"

type Team struct {
	name        string
	playerNames []string
}

type League struct {
	Teams []string
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) error {
	if score1 > score2 {
		l.Wins[team1] += 1
	} else if score2 > score1 {
		l.Wins[team2] += 1
	}
	return nil
}

func main() {
	t1 := Team{"t1Telecom", []string{"bob", "alice"}}
	t2 := Team{"g2", []string{"foo", "bar"}}
	l := League{[]string{t1.name, t2.name}, map[string]int{}}
	err := l.MatchResult(t1.name, 42, t2.name, 13)
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
}
