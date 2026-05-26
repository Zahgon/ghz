package gtime

import (
	"context"
	"time"
)

type TimeService struct {
	UnimplementedTimeServiceServer

	LastTimestamp time.Time
	LastDuration  time.Duration
}

func (s *TimeService) TestCall(ctx context.Context, req *CallRequest) (*CallReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
