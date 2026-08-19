package repository

import (
	"context"
	"database/sql"
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
	return fights,nil
}