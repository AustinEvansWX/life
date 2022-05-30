package neat

type Creature struct {
	Position float64
	Fitness  float64
	Genome   Genome
}

func (c *Creature) GetInputs() []float64 {
	inputs := []float64{}

	inputs = append(inputs, c.Position)

	return inputs
}
