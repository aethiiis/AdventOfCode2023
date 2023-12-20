package main

import (
	"fmt"
	"strings"
)

// false = low && high = true

func Part2(lines []string) int {
	var FlipFlops map[string]FlipFlop
	var Conjunctions map[string]Conjunction
	var broadcaster Broadcast
	FlipFlops = make(map[string]FlipFlop)
	Conjunctions = make(map[string]Conjunction)
	for _, line := range lines {
		if strings.HasPrefix(line, "%") {
			nameDest := strings.Split(line[1:], " -> ")
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
	var output string
	for key1, conjonction := range Conjunctions {
		for _, d := range conjonction.destination {
			if d == "rx" {
				output = key1
			}
		}
	}
	var cycle map[string]int
	cycle = make(map[string]int)
	var vus map[string]int
	vus = make(map[string]int)
	for key1, conjonction := range Conjunctions {
		for _, d := range conjonction.destination {
			if d == output {
				vus[key1] = 0
			}
		}
	}
	push := 0
	for {
		push++
		var queue []Signal
		queue = append(queue, Signal{
			pulse:  false,
			origin: "button",
			dest:   "broadcaster",
		})
		for len(queue) != 0 {
			signal := queue[0]
			queue = queue[1:]
			if signal.dest == output && signal.pulse {
				vus[signal.origin] += 1
				if _, ok := cycle[signal.origin]; !ok {
					cycle[signal.origin] = push
				} else {
					if push != vus[signal.origin]*cycle[signal.origin] {
						fmt.Println("Alerte ROUGE")
					}
				}
				var test bool
				test = true
				var rx int
				rx = 1
				for key, count := range vus {
					if count == 0 {
						test = false
					} else {
						rx *= cycle[key]
					}
				}
				if test {
					return rx
				}
			}
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
				// On ne fait rien
			}
		}
	}
}
