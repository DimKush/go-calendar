package Event

import "time"

type Event struct {
	EndTime   time.Time `json:"end_time"`
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	StartTime time.Time `json:"start_time"`
	Text      string    `json:"text"`
	Title     string    `json:"title"`
}
