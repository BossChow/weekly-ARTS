package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	topAndBottomSkyline := make([]int, len(grid))
	leftAndRightSkyline := make([]int, len(grid[0]))

	for i, arr := range grid {
		for j, item := range arr {
			if item > topAndBottomSkyline[j] {
				topAndBottomSkyline[j] = item

			}

			if item > leftAndRightSkyline[i] {
				leftAndRightSkyline[i] = item

			}

		}

	}

	counter := 0
	for i, arr := range grid {
		for j, item := range arr {
			if topAndBottomSkyline[j] >= leftAndRightSkyline[i] {
				counter += leftAndRightSkyline[i] - item

			} else {
				counter += topAndBottomSkyline[j] - item

			}

		}

	}

	return counter

}
