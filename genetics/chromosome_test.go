package genetics

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestFitness(t *testing.T) {
	tests := []struct {
		c1, c2   *Chromosome
		expected float64
	}{
		{
			c1: &Chromosome{
				fitness: 0,
				genes:   "abb",
			},
			c2: &Chromosome{
				fitness: 1,
				genes:   "abc",
			},
			expected: -0.49444444444444446,
		},
		{
			c1: &Chromosome{
				fitness: 0,
				genes:   "abd",
			},
			c2: &Chromosome{
				fitness: 1,
				genes:   "abc",
			},
			expected: -0.49444444444444446,
		},
		{
			c1: &Chromosome{
				fitness: 0,
				genes:   "abc",
			},
			c2: &Chromosome{
				fitness: 1,
				genes:   "abc",
			},
			expected: 1,
		},
	}

	for _, v := range tests {
		a := v.c1
		a.Fitness(v.c2)
		e := v.expected

		if a.fitness != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

func TestMate(t *testing.T) {
	tests := []struct {
		p1, p2 *Chromosome
		e1, e2 *Chromosome
	}{
		{
			p1: &Chromosome{
				fitness: 0,
				genes:   "parthalf",
			},
			p2: &Chromosome{
				fitness: 0,
				genes:   "halfpart",
			},
			e1: &Chromosome{
				fitness: 0,
				genes:   "partpart",
			},
			e2: &Chromosome{
				fitness: 0,
				genes:   "halfhalf",
			},
		},
	}

	for _, v := range tests {
		r1, r2 := v.p1.Mate(v.p2)
		c1, c2 := r1.genes, r2.genes
		e1, e2 := v.e1.genes, v.e2.genes

		if (e1 != c1 && e1 != c2) || (e2 != c1 && e2 != c2) {
			t.Errorf("Found %v, %v expected %v, %v\n", c1, c2, e1, e2)
		}
	}
}

// Repeatability is tested.
func TestMutate(t *testing.T) {
	tests := []struct {
		actual   *Chromosome
		expected *Chromosome
		seed     int64
	}{
		{
			actual: &Chromosome{
				fitness: 0,
				genes:   "test",
			},
			expected: &Chromosome{
				fitness: 0,
				genes:   "-est",
			},
			seed: 64,
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := v.actual
		e := v.expected.genes

		a.Mutate(0.75)
		g := a.genes

		if g != e {
			t.Errorf("Found %v expected %v\n", g, e)
		}
	}
}

func TestNewChromosome(t *testing.T) {
	tests := []struct {
		expected *Chromosome
		fitness  float64
		genes    string
	}{
		{
			expected: &Chromosome{
				fitness: 1,
				genes:   "genes",
			},
			fitness: 1,
			genes:   "genes",
		},
	}

	for _, v := range tests {
		a := NewChromosome(v.fitness, v.genes)
		e := v.expected

		if !reflect.DeepEqual(a, e) {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

// Repeatability is tested.
func TestRandom(t *testing.T) {
	tests := []struct {
		actual   *Chromosome
		expected string
		seed     int64
	}{
		{
			actual: &Chromosome{
				fitness: 0,
				genes:   "random",
			},
			expected: "9S8s",
			seed:     42,
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := v.actual
		e := v.expected

		a.Random(len(e))
		g := a.genes

		if g != e {
			t.Errorf("Found %v expected %v\n", g, e)
		}
	}
}

// Repeatability is tested.
func TestShuffle(t *testing.T) {
	tests := []struct {
		actual   *Chromosome
		expected string
		seed     int64
	}{
		{
			actual: &Chromosome{
				fitness: 0,
				genes:   "random",
			},
			expected: "oadrnm",
			seed:     42,
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := v.actual
		e := v.expected

		a.Shuffle()
		g := a.genes

		if g != e {
			t.Errorf("Found %v expected %v\n", g, e)
		}
	}
}
