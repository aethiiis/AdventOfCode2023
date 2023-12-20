package main

import (
	"strconv"
	"strings"
	"sync"
)

// Part1_2 : Solution élaborée après coup, avec le recul de discussion avec mes camarades et après avoir déja examiné certaines
// solutions sur Internet. Tentative d'implémentation des channels en go.
func Part1_2(lines []string) int {
	res := 0
	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	channel := make(chan int, len(lines))

	for _, line := range lines {
		go func(line string, channel chan int) {
			springsNumbers := strings.SplitN(line, " ", 2)
			springs := springsNumbers[0]
			numbersList := strings.SplitN(springsNumbers[1], ",", (len(springsNumbers[1])/2)+1)
			var numbers []int

			numHashtagInTotal := 0
			for _, value := range numbersList {
				num, err := strconv.Atoi(value)
				if err != nil {
					panic("Erreur de conversion string --> int")
				} else {
					numbers = append(numbers, num)
					numHashtagInTotal += num
				}
			}

			channel <- count(springs, numbers, 0)

			wg.Done()
		}(line, channel)
	}
	wg.Wait()

	close(channel)

	for num := range channel {
		res += num
	}
	return res
}

func count(springs string, numbers []int, start int) int {
	for i := start; i < len(springs); i++ {
		if springs[i] == '?' {
			copy1 := springs[:i] + "." + springs[i+1:]
			res := count(copy1, numbers, i+1)
			copy2 := springs[:i] + "#" + springs[i+1:]
			res += count(copy2, numbers, i+1)
			return res
		}
	}
	if check(springs, numbers) {
		return 1
	} else {
		return 0
	}
}

func check(springs string, groups []int) bool {
	sequence := 0
	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' {
			sequence++
		} else if sequence > 0 {
			if len(groups) > 0 && sequence == groups[0] {
				groups = groups[1:]
				sequence = 0
			} else {
				return false
			}
		}
	}
	return (len(groups) == 0 && sequence == 0) || (len(groups) == 1 && sequence == groups[0])
}
