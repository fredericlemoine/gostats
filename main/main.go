package main

import (
	"fmt"
	"math/rand"

	"github.com/fredericlemoine/gostats"
)

func main() {
	gsr := gostats.New(rand.New(rand.NewSource(10)))

	nvals := 1000
	for i := 0; i < 10000; i++ {
		d1, _ := gsr.Dirichlet1(float64(nvals), nvals)
		for _, v := range d1 {
			fmt.Println(v)
		}
	}
	//fmt.Println(d1)
	//fmt.Println(sum1)

	alphas := make([]float64, nvals)
	for i := 0; i < nvals; i++ {
		alphas[i] = 1.0
	}
	for i := 0; i < 10000; i++ {
		d2, _ := gsr.Dirichlet(float64(nvals), alphas...)
		for _, v := range d2 {
			fmt.Println(v)
		}
	}
	//fmt.Println(d2)
	//fmt.Println(sum2)
}
