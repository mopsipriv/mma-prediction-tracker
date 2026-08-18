package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"mma-prediction-tracker/internal/models"
)

type FighterRepository struct {
	db *sql.DB
}

func NewFighterRepository(db *sql.DB) *FighterRepository {
	return &FighterRepository{db: db}
}

func (r *FighterRepository) GetAll(ctx context.Context) ([]models.Fighter, error) {
	query := `
		SELECT id, first_name, last_name, nickname, weight_class, created_at 
		FROM fighters 
		ORDER BY id ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query fighters: %w", err)
	}
	defer rows.Close()

	var fighters []models.Fighter

	for rows.Next() {
		var f models.Fighter
		err := rows.Scan(
			&f.ID,
			&f.FirstName,
			&f.LastName,
			&f.Nickname,
			&f.WeightClass,
			&f.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan fighter: %w", err)
		}
		fighters = append(fighters, f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return fighters, nil
}

func (r *FighterRepository) GetByID(ctx context.Context, id int64) (*models.Fighter, error) {
	query := `
		SELECT id, first_name, last_name, nickname, weight_class, created_at
		FROM fighters
		WHERE id = $1`

	var f models.Fighter
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&f.ID,
		&f.FirstName,
		&f.LastName,
		&f.Nickname,
		&f.WeightClass,
		&f.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get fighter by id: %w", err)
	}

	return &f, nil
}

func (r *FighterRepository) Create(ctx context.Context, fighter *models.Fighter) error {
	query := `
		INSERT INTO fighters(first_name, last_name, nickname, weight_class)
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		fighter.FirstName,
		fighter.LastName,
		fighter.Nickname,
		fighter.WeightClass,
	).Scan(&fighter.ID, &fighter.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create fighter: %w", err)
	}

	return nil
}