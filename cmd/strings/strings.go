package strings

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"math"
	"strings"
	"unicode"
)

type Options struct {
	InsCost float64
	DelCost float64
	SubCost float64
}

var DefaultOptions = Options{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
}

type (
	Match struct {
		Percentage Distribution
	}
	Distribution struct {
		Levenshtein float64
		JaroWinkler float64
		Media       float64
	}
)

func GetSimilarity(srt1, str2 string) Match {
	var match Match

	levSim := GetLevenshteinSimilarity(srt1, str2, DefaultOptions)
	jaroWSim := GetJaroWinklerSimilarity(srt1, str2)

	match.Percentage = Distribution{
		Levenshtein: levSim,
		JaroWinkler: jaroWSim,
		Media:       (levSim + jaroWSim) / 2,
	}
	return match
}

func GetLevenshteinSimilarity(source, target string, options Options) float64 {

	sourceNorm, targetNorm := strNormalization(source, target)

	return 1 - normalized(sourceNorm, targetNorm, options)
}

func GetJaroWinklerSimilarity(source, target string) float64 {

	sourceNorm, targetNorm := strNormalization(source, target)

	return jaroWinklerDistance(sourceNorm, targetNorm)
}

func normalized(source, target string, options Options) float64 {
	d := levenshteinDistance([]rune(source), []rune(target), options)

	var m = len(source)
	var n = len(target)

	if m > n {
		return d / float64(m)
	}

	return d / float64(n)
}

func levenshteinDistance(source, target []rune, options Options) float64 {
	if len(source) == 0 || len(target) == 0 {
		return 0
	}

	var rows = len(source) + 1
	var columns = len(target) + 1

	distance := make([][]float64, rows)

	for i := range distance {
		distance[i] = make([]float64, columns)
	}

	for i := 0; i < rows; i++ {
		distance[i][0] = float64(i) * options.DelCost
	}

	for j := 0; j < columns; j++ {
		distance[0][j] = float64(j) * options.InsCost
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			deletion := distance[i-1][j] + options.DelCost
			insertion := distance[i][j-1] + options.InsCost
			substitutionOrEqual := distance[i-1][j-1]

			if source[i-1] != target[j-1] {
				substitutionOrEqual += options.SubCost
			}

			distance[i][j] = math.Min(deletion, math.Min(insertion, substitutionOrEqual))
		}
	}

	return distance[rows-1][columns-1]
}

func jaroWinklerDistance(s1, s2 string) float64 {

	s1Matches := make([]bool, len(s1)) // |s1|
	s2Matches := make([]bool, len(s2)) // |s2|

	var matchingCharacters = 0.0
	var transpositions = 0.0

	// sanity checks

	// return 0 if either one is empty string
	if len(s1) == 0 || len(s2) == 0 {
		return 0 // no similarity
	}

	// return 1 if both strings are empty
	if len(s1) == 0 && len(s2) == 0 {
		return 1 // exact match
	}

	if strings.EqualFold(s1, s2) { // case insensitive
		return 1 // exact match
	}

	// Two characters from s1 and s2 respectively,
	// are considered matching only if they are the same and not farther than
	// [ max(|s1|,|s2|) / 2 ] - 1
	matchDistance := len(s1)
	if len(s2) > matchDistance {
		matchDistance = len(s2)
	}
	matchDistance = matchDistance/2 - 1

	// Each character of s1 is compared with all its matching characters in s2
	for i := range s1 {
		low := i - matchDistance
		if low < 0 {
			low = 0
		}
		high := i + matchDistance + 1
		if high > len(s2) {
			high = len(s2)
		}
		for j := low; j < high; j++ {
			if s2Matches[j] {
				continue
			}
			if s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matchingCharacters++
			break
		}
	}

	if matchingCharacters == 0 {
		return 0 // no similarity, exit early
	}

	// Count the transpositions.
	// The number of matching (but different sequence order) characters divided by 2 defines the number of transpositions
	k := 0
	for i := range s1 {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++ // increase transpositions
		}
		k++
	}

	transpositions /= 2

	weight := (matchingCharacters/float64(len(s1)) + matchingCharacters/float64(len(s2)) + (matchingCharacters-transpositions)/matchingCharacters) / 3

	//  the length of common prefix at the start of the string up to a maximum of four characters
	l := 0

	// is a constant scaling factor for how much the score is adjusted upwards for having common prefixes.
	//The standard value for this constant in Winkler's work is {\displaystyle p=0.1}p=0.1
	p := 0.1

	if weight > 0.7 {
		for (l < 4) && s1[l] == s2[l] {
			l++
		}

		weight = weight + float64(l)*p*(1-weight)
	}

	return weight
}

func strNormalization(source string, target string) (string, string) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	sourceNorm, _, _ := transform.String(t, source)
	targetNorm, _, _ := transform.String(t, target)

	sourceNorm = strings.ToLower(sourceNorm)
	targetNorm = strings.ToLower(targetNorm)

	sourceNorm = trimSpace(sourceNorm)
	targetNorm = trimSpace(targetNorm)

	return sourceNorm, targetNorm
}

func trimSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
