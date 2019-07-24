package persistant

import "time"

type Clock struct {
	Time time.Time
}

func New() Clock {
	return Clock{Time: time.Now()}
}
