package posts

import (
	"context"
	"grpc_users_v1/internal/grpc/pb"
)

const (
	getPostsQuery = `select id, title, message from posts where userId=$1`
)

func (p *PostServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	var posts []*pb.Post
	rows, err := p.db.Query(getPostsQuery, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post pb.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Message); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	res := &pb.GetPostsResponse{Posts: posts}
	return res, nil
}
