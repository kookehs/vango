package genetics

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {
	tests := []struct {
		actual *Population
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 0,
						genes:   "baz",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "test",
				},
			},
		},
	}

	for _, v := range tests {
		a := v.actual
		a.Display()
	}
}

func TestGeneration(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected bool
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0,
						genes:   "bar",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "far",
				},
			},
			expected: true,
		},
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0,
						genes:   "bar",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "baz",
				},
			},
			expected: false,
		},
	}

	for _, v := range tests {
		a := v.actual
		e := v.expected

		if a.Generation() != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected int
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 0,
						genes:   "baz",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "test",
				},
			},
			expected: 3,
		},
	}

	for _, v := range tests {
		a := v.actual.Len()
		e := v.expected

		if a != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

func TestLess(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected bool
		x, y     int
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 1,
						genes:   "bar",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "bar",
				},
			},
			expected: true,
			x:        0,
			y:        1,
		},
	}

	for _, v := range tests {
		a := v.actual.Less(v.x, v.y)
		e := v.expected

		if a != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

func TestNewPopulation(t *testing.T) {
	tests := []struct {
		expected *Population
		target   *Chromosome
	}{
		{
			expected: &Population{
				generation: 0,
				members:    []*Chromosome{},
				target:     nil,
			},
			target: nil,
		},
	}

	for _, v := range tests {
		a := NewPopulation(nil)
		e := v.expected

		if !reflect.DeepEqual(a, e) {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

// Repeatability is tested.
func TestPopulate(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected *Population
		seed     int64
		target   *Chromosome
	}{
		{
			actual: &Population{
				generation: 0,
				members:    []*Chromosome{},
				target: &Chromosome{
					fitness: 1,
					genes:   "foo",
				},
			},
			expected: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "S8s",
					},
					&Chromosome{
						fitness: 0,
						genes:   "YN;",
					},
					&Chromosome{
						fitness: 0,
						genes:   "aB_",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "foo",
				},
			},
			seed: 42,
			target: &Chromosome{
				fitness: 1,
				genes:   "foo",
			},
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := v.actual
		e := v.expected

		a.Populate(3, v.target)

		if !reflect.DeepEqual(a, e) {
			t.Errorf("Found %v expected %v\n", a, e)
			a.Display()
		}
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected *Population
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0.5055555555555555,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 1,
						genes:   "baz",
					},
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "baz",
				},
			},
			expected: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0.5055555555555555,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 1,
						genes:   "baz",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "baz",
				},
			},
		},
	}

	for _, v := range tests {
		a := v.actual
		e := v.expected

		a.Sort()

		if !reflect.DeepEqual(a, e) {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		actual   *Population
		expected *Population
		x, y     int
	}{
		{
			actual: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0.5055555555555555,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 1,
						genes:   "baz",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "baz",
				},
			},
			expected: &Population{
				generation: 0,
				members: []*Chromosome{
					&Chromosome{
						fitness: 0,
						genes:   "foo",
					},
					&Chromosome{
						fitness: 0.5055555555555555,
						genes:   "bar",
					},
					&Chromosome{
						fitness: 1,
						genes:   "baz",
					},
				},
				target: &Chromosome{
					fitness: 1,
					genes:   "baz",
				},
			},
			x: 0,
			y: 1,
		},
	}

	for _, v := range tests {
		a := v.actual
		e := v.expected

		a.Swap(v.x, v.y)

		if !reflect.DeepEqual(a, e) {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}
