package mapper

import (
	"github.com/alancorleto/piu-tournament-manager/internal/database"
	"github.com/alancorleto/piu-tournament-manager/internal/http/dto"
	"github.com/google/uuid"
)

func ChartResponse(dbChart database.Chart) dto.ChartResponse {
	return dto.ChartResponse{
		ID:          dbChart.ID,
		SongID:      dbChart.SongID,
		Mode:        dbChart.Mode,
		Level:       dbChart.Level,
		PlayerCount: dbChart.PlayerCount,
	}
}

func CreateChartParams(req dto.CreateChartRequest) database.CreateChartParams {
	return database.CreateChartParams{
		SongID:      req.SongID,
		Mode:        req.Mode,
		Level:       req.Level,
		PlayerCount: req.PlayerCount,
	}
}

func UpdateChartParams(id uuid.UUID, req dto.UpdateChartRequest) database.UpdateChartParams {
	return database.UpdateChartParams{
		ID:          id,
		SongID:      toNullUUID(req.SongID),
		Mode:        toNullInt32(req.Mode),
		Level:       toNullInt32(req.Level),
		PlayerCount: toNullInt32(req.PlayerCount),
	}
}
