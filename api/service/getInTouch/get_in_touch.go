package getInTouch

import (
	"context"
	"portfolio/api/storage"
	gk "portfolio/api/gunk/v1/admin/getInTouch"
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
