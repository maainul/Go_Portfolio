package getInTouch

import (
	"context"
	gk "portfolio/api/gunk/v1/admin/getInTouch"
	"portfolio/api/storage"
)

type GetInTouchCoreFuncs struct {
	bc GetInTouchCore
	gk.UnimplementedGetInTouchServiceServer
}

type GetInTouchCore interface {
	CreateGetInTouch(context.Context, storage.GetInTouch) (string, error)
}

func GetInTouchCoreConn(bc GetInTouchCore) *GetInTouchCoreFuncs {
	return &GetInTouchCoreFuncs{bc: bc}
}
