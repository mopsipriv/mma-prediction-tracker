package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"mma-prediction-tracker/internal/models"
)

type FightRepository struct {
	db *sql.DB
}

func NewFightRepository(db *sql.DB) *FightRepository {
	return &FightRepository{db: db}
}

func (r *FightRepository) GetAll(ctx context.Context) ([]models.Fight, error){
	query := `
		SELECT id, event_id, fighter_1_id, fighter_2_id, winner_id, weight_class, created_at
		FROM fights
		ORDER BY id ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query fight: %w", err)
	}
	defer rows.Close()

	var fights []models.Fight

	for rows.Next() {
		var f models.Fight
		err := rows.Scan(
			&f.ID,
			&f.EventID,
			&f.Fighter1ID, 
			&f.Fighter2ID, 
			&f.WinnerID, 
			&f.WeightClass, 
			&f.CreatedAt, 
		)
		if err != nil{
			return nil, fmt.Errorf("failed to scan fight: %w", err)
		}
		fights = append(fights,f)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return fights,nil
}

func (r *FightRepository) GetByID(ctx context.Context, id int64) (*models.Fight, error) {
	query := `
		SELECT id, event_id, fighter_1_id, fighter_2_id, winner_id, weight_class, created_at
		FROM fights
		WHERE id = $1`

	var f models.Fight
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&f.ID,
		&f.EventID,
		&f.Fighter1ID, 
		&f.Fighter2ID, 
		&f.WinnerID, 
		&f.WeightClass, 
		&f.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get fight by id: %w", err)
	}

	return &f, nil
}

func (r *FightRepository) SetWinner(ctx context.Context, fightID int64, winnerID int64) error {
	query := `
		UPDATE fights
		SET winner_id = $1
		WHERE id = $2`

	result, err := r.db.ExecContext(ctx, query, winnerID, fightID)
	if err != nil {
		return fmt.Errorf("failed to set fight winner: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("fight with id %d not found", fightID)
	}

	return nil
}