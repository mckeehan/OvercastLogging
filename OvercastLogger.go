package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg OvercastConfig

func main() {
	var err error
	var x Opml
	var targetDate = time.Now()

	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		targetDate, err = time.Parse("2006-01-02", os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Fetching Overcast data")
	err = xml.Unmarshal([]byte(GetOpml(cfg)), &x)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Processing for date %s\n", targetDate.Format("2006-01-02"))
	results := processOpml(targetDate, x)
	if results != nil {
		fmt.Printf("Adding content to %s\n", targetDate.Format("2006-01-02") )
		// writeToFile(fmt.Sprintf("/Users/mckeehan/Sync/Obsidian/DailyAgenda/%s.md", targetDate.Format("2006-01-02")), results)
		for _, e	:= range results {
			fmt.Println(e)
		}
	}
}

func writeToFile(filename string, results []string) {
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, e	:= range results {
		if _, err = f.WriteString(e); err != nil {
			panic(err)
		}
	}
}

func processOpml(targetDate time.Time, opml Opml) []string {
	var results []string
	for _, group := range opml.Body.Outline {
		if( group.AttrText == "feeds" ) {
			for _, feed := range group.Podcasts {
				podcastResults := processFeed(targetDate, feed)
				results = append(results, podcastResults...)
			}
		}
	}
	return results
}

func processFeed(targetDate time.Time,  podcast Podcast ) []string {
	var results []string;

	// fmt.Printf("	%s\n", p.AttrText)
	for _, episode := range podcast.Episodes {
		episodeResults := processEpisode(targetDate, podcast, episode)
		if episodeResults != nil {
			results = append(results, episodeResults... )
		}
	}

	return results
}

func processEpisode(targetDate time.Time, p Podcast, e Episode) []string {
	var results []string

	episodeDate, error := time.Parse(time.RFC3339, e.UserUpdatedDate)

	if error != nil {
		fmt.Println(error)
		return nil
	}

	if e.Played == "1" && ( targetDate.Year() == episodeDate.Year() && targetDate.YearDay() == episodeDate.YearDay() ) {
		results = append(results, fmt.Sprintf("- log-podcast:: [%s -  %s](%s)", p.Title, e.Title, e.URL) )
	}

	return results
}
