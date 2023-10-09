package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Division int

func (d Division) String() string {
	switch d {
	case Flyweight:
		return "Flyweight"
	case Bantamweight:
		return "Bantamweight"
	case Featherweight:
		return "Featherweight"
	case Lightweight:
		return "Lightweight"
	case Welterweight:
		return "Welterweight"
	case Middleweight:
		return "Middleweight"
	case Lightheavyweight:
		return "Light Heavyweight"
	case Heavyweight:
		return "Heavyweight"
	case WomensStrawweight:
		return "Women's Strawweight"
	case WomensFlyweight:
		return "Women's Flyweight"
	case WomensBantamweight:
		return "Women's Bantamweight"
	case WomensFeatherweight:
		return "Women's Featherweight"
	default:
		return "Unknown"
	}
}

const (
	Flyweight Division = iota
	Bantamweight
	Featherweight
	Lightweight
	Welterweight
	Middleweight
	Lightheavyweight
	Heavyweight
	WomensStrawweight
	WomensFlyweight
	WomensBantamweight
	WomensFeatherweight
)

type FightersCollection struct {
	Fighters []Fighter
}

type FighterStats struct {
	TotalSigStrLandned   int
	TotalSigStrAttempted int
	StrAccuracy          int
	TotalTkdLanded       int
	TotalTkdAttempted    int
	TkdAccuracy          int
	SigStrLanded         float32
	SigStrAbs            float32
	SigStrDefense        int8
	TakedownDefense      int8
	TakedownAvg          float32
	SubmissionAvg        float32
	KnockdownAvg         float32
	AvgFightTime         string
	WinByKO              int
	WinBySub             int
	WinByDec             int
}

type Fighter struct {
	Name          string
	NickName      string
	Division      Division
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
	Stats         FighterStats
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

func (f *Fighter) SetDivision(d string) {
	switch d {
	case "Flyweight Division":
		f.Division = Flyweight
	case "Bantamweight Division":
		f.Division = Bantamweight
	case "Featherweight Division":
		f.Division = Featherweight
	case "Lightweight Division":
		f.Division = Lightweight
	case "Welterweight Division":
		f.Division = Welterweight
	case "Middleweight Division":
		f.Division = Middleweight
	case "Light Heavyweight Division":
		f.Division = Lightheavyweight
	case "Heavyweight Division":
		f.Division = Heavyweight
	case "Women's Strawweight Division":
		f.Division = WomensStrawweight
	case "Women's Flyweight Division":
		f.Division = WomensFlyweight
	case "Women's Bantamweight Division":
		f.Division = WomensBantamweight
	case "Women's Featerweight Division":
		f.Division = WomensFeatherweight
	}
}
