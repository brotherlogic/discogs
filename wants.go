package discogs

import (
	"context"

	pb "github.com/brotherlogic/discogs/proto"
)

func (p *prodClient) GetWants(ctx context.Context, page int32) ([]*pb.Want, *pb.Pagination, error) {
	return []*pb.Want{}, &pb.Pagination{}, nil
}
