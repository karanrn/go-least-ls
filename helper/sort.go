package helper

import(
	"sort"
	"time"
)

type Pair struct {
	Key string
	Value time.Time
}
  
type PairList []Pair

func (p PairList) Len() int {
    return len(p)
}

func (p PairList) Less(i, j int) bool {
    return p[i].Value.Before(p[j].Value)
}

func (p PairList) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

// Return sorted list based on decreasing timestamp value
func Sort(files map[string]time.Time) PairList{
	// sort files by last access time.
	pairs := make(PairList, len(files))
	i := 0
	for k, v := range files {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(pairs)

	return pairs
}