package dto

import (
	"github.com/google/uuid"
)

type CreateScoreRequest struct {
	PlayerID  uuid.UUID `json:"player_id"`
	ChartID   uuid.UUID `json:"chart_id"`
	Score     int32     `json:"score"`
	Perfect   *int32    `json:"perfect,omitempty"`
	Great     *int32    `json:"great,omitempty"`
	Good      *int32    `json:"good,omitempty"`
	Bad       *int32    `json:"bad,omitempty"`
	Miss      *int32    `json:"miss,omitempty"`
	MaxCombo  *int32    `json:"max_combo,omitempty"`
	Kcal      *float32  `json:"kcal,omitempty"`
	Grade     *Grade    `json:"grade,omitempty"`
	StagePass *bool     `json:"stage_pass,omitempty"`
	VideoURL  *string   `json:"video_url,omitempty"`
}

type ScoreResponse struct {
	ID        uuid.UUID `json:"id"`
	PlayerID  uuid.UUID `json:"player_id"`
	ChartID   uuid.UUID `json:"chart_id"`
	RoundID   uuid.UUID `json:"round_id"`
	Score     int32     `json:"score"`
	Perfect   *int32    `json:"perfect,omitempty"`
	Great     *int32    `json:"great,omitempty"`
	Good      *int32    `json:"good,omitempty"`
	Bad       *int32    `json:"bad,omitempty"`
	Miss      *int32    `json:"miss,omitempty"`
	MaxCombo  *int32    `json:"max_combo,omitempty"`
	Kcal      *float32  `json:"kcal,omitempty"`
	Grade     *Grade    `json:"grade,omitempty"`
	StagePass *bool     `json:"stage_pass,omitempty"`
	VideoURL  *string   `json:"video_url,omitempty"`
}

type UpdateScoreRequest struct {
	Score     *int32   `json:"score,omitempty"`
	Perfect   *int32   `json:"perfect,omitempty"`
	Great     *int32   `json:"great,omitempty"`
	Good      *int32   `json:"good,omitempty"`
	Bad       *int32   `json:"bad,omitempty"`
	Miss      *int32   `json:"miss,omitempty"`
	MaxCombo  *int32   `json:"max_combo,omitempty"`
	Kcal      *float32 `json:"kcal,omitempty"`
	Grade     *Grade   `json:"grade,omitempty"`
	StagePass *bool    `json:"stage_pass,omitempty"`
	VideoURL  *string  `json:"video_url,omitempty"`
}

type Grade int32

const (
	F Grade = iota
	D
	C
	B
	A
	A_P
	AA
	AA_P
	AAA
	AAA_P
	S
	S_P
	SS
	SS_P
	SSS
	SSS_P
)
