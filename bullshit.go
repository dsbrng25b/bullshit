package bullshit

import (
	"bytes"
	"math/rand"
	"time"
)

//go:generate go run gen.go

type WordSet struct {
	Words     []string
	Protocols []string
	Ends      []string
	Suffixes  []string
	Starts    []string
	NoEnds    []string
}

func Get() string {
	return GetWithWordSet(DefaultWordSet)
}

func GetWithWordSet(ws WordSet) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := bytes.Buffer{}
	last := ""
	out := 0
	total := int(rand.Float32()*7) + 3
	hasSuffix := false

	n := int(rand.Float32() * float32(minInt(total-out, 3)))

	for i := 0; i < n; i++ {
		b.WriteString(getRand(ws.Starts) + " ")
	}
	out += n
	hasSuffix = false
	n = int(rand.Float32() * float32(minInt(total-out, 3)))
	for i := 0; i < n; i++ {
		suffix := getSuffix(ws)
		if suffix != "" {
			hasSuffix = true
		}
		word := getRand(ws.Words)
		b.WriteString(word + suffix + " ")
		last = word
	}

	if rand.Float32() > 0.5 {
		n = int(rand.Float32() * 3)
		for i := 0; i < n; i++ {
			b.WriteString(getRand(ws.Protocols) + " ")
			if i != n-1 {
				b.WriteString("over ")
			}
		}
		out += n
		last = ""
		hasSuffix = false
	}
	n = int(rand.Float32() * float32(minInt(total-out, 3)))
	if out+n == 1 || last == "" {
		n += 2
	}
	for i := 0; i < n; i++ {
		suffix := getSuffix(ws)
		if suffix != "" {
			hasSuffix = true
		}
		word := getRand(ws.Words)
		b.WriteString(word + suffix + " ")
		last = word
	}
	if rand.Float32() < 0.1 || (last != "" && contains(ws.NoEnds, last)) || hasSuffix {
		b.WriteString(getRand(ws.Ends) + " ")
	}
	return b.String()
}

func contains(ws []string, word string) bool {
	for _, w := range ws {
		if w == word {
			return true
		}
	}
	return false
}

func getSuffix(ws WordSet) string {
	if rand.Float32() < 0.2 {
		return getRand(ws.Suffixes)
	}
	return ""
}

func getRand(words []string) string {
	return words[rand.Intn(len(words))]
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
