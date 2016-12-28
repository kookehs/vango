package genetics

import (
	"math/rand"

	"github.com/kookehs/vango/algorithms"
)

var Genes = []byte{
	' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';', '<', '=', '>', '?',
	'@', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']', '^', '_',
	'`', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '{', '|', '}', '~',
}

// Chromosome contains genes and the fitness of those genes.
type Chromosome struct {
	fitness float64
	genes   string
}

// NewChromosome returns a Chromosome with the specified fitness and genes.
func NewChromosome(f float64, g string) *Chromosome {
	return &Chromosome{fitness: f, genes: g}
}

// Fitness calculates the fitness of the calling Chromosome in relation to the provide target.
func (c *Chromosome) Fitness(t *Chromosome) {
	f := algorithm.FuzzyStringSearch(c.genes, t.genes)

	for i := 0; i < len(c.genes); i++ {
		f -= float64((c.genes[i] - t.genes[i])) * float64((c.genes[i] - t.genes[i]))
	}

	c.fitness = f
}

// Mate returns two Chromosome as a result of splitting and swapping genes.
func (c *Chromosome) Mate(m *Chromosome) (*Chromosome, *Chromosome) {
	g := c.genes
	pg := len(g) / 2

	o := m.genes
	po := len(o) / 2

	a := g[:pg] + o[po:]
	b := o[:po] + g[pg:]

	return &Chromosome{fitness: 0, genes: a}, &Chromosome{fitness: 0, genes: b}
}

// Mutate randomly swaps a gene for another.
// The provided argument f denotes the percent chance to mutate.
func (c *Chromosome) Mutate(f float64) {
	r := rand.Float64()

	if r < f {
		g := c.genes
		i := rand.Int31n(int32(len(g)))
		j := rand.Int31n(int32(len(Genes)))
		b := g[:i] + string(Genes[j]) + g[i+1:]
		c.genes = b
	}
}

// Random sets the genes of the calling Chromosome randomly
func (c *Chromosome) Random(n int) {
	c.genes = algorithm.StringRandom(n, string(Genes))
}

// Shuffle randomly shuffles the genes.
func (c *Chromosome) Shuffle() {
	c.genes = algorithm.StringShuffle(c.genes)
}
