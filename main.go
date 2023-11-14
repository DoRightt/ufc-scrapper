package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	model "projects/ufc-scrapper/models"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

var gc *colly.Collector
var detailsCollector *colly.Collector
var collection = model.FightersCollection{}
var wg sync.WaitGroup

func main() {
	gc = colly.NewCollector()
	detailsCollector = gc.Clone()
	url := "https://www.ufc.com/athletes/all"

	gc.OnHTML("div[class*='flipcard__action'] a[href]", parseAthletesListing)
	gc.OnHTML("li.pager__item a[href]", moveNextPage)
	detailsCollector.OnHTML("div[class='hero-profile-wrap']", getData)

	err := gc.Visit(url)
	if err != nil {
		log.Fatalf("Error while request: %v", err)
	}

	wg.Wait()

	fmt.Println("DONE")

	saveToJSON(collection)
}

func parseAthletesListing(e *colly.HTMLElement) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		athleteURL := e.Attr("href")
		athleteURL = e.Request.AbsoluteURL(athleteURL)

		fmt.Println("Athlete link:", athleteURL)

		detailsCollector.Visit(athleteURL)
	}()
}

func getData(e *colly.HTMLElement) {
	wg.Add(1)
	defer wg.Done()

	fighterEl := e.DOM.Parent()

	profileEl := fighterEl.Find("div.hero-profile-wrap")
	statString := profileEl.Find("p.hero-profile__division-body").Text()

	fighter := model.Fighter{
		Name:       profileEl.Find("h1.hero-profile__name").Text(),
		NickName:   profileEl.Find("p.hero-profile__nickname").Text(),
		FighterURL: e.Request.URL.String(),
		ImageUrl:   profileEl.Find(".hero-profile__image-wrap img").AttrOr("src", ""),
	}

	fighter.SetDivision(profileEl.Find("p.hero-profile__division-title").Text())
	fighter.SetStatistic(statString)

	parseData(&fighter, fighterEl)

	collection.Fighters = append(collection.Fighters, fighter)
}

func parseData(f *model.Fighter, fighterEl *goquery.Selection) {
	parseBioFields(f, fighterEl)
	parseMainStats(f, fighterEl)
	parseSpecialStats(f, fighterEl)
	parseWinMethodStats(f, fighterEl)
}

func parseBioFields(f *model.Fighter, fighterEl *goquery.Selection) {
	fields := fighterEl.Find("div.c-bio__info-details")
	fields.Find("div.c-bio__info-details .c-bio__field").Each(func(index int, bioField *goquery.Selection) {
		fieldLabel := bioField.Find(".c-bio__label").Text()
		fieldValue := strings.TrimSpace(bioField.Find(".c-bio__text").Text())

		switch fieldLabel {
		case "Age":
			v, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Age conversion error:", err)
			} else {
				f.Age = int8(v)
			}
		case "Status":
			f.Status = fieldValue
		case "Hometown":
			f.Hometown = fieldValue
		case "Trains at":
			f.TrainsAt = fieldValue
		case "Fighting style":
			f.FightingStyle = fieldValue
		case "Height":
			v, err := strconv.ParseFloat(fieldValue, 32)
			if err != nil {
				fmt.Println("Height conversion error:", err)
			} else {
				f.Height = float32(v)
			}
		case "Weight":
			v, err := strconv.ParseFloat(fieldValue, 32)
			if err != nil {
				fmt.Println("Weight conversion error:", err)
			} else {
				f.Weight = float32(v)
			}
		case "Octagon Debut":
			f.OctagonDebut = fieldValue
		case "Reach":
			v, err := strconv.ParseFloat(fieldValue, 32)
			if err != nil {
				fmt.Println("Reach conversion error:", err)
			} else {
				f.Reach = float32(v)
			}
		case "Leg reach":
			v, err := strconv.ParseFloat(fieldValue, 32)
			if err != nil {
				fmt.Println("Leg Reach conversion error:", err)
			} else {
				f.LegReach = float32(v)
			}
		}
	})
}

