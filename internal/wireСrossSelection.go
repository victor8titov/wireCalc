package calculation

import "math"

// Расчет сечения по его диаметру
func CrossSectionByDiameter(d float64) float64 {
	s := (math.Pi * math.Pow(d, 2)) / 4
	return math.Round(s*100) / 100
}
