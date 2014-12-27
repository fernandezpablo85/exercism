package clock

import "fmt"

const TestVersion = 1

type Clock struct {
	minutes int
}

func New(h int, m int) Clock {
	minutes := m + h*60
	return Clock{normalize(minutes)}
}

func (c Clock) String() string {
	h, m := c.toHoursAndMinutes()
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (c Clock) Add(m int) Clock {
	return Clock{normalize(c.minutes + m)}
}

func (c Clock) toHoursAndMinutes() (int, int) {
	return (c.minutes / 60) % 24, c.minutes % 60
}

func normalize(minutes int) int {
	mod := minutes % 1440
	if mod < 0 {
		return mod + 1440
	} else {
		return mod
	}
}
