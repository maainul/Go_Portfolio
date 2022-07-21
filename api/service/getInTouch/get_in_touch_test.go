package getInTouch

import (
	"context"
	"log"
	"os"
	"path/filepath"
	bcr "portfolio/api/core/getInTouch"
	gk "portfolio/api/gunk/v1/admin/getInTouch"
	"portfolio/api/storage/postgres"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetInTouchCoreConn(t *testing.T) {
	dbconn, cleanup := postgres.NewTestStorage(os.Getenv("DATABASE_CONNECTION"), filepath.Join("..", "..", "migrations", "sql"))
	t.Cleanup(cleanup)
	svc := GetInTouchCoreConn(bcr.ConnWithStorage(dbconn))
	tests := []struct {
		name    string
		in      gk.GetInTouch
		want    *gk.GetInTouch
		WantErr bool
	}{
		{
			name: "CREATE_GetInTouch_SUCCESS",
			in: gk.GetInTouch{
				Name:     "New GetInTouch Title",
				CreatedAt: timestamppb.Now(),
				CreatedBy: "12345",
			},
			want: &gk.GetInTouch{
				Name:     "New GetInTouch Title",
				CreatedAt: timestamppb.Now(),
				CreatedBy: "12345",
			},
			WantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			id, err := svc.CreateGetInTouch(context.TODO(), &gk.CreateGetInTouchRequest{
				GetInTouch: &tt.in,
			})
			if err != nil {
				t.Errorf("Storage.CreateGetInTouch() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if id.ID == "" {
				log.Fatalf("Create GetInTouch Failed")
			}

		})
	}
}
