package repository

import (
	"azuk774/misskey-dev-tools/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type MClient struct {
	Domain string // ex. 'misskey.io'
	UserID string
	I      string // secret
}

func NewMClient() *MClient {
	return &MClient{
		Domain: os.Getenv("DOMAIN"),
		UserID: os.Getenv("USERID"),
		I:      os.Getenv("I"),
	}
}

func (c *MClient) GetRecentReactions(ctx context.Context, num int) (nrs []model.NoteReaction, err error) {
	endpoint := "https://" + c.Domain + "/api/users/reactions"
	reqData := model.ReactionsReq{
		UserID: c.UserID,
		I:      c.I,
		Limit:  num,
	}

	reqJson, err := json.Marshal(reqData)
	if err != nil {
		return []model.NoteReaction{}, err
	}

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return []model.NoteReaction{}, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return []model.NoteReaction{}, err
	}

	err = json.Unmarshal(b, &nrs)
	if err != nil {
		return []model.NoteReaction{}, err
	}

	return nrs, nil
}

func (c *MClient) PostNote(ctx context.Context, text string) (err error) {
	// TODO
	return nil
}
