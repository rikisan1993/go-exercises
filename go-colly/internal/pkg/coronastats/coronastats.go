package coronastats

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// CoronaStats singular corona stats
type CoronaStats struct {
	Country string `json:"country"`
	TotalCase int `json:"total_case"`
	NewCase int `json:"new_case"`
	TotalDeath int `json:"total_death"`
	NewDeath int `json:"new_death"`
	TotalRecovered int `json:"total_recovered"`
	ActiveCase int `json:"active_case"`
	SeriousCase int `json:"serious_case"`
	CaseInMillion float64 `json:"case_in_million"`
}

// Get run scrapper to collect current stats
func Get() (*[]CoronaStats, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("worldometers.info", "www.worldometers.info"),
		colly.CacheDir("./corona_cache"),
	)

	stats := []CoronaStats{}

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("table#main_table_countries_today tbody tr", func(_ int, el *colly.HTMLElement) {

			totalCase, err := parseStringNumber(el.ChildText("td:nth-child(2)"))
			if err != nil {
				return nil, err
			}

			newCase, err := parseStringNumber(el.ChildText("td:nth-child(3)"))
			if err != nil {
				return nil, err
			}

			totalDeath, err := parseStringNumber(el.ChildText("td:nth-child(4)"))
			if err != nil {
				return nil, err
			}

			newDeath, err := parseStringNumber(el.ChildText("td:nth-child(5)"))
			if err != nil {
				return nil, err
			}

			totalRecovered, err := parseStringNumber(el.ChildText("td:nth-child(6)"))
			if err != nil {
				return nil, err
			}

			activeCase, err := parseStringNumber(el.ChildText("td:nth-child(7)"))
			if err != nil {
				return nil, err
			}

			seriousCase, err := parseStringNumber(el.ChildText("td:nth-child(8)"))
			if err != nil {
				return nil, err
			}

			caseInMillion, err := parseStringNumber(el.ChildText("td:nth-child(9)"))
			if err != nil {
				return nil, err
			}

			stat := CoronaStats{
				Country : el.ChildText("td:nth-child(1)"),
				TotalCase : totalCase,
				NewCase : newCase,
				TotalDeath : totalDeath,
				NewDeath : newDeath,
				TotalRecovered : totalRecovered,
				ActiveCase : activeCase,
				SeriousCase : seriousCase,
				CaseInMillion : caseInMillion,
			}

			stats = append(stats, stat)
		})

		return stats, nil
	})

	c.Visit("https://www.worldometers.info/coronavirus/")
}

func parseStringNumber(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	i, err := strconv.Atoi(strings.Replace(s, ",", "", -1))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func parseStringFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}

	f,err := strconv.ParseFloat(strings.Replace(s, ",", "", -1), 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}