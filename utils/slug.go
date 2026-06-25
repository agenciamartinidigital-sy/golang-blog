package utils

import (
	"regexp"
	"strings"
)

var replacements = strings.NewReplacer(
	"á", "a",
	"ã", "a",
	"à", "a",
	"ä", "a",
	"â", "a",
	"é", "e",
	"ê", "e",
	"ë", "e",
	"í", "i",
	"ó", "o",
	"ô", "o",
	"ú", "u",
	"Á", "a",
	"Â", "a",
	"À", "a",
	"Ä", "a",
	"Ã", "a",
	"É", "e",
	"Ê", "e",
	"Ë", "e",
	"Í", "i",
	"Ó", "o",
	"Ô", "o",
	"Ú", "u",
	"ç", "c",
	"ñ", "n", // usado em espanhol
)

var rNunAlunum = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(s string) string {
	s = replacements.Replace(s)
	s = strings.ToLower(s)
	s = rNunAlunum.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
