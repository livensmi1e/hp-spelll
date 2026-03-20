package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// To be updated
var HARRY_POTTER_SPELLS = map[string]string{
	"lumos":                 "Turns on light at wand tip",
	"alohomora":             "Unlocks doors or objects",
	"wingardium leviosa":    "Makes objects levitate",
	"expelliarmus":          "Disarms opponent's weapon or wand",
	"rictusempra":           "Causes uncontrollable laughter",
	"serpensortia":          "Conjures a snake from wand",
	"expecto patronum":      "Summons a Patronus to repel Dementors",
	"lumos maxima":          "Creates a very bright light",
	"accio":                 "Summons an object toward caster",
	"oculus reparo":         "Repairs glasses or broken items",
	"prior incantato":       "Reveals previous spells cast by wand",
	"stupefy":               "Stuns or knocks out opponent",
	"impedimenta":           "Slows or stops a moving object",
	"protego":               "Creates a protective shield",
	"reducto":               "Destroys or shrinks an object",
	"levicorpus":            "Hoists a person into the air by their legs",
	"liberacorpus":          "Releases someone from levicorpus",
	"sectumsempra":          "Inflicts deep wounds on opponent",
	"muffliato":             "Fills ears of nearby people with buzzing to keep conversation private",
	"aparecium":             "Reveals invisible ink or hidden messages",
	"homenum revelio":       "Detects human presence nearby",
	"imperio":               "Controls actions of another person",
	"fiendfyre":             "Conjures uncontrollable magical fire",
	"water to rum spell":    "Transforms water into rum",
	"avada kedavra":         "Instantly kills opponent (Unforgivable Curse)",
	"crucio":                "Inflicts unbearable pain (Unforgivable Curse)",
	"confundo":              "Confuses target",
	"obliviate":             "Erases memories",
	"colloportus":           "Magically locks doors",
	"diffindo":              "Cuts or rips objects",
	"incendio":              "Creates fire",
	"nox":                   "Turns off light at wand tip",
	"alohomora maximus":     "Stronger version of unlocking spell",
	"engorgio":              "Enlarges an object",
	"reparo":                "Repairs broken objects",
	"arresto momentum":      "Slows falling objects or people",
	"glacius":               "Freezes objects",
	"sonorus":               "Amplifies voice",
	"silencio":              "Silences target",
	"tarantallegra":         "Makes target's legs dance uncontrollably",
	"protego maxima":        "Stronger protective shield",
	"sectumsempra advanced": "More powerful cutting spell",
}

var (
	prev   []int
	curr   []int
	maxLen = maxSpellLen()
)

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Go back to Hogwarts and learn some spells!")
		return
	}
	spell := args[1]
	result := doCast(spell)
	cast(result)
}

// Sort is overkill
func doCast(spell string) string {
	threshold := threshold(spell)
	candidates := []string{}
	prev = make([]int, maxLen+1)
	curr = make([]int, maxLen+1)
	for s := range HARRY_POTTER_SPELLS {
		d := levenshtein(spell, s)
		if d == 0 {
			return HARRY_POTTER_SPELLS[s]
		}
		if d <= threshold {
			candidates = append(candidates, s)
		}
	}
	if len(candidates) == 0 {
		return "Avada Kedavra"
	}
	top := min(2, len(candidates))
	slices.Sort(candidates)
	return suggest(candidates[:top]...)
}

func levenshtein(a, b string) int {
	Na := len(a)
	Nb := len(b)
	if cap(prev) < Nb+1 {
		prev = make([]int, Nb+1)
		curr = make([]int, Nb+1)
	}
	prev = prev[:Nb+1]
	curr = curr[:Nb+1]
	for j := 0; j <= Nb; j++ {
		prev[j] = j
	}
	for i := 1; i <= Na; i++ {
		curr[0] = i
		for j := 1; j <= Nb; j++ {
			s := 0
			if a[i-1] != b[j-1] {
				s = 1
			}
			curr[j] = min(prev[j-1]+s, min(prev[j]+1, curr[j-1]+1))
		}
		prev, curr = curr, prev
	}
	return prev[Nb]
}

func suggest(spells ...string) string {
	var b strings.Builder
	b.WriteString("The most similar spell")
	if len(spells) > 1 {
		b.WriteString("s are: \n")
	} else {
		b.WriteString(" is: \n")
	}
	for _, spell := range spells {
		b.WriteString("\t")
		b.WriteString(spell)
		b.WriteString("\n")
	}
	return b.String()
}

func threshold(spell string) int {
	return max(2, len(spell)/3)
}

func maxSpellLen() int {
	m := 0
	for s := range HARRY_POTTER_SPELLS {
		m = max(m, len(s))
	}
	return m
}

func cast(result string) {
	fmt.Println(result)
}
