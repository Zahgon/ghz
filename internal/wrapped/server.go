package wrapped

import (
	"context"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

type WrappedService struct{}

func (s *WrappedService) GetMessage(ctx context.Context, req *wrappers.StringValue) (*wrappers.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *WrappedService) GetBytesMessage(ctx context.Context, req *wrappers.BytesValue) (*wrappers.BytesValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *WrappedService) mustEmbedUnimplementedWrappedServiceServer() {
	_ = "STUB: not implemented"
	return
}
