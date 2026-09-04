package main

import (
	"math/rand/v2"
	"slices"
)

//select elite members of the population
func SelectElites(n int, population []chromosome) (elites []chromosome) {
	if n == 0 {
		return
	}

    //avoid issues with referencing slices by duplicating the population in memory
	var populationByValue []chromosome = make([]chromosome, len(population))
	copy(populationByValue, population)

	elites = populationByValue[:n]
	elites = Quicksort(elites, 0, n-1) // ascending quicksort on fitness
	for n = n; n < len(population); n++ {
		if population[n].fitness > elites[0].fitness {
			elites = Insert(elites, population[n]) // insertion sort removing the 1st item
		}
	}

	return
}

//quicksort the elites
func SelectElitesQsort(n int, population []chromosome) (elites []chromosome) {
	elites = Quicksort(population, 0, len(population)-1)[:n]
	return
}

//insert single elite
func Insert(elites []chromosome, newElite chromosome) (Nelites []chromosome) {
	Nelites = elites
	Nelites[0] = newElite
	for i := 0; i < len(elites)-1; i++ {
		if elites[i].fitness > elites[i+1].fitness {
			elites[i], elites[i+1] = elites[i+1], elites[i]
		}
	}
	return
}

//standard quicksort operation
func Quicksort(elites []chromosome, min int, max int) (outputElites []chromosome) {
	var (
		left       int
		right      int
		pivotIndex int
		pivot      float64
	)

	left, right = min, max
	pivotIndex = (min + max) / 2 // no need for div operator with type int
	pivot = elites[pivotIndex].fitness

	for left <= right {
		for elites[left].fitness < pivot && left < max {
			left++
		}

		for elites[right].fitness > pivot && right > min {
			right--
		}

		if left <= right {
			elites[right], elites[left] = elites[left], elites[right]
			left++
			right--
		}
	}

	if min < right {
		elites = Quicksort(elites, min, right)
	}
	if max > left {
		elites = Quicksort(elites, left, max)
	}

	outputElites = elites
	return
}

//create mutations within the population
func CreateMutations(population []chromosome, elites []chromosome) {
	var fitness float64
	var mutated chromosome
	var monteCarloException int = 0
	var mutations int = 0

	for i, chromo := range population {
	    //check if in elites
		if !slices.Contains(elites, chromo) {
		    //mutate and check fitness
			mutated = MutateChromosome(chromo)
			fitness = GetFitness(mutated, population)

            //if improvement or a random chance given that it is not out of bounds
            //keep the changes
			if fitness < chromo.fitness && rand.Float64() <= monteCarloConstant { //technique is part of monte carlo markov chains
				monteCarloException++
				continue
			} else if GetBoundsPenalty(mutated, 1) > GetBoundsPenalty(chromo, 1) {
				continue
			}

			mutations++
			mutated.fitness = fitness
			population[i] = mutated
		}
	}
}

func MutateChromosome(chromo chromosome) chromosome {
    //apply random changes on non protected areas of the chromosome, subject and class do not change
	var bit uint32
	for bit == 0 {
		bit = (^protectedGene) & (1 << rand.IntN(28))
	}

	chromo.gene = chromo.gene ^ bit

	return chromo
}
