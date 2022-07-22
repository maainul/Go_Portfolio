package postgres

import (
	"context"
	"portfolio/api/storage"
	"portfolio/svcUtils/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const crtgit = `
INSERT INTO get_in_touch (
	name,
	email,
	message,
	created_at,
	created_by
) VALUES (
	:name,
	:email,
	:message,
	:created_at,
	:created_by
) RETURNING
	id
`

func (s *Storage) CreateGetInTouch(ctx context.Context, git storage.GetInTouch) (string, error) {
	log := logging.FromContext(ctx)
	stmt, err := s.db.PrepareNamed(crtgit)
	if err != nil {
		logging.WithError(err, log).Error("Failed to create GetInTouch")
		return "", err
	}
	defer stmt.Close()
	var id string
	if err := stmt.Get(&id, git); err != nil {
		logging.WithError(err, log).Error("Failed to insert GetInTouch")
		return "", status.Errorf(codes.Internal, "%d", err)

	}
	return id, nil

}
