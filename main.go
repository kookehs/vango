package main

import (
	"math/rand"
	"time"

	"github.com/kookehs/vango/genetics"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	t := genetics.NewChromosome(1, "Genetic Algorithm")
	p := genetics.NewPopulation(t)
	p.Populate(10, t)

	for i := false; i != true; i = p.Generation() {
	}
}
