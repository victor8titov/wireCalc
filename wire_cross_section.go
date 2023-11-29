package wire_cross_section

import "math"

// Расчет сечения по его диаметру
func CrossSectionByDiameter(d) {
	s := (math.Pi * math.Pow(d, 2)) / 4
	return math.Round(s*100) / 100
}
