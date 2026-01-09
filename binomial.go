package gostats

func (gsr *GoStatRand) Binomial(p float64, nb int) int {
	var binom int = 0
	for i := 0; i < nb; i++ {
		if gsr.rand.Float64() < p {
			binom++
		}
	}
	return (binom)
}
