package getInTouch

import (
	"context"
	"portfolio/api/storage"
	"portfolio/svcUtils/logging"
)

func (stgGetInTouch *GetInTouchStrgFuncs) CreateGetInTouch(ctx context.Context, GetInTouch storage.GetInTouch) (string, error) {
	log := logging.FromContext(ctx).WithField("method", "CreateGetInTouch")
	id, err := stgGetInTouch.st.CreateGetInTouch(ctx, GetInTouch)
	if err != nil {
		logging.WithError(err, log).Error("Failed to create get in touch")
	}
	return id, nil
}
