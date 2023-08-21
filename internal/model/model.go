package model

import (
	"fmt"
	"strings"
)

// NoteReactionSlice: count, reaction_name の構造体
type NoteReactionSlice struct {
	ReactionName string
	Count        int
}

func NoteCountGetText(nrsc []NoteReactionSlice) string {
	t := ""
	for _, n := range nrsc {
		// "@." を取る :aruaru:@. -> :aruaru:
		n.ReactionName = strings.Replace(n.ReactionName, "@.:", ":", 1)
		t += fmt.Sprintf("%s x%d\n", n.ReactionName, n.Count)
	}
	return t
}
