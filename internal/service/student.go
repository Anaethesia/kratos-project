package service

import (
	"context"
	"fmt"
	"helloworld/internal/biz"

	pb "helloworld/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	uc *biz.StudentUseCase
}

func NewStudentService(uc *biz.StudentUseCase) *StudentService {
	return &StudentService{uc: uc}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResp, error) {
	msg := fmt.Sprintf("name: %s, age: %d", req.Name, req.Age)
	return &pb.StudentResp{Messsage: msg}, nil
}
