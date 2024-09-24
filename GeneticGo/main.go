package main

import "fmt"

type chromosome struct{
	gene uint32
	fitness float64
}

func GetBoundsPenalty(c chromosome, populationSize int)(penalty int){
	// punish the individual so that it does not survive if any of its values are outside the available constraints
	// ensures teacher, classroom and timeslot exist

	boundsPenalty := populationSize * len(comparisonChecks) + 1

	// for each of them, mask so that it is only that section then bitshift to get the actual value, then check with constraints
	for region, mask := range masks{
		if (c.gene & mask) >> startingMotion[region] >= constraints[region]{
			penalty += boundsPenalty
		}
	}
	return
}


func GetOverlapPortions(chromosome1 chromosome, chromosome2 chromosome)(overlapPortions uint32){
	chromosomeOverlap := chromosome1.gene & chromosome2.gene

	for category, mask := range masks{
		if chromosomeOverlap & mask == chromosome1.gene & mask{
			overlapPortions |= checkMap[category]
		}
	}
	return
}

func CheckOverlapPortions(overlapPortions uint32)(checksFailed int, checkCount int){
	for j := 0; j < len(comparisonChecks); j++ {
		checkCount++
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

	for i := 0; i < len(population); i++ {
		overlapPortions = GetOverlapPortions(c, population[i])
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
		total            float64
	)

	overlapCount, totalChecks = GetOverlap(c, population)
	overlapCount += GetBoundsPenalty(c, len(population))
	overlap, total = float64(overlapCount), float64(totalChecks)
	fitness = (total - overlap)/ total
	c.fitness = fitness

	return

}

var constraints    map[string]uint32 = map[string]uint32{"timeSlot": 35, "teacher": 10, "class":1024, "classroom": 20, "subject": 1024}
var startingMotion map[string]uint32 = map[string]uint32{"timeSlot": 26, "teacher": 18, "class": 14,  "classroom":7,   "subject": 0}
var masks          map[string]uint32 = map[string]uint32{"timeSlot": 4227858432, "teacher": 133693440, "class": 491520, "classroom": 32512, "subject": 255}
var checkMap       map[string]uint32 = map[string]uint32{"timeSlot": 1,  "teacher": 2,  "class": 4,   "classroom": 8,  "subject": 16}
var comparisonChecks [3]uint32 = [3]uint32{checkMap["teacher"] | checkMap["timeSlot"], checkMap["class"] | checkMap["timeSlot"], checkMap["classroom"] | checkMap["timeSlot"]}


func main(){
	allmask := masks["timeSlot"] | masks["teacher"] | masks["classroom"] | masks["class"] | masks["subject"]
	fmt.Print(allmask)
}