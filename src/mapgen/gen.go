package mapgen

import (
	"time"

	"github.com/aquilax/go-perlin"
)

// 2D array float64 is generated
func Gen2DArray(width, height int, seed int64) [][]float64 {
	perlin := perlin.NewPerlin(3, 2.5, 3, seed)
	var array [][]float64
	for x := 0; x < width; x++ {
		var row []float64
		for y := 0; y < height; y++ {
			row = append(row, perlin.Noise2D(float64(x)/10.0, float64(y)/10))
		}
		array = append(array, row)
	}
	return array
}

// Gen2DArray generates a 2D array of float64
func GenMap(num_mats int, mapSize int) [][]int8 {
	rand := time.Now().UnixNano()
	gen := Gen2DArray(mapSize, mapSize, rand)
	var mapArray [][]int8
	for x := 0; x < mapSize; x++ {
		var row []int8
		for y := 0; y < mapSize; y++ {
			if gen[x][y] < 0 {
				gen[x][y] *= -1
			}
			value := int8(gen[x][y] * float64(num_mats))
			// fmt.Fprintf(os.Stderr, "%d ", value)
			row = append(row, value)
		}
		mapArray = append(mapArray, row)
	}
	return mapArray
}
