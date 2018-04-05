package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Season struct {
	Name   string  `json:"name"`
	Rounds []Round `json:"rounds"`
}

type Round struct {
	Name    string  `json:"name"`
	Matches []Match `json:"matches"`
}

type Match struct {
	Date  string `json:"date"`
	Team1 Team   `json:"team1"`
	Team2 Team   `json:"team2"`
}

type Team struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func readSeason() Season {
	b, err := ioutil.ReadFile("football.json")
	if err != nil {
		log.Panicf("Failed to read football.json\n")
	}

	var season Season
	err = json.Unmarshal(b, &season)
	if err != nil {
		log.Panicf("Failed to unmarshal football.json\n")
	}

	return season
}
