package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/argusdusty/Ferret"
	"github.com/codegangsta/martini"
)

type SearchResponse struct {
	Query   string        `json:"query"`
	Timing  string        `json:"timing"`
	Results []string      `json:"results"`
	Values  []interface{} `json:"values"`
}

func SearchEngine() martini.Handler {
	search := buildFerret()

	return func(c martini.Context) {
		c.Map(search)
		c.Next()
	}
}

const (
	ContextSize  = 12
	ContextAfter = 6
)

var (
	Words          []string
	Values         []interface{}
	RunningContext []string
)

func buildFerret() *ferret.InvertedSuffix {
	t := time.Now()
	// FILE
	file, err := os.Open(*fileFlag)
	if err != nil {
		panic(err)
	}

	// SCAN FILE
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		eachWord(scanner.Text())
	}

	// STATS
	fmt.Println("Loaded document in:", time.Now().Sub(t))
	fmt.Printf("There are %v words.\n", len(Words))

	// INDEX
	t = time.Now()
	SearchEngine := ferret.New(Words, Words, Values, ferret.UnicodeToLowerASCII)
	fmt.Println("Created index in:", time.Now().Sub(t))

	return SearchEngine
}

func eachWord(scannedWord string) {
	contextWord, cleanedWord := cleanWord(scannedWord)
	Words = append(Words, cleanedWord)
	if len(RunningContext) > ContextSize {
		RunningContext = RunningContext[1:]
	}
	RunningContext = append(RunningContext, contextWord)
	context := strings.Join(RunningContext, " ")
	Values = append(Values, context)
	length := len(Values)
	for i := 1; i <= ContextAfter; i++ {
		if length > i {
			offset := i + 1
			Values[length-offset] = context
		}
	}
}

func cleanWord(rawWord string) (before, after string) {
	before = rawWord
	r := strings.NewReplacer("*", "", ",", "", ".", "")
	after = strings.ToLower(r.Replace(rawWord))
	return
}
