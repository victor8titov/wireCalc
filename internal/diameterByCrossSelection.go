package calculation

import "math"

// Расчет диаметра по сечению
func DiameterByCrossSelection(s float64) float64 {
	d := math.Sqrt((4 * s) / math.Pi)
	return math.Round(d*100) / 100
}
