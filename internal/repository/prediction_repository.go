package repository

import (
	"context"
	"database/sql"
	"fmt"

	"mma-prediction-tracker/internal/models"
)

type PredictionRepository struct {
	db *sql.DB
}

func NewPredictionRepository(db *sql.DB) *PredictionRepository {
	return &PredictionRepository{db: db}
}

func (r *PredictionRepository) Create(ctx context.Context, prediction *models.Prediction) error {
	query := `
		INSERT INTO predictions (predictor_id, fight_id, predicted_winner_id, probability) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		prediction.PredictorID,
		prediction.FightID,
		prediction.PredictedWinnerID,
		prediction.Probability,
	).Scan(&prediction.ID, &prediction.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create prediction: %w", err)
	}

	return nil
}