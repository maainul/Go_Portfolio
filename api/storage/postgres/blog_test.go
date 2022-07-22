package postgres

import (
	"context"
	"portfolio/api/storage"
	"testing"
)

func TestStorage_CreateBlog(t *testing.T) {
	ts := NewStorageTest(t)

	type args struct {
		ctx context.Context
		blg storage.Blog
	}
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.CreateBlog(context.TODO(), storage.Blog{
				Name:      "Test Blog",
				CreatedBy: "123",
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.CreateBlog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.CreateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}
