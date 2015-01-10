package house

import "strings"

func Verse(a string, b []string, c string) string {
	verse := make([]string, 0, 70)
	verse = append(verse, a)
	if len(b) > 0 {
		verse = append(verse, strings.Join(b, " "))
	}
	verse = append(verse, c)
	return strings.Join(verse, " ")
}

func Embed(a string, b string) string {
	return a + " " + b
}

func Song() string {
	ps := []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	start := "This is"
	nounPhrase := "the house that Jack built."
	s := start + " " + nounPhrase
	for c := len(ps) - 1; c >= 0; c-- {
		s += "\n\n" + Verse(start, ps[c:], nounPhrase)
	}
	return s
}
