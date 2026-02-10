package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

// RoundResponse maps a database.Round to a dto.RoundResponse.
func RoundResponse(dbRound database.Round) dto.RoundResponse {
	return dto.RoundResponse{
		ID:              dbRound.ID,
		Name:            fromNullString(dbRound.Name),
		Format:          dbRound.Format,
		Levels:          fromNullString(dbRound.Levels),
		QualifiersCount: dbRound.QualifiersCount,
		State:           int32(RoundStateToDTO(dbRound.CurrentState)),
		CategoryID:      dbRound.CategoryID,
		OrderIndex:      dbRound.OrderIndex,
	}
}

func CreateRoundParams(req dto.CreateRoundRequest, categoryID uuid.UUID) database.CreateRoundParams {
	return database.CreateRoundParams{
		Name:            toNullString(req.Name),
		Format:          req.Format,
		Levels:          toNullString(req.Levels),
		QualifiersCount: req.QualifiersCount,
		CurrentState:    RoundStateFromDTO(dto.RoundState(req.State)),
		CategoryID:      categoryID,
	}
}

func UpdateRoundParams(id uuid.UUID, req dto.UpdateRoundRequest) database.UpdateRoundParams {
	return database.UpdateRoundParams{
		ID:              id,
		Name:            toNullString(req.Name),
		Format:          toNullInt32(req.Format),
		Levels:          toNullString(req.Levels),
		QualifiersCount: toNullInt32(req.QualifiersCount),
		CurrentState:    toNullRoundState(req.State),
	}
}

func AddPlayersToRoundBulkParams(roundID uuid.UUID, req dto.AddPlayersToRoundBulkRequest) database.AddPlayersToRoundBulkParams {
	return database.AddPlayersToRoundBulkParams{
		RoundID:   roundID,
		PlayerIds: req.PlayerIDs,
	}
}

func RemovePlayerFromRoundParams(roundID, playerID uuid.UUID) database.RemovePlayerFromRoundParams {
	return database.RemovePlayerFromRoundParams{
		RoundID:  roundID,
		PlayerID: playerID,
	}
}

func FixMissingIndexFromRoundPlayerOrderParams(roundID uuid.UUID, removedOrderIndex int32) database.FixMissingIndexFromRoundPlayerOrderParams {
	return database.FixMissingIndexFromRoundPlayerOrderParams{
		RoundID:    roundID,
		OrderIndex: removedOrderIndex,
	}
}

func UpdateRoundPlayersOrderBulkParams(roundID uuid.UUID, req dto.UpdateRoundPlayersOrderBulkRequest) database.UpdateRoundPlayersOrderBulkParams {
	return database.UpdateRoundPlayersOrderBulkParams{
		RoundID:   roundID,
		PlayerIds: req.PlayerIDs,
	}
}

func AddChartToRoundParams(roundID, chartID uuid.UUID) database.AddChartToRoundParams {
	return database.AddChartToRoundParams{
		RoundID: roundID,
		ChartID: chartID,
	}
}

func RemoveChartFromRoundParams(roundID, chartID uuid.UUID) database.RemoveChartFromRoundParams {
	return database.RemoveChartFromRoundParams{
		RoundID: roundID,
		ChartID: chartID,
	}
}

func FixMissingIndexFromRoundChartOrderParams(roundID uuid.UUID, removedOrderIndex int32) database.FixMissingIndexFromRoundChartOrderParams {
	return database.FixMissingIndexFromRoundChartOrderParams{
		RoundID:    roundID,
		OrderIndex: removedOrderIndex,
	}
}

func ReplaceRoundChartParams(roundID uuid.UUID, orderIndex int32, chartID uuid.UUID) database.ReplaceRoundChartParams {
	return database.ReplaceRoundChartParams{
		RoundID:    roundID,
		OrderIndex: orderIndex,
		ChartID:    chartID,
	}
}

func RoundStateFromDTO(state dto.RoundState) database.RoundState {
	switch state {
	case dto.RoundStateNotStarted:
		return database.RoundStateNotStarted
	case dto.RoundStateInProgress:
		return database.RoundStateInProgress
	case dto.RoundStatePaused:
		return database.RoundStatePaused
	case dto.RoundStateFinished:
		return database.RoundStateFinished
	default:
		return database.RoundStateNotStarted // Default to not started if unknown value
	}
}

func RoundStateToDTO(state database.RoundState) dto.RoundState {
	switch state {
	case database.RoundStateNotStarted:
		return dto.RoundStateNotStarted
	case database.RoundStateInProgress:
		return dto.RoundStateInProgress
	case database.RoundStatePaused:
		return dto.RoundStatePaused
	case database.RoundStateFinished:
		return dto.RoundStateFinished
	default:
		return dto.RoundStateNotStarted // Default to not started if unknown value
	}
}
