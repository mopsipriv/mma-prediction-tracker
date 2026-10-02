package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"mma-prediction-tracker/internal/models"
)

type PredictorRepository struct {
	db *sql.DB
}

func NewPredictorRepository(db *sql.DB) *PredictorRepository {
	return &PredictorRepository{db: db}
}

func (r *PredictorRepository) Create(ctx context.Context, predictor *models.Predictor) error {
	query := `
		INSERT INTO predictors (name)
		VALUES ($1)
		RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		predictor.Name,
	).Scan(&predictor.ID, &predictor.CreatedAt)
	
	if err != nil {
		return fmt.Errorf("failed to create predictor: %w", err)
	}

	return nil
}

func (r *PredictorRepository) GetByID(ctx context.Context, id int64) (*models.Predictor, error) {
	query := `
		SELECT id, name, created_at
		FROM predictors
		WHERE id = $1`
	
	var p models.Predictor
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Name,
		&p.CreatedAt,
	)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil 
		}
		return nil, fmt.Errorf("failed to get predictor by id: %w", err)
	}

	return &p, nil
}

func (r *PredictorRepository) GetAll(ctx context.Context) ([]models.Predictor, error) {
	query := `
		SELECT id, name, created_at
		FROM predictors
		ORDER BY id ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query predictors: %w", err)
	}
	defer rows.Close()

	var predictors []models.Predictor

	for rows.Next() {
		var p models.Predictor
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan predictor: %w", err)
		}
		predictors = append(predictors, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return predictors, nil
}