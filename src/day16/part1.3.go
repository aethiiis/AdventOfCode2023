package main

import (
	"fmt"
)

func Part1_3(grid []string) int {

	a := [][4]int{{0, -1, 0, 1}}
	seen := make(map[[4]int]bool)
	q := make([][4]int, 0)

	q = append(q, a[0])

	for len(q) > 0 {
		r, c, dr, dc := q[0][0], q[0][1], q[0][2], q[0][3]
		q = q[1:]

		r += dr
		c += dc

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}

		ch := grid[r][c]

		if ch == '.' || (ch == '-' && dc != 0) || (ch == '|' && dr != 0) {
			if _, ok := seen[[4]int{r, c, dr, dc}]; !ok {
				seen[[4]int{r, c, dr, dc}] = true
				q = append(q, [4]int{r, c, dr, dc})
			}
		} else if ch == '/' {
			dr, dc = -dc, -dr
			if _, ok := seen[[4]int{r, c, dr, dc}]; !ok {
				seen[[4]int{r, c, dr, dc}] = true
				q = append(q, [4]int{r, c, dr, dc})
			}
		} else if ch == '\\' {
			dr, dc = dc, dr
			if _, ok := seen[[4]int{r, c, dr, dc}]; !ok {
				seen[[4]int{r, c, dr, dc}] = true
				q = append(q, [4]int{r, c, dr, dc})
			}
		} else {
			for _, d := range [2][2]int{{1, 0}, {-1, 0}} {
				dr, dc := d[0], d[1]
				if _, ok := seen[[4]int{r, c, dr, dc}]; !ok {
					seen[[4]int{r, c, dr, dc}] = true
					q = append(q, [4]int{r, c, dr, dc})
				}
			}
		}
	}

	coords := make(map[[2]int]bool)

	for k := range seen {
		coords[[2]int{k[0], k[1]}] = true
	}

	fmt.Println(len(coords))
	return len(coords)
}
