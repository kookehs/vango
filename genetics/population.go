package genetics

import (
	"fmt"
	"sort"
)

// Population contains informatio necessary for simulation.
type Population struct {
	generation int
	members    []*Chromosome
	target     *Chromosome
}

// NewPopulation returns an empty Population with the given target.
func NewPopulation(t *Chromosome) *Population {
	return &Population{
		generation: 0,
		target:     t,
		members:    []*Chromosome{},
	}
}

// Display prints the generation, genes, and fitnesses.
func (p *Population) Display() {
	fmt.Printf("Generation: %d\n", p.generation)

	for _, v := range p.members {
		fmt.Printf("%s - %g\n", v.genes, v.fitness)
	}
}

// Generation simulates a round of crossover and mutation.
func (p *Population) Generation() bool {
	t := p.target

	for _, v := range p.members {
		v.Fitness(t)
	}

	p.Sort()
	p.Display()

	// Mate the two chromosomes with the highest fitness
	l := len(p.members)
	p1, p2 := p.members[l-1], p.members[l-2]
	c1, c2 := p1.Mate(p2)
	// Remove the two with lowest fitness and append the children
	p.members = p.members[2:]
	p.members = append(p.members, c1, c2)
	f := false

	for _, v := range p.members {
		v.Mutate(0.25)
		v.Fitness(t)

		if v.genes == t.genes {
			f = true
		}
	}

	p.generation++

	if f {
		p.Sort()
		p.Display()
		return true
	}

	return false
}

// Len returns the number of members.
func (p *Population) Len() int {
	return len(p.members)
}

// Less returns whether the value at i is less than the value at j.
func (p *Population) Less(i, j int) bool {
	m := p.members
	return m[i].fitness < m[j].fitness
}

// Populate fills the Population's members with n amount of Chromosomes.
func (p *Population) Populate(n int, t *Chromosome) {
	for i := n; i > 0; i-- {
		c := NewChromosome(0, "")
		c.Random(len(t.genes))
		p.members = append(p.members, c)
	}
}

// Sort sorts the members of the Population using the sort package.
func (p *Population) Sort() {
	sort.Sort(p)
}

// Swap swaps the value at i with the value at j.
func (p *Population) Swap(i, j int) {
	p.members[i], p.members[j] = p.members[j], p.members[i]
}
