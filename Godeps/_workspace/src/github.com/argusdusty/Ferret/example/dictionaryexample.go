// Copyright 2013 Mark Canning
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
// Author: Mark Canning
// Developed at: Tamber, Inc. (http://www.tamber.com/).
//
// Tamber also has this really cool recommendation engine for music
// (also development by me) which prioritizes up-and-coming artists, so
// it doesn't succomb to the popularity biases that plague modern
// recommendation engines, and still produces excellent personalized
// recommendations! Make sure to check us out at http://www.tamber.com
// or https://itunes.apple.com/us/app/tamber-concerts/id658240483

package main

import (
	"bytes"
	"fmt"
	"github.com/argusdusty/Ferret"
	"io/ioutil"
	"strconv"
	"time"
)

var Correction = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
var LengthSorter = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
var FreqSorter = func(s string, v interface{}, l int, i int) float64 { return float64(v.(uint64)) }
var Converter = ferret.UnicodeToLowerASCII

func main() {
	t := time.Now()
	Data, err := ioutil.ReadFile("dictionary.dat")
	if err != nil {
		panic(err)
	}
	Words := make([]string, 0)
	Values := make([]interface{}, 0)
	for _, Vals := range bytes.Split(Data, []byte("\n")) {
		Vals = bytes.TrimSpace(Vals)
		WordFreq := bytes.Split(Vals, []byte(" "))
		if len(WordFreq) != 2 {
			continue
		}
		Freq, err := strconv.ParseUint(string(WordFreq[1]), 10, 64)
		if err != nil {
			continue
		}
		Words = append(Words, string(WordFreq[0]))
		Values = append(Values, Freq)
	}
	fmt.Println("Loaded dictionary in:", time.Now().Sub(t))
	t = time.Now()

	SearchEngine := ferret.New(Words, Words, Values, Converter)
	fmt.Println("Created index in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("ar", 5))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("test", 5))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.ErrorCorrectingQuery("tsst", 5, Correction))
	fmt.Println("Performed error correcting search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.SortedErrorCorrectingQuery("tssst", 5, Correction, LengthSorter))
	fmt.Println("Performed sorted error correcting search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.SortedErrorCorrectingQuery("tssst", 5, Correction, FreqSorter))
	fmt.Println("Performed sorted error correcting search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.SortedQuery("a", 5, LengthSorter))
	fmt.Println("Performed sorted search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.SortedQuery("a", 5, FreqSorter))
	fmt.Println("Performed sorted search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("a", 5))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("the", 25))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.SortedQuery("the", 25, FreqSorter))
	fmt.Println("Performed sorted search in:", time.Now().Sub(t))
	t = time.Now()
	SearchEngine.Insert("asdfghjklqwertyuiopzxcvbnm", "asdfghjklqwertyuiopzxcvbnm", uint64(0))
	fmt.Println("Performed insert in:", time.Now().Sub(t))
	t = time.Now()
	fmt.Println(SearchEngine.Query("sdfghjklqwert", 5))
	fmt.Println("Performed search in:", time.Now().Sub(t))
	fmt.Println("Running benchmarks...")
	t = time.Now()
	n := 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.Query(string(Query), 5)
		n++
	}
	fmt.Println("Performed", n, "limit-5 searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.Query(string(Query), 25)
		n++
	}
	fmt.Println("Performed", n, "limit-25 searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.SortedQuery(string(Query), 5, LengthSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-5 length sorted searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.SortedQuery(string(Query), 25, LengthSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-25 length sorted searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.SortedQuery(string(Query), 5, FreqSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-5 frequency sorted searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words {
		SearchEngine.SortedQuery(string(Query), 25, FreqSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-25 frequency sorted searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.ErrorCorrectingQuery(string(Query)+"0", 5, Correction)
		n++
	}
	fmt.Println("Performed", n, "limit-5 error correcting searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.ErrorCorrectingQuery(string(Query)+"0", 25, Correction)
		n++
	}
	fmt.Println("Performed", n, "limit-25 error correcting searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.SortedErrorCorrectingQuery(string(Query)+"0", 5, Correction, LengthSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-5 length sorted error correcting searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.SortedErrorCorrectingQuery(string(Query)+"0", 25, Correction, LengthSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-25 length sorted error correcting searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.SortedErrorCorrectingQuery(string(Query)+"0", 5, Correction, FreqSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-5 frequency sorted error correcting searches in:", time.Now().Sub(t))
	t = time.Now()
	n = 0
	for _, Query := range SearchEngine.Words[:2048] {
		SearchEngine.SortedErrorCorrectingQuery(string(Query)+"0", 25, Correction, FreqSorter)
		n++
	}
	fmt.Println("Performed", n, "limit-25 frequency sorted error correcting searches in:", time.Now().Sub(t))
}
