package posts

import (
	"context"
	"errors"
	"grpc_users_v1/internal/grpc/pb"
)

const (
	createPostQuery = `insert into posts (title, message, userId) values ($1, $2, $3) returning id`
)

func (p *PostServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	var post pb.CreatePostResponse

	if errTitle := EmptyTitle(req.Title); errTitle != true {
		return nil, errors.New("title can not be empty")
	}

	if errMessage := EmptyMessage(req.Message); errMessage != true {
		return nil, errors.New("message can not be empty")
	}

	rows, err := p.db.Query(createPostQuery, req.Title, req.Message, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&post.Id); err != nil {
			return nil, err
		}
	}
	createPost := &pb.CreatePostResponse{Id: post.GetId(), Title: req.Title, Message: req.Message, UserId: post.GetUserId()}
	return createPost, nil
}
