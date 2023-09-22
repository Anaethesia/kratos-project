package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentrRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) Save(ctx context.Context, g *biz.Student) (*biz.Student, error) {
return g, nil
}