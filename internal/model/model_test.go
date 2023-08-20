package model

import "testing"

func TestNoteCountGetText(t *testing.T) {
	type args struct {
		nrsc []NoteReactionSlice
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				nrsc: []NoteReactionSlice{
					{
						ReactionName: ":test_reactionC@.:",
						Count:        3,
					},
					{
						ReactionName: ":test_reactionB@.:",
						Count:        2,
					},
				},
			},
			want: ":test_reactionC: x3\n:test_reactionB: x2\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteCountGetText(tt.args.nrsc); got != tt.want {
				t.Errorf("NoteCountGetText() = %v, want %v", got, tt.want)
			}
		})
	}
}
