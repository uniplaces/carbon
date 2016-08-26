Carbon
======

A simple extension for [Time](https://golang.org/pkg/time/#Time) based on php's [carbon](http://carbon.nesbot.com) library.

# Example
A simple example to get you started:
```go
package main

import (
	"fmt"
	"time"

	carbon "github.com/uniplaces/carbon"
)

func main() {
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	today, _ := carbon.NowInLocation("America/Vancouver")
	fmt.Printf("Right now in Vancouver is %s\n", today)

	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	nextOlmypics := carbon.CreateFromDate(2016, time.August, 5, time.UTC).AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlmypics.Year())

	if carbon.Now().IsWeekend() {
		fmt.Printf("Party time!")
	}
}
```

# Installation
To install carbon run:
```
go get github.com/uniplaces/carbon
```

#Contributing
Please feel free to make suggestions, create issues, fork the repository and send pull requests!
