package examples

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func formats() {
	now := carbon.Now()

	fmt.Printf("Date in Atom String: %s\n", now.AtomString())
	fmt.Printf("Date in Cookie String: %s\n", now.CookieString())
	fmt.Printf("Date in ISO 8601 String: %s\n", now.ISO8601String())
	fmt.Printf("Date in Rss String: %s\n", now.RSSString())
	fmt.Printf("Date in W3C String: %s\n", now.W3CString())
}
