package postgres

import (
	"context"
	"portfolio/api/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateblog(t *testing.T) {
	ts := NewStorageTest(t)
	arg := storage.Blog{
		Name:      "Mainul",
		CreatedBy: "12345",
	}
	blog, err := ts.CreateBlog(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, blog)
	require.NotZero(t, blog)

}
