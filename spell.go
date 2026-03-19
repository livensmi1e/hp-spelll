package main

import (
	"fmt"
	"os"
	"slices"
)

// To be updated
var HARRY_POTTER_SPELLS = map[string]string{
	"lumos":              "",
	"alohomora":          "",
	"wingardium leviosa": "",
	"expelliarmus":       "",
	"rictusempra":        "",
	"serpensortia":       "",
	"expecto patronum":   "",
	"lumos maxima":       "",
	"accio":              "summon object toward caster",
	"oculus reparo":      "",
	"prior incantato":    "",
	"stupefy":            "",
	"impedimenta":        "",
	"protego":            "",
	"reducto":            "",
	"levicorpus":         "",
	"liberacorpus":       "",
	"sectumsempra":       "",
	"muffliato":          "",
	"aparecium":          "",
	"homenum revelio":    "",
	"imperio":            "",
	"fiendfyre":          "",
	"water to rum spell": "",
}

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Go back to Hogwarts and learn some spells!")
		return
	}
	spell := args[1]
	cast(spell)
}

// O(n*max(len(s)^2))
// len(s) <= 50
// n <= 1000
func cast(spell string) {
	N := threshold(spell)
	S := []string{}
	for s := range HARRY_POTTER_SPELLS {
		d := levenshtein(spell, s)
		if d == 0 {
			fmt.Println(HARRY_POTTER_SPELLS[s])
			return
		}
		if d <= N {
			S = append(S, s)
		}
	}
	T := min(2, len(S))
	slices.Sort(S)
	suggest(S[:T]...)
}

func levenshtein(a, b string) int {
	Na := len(a)
	Nb := len(b)
	if Na == 0 {
		return Nb
	}
	if Nb == 0 {
		return Na
	}
	d := make([][]int, Na+1)
	for i := 0; i <= Na; i++ {
		d[i] = make([]int, Nb+1)
	}
	for i := 1; i <= Na; i++ {
		d[i][0] = i
	}
	for j := 1; j <= Nb; j++ {
		d[0][j] = j
	}
	for i := 1; i <= Na; i++ {
		for j := 1; j <= Nb; j++ {
			s := 0
			if a[i-1] != b[j-1] {
				s = 1
			}
			d[i][j] = min(d[i-1][j-1]+s, min(d[i-1][j]+1, d[i][j-1]+1))
		}
	}
	return d[Na][Nb]
}

func suggest(spells ...string) {
	fmt.Print("The most similar spell")
	if len(spells) > 1 {
		fmt.Println("s are: ")
	} else {
		fmt.Println(" is: ")
	}
	for _, spell := range spells {
		fmt.Println("\t", spell)
	}
}

func threshold(spell string) int {
	return max(2, len(spell)/3)
}
