package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) (c Clock) {
	mHours := m / 60
	c.m = m % 60

	if c.m < 0 {
		mHours--
		c.m = 60 + c.m
	}

	c.h = (h + mHours) % 24

	if c.h < 0 {
		c.h = 24 + c.h
	}

	return
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.h, c.m)
}
