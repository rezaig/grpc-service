package grpcsvc

import (
	"context"

	pb "github.com/rezaig/grpc-service/story-service/pb/story"
)

type Service struct {
	pb.UnimplementedStoryServiceServer
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) FindByID(ctx context.Context, in *pb.FindByIDRequest) (out *pb.Story, err error) {
	return &pb.Story{
		Id:    1,
		Title: "Example title",
	}, nil
}

func (s *Service) FindAll(ctx context.Context, in *pb.FindAllRequest) (out *pb.Stories, err error) {
	return &pb.Stories{Stories: []*pb.Story{
		{
			Id:    1,
			Title: "Example title",
		},
		{
			Id:    2,
			Title: "Example title",
		},
	}}, nil
}
