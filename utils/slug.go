package utils

import "strings"

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

func Slugify(s string) error {
	strings.Replace(s string, old string, new string) string {
		text := Slug.make()
	}

}
