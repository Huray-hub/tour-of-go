package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	grid := make([][]uint8, dy)

	for i := range grid {
        grid[i] = make([]uint8, dx)
	}

    for _, col := range grid {
        for i, pixel := range col{
            col[i] = applyFilter(pixel) 
        }
    }
            
    return grid
}

func applyFilter(pixel uint8) uint8 {
   return pixel+30/2 
}

func main() {
	pic.Show(Pic)
}
