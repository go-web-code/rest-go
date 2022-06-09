package common

import (
	"fmt"
	"time"
)

type JSONTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormart))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	time, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JSONTime(time)
	return
}

func (t JSONTime) String() string {
	return fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormart))
}
