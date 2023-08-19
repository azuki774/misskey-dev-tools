package repository

import (
	"azuk774/misskey-dev-tools/internal/model"
	"context"
)

type MClient struct{}

func NewMClient() *MClient {
	return &MClient{}
}

func (c *MClient) GetRecentReactions(ctx context.Context, num int) (nrs []model.NoteReaction, err error) {
	// TODO
	return []model.NoteReaction{}, nil
}

func (c *MClient) PostNote(ctx context.Context, text string) (err error) {
	// TODO
	return nil
}
