package blog

import (
	"context"
	gk "portfolio/api/gunk/v1/admin/blog"
	"portfolio/api/storage"
)

type BlogCoreFuncs struct {
	bc BlogCore
	gk.UnimplementedBlogServiceServer
}

type BlogCore interface {
	CreateBlog(context.Context, storage.Blog) (string, error)
}

func BlogCoreConn(bc BlogCore) *BlogCoreFuncs {
	return &BlogCoreFuncs{bc: bc}
}
