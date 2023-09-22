package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	name string
	age int64
}

type StudentRepo interface {
	Save(ctx context.Context, student *Student) (*Student, error)
}

type StudentUseCase struct {
	repo StudentRepo
	log  *log.Helper
}

func NewStudentUseCase(repo StudentRepo, logger log.Logger) *StudentUseCase {
	return &StudentUseCase{repo: repo, log:log.NewHelper(logger)}
}


