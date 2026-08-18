import (
	"context",
	"database/sql",
	"fmt",
	"mma-prediction-tracker/internal/models"
)

func FighterRepository struct {
	db *sql.DB
}

func NewFighterRepository(db *sql.DB) *FighterRepository {
	return &FighterRepository{db:db}
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