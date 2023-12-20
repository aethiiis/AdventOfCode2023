package main

import (
	"strconv"
	"strings"
	"sync"
)

// Part2 : Solution élaborée après coup, avec le recul de discussion avec mes camarades et après avoir déja examiné certaines
// solutions sur Internet. Tentative d'implémentation des channels en go.
func Part2(lines []string) int {
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
			ogSprings := springs
			ogNumbers := numbers
			for i := 0; i < 4; i++ {
				springs += "?"
				springs += ogSprings
				numbers = append(numbers, ogNumbers...)
			}
			//fmt.Println(springs, numbers)
			channel <- count2(springs, numbers, map[string]int{})

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

func count2(springs string, numbers []int, cache map[string]int) int {
	clef := springs + strconv.Itoa(len(numbers))
	if value, ok := cache[clef]; ok {
		return value
	}

	// Toutes les permutations ont été checkées
	if len(springs) == 0 {
		// Tous les nombres ont été assignés
		if len(numbers) == 0 {
			cache[clef] = 1
			return 1
			// Certains nombres ne sont pas assignés
		} else {
			cache[clef] = 0
			return 0
		}
		// Tous les nombres ont été assignés
	} else if len(numbers) == 0 {
		// On examine toutes les sources
		for _, spring := range springs {
			// S'il reste une source cassée alors que tous les nombres ont été assignés
			if spring == '#' {
				cache[clef] = 0
				return 0
			}
		}
		// Sinon c'est que tous les nombres et les sources sont assignés
		cache[clef] = 1
		return 1
		// S'il reste encore des nombres à assigner et des permutations à checker
	} else {
		sum := 0
		for _, number := range numbers {
			sum += number
		}

		// S'il ne reste plus assez de sources pour tous les nombres restants
		if len(springs) < sum+len(numbers)-1 {
			cache[clef] = 0
			return 0
		}
	}

	res := 0
	// Cas '?'
	if springs[0] == '?' {
		res += count2(springs[1:], numbers, cache)
		res += count2("#"+springs[1:], numbers, cache)

		cache[clef] = res
		return res
	}

	// Cas '.'
	if springs[0] == '.' {
		return count2(springs[1:], numbers, cache)
	}

	// Cas '#'
	if len(numbers) == 0 {
		cache[clef] = 0
		return 0
	}

	if springs[0] == '#' {
		index := 1
		for index < len(springs) && springs[index] != '.' && !(springs[index] == '?' && index == numbers[0]) {
			index++
		}
		if index == numbers[0] {
			// S'il reste des sources
			if index < len(springs) {
				// On laisse un '.'
				res += count2(springs[index+1:], numbers[1:], cache)
				cache[clef] = res
				return res
			} else {
				res += count2(springs[index:], numbers[1:], cache)
				cache[clef] = res
				return res
			}
			// Si la séquence ne matche pas
		} else {
			cache[clef] = 0
			return 0
		}
	}
	cache[clef] = 0
	return 0
}
