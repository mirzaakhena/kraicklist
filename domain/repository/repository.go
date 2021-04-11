package repository

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/domain/entity"
)

type SaveDataRepo interface {
	SaveData(ctx context.Context, obj *entity.RawData) error
}

type FindDataByKeywordRepo interface {
	FindDataByKeyword(ctx context.Context, someID string) (*FindDataByKeywordRepoResult, error)
}

type FindDataByKeywordRepoResult struct {
	Result           []*entity.RawData
	ContainedKeyword []string
}
