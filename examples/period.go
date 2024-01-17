package examples

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func period() {
	t1, _ := carbon.Create(2012, 1, 1, 12, 0, 0, 0, "Lisbon/Rome")
	t2, _ := carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Lisbon/Rome")
	days := 7

	periods, err := carbon.Period(t1, days, t2)
	if err != nil {
		return
	}

	for _, val := range periods {
		fmt.Println(val)
	}
}
