package storage

import "sort"

// FreqTable defines a frequency map of words
type FreqTable map[string]int

// model defines a map of users to a frequency map of words
type model map[string]FreqTable

// WordKeyValue holds a key-value pair of word and frequecies
// used for sorting
type WordKeyValue struct {
	Word string
	Freq int
}

type orderedWords []WordKeyValue

func (w orderedWords) Len() int {
	return len(w)
}

func (w orderedWords) Less(i, j int) bool {
	return w[i].Freq > w[j].Freq
}

func (w orderedWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

// Ordered returns the words in order of most frequent
func (f FreqTable) Ordered() []WordKeyValue {
	ranked := make(orderedWords, 0, len(f))
	for word, freq := range f {
		ranked = append(ranked, WordKeyValue{word, freq})
	}
	sort.Sort(ranked)
	return ranked
}
