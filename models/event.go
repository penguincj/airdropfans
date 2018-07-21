package models

import (
	"container/list"

	"github.com/gorilla/websocket"
)

type EventType int

const (
	EVENT_MARKET = iota
	EVENT_TEMPERATURE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType // EVETNT_PRICE, LEAVE, MESSAGE
	Id        int64
	Conn      *websocket.Conn
	Timestamp int // Unix timestamp (secs)
	Data      interface{}
}

const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

// GetEvents returns all events after lastReceived.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
