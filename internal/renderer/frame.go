package renderer

import (
	"image"
	"image/color"
	"life/internal/neat"
	"life/internal/utils"
	"math"
)

func CreateFrame(population neat.Population) *image.Paletted {
	width := 500
	height := 300

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	bg := color.RGBA{46, 52, 64, 255}
	barColor := color.RGBA{67, 76, 94, 255}
	goalColor := color.RGBA{163, 190, 140, 255}
	oneDeviation := color.RGBA{235, 203, 139, 255}
	twoDeviation := color.RGBA{208, 135, 112, 255}
	threeDeviation := color.RGBA{191, 97, 106, 255}
	img := image.NewPaletted(image.Rectangle{upLeft, lowRight}, color.Palette{bg, barColor, goalColor, oneDeviation, twoDeviation, threeDeviation})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, bg)
		}
	}

	distribution := map[int]int{}

	for i := 0; i < 50; i++ {
		distribution[i] = 0
	}

	for _, c := range population.Creatures {
		barIndex := int(math.Min(49.0, utils.Round(c.Position, 0.02)*50))
		distribution[barIndex]++
	}

	goalBarIndex := int(utils.Round(population.Goal, 0.02) * 50)
	barWidth := width / 50

	for i := 0; i < 50; i++ {
		deviation := int(math.Abs(float64(i - goalBarIndex)))
		distributionPercentage := float64(distribution[i]) / float64(population.PopulationSize)

		for x := i * barWidth; x < i*barWidth+barWidth; x++ {
			for y := 0; y < int(float64(height)*distributionPercentage); y++ {
				switch deviation {
				case 0:
					img.Set(x, height-y, goalColor)
				case 1:
					img.Set(x, height-y, oneDeviation)
				case 2:
					img.Set(x, height-y, twoDeviation)
				case 3:
					img.Set(x, height-y, threeDeviation)
				default:
					img.Set(x, height-y, barColor)
				}
			}
		}
	}

	return img
}
