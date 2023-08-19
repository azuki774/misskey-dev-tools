package model

import "time"

type NoteReaction struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	User      struct {
		ID             string `json:"id,omitempty"`
		Name           string `json:"name,omitempty"`
		Username       string `json:"username,omitempty"`
		Host           any    `json:"host,omitempty"`
		AvatarURL      string `json:"avatarUrl,omitempty"`
		AvatarBlurhash string `json:"avatarBlurhash,omitempty"`
		IsBot          bool   `json:"isBot,omitempty"`
		IsCat          bool   `json:"isCat,omitempty"`
		Emojis         struct {
		} `json:"emojis,omitempty"`
		OnlineStatus string `json:"onlineStatus,omitempty"`
		BadgeRoles   []struct {
			Name         string `json:"name,omitempty"`
			IconURL      string `json:"iconUrl,omitempty"`
			DisplayOrder int    `json:"displayOrder,omitempty"`
		} `json:"badgeRoles,omitempty"`
	} `json:"user,omitempty"`
	Type string `json:"type,omitempty"`
	Note struct {
		ID        string    `json:"id,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`
		UserID    string    `json:"userId,omitempty"`
		User      struct {
			ID             string `json:"id,omitempty"`
			Name           string `json:"name,omitempty"`
			Username       string `json:"username,omitempty"`
			Host           any    `json:"host,omitempty"`
			AvatarURL      string `json:"avatarUrl,omitempty"`
			AvatarBlurhash string `json:"avatarBlurhash,omitempty"`
			IsBot          bool   `json:"isBot,omitempty"`
			IsCat          bool   `json:"isCat,omitempty"`
			Emojis         struct {
			} `json:"emojis,omitempty"`
			OnlineStatus string `json:"onlineStatus,omitempty"`
			BadgeRoles   []struct {
				Name         string `json:"name,omitempty"`
				IconURL      string `json:"iconUrl,omitempty"`
				DisplayOrder int    `json:"displayOrder,omitempty"`
			} `json:"badgeRoles,omitempty"`
		} `json:"user,omitempty"`
		Text               string `json:"text,omitempty"`
		Cw                 any    `json:"cw,omitempty"`
		Visibility         string `json:"visibility,omitempty"`
		LocalOnly          bool   `json:"localOnly,omitempty"`
		ReactionAcceptance string `json:"reactionAcceptance,omitempty"`
		RenoteCount        int    `json:"renoteCount,omitempty"`
		RepliesCount       int    `json:"repliesCount,omitempty"`
		Reactions          struct {
		} `json:"reactions,omitempty"`
		ReactionEmojis struct {
		} `json:"reactionEmojis,omitempty"`
		FileIds    []any  `json:"fileIds,omitempty"`
		Files      []any  `json:"files,omitempty"`
		ReplyID    any    `json:"replyId,omitempty"`
		RenoteID   any    `json:"renoteId,omitempty"`
		MyReaction string `json:"myReaction,omitempty"`
	} `json:"note,omitempty"`
}
