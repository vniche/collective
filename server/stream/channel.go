package stream

import (
	"github.com/google/uuid"
	serversv1 "github.com/vniche/collective/protocol/v1"
)

// Stream represents a stream for events
type Stream struct {
	ID      string                `json:"id"`
	Channel chan *serversv1.Event `json:"channel"`
}

var managedChannels = map[string]*Stream{}

// New is a constructor for new Stream
func New() *Stream {
	ID := uuid.New().String()
	channel := make(chan *serversv1.Event)
	stream := &Stream{
		ID:      ID,
		Channel: channel,
	}
	managedChannels[ID] = stream

	return stream
}

func PublishToAll(event *serversv1.Event) {
	for _, current := range managedChannels {
		current.Channel <- event
	}
}

// IsClosed returns if the stream is closed or not
func (stream *Stream) IsClosed() bool {
	return managedChannels[stream.ID] == nil
}

// Close tries to close a managed stream channel
func (stream *Stream) Close() {
	if managedChannels[stream.ID] == nil {
		// stream already closed
		return
	}

	// closes stream channel
	close(stream.Channel)

	// deletes stream from managed channels map
	delete(managedChannels, stream.ID)
}

// Shutdown tries to close every managed channel and delete it from managed channels slice
func Shutdown() {
	for _, current := range managedChannels {
		current.Close()
	}
}
