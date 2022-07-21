package getInTouch


import "portfolio/api/storage/postgres"

type GetInTouchStrgFuncs struct {
	st *postgres.Storage
}

func ConnWithStorage(ps *postgres.Storage) *GetInTouchStrgFuncs {
	return &GetInTouchStrgFuncs{
		st: ps,
	}
}
