package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Fighter struct {
	Name          string
	Status        string
	Hometown      string
	TrainsAt      string
	FightingStyle string
	Age           string
	Height        string
	Weight        string
	OctagonDebut  string
	Reach         string
	LegReach      string
	Wins          int
	Loses         int
	Draw          int
}

func (f *Fighter) SetStatistic(stat string) {
	parts := strings.Split(strings.Split(stat, " ")[0], "-")

	var scores []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Conversion error:", err)
			scores = append(scores, 0)
			return
		}
		scores = append(scores, num)
	}

	f.Wins = scores[0]
	f.Loses = scores[1]
	f.Draw = scores[2]
}
