package main

type Team struct {
	teamName    string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[int]Team
}
