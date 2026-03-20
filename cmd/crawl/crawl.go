// I don't know why it gets status code 403. Maybe because CF layer.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type spell struct {
	Name          string   `json:"name"`
	Type          []string `json:"type,omitempty"`
	Pronunciation string   `json:"pronunciation,omitempty"`
	Description   string   `json:"description,omitempty"`
}

func main() {
	url := "https://harrypotter.fandom.com/wiki/List_of_spells?action=render"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status: %d", resp.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	spells := parse(doc)
	fmt.Println(spells)
}

func parse(doc *goquery.Document) []spell {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var spells []spell
	for _, l := range letters {
		start := doc.Find("span.mw-headline#" + string(l)).First()
		if start.Length() == 0 {
			continue
		}
		h2 := start.Parent()
		for s := h2.Next(); s.Length() > 0; s = s.Next() {
			if goquery.NodeName(s) == "h2" {
				break
			}
			if goquery.NodeName(s) != "h3" {
				continue
			}
			var spell spell
			name := strings.TrimSpace(s.Find("i").Text())
			spell.Name = name
			dl := s.NextAllFiltered("dl").First()
			dl.Find("dd").Each(func(_ int, dd *goquery.Selection) {
				text := clean(dd.Text())
				switch {
				case strings.HasPrefix(text, "Type:"):
					val := strings.TrimPrefix(text, "Type:")
					spell.Type = split(val)
				case strings.HasPrefix(text, "Pronunciation:"):
					spell.Pronunciation = strings.TrimPrefix(text, "Pronunciation:")
				case strings.HasPrefix(text, "Description:"):
					spell.Description = strings.TrimPrefix(text, "Description:")
				}
			})
			spells = append(spells, spell)
		}
	}
	return spells
}

func clean(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	return strings.TrimSpace(s)
}

func split(s string) []string {
	parts := strings.Split(s, ",")
	var res []string
	for _, p := range parts {
		res = append(res, strings.TrimSpace(p))
	}
	return res
}
