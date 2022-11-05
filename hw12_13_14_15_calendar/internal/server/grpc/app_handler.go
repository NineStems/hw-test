package grpc

import (
	"context"
	"time"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors"
)

func (s *ServerGRPC) CreateEvent(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	event, err := CreateEventToDomain(req)
	if err != nil {
		return nil, errors.Wrap(err, "CreateEventToDomain")
	}

	id, err := s.app.CreateEvent(ctx, event)
	if err != nil {
		return nil, errors.Wrap(err, "s.app.CreateEvent")
	}

	return &v1.CreateResponse{
		Id: id,
	}, nil
}

func (s *ServerGRPC) UpdateEvent(ctx context.Context, req *v1.UpdateRequest) (*v1.Empty, error) {
	event, err := UpdateEventToDomain(req)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateEventToDomain")
	}

	err = s.app.UpdateEvent(ctx, event)
	if err != nil {
		return nil, errors.Wrap(err, "s.app.UpdateEvent")
	}

	return &v1.Empty{}, nil
}

func (s *ServerGRPC) DeleteEvent(ctx context.Context, req *v1.DeleteRequest) (*v1.Empty, error) {
	err := s.app.DeleteEvent(ctx, []string{req.Id})
	if err != nil {
		return nil, errors.Wrap(err, "s.app.DeleteEvent")
	}

	return &v1.Empty{}, nil
}

func (s *ServerGRPC) DeleteEvents(ctx context.Context, req *v1.DeleteEventsRequest) (*v1.Empty, error) {
	err := s.app.DeleteEvent(ctx, req.Ids)
	if err != nil {
		return nil, errors.Wrap(err, "s.app.DeleteEvent")
	}

	return &v1.Empty{}, nil
}

func (s *ServerGRPC) ReadEvents(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResult, error) {
	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, "time.Parse")
	}

	events, err := s.app.ReadEvents(ctx, date, int(req.Condition))
	if err != nil {
		return nil, errors.Wrap(err, "s.app.ReadEvents")
	}

	return &v1.ReadResult{Events: EventsFromDomain(events)}, nil
}
