package examples

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func comparison() {
	t1, _ := carbon.CreateFromDate(2010, 10, 1, "Europe/Paris")
	t2, _ := carbon.CreateFromDate(2011, 10, 20, "Europe/Paris")

	fmt.Printf("t1 equal to t2: %t", t1.Eq(t2))
	fmt.Printf("t1 not equal to t2: %t", t1.Ne(t2))

	fmt.Printf("t1 greater than t2: %t", t1.Gt(t2))
	fmt.Printf("t1 lesser than t2: %t", t1.Lt(t2))

	t3, _ := carbon.CreateFromDate(2011, 1, 20, "Europe/Paris")
	fmt.Printf("t3 between t1 and t2: %t", t3.Between(t1, t2, true))

	now := carbon.Now()
	fmt.Printf("Weekday? %t", now.IsWeekday())
	fmt.Printf("Weekend? %t", now.IsWeekend())
	fmt.Printf("LeapYear? %t", now.IsLeapYear())
	fmt.Printf("Past? %t", now.IsPast())
	fmt.Printf("Future? %t", now.IsFuture())
}
