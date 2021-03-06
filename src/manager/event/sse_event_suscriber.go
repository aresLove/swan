package event

import (
	"net/http"

	"github.com/manucorporat/sse"
)

type SSESubscriber struct {
	Key    string
	rw     http.ResponseWriter
	doneCh chan struct{}
}

func NewSSESubscriber(key string, rw http.ResponseWriter) (*SSESubscriber, chan struct{}) {
	sseSubscriber := &SSESubscriber{
		Key:    key,
		rw:     rw,
		doneCh: make(chan struct{}),
	}

	if c, ok := sseSubscriber.rw.(http.CloseNotifier); ok {
		go func() {
			for {
				select {
				case <-c.CloseNotify():
					sseSubscriber.doneCh <- struct{}{}
				}
			}
		}()
	} else { // exit fast
		sseSubscriber.doneCh <- struct{}{}
	}

	return sseSubscriber, sseSubscriber.doneCh
}

func (sse *SSESubscriber) Subscribe(bus *EventBus) error {
	bus.Subscribers[sse.Key] = sse
	return nil
}

func (sse *SSESubscriber) Unsubscribe(bus *EventBus) error {
	delete(bus.Subscribers, sse.Key)
	return nil
}

func (sses *SSESubscriber) Write(e *Event) error {
	sse.Encode(sses.rw, sse.Event{
		Id:    e.Id,
		Event: e.Type,
		Data:  e.Payload,
	})

	// TODO
	if f, ok := sses.rw.(http.Flusher); ok {
		f.Flush()
	}

	return nil
}

func (sse *SSESubscriber) InterestIn(e *Event) bool {
	if e.Type == EventTypeTaskAdd {
		return true
	}

	if e.Type == EventTypeTaskRm {
		return true
	}

	return false
}
