package main

import (
	"fmt"
	"log"
	model "projects/ufc-scrapper/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

var gc *colly.Collector

func main() {
	gc = colly.NewCollector()
	url := "https://www.ufc.com/athletes/all"

	gc.OnHTML("div[class*='flipcard__action'] a[href]", parseAthletesListing)
	gc.OnHTML("div[class='hero-profile-wrap']", getData)
	// c.OnHTML("li.pager__item a[href]", moveNextPage)

	err := gc.Visit(url)
	if err != nil {
		log.Fatalf("Error while request: %v", err)
	}
}

func parseAthletesListing(e *colly.HTMLElement) {
	athleteURL := e.Attr("href")
	athleteURL = e.Request.AbsoluteURL(athleteURL)

	fmt.Println("Athlete link:", athleteURL)

	e.Request.Visit(athleteURL)
}

func getData(e *colly.HTMLElement) {
	fighterEl := e.DOM.Parent()

	profileEl := fighterEl.Find("div.hero-profile-wrap")
	statString := profileEl.Find("p.hero-profile__division-body").Text()

	// statsEl := fighterEl.Find("div.stats-record-wrap")
	// recordEl := fighterEl.Find("div.athlete-record")
	bioFields := fighterEl.Find("div.c-bio__info-details")

	fighter := model.Fighter{
		Name: profileEl.Find("h1.hero-profile__name").Text(),
	}

	fighter.SetStatistic(statString)

	bioFields.Find("div.c-bio__info-details .c-bio__field").Each(func(index int, bioField *goquery.Selection) {
		fieldLabel := bioField.Find(".c-bio__label").Text()
		fieldValue := bioField.Find(".c-bio__text").Text()

		switch fieldLabel {
		case "Age":
			fighter.Age = fieldValue
		case "Status":
			fighter.Status = fieldValue
		case "Hometown":
			fighter.Hometown = fieldValue
		case "Trains at":
			fighter.TrainsAt = fieldValue
		case "Fighting style":
			fighter.FightingStyle = fieldValue
		case "Height":
			fighter.Height = fieldValue
		case "Weight":
			fighter.Weight = fieldValue
		case "Octagon Debut":
			fighter.OctagonDebut = fieldValue
		case "Reach":
			fighter.Reach = fieldValue
		case "Leg reach":
			fighter.LegReach = fieldValue
		}
	})

	fmt.Println(fighter)
}

func moveNextPage(e *colly.HTMLElement) {
	nextUrl := e.Attr("href")
	nextUrl = e.Request.AbsoluteURL(nextUrl)
	fmt.Println("Next page:", nextUrl)

	e.Request.Visit(nextUrl)
}
