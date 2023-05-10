package model

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrNotFound        = errors.New("event not found")
	ErrInvalidInterval = errors.New("interval invalid")
)

type Event struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        JSONTime `json:"date"`
}

type JSONTime time.Time

const (
	timeFormat = "02-01-2006"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), `"`)
	newTime, err := time.ParseInLocation(timeFormat, string(s), time.Local)
	*t = JSONTime(newTime)
	return
}
func (t JSONTime) String() string {
	return time.Time(t).Format(timeFormat)
}
