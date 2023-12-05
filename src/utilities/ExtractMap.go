package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractMap(lines []string) ([]int, [][]int, [][]int, [][]int, [][]int, [][]int, [][]int, [][]int) {
	var seed2Soil [][]int
	var soil2Fertiliser [][]int
	var fertiliser2Water [][]int
	var water2Light [][]int
	var light2Temperature [][]int
	var temperature2Humidity [][]int
	var humidity2Location [][]int

	// Traitement de la première ligne - Extraction de la liste de seeds
	seedsString := lines[0]
	splitSeedsNumbers := strings.SplitN(seedsString, ":", 2)

	numSeeds := strings.Count(splitSeedsNumbers[1], " ")

	splitSeeds := strings.SplitN(splitSeedsNumbers[1], " ", numSeeds+1)

	splitSeeds = splitSeeds[1:]

	var seeds []int

	for i := range splitSeeds {
		chiffre, err := strconv.Atoi(splitSeeds[i])
		if err != nil {
			fmt.Println("Erreur de conversion :", err)
		}
		seeds = append(seeds, chiffre)
	}

	for i := range seeds {
		fmt.Printf("Les seeds sont %d\n", seeds[i])
	}
	fmt.Println("seeds fait !")

	// Traitement du premier paragraphe - Extraction de la seed-to-soil map
	startSeed2Soil := 2
	endSeed2Soil := 2

	lineRead := lines[endSeed2Soil]
	for len(lineRead) > 0 {
		endSeed2Soil++
		lineRead = lines[endSeed2Soil]
	}

	//mapSeed2Soil := make(map[int]int)

	for i := startSeed2Soil + 1; i < endSeed2Soil; i++ {
		seed2SoilStrings := strings.SplitN(lines[i], " ", 3)
		var seed2SoilNumbers []int
		for j := range seed2SoilStrings {
			chiffre, err := strconv.Atoi(seed2SoilStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			seed2SoilNumbers = append(seed2SoilNumbers, chiffre)
		}
		seed2Soil = append(seed2Soil, seed2SoilNumbers)
		/*for j := 0; j < seed2SoilNumbers[2]; j++ {
			mapSeed2Soil[seed2SoilNumbers[1]+j] = seed2SoilNumbers[0] + j
			fmt.Printf("%.2f%% de seed2Soil fait\n", float64(j)/float64(seed2SoilNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", seed2SoilNumbers[1]+j, seed2SoilNumbers[0]+j)
		}*/
	}
	fmt.Println("seed2Soil fait !")

	// Traitement du deuxième paragraphe - Extraction de la soil-to-fertiliser map
	startSoil2Fertiliser := endSeed2Soil + 1
	endSoil2Fertiliser := endSeed2Soil + 1

	lineRead = lines[endSoil2Fertiliser]
	for len(lineRead) > 0 {
		endSoil2Fertiliser++
		lineRead = lines[endSoil2Fertiliser]
	}

	//mapSoil2Fertiliser := make(map[int]int)

	for i := startSoil2Fertiliser + 1; i < endSoil2Fertiliser; i++ {
		Soil2FertiliserStrings := strings.SplitN(lines[i], " ", 3)
		var Soil2FertiliserNumbers []int
		for j := range Soil2FertiliserStrings {
			chiffre, err := strconv.Atoi(Soil2FertiliserStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Soil2FertiliserNumbers = append(Soil2FertiliserNumbers, chiffre)
		}
		soil2Fertiliser = append(soil2Fertiliser, Soil2FertiliserNumbers)
		/*for j := 0; j < Soil2FertiliserNumbers[2]; j++ {
			mapSoil2Fertiliser[Soil2FertiliserNumbers[1]+j] = Soil2FertiliserNumbers[0] + j
			fmt.Printf("%.2f%% de soil2Fertiliser fait\n", float64(j)/float64(Soil2FertiliserNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Soil2FertiliserNumbers[1]+j, Soil2FertiliserNumbers[0]+j)
		}*/
	}
	fmt.Println("Soil2Fertiliser fait !")

	// Traitement du troisième paragraphe - Extraction de la fertilizer-to-water map
	startFertiliser2Water := endSoil2Fertiliser + 1
	endFertiliser2Water := endSoil2Fertiliser + 1

	lineRead = lines[endFertiliser2Water]
	for len(lineRead) > 0 {
		endFertiliser2Water++
		lineRead = lines[endFertiliser2Water]
	}

	//mapFertiliser2Water := make(map[int]int)

	for i := startFertiliser2Water + 1; i < endFertiliser2Water; i++ {
		Fertiliser2WaterStrings := strings.SplitN(lines[i], " ", 3)
		var Fertiliser2WaterNumbers []int
		for j := range Fertiliser2WaterStrings {
			chiffre, err := strconv.Atoi(Fertiliser2WaterStrings[j])
			if err != nil {
				//fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Fertiliser2WaterNumbers = append(Fertiliser2WaterNumbers, chiffre)
		}
		fertiliser2Water = append(fertiliser2Water, Fertiliser2WaterNumbers)
		/*for j := 0; j < Fertiliser2WaterNumbers[2]; j++ {
			mapFertiliser2Water[Fertiliser2WaterNumbers[1]+j] = Fertiliser2WaterNumbers[0] + j
			fmt.Printf("%.2f%% de fertiliser2Water fait\n", float64(j)/float64(Fertiliser2WaterNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Fertiliser2WaterNumbers[1]+j, Fertiliser2WaterNumbers[0]+j)
		}*/
	}
	fmt.Println("Fertiliser2Water fait !")

	// Traitement du quatrième paragraphe - Extraction de la water-to-light map
	startWater2Light := endFertiliser2Water + 1
	endWater2Light := endFertiliser2Water + 1

	lineRead = lines[endWater2Light]
	for len(lineRead) > 0 {
		endWater2Light++
		lineRead = lines[endWater2Light]
	}

	//mapWater2Light := make(map[int]int)

	for i := startWater2Light + 1; i < endWater2Light; i++ {
		Water2LightStrings := strings.SplitN(lines[i], " ", 3)
		var Water2LightNumbers []int
		for j := range Water2LightStrings {
			chiffre, err := strconv.Atoi(Water2LightStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Water2LightNumbers = append(Water2LightNumbers, chiffre)
		}
		water2Light = append(water2Light, Water2LightNumbers)
		/*for j := 0; j < Water2LightNumbers[2]; j++ {
			mapWater2Light[Water2LightNumbers[1]+j] = Water2LightNumbers[0] + j
			fmt.Printf("%.2f%% de water2Light fait\n", float64(j)/float64(Water2LightNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Water2LightNumbers[1]+j, Water2LightNumbers[0]+j)
		}*/
	}
	fmt.Println("water2Light fait !")

	// Traitement du cinquième paragraphe - Extraction de la light-to-temperature map
	startLight2Temperature := endWater2Light + 1
	endLight2Temperature := endWater2Light + 1

	lineRead = lines[endLight2Temperature]
	for len(lineRead) > 0 {
		endLight2Temperature++
		lineRead = lines[endLight2Temperature]
	}

	//mapLight2Temperature := make(map[int]int)

	for i := startLight2Temperature + 1; i < endLight2Temperature; i++ {
		Light2TemperatureStrings := strings.SplitN(lines[i], " ", 3)
		var Light2TemperatureNumbers []int
		for j := range Light2TemperatureStrings {
			chiffre, err := strconv.Atoi(Light2TemperatureStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Light2TemperatureNumbers = append(Light2TemperatureNumbers, chiffre)
		}
		light2Temperature = append(light2Temperature, Light2TemperatureNumbers)
		/*for j := 0; j < Light2TemperatureNumbers[2]; j++ {
			mapLight2Temperature[Light2TemperatureNumbers[1]+j] = Light2TemperatureNumbers[0] + j
			fmt.Printf("%.2f%% de light2Temperature fait\n", float64(j)/float64(Light2TemperatureNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Light2TemperatureNumbers[1]+j, Light2TemperatureNumbers[0]+j)
		}*/
	}
	fmt.Println("light2Temperature fait !")

	// Traitement du sixième paragraphe - Extraction de la temperature-to-humidity map
	startTemperature2Humidity := endLight2Temperature + 1
	endTemperature2Humidity := endLight2Temperature + 1

	lineRead = lines[endTemperature2Humidity]
	for len(lineRead) > 0 {
		endTemperature2Humidity++
		lineRead = lines[endTemperature2Humidity]
	}

	//mapTemperature2Humidity := make(map[int]int)

	for i := startTemperature2Humidity + 1; i < endTemperature2Humidity; i++ {
		Temperature2HumidityStrings := strings.SplitN(lines[i], " ", 3)
		var Temperature2HumidityNumbers []int
		for j := range Temperature2HumidityStrings {
			chiffre, err := strconv.Atoi(Temperature2HumidityStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Temperature2HumidityNumbers = append(Temperature2HumidityNumbers, chiffre)
		}
		temperature2Humidity = append(temperature2Humidity, Temperature2HumidityNumbers)
		/*for j := 0; j < Temperature2HumidityNumbers[2]; j++ {
			mapTemperature2Humidity[Temperature2HumidityNumbers[1]+j] = Temperature2HumidityNumbers[0] + j
			fmt.Printf("%.2f%% de temperature2Humidity fait\n", float64(j)/float64(Temperature2HumidityNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Temperature2HumidityNumbers[1]+j, Temperature2HumidityNumbers[0]+j)
		}*/
	}
	fmt.Println("temperature2Humidity fait !")

	// Traitement du septième paragraphe - Extraction de la humidity-to-location map
	startHumidity2Location := endTemperature2Humidity + 1
	endHumidity2Location := endTemperature2Humidity + 1

	lineRead = lines[endHumidity2Location]
	for len(lineRead) > 0 {
		endHumidity2Location++
		lineRead = lines[endHumidity2Location]
	}

	//mapHumidity2Location := make(map[int]int)

	for i := startHumidity2Location + 1; i < endHumidity2Location; i++ {
		Humidity2LocationStrings := strings.SplitN(lines[i], " ", 3)
		var Humidity2LocationNumbers []int
		for j := range Humidity2LocationStrings {
			chiffre, err := strconv.Atoi(Humidity2LocationStrings[j])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			//fmt.Printf("Le chiffre append est %d\n", chiffre)
			Humidity2LocationNumbers = append(Humidity2LocationNumbers, chiffre)
		}
		humidity2Location = append(humidity2Location, Humidity2LocationNumbers)
		/*for j := 0; j < Humidity2LocationNumbers[2]; j++ {
			mapHumidity2Location[Humidity2LocationNumbers[1]+j] = Humidity2LocationNumbers[0] + j
			fmt.Printf("%.2f%%  de humidity2Location fait\n", float64(j)/float64(Humidity2LocationNumbers[2]))
			//fmt.Printf("Clé: %d, Valeur: %d\n", Humidity2LocationNumbers[1]+j, Humidity2LocationNumbers[0]+j)
		}*/
	}
	fmt.Println("humidity2Location fait !")
	return seeds, seed2Soil, soil2Fertiliser, fertiliser2Water, water2Light, light2Temperature, temperature2Humidity, humidity2Location
}
