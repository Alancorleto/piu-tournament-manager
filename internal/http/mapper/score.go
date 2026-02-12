package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// ScoreResponse maps a database.Score to a dto.ScoreResponse.
func ScoreResponse(dbScore database.Score) dto.ScoreResponse {
	return dto.ScoreResponse{
		ID:        dbScore.ID,
		PlayerID:  dbScore.PlayerID,
		ChartID:   dbScore.ChartID,
		RoundID:   dbScore.RoundID,
		Score:     dbScore.Score,
		Perfect:   fromNullInt32(dbScore.Perfect),
		Great:     fromNullInt32(dbScore.Great),
		Good:      fromNullInt32(dbScore.Good),
		Bad:       fromNullInt32(dbScore.Bad),
		Miss:      fromNullInt32(dbScore.Miss),
		MaxCombo:  fromNullInt32(dbScore.MaxCombo),
		Kcal:      fromNullFloat32(dbScore.Kcal),
		Grade:     NullGradeToDTO(dbScore.Grade),
		StagePass: fromNullBool(dbScore.StagePass),
		VideoURL:  fromNullString(dbScore.VideoUrl),
	}
}

// CreateScoreParams maps a dto.CreateScoreRequest to database.CreateScoreParams.
func CreateScoreParams(req dto.CreateScoreRequest, roundID uuid.UUID) database.CreateScoreParams {
	return database.CreateScoreParams{
		PlayerID:  req.PlayerID,
		ChartID:   req.ChartID,
		RoundID:   roundID,
		Score:     req.Score,
		Perfect:   toNullInt32(req.Perfect),
		Great:     toNullInt32(req.Great),
		Good:      toNullInt32(req.Good),
		Bad:       toNullInt32(req.Bad),
		Miss:      toNullInt32(req.Miss),
		MaxCombo:  toNullInt32(req.MaxCombo),
		Kcal:      toNullFloat32(req.Kcal),
		Grade:     NullGradeFromDTO(req.Grade),
		StagePass: toNullBool(req.StagePass),
		VideoUrl:  toNullString(req.VideoURL),
	}
}

// UpdateScoreParams maps a dto.UpdateScoreRequest to database.UpdateScoreParams.
func UpdateScoreParams(id uuid.UUID, req dto.UpdateScoreRequest) database.UpdateScoreParams {
	return database.UpdateScoreParams{
		ID:        id,
		Score:     toNullInt32(req.Score),
		Perfect:   toNullInt32(req.Perfect),
		Great:     toNullInt32(req.Great),
		Good:      toNullInt32(req.Good),
		Bad:       toNullInt32(req.Bad),
		Miss:      toNullInt32(req.Miss),
		MaxCombo:  toNullInt32(req.MaxCombo),
		Kcal:      toNullFloat32(req.Kcal),
		Grade:     NullGradeFromDTO(req.Grade),
		StagePass: toNullBool(req.StagePass),
		VideoUrl:  toNullString(req.VideoURL),
	}
}

func NullGradeToDTO(grade database.NullGrade) *dto.Grade {
	if !grade.Valid {
		return nil
	}
	dtoGrade := GradeToDTO(grade.Grade)
	return &dtoGrade
}

func NullGradeFromDTO(grade *dto.Grade) database.NullGrade {
	if grade == nil {
		return database.NullGrade{Valid: false}
	}
	return database.NullGrade{Grade: GradeFromDTO(*grade), Valid: true}
}

func GradeToDTO(grade database.Grade) dto.Grade {
	switch grade {
	case database.GradeF:
		return dto.F
	case database.GradeD:
		return dto.D
	case database.GradeC:
		return dto.C
	case database.GradeB:
		return dto.B
	case database.GradeA:
		return dto.A
	case database.GradeAP:
		return dto.A_P
	case database.GradeAA:
		return dto.AA
	case database.GradeAAP:
		return dto.AA_P
	case database.GradeAAA:
		return dto.AAA
	case database.GradeAAAP:
		return dto.AAA_P
	case database.GradeS:
		return dto.S
	case database.GradeSP:
		return dto.S_P
	case database.GradeSS:
		return dto.SS
	case database.GradeSSP:
		return dto.SS_P
	case database.GradeSSS:
		return dto.SSS
	case database.GradeSSSP:
		return dto.SSS_P
	default:
		panic("unexpected grade value")
	}
}

func GradeFromDTO(grade dto.Grade) database.Grade {
	switch grade {
	case dto.F:
		return database.GradeF
	case dto.D:
		return database.GradeD
	case dto.C:
		return database.GradeC
	case dto.B:
		return database.GradeB
	case dto.A:
		return database.GradeA
	case dto.A_P:
		return database.GradeAP
	case dto.AA:
		return database.GradeAA
	case dto.AA_P:
		return database.GradeAAP
	case dto.AAA:
		return database.GradeAAA
	case dto.AAA_P:
		return database.GradeAAAP
	case dto.S:
		return database.GradeS
	case dto.S_P:
		return database.GradeSP
	case dto.SS:
		return database.GradeSS
	case dto.SS_P:
		return database.GradeSSP
	case dto.SSS:
		return database.GradeSSS
	case dto.SSS_P:
		return database.GradeSSSP
	default:
		panic("unexpected grade value")
	}
}
