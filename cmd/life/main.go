package main

import (
	"image"
	"life/internal/galg"
	"life/internal/logger"
	"life/internal/renderer"
	"life/internal/utils"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	population := galg.Population{
		PopulationSize:  200,
		GenomeSize:      6,
		GenerationSteps: 100,
		Goal:            utils.RandFloat(0.0, 1.0),
		//Goal:         0.85,
		MutationRate: 0.00,
		Network: galg.Network{
			Input:    2,
			Internal: 1,
			Output:   2,
		},
	}

	population.Spawn()

	avgFitness := population.GetAverageFitness()
	generation := 1
	frames := []*image.Paletted{}

	//for avgFitness < 0.97 {
	for generation < 200 {
		rand.Seed(time.Now().UnixNano())

		population.RunGeneration()
		population.RankCreatures()

		bestFitness := population.Creatures[0].Fitness
		avgFitness = population.GetAverageFitness()

		logger.Info("Generation: %d | Goal: %f", generation, population.Goal)
		logger.Info("Best Fitness: %f", bestFitness)
		logger.Info("Average Fitness: %f\n", avgFitness)

		frames = append(frames, renderer.CreateFrame(population))

		population.Reproduce()
		population.Goal = utils.RandFloat(0.0, 1.0)

		generation++
	}

	renderer.CreateGif(frames, "result.gif")
}
