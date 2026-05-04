package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	epoch := flag.Int64("epoch", 0, "Unix timestamp to convert")
	now := flag.Bool("now", false, "Show current time in all formats")
	from := flag.String("from", "", "Parse a date string (RFC3339 or common formats)")
	flag.Parse()

	var t time.Time

	switch {
	case *now:
		t = time.Now()
	case *epoch != 0:
		t = time.Unix(*epoch, 0)
	case *from != "":
		var err error
		formats := []string{
			time.RFC3339,
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
			"2006-01-02",
			"01/02/2006",
			"02-Jan-2006",
			time.RFC822,
			time.RFC1123,
		}
		for _, f := range formats {
			t, err = time.Parse(f, *from)
			if err == nil {
				break
			}
		}
		if err != nil {
			// Try as raw epoch
			if v, e := strconv.ParseInt(*from, 10, 64); e == nil {
				t = time.Unix(v, 0)
			} else {
				fmt.Fprintf(os.Stderr, "cannot parse date: %s\n", *from)
				os.Exit(1)
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Usage: timestamp [-now] [-epoch <unix>] [-from <date>]")
		os.Exit(1)
	}

	utc := t.UTC()
	local := t.Local()

	fmt.Printf("Unix Epoch:   %d\n", t.Unix())
	fmt.Printf("Unix Milli:   %d\n", t.UnixMilli())
	fmt.Printf("UTC:          %s\n", utc.Format(time.RFC3339))
	fmt.Printf("Local:        %s\n", local.Format(time.RFC3339))
	fmt.Printf("RFC822:       %s\n", utc.Format(time.RFC822))
	fmt.Printf("RFC1123:      %s\n", utc.Format(time.RFC1123))
	fmt.Printf("Human:        %s\n", utc.Format("Monday, 02 January 2006 15:04:05 UTC"))
	fmt.Printf("Log Format:   %s\n", utc.Format("02/Jan/2006:15:04:05 +0000"))
}
