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

func (r *PredictionRepository) GetByPredictorID(ctx context.Context, predictorID int64) ([]models.Prediction, error) {
	query := `
		SELECT id, predictor_id, fight_id, predicted_winner_id, probability, created_at
		FROM predictions
		WHERE predictor_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, predictorID)
	if err != nil {
		return nil, fmt.Errorf("failed to query predictions by predictor: %w", err)
	}
	defer rows.Close()

	var predictions []models.Prediction

	for rows.Next() {
		var p models.Prediction
		err := rows.Scan(
			&p.ID,
			&p.PredictorID,
			&p.FightID,
			&p.PredictedWinnerID,
			&p.Probability,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan prediction: %w", err)
		}
		predictions = append(predictions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return predictions, nil
}