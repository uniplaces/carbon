package examples

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func instantiation() {
	now := carbon.NewCarbon(time.Now())
	fmt.Printf("New carbon from time instance: %s\n", now)

	now = carbon.Now()
	fmt.Printf("New carbon now: %s\n", now)

	fromDate, _ := carbon.CreateFromDate(2000, 1, 1, "Europe/London")
	fmt.Printf("Created from date: %s\n", fromDate)

	fromTime, _ := carbon.CreateFromTime(9, 16, 11, 0, "Europe/Madrid")
	fmt.Printf("Created from time: %s\n", fromTime)

	parsed, _ := carbon.Parse(carbon.DateFormat, "2000-08-20", "Europe/Berlin")
	fmt.Printf("Parsed time: %s\n", parsed)
}
