package service

func CalculateBrierScore(probability float64, didWin bool) float64{
    var outcome float64
    if didWin{
        outcome = 1.0
    } else {
        outcome = 0.0
    }
    difference := probability - outcome
    BS := difference * difference
    return BS

}