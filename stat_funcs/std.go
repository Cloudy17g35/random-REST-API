package stat_funcs

import (
	"log"
	"github.com/montanaflynn/stats"
)


func StdDev(numbers []float64) (float64, error) {
	r, err := stats.StandardDeviation(numbers)
	if err != nil {
		log.Fatal(err)
	}
	return r, nil

}

