package types

import "time"

type Time int64

func (t Time) String() string {
	return t.Time().Format("2006-01-02 15:04:05")
}

func (t Time) Unix() int64 {
	return int64(t)
}

func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0)
}
