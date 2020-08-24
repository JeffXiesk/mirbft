/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package testengine

import (
	"container/list"
	"io"
	"math/rand"

	"github.com/pkg/errors"

	pb "github.com/IBM/mirbft/mirbftpb"
	"github.com/IBM/mirbft/recorder"
	rpb "github.com/IBM/mirbft/recorder/recorderpb"
)

type EventLog struct {
	// List is a list of *rpb.RecordedEvent messages, in order of time.
	List *list.List

	// LastConsumed is the message most recently removed from the front of the list for consumption.
	LastConsumed *rpb.RecordedEvent

	// FakeTime is the current 'time' according to this log.
	FakeTime int64

	// Rand is a source of randomness for the manglers
	Rand *rand.Rand

	// Mangler give the ability to filter / managle events as they are inserted
	Mangler Mangler

	// Output is optionally a place to serialize RecordedEvents when consumed.
	Output io.Writer
}

func ReadEventLog(source io.Reader) (el *EventLog, err error) {
	reader := recorder.NewReader(source)
	eventLog := &EventLog{
		List: list.New(),
	}

	for {
		stateEvent, err := reader.ReadEvent()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		eventLog.List.PushBack(stateEvent)
	}

	return eventLog, nil
}

func (l *EventLog) ReadEvent() (*rpb.RecordedEvent, error) {
	nele := l.List.Front()
	if nele == nil {
		return nil, io.EOF
	}

	l.LastConsumed = l.List.Remove(nele).(*rpb.RecordedEvent)
	l.FakeTime = l.LastConsumed.Time
	if l.Output != nil {
		err := recorder.WriteRecordedEvent(l.Output, l.LastConsumed)
		if err != nil {
			return nil, errors.WithMessage(err, "could not write event before processing")
		}
	}

	return l.LastConsumed, nil
}

func (l *EventLog) InsertTickEvent(target uint64, fromNow int64) {
	l.InsertStateEvent(
		target,
		&pb.StateEvent{
			Type: &pb.StateEvent_Tick{
				Tick: &pb.StateEvent_TickElapsed{},
			},
		},
		fromNow,
	)
}

func (l *EventLog) InsertProposeEvent(target uint64, req *pb.Request, fromNow int64) {
	l.InsertStateEvent(
		target,
		&pb.StateEvent{
			Type: &pb.StateEvent_Propose{
				Propose: &pb.StateEvent_Proposal{
					Request: req,
				},
			},
		},
		fromNow,
	)
}

func (l *EventLog) InsertStepEvent(target uint64, stepEvent *pb.StateEvent_InboundMsg, fromNow int64) {
	l.InsertStateEvent(
		target,
		&pb.StateEvent{
			Type: &pb.StateEvent_Step{
				Step: stepEvent,
			},
		},
		fromNow,
	)
}

func (l *EventLog) InsertStateEvent(target uint64, stateEvent *pb.StateEvent, fromNow int64) {
	l.Insert(&rpb.RecordedEvent{
		NodeId:     target,
		Time:       l.FakeTime + fromNow,
		StateEvent: stateEvent,
	})
}

func (l *EventLog) InsertProcess(target uint64, fromNow int64) {
	l.InsertStateEvent(
		target,
		&pb.StateEvent{
			Type: &pb.StateEvent_ActionsReceived{
				ActionsReceived: &pb.StateEvent_Ready{},
			},
		},
		fromNow,
	)
}

func (l *EventLog) Insert(initialEvent *rpb.RecordedEvent) {
	var events []*rpb.RecordedEvent
	if l.Mangler != nil {
		events = l.Mangler.Mangle(l.Rand.Int(), initialEvent)
	} else {
		events = []*rpb.RecordedEvent{initialEvent}
	}

	for _, event := range events {
		if event.Time < l.FakeTime {
			panic("attempted to modify the past")
		}

		for el := l.List.Front(); el != nil; el = el.Next() {
			if el.Value.(*rpb.RecordedEvent).Time > event.Time {
				l.List.InsertBefore(event, el)
				return
			}
		}

		l.List.PushBack(event)
	}
}

func (l *EventLog) Count() int {
	return l.List.Len()
}
