package service

import (
	"azuk774/misskey-dev-tools/internal/model"
	"reflect"
	"testing"
)

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
