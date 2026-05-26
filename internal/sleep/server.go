package sleep

import (
	"context"
)

type SleepService struct{}

func (s *SleepService) SleepFor(ctx context.Context, req *SleepRequest) (*SleepResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
