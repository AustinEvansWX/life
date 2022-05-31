package galg

import (
	"life/internal/utils"
)

type Genome []Gene

type Gene struct {
	Weight     float64
	SourceType int
	SourceId   int
	SinkId     int
}

func RandomGenome(size, input, internal, output int) Genome {
	genome := Genome{}

	for j := 0; j < size; j++ {
		genome = append(genome, RandomGene(input, internal, output))
	}

	return genome
}

func RandomGene(input, internal, output int) Gene {
	weight := utils.RandFloat(-4.0, 4.0)

	gene := Gene{Weight: weight}
	gene.SourceType = utils.RandInt(0, 2)

	switch gene.SourceType {
	case 0:
		gene.SourceId = utils.RandInt(0, input)
		gene.SinkId = utils.RandInt(0, internal)
	case 1:
		gene.SourceId = utils.RandInt(0, internal)
		gene.SinkId = utils.RandInt(0, output)
	}

	return gene
}
