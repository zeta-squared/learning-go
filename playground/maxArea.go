package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0

	for i < j {
		h1, h2 := height[i], height[j]
		if h1 < h2 {
			if h1*(j-i) > area {
				area = h1 * (j - i)
			}
			i++
		} else {
			if h2*(j-i) > area {
				area = h2 * (j - i)
			}
			j--
		}
	}

	return area
}
