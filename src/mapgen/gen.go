package mapgen

import (
	"time"

	"github.com/aquilax/go-perlin"
)

type Pair struct {
	First  interface{}
	Second interface{}
}

const (
	MAXTHREADS = 10
)

var (
	chanArray chan []Pair
)

// 2D array float64 is generated
func Gen2DArray(width, height int, seed int64) [][]float64 {
	perlin := perlin.NewPerlin(3, 2.5, 3, seed)

	array := make([][]float64, width)
	for x := 0; x < width; x++ {
		array[x] = make([]float64, height)
	}
	pairs := make([]Pair, 0)

	chanArray = make(chan []Pair, 1)
	defer close(chanArray)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			array[x][y] = perlin.Noise2D(float64(x)/10, float64(y)/10)
			pairs = append(pairs, Pair{First: x, Second: y})
		}
	}

	for i := 0; i < MAXTHREADS; i++ {
		pairs = append(pairs, Pair{First: -1, Second: -1})
	}

	chanArray <- pairs

	for i := 0; i < MAXTHREADS; i++ {
		go func() {
			for {
				pairs := <-chanArray
				if pairs[0].First == -1 {
					break
				}
				array[pairs[0].First.(int)][pairs[0].Second.(int)] = perlin.Noise2D(float64(pairs[0].First.(int))/10, float64(pairs[0].Second.(int))/10)
				pairs = pairs[1:]
				chanArray <- pairs
			}
		}()
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
