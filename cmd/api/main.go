package main
package repository
import ("fmt")

type Prediction struct {
	ID int64 `json:"id"`
	PredictedProb float64 `json:"predicted_prob"`
	IsWin bool `json:"is_win"`
}

func (p *Prediction) BrierScore() float64{
	actual := 0.0
	if p.IsWin{
		actual = 1.0
	}
	diff := p.PredictedProb - actual
	return diff * diff
}

func (r *PredictionRepository) Save(p *models.Prediction) error {
    query := `INSERT INTO predictions (predicted_prob, is_win) 
              VALUES ($1, $2) 
              RETURNING id`

	row := r.db.QueryRow(query, p.PredictedProb,p.IsWin)

	err :=row.Scan(&p.ID)
	if err !=nil{
		return err
	}

	return nil
}
    // Твоя задача:
    // 1. Вызвать r.db.QueryRow(query, p.PredictedProb, p.IsWin)
    // 2. Вызвать .Scan(&p.ID), чтобы записать полученный id обратно в p.ID
    // 3. Вернуть ошибку, если она возникнет

    return nil
}