func parseMainStats(f *model.Fighter, fighterEl *goquery.Selection) {
	reg := regexp.MustCompile("[^0-9]+")
	fields := fighterEl.Find("div.stats-records-inner-wrap")
	fields.Find("div.c-stat-compare__group").Each(func(index int, bioField *goquery.Selection) {
		fieldLabel := bioField.Find(".c-stat-compare__label").Text()
		fieldValue := strings.TrimSpace(bioField.Find(".c-stat-compare__number").Text())

		switch fieldLabel {
		case "Sig. Str. Landed":
			if fieldValue != "" {
				v, err := strconv.ParseFloat(fieldValue, 32)
				if err != nil {
					fmt.Println("Sig. Str. Landed conversion error:", err)
				} else {
					f.Stats.SigStrLanded = float32(v)
				}
			}
		case "Sig. Str. Absorbed":
			if fieldValue != "" {
				v, err := strconv.ParseFloat(fieldValue, 32)
				if err != nil {
					fmt.Println("Sig. Str. Absorbed conversion error:", err)
				} else {
					f.Stats.SigStrAbs = float32(v)
				}
			}
		case "Sig. Str. Defense":
			numericString := reg.ReplaceAllString(fieldValue, "")
			if numericString != "" {
				v, err := strconv.Atoi(numericString)
				if err != nil {
					fmt.Println("Sig. Str. Defense conversion error:", err)
				} else {
					f.Stats.SigStrDefense = int8(v)
				}
			}
		case "Takedown Defense":
			numericString := reg.ReplaceAllString(fieldValue, "")
			v, err := strconv.Atoi(numericString)
			if err != nil {
				if fieldValue != "" {
					fmt.Println("Takedown Defense conversion error:", err)
				}
			} else {
				f.Stats.TakedownDefense = int8(v)
			}
		case "Takedown avg":
			if fieldValue != "" {
				v, err := strconv.ParseFloat(fieldValue, 32)
				if err != nil {
					fmt.Println("Takedown avg conversion error:", err)
				} else {
					f.Stats.TakedownAvg = float32(v)
				}
			}
		case "Submission avg":
			if fieldValue != "" {
				v, err := strconv.ParseFloat(fieldValue, 32)
				if err != nil {
					fmt.Println("Submission avg conversion error:", err)
				} else {
					f.Stats.SubmissionAvg = float32(v)
				}
			}
		case "Knockdown Avg":
			if fieldValue != "" {
				v, err := strconv.ParseFloat(fieldValue, 32)
				if err != nil {
					fmt.Println("Knockdown Avg conversion error:", err)
				} else {
					f.Stats.KnockdownAvg = float32(v)
				}
			}
		case "Average fight time":
			f.Stats.AvgFightTime = fieldValue
		}
	})
}

func parseSpecialStats(f *model.Fighter, fighterEl *goquery.Selection) {
	fields := fighterEl.Find("div.stats-records-inner-wrap")

	fields.Find("div.c-overlap__inner .c-overlap__stats").Each(func(index int, bioField *goquery.Selection) {
		fieldLabel := bioField.Find("dt.c-overlap__stats-text").Text()
		fieldValue := strings.TrimSpace(bioField.Find("dd.c-overlap__stats-value").Text())

		switch fieldLabel {
		case "Sig. Strikes Landed":
			v, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Total Sig. Strikes Landed conversion error:", err)
			} else {
				f.Stats.TotalSigStrLandned = v
			}
		case "Sig. Strikes Attempted":
			v, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Total Sig. Strikes Attempted conversion error:", err)
			} else {
				f.Stats.TotalSigStrAttempted = v
			}
		case "Takedowns Landed":
			v, err := strconv.Atoi(fieldValue)
			if err != nil {
				if fieldValue != "" {
					fmt.Println("Total Takedowns Landed conversion error:", err)
				}
			} else {
				f.Stats.TotalTkdLanded = v
			}
		case "Takedowns Attempted":
			v, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("Total Takedowns Attempted conversion error:", err)
			} else {
				f.Stats.TotalTkdAttempted = v
			}
		}
	})

	if f.Stats.TotalTkdAttempted != 0 {
		f.Stats.TkdAccuracy = int(float64(f.Stats.TotalTkdLanded) / float64(f.Stats.TotalTkdAttempted) * 100)
	}

	if f.Stats.TotalSigStrAttempted != 0 {
		f.Stats.StrAccuracy = int(float64(f.Stats.TotalSigStrLandned) / float64(f.Stats.TotalSigStrAttempted) * 100)
	}
}

func parseWinMethodStats(f *model.Fighter, el *goquery.Selection) {
	fields := el.Find("div.stats-records-inner-wrap")

	fields.Find("div.stats-records:last-of-type div.stats-records-inner .c-stat-3bar__group").Each(func(index int, bioField *goquery.Selection) {
		fieldLabel := strings.TrimSpace(bioField.Find("div.c-stat-3bar__label").Text())
		fieldValue := strings.TrimSpace(bioField.Find("div.c-stat-3bar__value").Text())

		switch fieldLabel {
		case "KO/TKO":
			v, err := strconv.Atoi(strings.Split(fieldValue, " ")[0])

			if err != nil {
				fmt.Println("KO/TKO data conversion error:", err)
			} else {
				f.Stats.WinByKO = v
			}
		case "DEC":
			v, err := strconv.Atoi(strings.Split(fieldValue, " ")[0])
			if err != nil {
				fmt.Println("DEC data conversion error:", err)
			} else {
				f.Stats.WinByDec = v
			}

		case "SUB":
			v, err := strconv.Atoi(strings.Split(fieldValue, " ")[0])
			if err != nil {
				fmt.Println("SUB data conversion error:", err)
			} else {
				f.Stats.WinBySub = v
			}
		}
	})

}

func moveNextPage(e *colly.HTMLElement) {
	nextUrl := e.Attr("href")
	nextUrl = e.Request.AbsoluteURL(nextUrl)
	fmt.Println("Next page:", nextUrl)

	e.Request.Visit(nextUrl)
}

func saveToJSON(c model.FightersCollection) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error while marshalling:", err)
		return
	}

	file, err := os.Create("fighters.json")
	if err != nil {
		fmt.Println("File openning error:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("File writting error:", err)
		return
	}
}
