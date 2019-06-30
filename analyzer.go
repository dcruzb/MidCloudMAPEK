package MidCloudMAPEK

import (
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"sort"
	"time"
)

func Analyze(chanAnalyzer chan []CloudService, chanPlanner chan CloudService) {
	var first CloudService
	for {
		services := <-chanAnalyzer

		if len(services) > 0 {
			before := fmt.Sprintf("%v", services)

			sort.Sort(SortByPriceAndAvailability(services))

			lib.PrintlnInfo("Analyzer:", before, "=>", services)

			if first != services[0] {
				first = services[0]
				chanPlanner <- first
			}
		}

		// Todo put sleep time of the monitor in the config file
		time.Sleep(5 * time.Second)
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
	return (s[i].Status && s[i].Price < s[j].Price) || !s[j].Status || s[j].Removed
}

// Swap swaps the elements with indexes i and j.
func (s SortByPriceAndAvailability) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
