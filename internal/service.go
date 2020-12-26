package internal

import (
	"context"

	"github.com/boris-wtf/secret-server/internal/notes"

	pb "github.com/boris-wtf/secret-server/pkg/github.com/boris-wtf/apis/secret"
)

type SecretServiceServer struct {
	pb.UnimplementedSecretServiceServer
	Notes notes.Registry
}

func (s *SecretServiceServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteReply, error) {
	id, err := s.Notes.Create(req.Text)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteReply{ID: id}, nil
}

func (s *SecretServiceServer) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteReply, error) {
	text, err := s.Notes.Delete(req.ID)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteNoteReply{Text: text}, nil
}
