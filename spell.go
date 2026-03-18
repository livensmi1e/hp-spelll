package main

import "os"

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
	"accio":              "",
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

func cast(spell string) {
	N := threshold(spell)
	S := make([]string, 0, N)
	T := 5
	for _, s := range HARRY_POTTER_SPELLS {
		d := levenshtein(spell, s)
		if d == 0 {
			println(HARRY_POTTER_SPELLS[s])
			return
		}
		if d <= N {
			S = append(S, s)
		}
	}
	suggest(S[:T]...)
}

func levenshtein(a, b string) int {
	return 0
}

func suggest(spells ...string) {
	print("The most similar spell")
	if len(spells) > 1 {
		println("s are: ")
	} else {
		println(" is: ")
	}
	for _, spell := range spells {
		println("\t" + spell)
	}
}

func threshold(spell string) int {
	return max(2, len(spell)/3)
}
