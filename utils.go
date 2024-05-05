package main

import (
	"encoding/json"
	"math/rand"
	"os"
	model "projects/ufc-scrapper/models"
	"strconv"
	"strings"
	"time"
)

func createNewCollection(c model.FightersCollection) {
	file, err := os.Create("fighters.json")
	if err != nil {
		l.Error("File creation error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		l.Error("Error encoding JSON:", err)
		return
	}
}

func addToExistedCollection(c model.FightersCollection) {
	file, err := os.Open("fighters.json")
	if err != nil {
		l.Error("File opening error:", err)
		return
	}
	defer file.Close()

	var existingFighters model.FightersCollection
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&existingFighters); err != nil {
		l.Error("Error decoding JSON:", err)
		return
	}

	existingFighters.Fighters = append(existingFighters.Fighters, c.Fighters...)
	collection := getUniqueCollection(existingFighters)

	file, err = os.Create("fighters.json")
	if err != nil {
		l.Error("File creation error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(collection); err != nil {
		l.Error("Error encoding JSON:", err)
		return
	}
}

func getUniqueCollection(c model.FightersCollection) model.FightersCollection {
	uniqueFightersMap := make(map[string]model.Fighter)
	uniqueFighters := make([]model.Fighter, 0, 500)

	for _, fighter := range c.Fighters {
		key := fighter.Name + fighter.NickName + strconv.Itoa(fighter.DebutTimestamp)
		if _, exists := uniqueFightersMap[key]; !exists {
			uniqueFightersMap[key] = fighter
			uniqueFighters = append(uniqueFighters, fighter)
		}
	}

	return model.FightersCollection{
		Fighters: uniqueFighters,
	}
}

func getDebutTimestamp(octagonDebut string) int {
	layout := "Jan. 2, 2006"

	parsedTime, err := time.Parse(layout, octagonDebut)
	if err != nil {
		l.Error("Error while date parsing:", err)
		return 0
	}

	return int(parsedTime.Unix())
}

func getProxy() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(proxys))

	return proxys[idx]
}

func setProxys() {
	envProxys := os.Getenv("PROXYS")
	proxys = strings.Split(envProxys, "/")
}

func getLoggerFlag(toAdd bool) int {
	if toAdd {
		return os.O_APPEND
	} else {
		return os.O_TRUNC
	}
}
