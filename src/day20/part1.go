package main

import (
	"fmt"
	"strings"
)

// false = low && high = true

type Signal struct {
	pulse  bool
	origin string
	dest   string
}

type FlipFlop struct {
	status      bool
	destination []string
}

type Conjunction struct {
	recent      map[string]bool
	destination []string
}

type Broadcast struct {
	destination []string
}

func Part1(lines []string) int {
	var FlipFlops map[string]FlipFlop
	var Conjunctions map[string]Conjunction
	var broadcaster Broadcast

	FlipFlops = make(map[string]FlipFlop)
	Conjunctions = make(map[string]Conjunction)

	for _, line := range lines {
		if strings.HasPrefix(line, "%") {
			nameDest := strings.Split(line[1:], " -> ")
			//fmt.Println(nameDest)
			name := nameDest[0]
			dest := strings.Split(strings.Trim(nameDest[1], " "), ", ")
			for i := range dest {
				dest[i] = strings.Trim(dest[i], ",")
			}
			FlipFlops[name] = FlipFlop{
				status:      false,
				destination: dest,
			}
		}
		if strings.HasPrefix(line, "&") {
			nameDest := strings.Split(line[1:], " -> ")
			//fmt.Println(nameDest)
			name := nameDest[0]
			dest := strings.Split(strings.Trim(nameDest[1], " "), ", ")
			for i := range dest {
				dest[i] = strings.Trim(dest[i], ",")
			}
			Conjunctions[name] = Conjunction{
				recent:      make(map[string]bool),
				destination: dest,
			}
		}
		if strings.HasPrefix(line, "broadcaster") {
			nameDest := strings.Split(line, " -> ")
			//fmt.Println(nameDest)
			dest := strings.Split(strings.Trim(nameDest[1], " "), ", ")
			for i := range dest {
				dest[i] = strings.Trim(dest[i], ",")
			}
			broadcaster = Broadcast{destination: dest}
		}
	}

	// Enregistrement des modules connectés aux conjonctions
	for key1, conjonction := range Conjunctions {
		for key2, flipFlop := range FlipFlops {
			for _, d := range flipFlop.destination {
				if d == key1 {
					conjonction.recent[key2] = false
				}
			}
		}
		for key3, autreConjonction := range Conjunctions {
			for _, d := range autreConjonction.destination {
				if d == key1 {
					conjonction.recent[key3] = false
				}
			}
		}
		Conjunctions[key1] = conjonction
	}

	/*fmt.Println("Voici FlipFlops")
	for key, value := range FlipFlops {
		fmt.Println(key, value)
	}
	fmt.Println("Voici Conjunctions")
	for key, value := range Conjunctions {
		fmt.Println(key, value)
	}
	fmt.Println("Voici broadcaster")
	fmt.Println(broadcaster)*/

	low := 0
	high := 0
	rx := false
	for !rx {
		var queue []Signal
		queue = append(queue, Signal{
			pulse:  false,
			origin: "button",
			dest:   "broadcaster",
		})
		for len(queue) != 0 {
			signal := queue[0]
			queue = queue[1:]

			if flipFlop, ok1 := FlipFlops[signal.dest]; ok1 {
				if !signal.pulse && !flipFlop.status {
					flipFlop.status = true
					for _, d := range flipFlop.destination {
						queue = append(queue, Signal{
							pulse:  true,
							origin: signal.dest,
							dest:   d,
						})
					}
				} else if !signal.pulse && flipFlop.status {
					flipFlop.status = false
					for _, d := range flipFlop.destination {
						queue = append(queue, Signal{
							pulse:  false,
							origin: signal.dest,
							dest:   d,
						})
					}
				}
				FlipFlops[signal.dest] = flipFlop
			} else if conjunction, ok2 := Conjunctions[signal.dest]; ok2 {
				/*if _, ok3 := conjunction.recent[signal.origin]; !ok3 {
					conjunction.recent[signal.origin] = false
				}*/
				conjunction.recent[signal.origin] = signal.pulse
				pulseToSend := false
				for _, value := range conjunction.recent {
					if value == false {
						pulseToSend = true
						break
					}
				}
				for _, d := range conjunction.destination {
					queue = append(queue, Signal{
						pulse:  pulseToSend,
						origin: signal.dest,
						dest:   d,
					})
				}
			} else if signal.dest == "broadcaster" {
				for _, d := range broadcaster.destination {
					queue = append(queue, Signal{
						pulse:  signal.pulse,
						origin: signal.dest,
						dest:   d,
					})
				}
			} else {
				//fmt.Printf("La destination est : %s\n", signal.dest)
				if signal.dest == "rx" && !signal.pulse {
					rx = true
					fmt.Printf("La destination est pour un signal %t : %s\n", signal.pulse, signal.dest)
				}
			}
			//fmt.Printf("%s -%t-> %s\n", signal.origin, signal.pulse, signal.dest)
			if signal.pulse {
				high++
			} else {
				low++
			}
			/*if len(queue) == 0 && count < 1000 {
				count++
				fmt.Println("On appuie sur le bouton une nouvelle fois")
				queue = append(queue, Signal{
					pulse:  false,
					origin: "button",
					dest:   "broadcaster",
				})
			}*/
		}
	}

	//fmt.Printf("Voici high %d et low %d\n", high, low)
	return high * low
}
