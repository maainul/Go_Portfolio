package getInTouch

import (
	"context"
	gk "portfolio/api/gunk/v1/admin/getInTouch"
	"portfolio/api/storage"
	"portfolio/svcUtils/logging"
	"time"
)

func (s *GetInTouchCoreFuncs) CreateGetInTouch(ctx context.Context, req *gk.CreateGetInTouchRequest) (*gk.CreateGetInTouchResponse, error) {
	log := logging.FromContext(ctx).WithField("method", "service.GetInTouch.CreateGetInTouch")
	id, err := s.bc.CreateGetInTouch(ctx, storage.GetInTouch{
		Name:      req.GetInTouch.Name,
		Email:     req.GetInTouch.Email,
		Message:   req.GetInTouch.Message,
		CreatedAt: time.Now(),
		CreatedBy: req.GetInTouch.CreatedBy,
	})
	if err != nil {
		logging.WithError(err, log).Error("Failed to create GetInTouch")
		return nil, err
	}
	return &gk.CreateGetInTouchResponse{
		ID: id,
	}, nil
}
