package service

import (
	"azuk774/misskey-dev-tools/internal/model"
	"context"
	"log/slog"
	"sort"
)

type ISendReactionRepository interface {
	GetRecentReactions(ctx context.Context, num int) (nrs []model.NoteReaction, err error)
	PostNote(ctx context.Context, text string) (err error)
}

type sendReactionCountService struct {
	Repo            ISendReactionRepository
	FetchNoteNum    int // Fetch note number
	ReactionkindNum int // the number of sending category
}

func NewSendReactionCountService(Repo ISendReactionRepository) *sendReactionCountService {
	const fetchNum = 100
	const reactionkindNum = 5
	return &sendReactionCountService{FetchNoteNum: fetchNum, ReactionkindNum: reactionkindNum, Repo: Repo}
}

// countMyReactions: リアクション名 -> 個数 の mapping を作成
func countMyReactions(nrs []model.NoteReaction) (countf map[string]int, err error) {
	countf = make(map[string]int)
	for _, nr := range nrs {
		countf[nr.Note.MyReaction]++
	}

	return countf, nil
}

// pickReactionsFromCountf は countf をもとに実際に note するスライスを抽出・並び替えする
func pickReactionsFromCountf(countf map[string]int, pickNum int) (nrsc []model.NoteReactionSlice, err error) {
	for reactionName, num := range countf {
		nrsc = append(nrsc, model.NoteReactionSlice{ReactionName: reactionName, Count: num})
	}

	// Count が大きい順にソート
	sort.Slice(nrsc, func(i, j int) bool {
		return nrsc[i].Count >= nrsc[j].Count
	})

	if pickNum > len(nrsc) {
		pickNum = len(nrsc)
	}
	return nrsc[0:pickNum], nil
}

func (s *sendReactionCountService) Run(ctx context.Context) (err error) {
	slog.Info("sendReactionCount start")
	nrs, err := s.Repo.GetRecentReactions(ctx, s.FetchNoteNum)
	if err != nil {
		slog.Error("failed to get recent reactions", err)
		return err
	}
	slog.Info("fetch recent my reactions", "fetch_reactions_num", s.FetchNoteNum)

	countf, err := countMyReactions(nrs)
	if err != nil {
		slog.Error("failed to count my reactions", err)
		return err
	}
	slog.Info("count recent my reactions")

	_, err = pickReactionsFromCountf(countf, s.ReactionkindNum)
	if err != nil {
		slog.Error("failed to pick my reactions", err)
		return err
	}
	slog.Info("pick my reactions", "pick_kind_category", s.ReactionkindNum)

	return nil
}
