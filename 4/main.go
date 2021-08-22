package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := arrangeAnagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"})

	jsonstring, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(jsonstring))
	} else {
		fmt.Println("error")
	}
}

type Anagram struct {
	Origin []string `json:"words"`
	Sorted string   `json:"-"`
}

func arrangeAnagram(input []string) [][]string {
	var compareable []Anagram

	for _, word := range input {
		newChar := strings.Split(word, "")
		sort.Strings(newChar)
		sorted := strings.Join(newChar, "")
		included := false
		indexed := 0

		for i, an := range compareable {
			if an.Sorted == sorted {
				indexed = i
				included = true
			}
		}

		if !included {
			compareable = append(compareable, Anagram{
				Origin: []string{word},
				Sorted: sorted,
			})
		} else {
			compareable[indexed].Origin = append(compareable[indexed].Origin, word)
		}
	}

	prettyPrint := make([][]string, len(compareable))
	for i, an := range compareable {
		prettyPrint[i] = an.Origin
	}

	return prettyPrint
}
