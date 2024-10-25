package timetable

func GetBoundsPenalty(c chromosome, populationSize int)(penalty int){
	// punish the individual so that it does not survive if any of its values are outside the available constraints
	// ensures teacher, classroom and timeslot exist

	boundsPenalty := populationSize * len(comparisonChecks) + 1

	// for each of them, mask so that it is only that section then bitshift to get the actual value, then check with constraints
	for region, _ := range masks{
		if OutsideRegion(c.gene, region){
			penalty += boundsPenalty
		}
	}
	return
}

func OutsideRegion(gene uint32, region string)(isWithinRegion bool){
	isWithinRegion = (gene & masks[region]) >> startingMotion[region] >= constraints[region]
	return
}


func GetOverlapPortions(chromosome1 chromosome, chromosome2 chromosome)(overlapPortions uint32){
	// bitwise & to find overlapping sections
	// chromosomeOverlap := chromosome1.gene & chromosome2.gene

	// go through all masks to check for identical
	for category, mask := range masks{
		// apply masks and check if region is identical
		if chromosome1.gene & mask == chromosome2.gene & mask{
			// if conditions are met note it in overlap
			overlapPortions |= checkMap[category]
		}
	}

	return

}

func CheckOverlapPortions(overlapPortions uint32)(checksFailed int, checkCount int){
	checkCount = len(comparisonChecks)
	// got through all checks
	for j := 0; j < len(comparisonChecks); j++ {

		// isolate areas in check and check if equal to the check
		if (comparisonChecks[j] & overlapPortions) == comparisonChecks[j]{
			checksFailed++
		}
	}

	return

}

func GetOverlap(c chromosome, population []chromosome)(overlapCount int, totalCount int){
	var(
		overlapPortions   uint32
		checksFailed      int
		checkCount        int
	)

	// for each member of the population
	for i := 0; i < len(population); i++ {
		// get the overlap between the member and the current chromosome
		overlapPortions = GetOverlapPortions(c, population[i])

		// check if the overlap is in any crucial regions
		checksFailed, checkCount = CheckOverlapPortions(overlapPortions)

		overlapCount += checksFailed
		totalCount += checkCount
	}

	return

}

func GetFitness(c chromosome, population []chromosome)(fitness float64){
	var(
		overlapCount      int
		totalChecks       int
		overlap           float64
		total             float64
	)

	// get the overlap
	overlapCount, totalChecks = GetOverlap(c, population)

	// apply bounds penalty
	overlapCount += GetBoundsPenalty(c, len(population))

	// calculate GetFitness
	overlap, total = float64(overlapCount), float64(totalChecks)
	fitness = (total - overlap)/ total

	return

}

// calculate fitness for population

func GetMeanPopulationFitness(population []chromosome)(fitness float64){

	for _, chromo := range population{
		chromo.fitness = GetFitness(chromo, population)
		fitness += chromo.fitness
	}
	fitness /= float64(len(population))

	return
}