package examples

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func addSub() {
	t1, _ := carbon.Create(2010, 1, 4, 19, 10, 10, 0, "UTC")

	t1 = t1.AddYear()     // 2011-01-04 19:10:10
	t1 = t1.AddMonth()    // 2011-02-04 19:10:10
	t1 = t1.SubDay()      // 2011-02-03 19:10:10
	t1 = t1.SubYears(10)  // 2001-02-03 19:10:10
	t1 = t1.AddMinutes(5) // 2001-02-03 19:15:10
	t1 = t1.AddWeekday()  // 2001-02-05 19:15:10
	t1 = t1.SubHours(10)  // 2001-02-05 10:15:10

	fmt.Println(t1)
}
