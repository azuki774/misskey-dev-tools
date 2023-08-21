package service

import (
	"azuk774/misskey-dev-tools/internal/model"
	"context"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
func Test_countMyReactions(t *testing.T) {
	// create test
	t1 := model.NoteReaction{}
	t1.Note.MyReaction = ":test_reactionA@.:"

	t2 := model.NoteReaction{}
	t2.Note.MyReaction = ":test_reactionB@.:"

	type args struct {
		nrs []model.NoteReaction
	}
	tests := []struct {
		name       string
		args       args
		wantCountf map[string]int
		wantErr    bool
	}{
		{
			name: "normal",
			args: args{
				nrs: []model.NoteReaction{
					t1,
					t2,
					t2,
				},
			},
			wantCountf: map[string]int{
				":test_reactionA@.:": 1,
				":test_reactionB@.:": 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCountf, err := countMyReactions(tt.args.nrs)
			if (err != nil) != tt.wantErr {
				t.Errorf("countMyReactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountf, tt.wantCountf) {
				t.Errorf("countMyReactions() = %v, want %v", gotCountf, tt.wantCountf)
			}
		})
	}
}

func Test_pickReactionsFromCountf(t *testing.T) {
	type args struct {
		countf  map[string]int
		pickNum int
	}
	tests := []struct {
		name     string
		args     args
		wantNrsc []model.NoteReactionSlice
		wantErr  bool
	}{
		{
			name: "ok #1",
			args: args{
				countf:  map[string]int{":test_reactionA@.:": 1, ":test_reactionB@.:": 2, ":test_reactionC@.:": 3},
				pickNum: 2,
			},
			wantNrsc: []model.NoteReactionSlice{
				{
					ReactionName: ":test_reactionC@.:",
					Count:        3,
				},
				{
					ReactionName: ":test_reactionB@.:",
					Count:        2,
				},
			},
			wantErr: false,
		},
		{
			name: "ok #2",
			args: args{
				countf:  map[string]int{":test_reactionA@.:": 1, ":test_reactionB@.:": 2, ":test_reactionC@.:": 3},
				pickNum: 1,
			},
			wantNrsc: []model.NoteReactionSlice{
				{
					ReactionName: ":test_reactionC@.:",
					Count:        3,
				},
			},
			wantErr: false,
		},
		{
			name: "ok #3 (not sufficient)",
			args: args{
				countf:  map[string]int{":test_reactionA@.:": 1, ":test_reactionB@.:": 2, ":test_reactionC@.:": 3},
				pickNum: 10,
			},
			wantNrsc: []model.NoteReactionSlice{
				{
					ReactionName: ":test_reactionC@.:",
					Count:        3,
				},
				{
					ReactionName: ":test_reactionB@.:",
					Count:        2,
				},
				{
					ReactionName: ":test_reactionA@.:",
					Count:        1,
				},
			},
			wantErr: false,
		},
		{
			name: "none",
			args: args{
				countf:  map[string]int{},
				pickNum: 1,
			},
			wantNrsc: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNrsc, err := pickReactionsFromCountf(tt.args.countf, tt.args.pickNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("pickReactionsFromCountf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNrsc, tt.wantNrsc) {
				t.Errorf("pickReactionsFromCountf() = %v, want %v", gotNrsc, tt.wantNrsc)
			}
		})
	}
}

func Test_sendReactionCountService_Run(t *testing.T) {
	type fields struct {
		Repo            ISendReactionRepository
		FetchNoteNum    int
		ReactionkindNum int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				Repo:            &mockSendReactionRepository{},
				FetchNoteNum:    100,
				ReactionkindNum: 5,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "fetch error",
			fields: fields{
				Repo:            &mockSendReactionRepository{ErrGetRecentReactions: fmt.Errorf("error text")},
				FetchNoteNum:    100,
				ReactionkindNum: 5,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sendReactionCountService{
				Repo:            tt.fields.Repo,
				FetchNoteNum:    tt.fields.FetchNoteNum,
				ReactionkindNum: tt.fields.ReactionkindNum,
			}
			if err := s.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("sendReactionCountService.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
