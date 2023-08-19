package service

import (
	"azuk774/misskey-dev-tools/internal/model"
	"context"
	"log/slog"
)

type ISendReactionRepository interface {
	GetRecentReactions(ctx context.Context, num int) (nrs []model.NoteReaction, err error)
	PostNote(ctx context.Context, text string) (err error)
}

type sendReactionCountService struct {
	Repo ISendReactionRepository
	Num  int // Fetch note number
}

// countMyReactions: リアクション名 -> 個数 の mapping を作成
func countMyReactions(nrs []model.NoteReaction) (countf map[string]int, err error) {
	countf = make(map[string]int)
	for _, nr := range nrs {
		countf[nr.Note.MyReaction]++
	}

	return countf, nil
}

func (s *sendReactionCountService) Run(ctx context.Context) (err error) {
	slog.Info("sendReactionCount start")
	_, err = s.Repo.GetRecentReactions(ctx, s.Num)
	if err != nil {
		slog.Error("failed to get recent reactions", err)
		return err
	}
	slog.Info("fetch recent my reactions")
	return nil
}
