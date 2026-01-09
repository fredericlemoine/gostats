package gostats

func (gsr *GoStatRand) Float64Range(a, b int) float64 {
	return float64(a) + gsr.rand.Float64()*float64(b-a)
}

func (gsr *GoStatRand) Float64RangeF(a, b float64) float64 {
	return a + gsr.rand.Float64()*(b-a)
}

func (gsr *GoStatRand) Proba(p float64) bool {
	return (gsr.rand.Float64() < p)
}

func Dunif(value, min, max float64) float64 {
	if value < min || value > max {
		return 0.0
	}
	return 1.0 / (max - min)
}
