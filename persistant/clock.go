package persistant

import "time"

//type Clock struct {
//	Time time.Time
//}
//
//func New() Clock {
//	return Clock{Time: time.Now()}
//}
type Clock interface {
	Now() time.Time
	//After(d time.Duration) <-chan time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

//func (RealClock) After(d time.Duration) <-chan time.Time { return time.After(d) }
