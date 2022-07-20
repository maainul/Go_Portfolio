package blog

import "portfolio/api/storage/postgres"

type BlgStrgFuncs struct {
	st *postgres.Storage
}

func ConnWithStorage(ps *postgres.Storage) *BlgStrgFuncs {
	return &BlgStrgFuncs{
		st: ps,
	}
}
