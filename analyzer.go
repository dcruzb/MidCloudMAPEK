package MidCloudMAPEK

import (
	"fmt"
	"sort"
	"time"
)

func Analyze(chanAnalyzer chan []CloudService) {
	for {
		services := <-chanAnalyzer

		fmt.Println(services)

		sort.Sort(SortByPriceAndAvailability(services))

		fmt.Println(services)

		time.Sleep(1 * time.Minute)
	}
}

// Type used to sort []CloudService by Price and Availability. Implements sort.Interface
type SortByPriceAndAvailability []CloudService

// Len is the number of elements in the collection.
func (s SortByPriceAndAvailability) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s SortByPriceAndAvailability) Less(i, j int) bool {
	return (s[i].Status && s[i].Price < s[j].Price) || !s[j].Status
}

// Swap swaps the elements with indexes i and j.
func (s SortByPriceAndAvailability) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
