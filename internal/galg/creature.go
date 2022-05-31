package galg

type Creature struct {
	Position float64
	Fitness  float64
	Genome   Genome
}

func (c *Creature) GetInputs(goal float64) []float64 {
	inputs := []float64{}

	if goal < c.Position {
		inputs = append(inputs, 1)
		inputs = append(inputs, 0)
	} else {
		inputs = append(inputs, 0)
		inputs = append(inputs, 1)
	}

	return inputs
}
