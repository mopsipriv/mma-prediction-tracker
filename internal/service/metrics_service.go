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

func CalculateAccuracy(correct int, total int) float64{
    if total == 0{
        return 0.0
    }
    return (float64(correct) / float64(total)) * 100.0
}