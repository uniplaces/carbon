package examples

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
)

func modifiers() {
	t1, _ := carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Lisbon/Rome")
	fmt.Printf("Start of day:%s\n", t1.StartOfDay()) // 2012-01-31 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of day:%s\n", t1.EndOfDay()) // 2012-01-31 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Start of month:%s\n", t1.StartOfMonth()) // 2012-01-01 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of month:%s\n", t1.EndOfMonth()) // 2012-01-31 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Start of year:%s\n", t1.StartOfYear()) // 2012-01-01 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of year:%s\n", t1.EndOfYear()) // 2012-12-31 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Start of decade:%s\n", t1.StartOfDecade()) // 2010-01-01 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of decade:%s\n", t1.EndOfDecade()) // 2019-12-31 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Start of century:%s\n", t1.StartOfCentury()) // 2000-01-01 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of century:%s\n", t1.EndOfCentury()) // 2099-12-31 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Start of week:%s\n", t1.StartOfWeek()) // 2012-01-30 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("End of week:%s\n", t1.EndOfWeek()) // 2012-02-05 23:59:59

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Next:%s\n", t1.Next(time.Wednesday)) // 2012-02-01 00:00:00

	t1, _ = carbon.Create(2012, 1, 31, 12, 0, 0, 0, "Europe/Rome")
	fmt.Printf("Previous:%s\n", t1.Previous(time.Wednesday)) // 2012-01-25 00:00:00
}
