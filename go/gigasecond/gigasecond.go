package gigasecond

import (
	"log"
	"time"
)

const TestVersion = 1

var Birthday, err = time.Parse(fmtD, "1985-01-17")

func AddGigasecond(when time.Time) time.Time {
	gigasecond, err := time.ParseDuration("1000000000s")
	if err != nil {
		log.Fatal("cant happen")
	}
	return when.Add(gigasecond)
}
