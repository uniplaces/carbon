package examples

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func diff() {
	vancouver, _ := carbon.Today("America/Vancouver")
	london, _ := carbon.Today("Europe/London")
	fmt.Println(vancouver.DiffInSeconds(london, true)) // 0

	ottawa, _ := carbon.CreateFromDate(2000, 1, 1, "America/Toronto")
	vancouver, _ = carbon.CreateFromDate(2000, 1, 1, "America/Vancouver")
	fmt.Println(ottawa.DiffInHours(vancouver, true)) // 3

	fmt.Println(ottawa.DiffInHours(vancouver, false)) // 3
	fmt.Println(vancouver.DiffInHours(ottawa, false)) // -3

	t, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	fmt.Println(t.DiffInDays(t.Copy().AddMonth(), true))  // 31
	fmt.Println(t.DiffInDays(t.Copy().SubMonth(), false)) // -31

	t, _ = carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(t.DiffInDays(t.Copy().AddMonth(), true)) // 30
	fmt.Println(t.DiffInDays(t.Copy().AddWeek(), true))  // 7

	t, _ = carbon.CreateFromTime(10, 1, 1, 0, "UTC")
	fmt.Println(t.DiffInMinutes(t.Copy().AddSeconds(59), true))  // 0
	fmt.Println(t.DiffInMinutes(t.Copy().AddSeconds(60), true))  // 1
	fmt.Println(t.DiffInMinutes(t.Copy().AddSeconds(119), true)) // 1
	fmt.Println(t.DiffInMinutes(t.Copy().AddSeconds(120), true)) // 2
}
