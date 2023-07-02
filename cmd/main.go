package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/alextsa22/word-frequency-counter/internal/common"
	"github.com/alextsa22/word-frequency-counter/internal/models"
)

const (
	topWordsCount = 20
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: main <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open file: %v", err)
		return
	}
	defer file.Close()

	trie := models.NewTrie()

	scanner := bufio.NewScanner(file)
	scanner.Split(common.ScanWords)
	for scanner.Scan() {
		w := common.BytesToLower(scanner.Bytes())
		trie.Inc(w)
	}

	words := trie.ToSlice()
	sort.Slice(words, func(i, j int) bool {
		return words[i].Frequency > words[j].Frequency
	})

	for i := 0; i < topWordsCount && i < len(words); i++ {
		fmt.Printf("%d %s\n", words[i].Frequency, words[i].Bytes)
	}
}
