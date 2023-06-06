package events

import "github.com/r3labs/sse/v2"

const (
	LOGS   = "logs"   // Logs event channel, contains daemon logs
	STATUS = "status" // status event channel, contains VM load info
)

type EventPublisher interface {
	Publish(event *sse.Event)
}

type EventProducer interface {
	Start(publisher EventPublisher)
	Stop()
}

type eventPublisher struct {
	streamId  string
	sseServer *sse.Server
}

func newEventPublisher(streamId string, server *sse.Server) EventPublisher {
	return &eventPublisher{
		streamId:  streamId,
		sseServer: server,
	}
}

func (ep *eventPublisher) Publish(event *sse.Event) {
	ep.sseServer.Publish(ep.streamId, event)
}
