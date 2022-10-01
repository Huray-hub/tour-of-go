package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    word_counters := make(map[string]int)

    for _,word:= range words{
        _, ok := word_counters[word]
        if ok {
            word_counters[word]++
            continue
        } 

        word_counters[word]=1
    }
   return word_counters
}

func main() {
	wc.Test(WordCount)
}
