package service

import (
	"azuk774/misskey-dev-tools/internal/model"
	"context"
)

type mockSendReactionRepository struct {
	ErrGetRecentReactions error
	ErrPostNote           error
}

func (m *mockSendReactionRepository) GetRecentReactions(ctx context.Context, num int) (nrs []model.NoteReaction, err error) {
	t1 := model.NoteReaction{}
	t1.Note.MyReaction = ":test_reactionA@.:"

	t2 := model.NoteReaction{}
	t2.Note.MyReaction = ":test_reactionB@.:"

	t3 := model.NoteReaction{}
	t3.Note.MyReaction = ":test_reactionC@.:"
	return []model.NoteReaction{}, m.ErrGetRecentReactions
}

func (m *mockSendReactionRepository) PostNote(ctx context.Context, text string) (err error) {
	return m.ErrPostNote
}
