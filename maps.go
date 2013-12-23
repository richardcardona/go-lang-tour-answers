package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wordmap := make(map[string]int)
    components := strings.Fields(s)
    for _,word := range components {
    	wordmap[word] += 1    
    }
    return wordmap
}

func main() {
    wc.Test(WordCount)
}