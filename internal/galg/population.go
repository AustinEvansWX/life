package galg

import (
	"life/internal/utils"
	"math"
	"sort"
)

type Population struct {
	PopulationSize  int
	GenomeSize      int
	GenerationSteps int
	MutationRate    float64
	Goal            float64
	Network         Network
	Creatures       []*Creature
}

func (p *Population) Spawn() {
	creatures := []*Creature{}

	for i := 0; i < p.PopulationSize; i++ {
		creatures = append(creatures, &Creature{
			Position: utils.RandFloat(0.0, 1.0),
			Genome:   RandomGenome(p.GenomeSize, p.Network.Input, p.Network.Internal, p.Network.Output),
		})
	}

	p.Creatures = creatures
}

func (p *Population) RunGeneration() {
	for i := 0; i < p.GenerationSteps; i++ {
		for _, creature := range p.Creatures {
			inputs := creature.GetInputs(p.Goal)

			internal := make([]float64, p.Network.Internal)

			for _, gene := range creature.Genome {
				if gene.SourceType != 0 {
					continue
				}

				internal[gene.SinkId] += inputs[gene.SourceId] * gene.Weight
			}

			for i := range internal {
				internal[i] = math.Tanh(internal[i])
			}

			output := make([]float64, p.Network.Output)

			for _, gene := range creature.Genome {
				if gene.SourceType != 1 {
					continue
				}

				output[gene.SinkId] += internal[gene.SourceId] * gene.Weight
			}

			for i := range output {
				output[i] = math.Tanh(output[i])
			}

			for i, val := range output {
				if val <= 0.0 {
					continue
				}

				rand := utils.RandFloat(0.0, 1.0)

				if rand <= val {
					switch i {
					case 0:
						creature.Position = math.Max(0.0, creature.Position-0.01)
					case 1:
						creature.Position = math.Min(1.0, creature.Position+0.01)
					}
				}
			}

			creature.Fitness = 1 - math.Abs(creature.Position-p.Goal)
		}
	}
}

func (p *Population) RankCreatures() {
	sort.SliceStable(p.Creatures, func(i, j int) bool {
		return p.Creatures[i].Fitness > p.Creatures[j].Fitness
	})
}

func (p *Population) GetAverageFitness() float64 {
	sum := 0.0

	for _, creature := range p.Creatures {
		sum += creature.Fitness
	}

	return sum / float64(p.PopulationSize)
}

func (p *Population) Reproduce() {
	bestCandidates := p.Creatures[:len(p.Creatures)/2]
	newCreatures := []*Creature{}

	i := 0

	for len(newCreatures) < p.PopulationSize {
		mom := bestCandidates[i]
		dad := bestCandidates[i+1]

		child := &Creature{
			Position: utils.RandFloat(0.0, 1.0),
			Genome:   Genome{},
		}

		for i := 0; i < p.GenomeSize/2; i++ {
			child.Genome = append(child.Genome, mom.Genome[i])
			child.Genome = append(child.Genome, dad.Genome[i+p.GenomeSize/2])
		}

		rand := utils.RandFloat(0.0, 1.0)

		if rand <= p.MutationRate {
			child.Genome[utils.RandInt(0, len(child.Genome)-1)] = RandomGene(p.Network.Input, p.Network.Internal, p.Network.Output)
		}

		newCreatures = append(newCreatures, child)

		i = (i + 1) % (p.PopulationSize/2 - 1)
	}

	p.Creatures = newCreatures
}
