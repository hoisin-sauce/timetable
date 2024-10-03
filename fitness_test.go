package timetable

import (
	"testing"
	"math/rand/v2"
	"strconv"
)

func TestGetOverlapPortions(t *testing.T){
	var (
		tests map[[2]chromosome]uint32
	)
	tests = GetOverlapPortionsTests()
	for params, expected := range tests{
		actual := GetOverlapPortions(params[0], params[1])
		if actual != expected{
			t.Error(`GetOverlapPortions(`, params[0], `,`, params[1],`) = `, actual,` EXPECTED ` ,expected)
		}
	}
}

func GetOverlapPortionsTests()(tests map[[2]chromosome]uint32){
	var i uint32
	// define cases
	// unique overlap combinations 2^5

	// get chromosome
	var chromo chromosome = GetRandomChromosome()
	tests = make(map[[2]chromosome]uint32)
	for i = 0; i < (1 << 5); i++{
		tests[[2]chromosome{chromo, CreateOverlap(chromo, i)}] = i
	}


	return
}

func GetRandomChromosome()chromosome{
	var (
		gene uint32
		fitness float64
		individual chromosome
	)

	gene = rand.Uint32()
	fitness = 0
	gene = 0
	individual = chromosome{gene: gene, fitness:fitness}
	
	return individual
}

func CreateOverlap(inputChromosome chromosome, overlapPortions uint32)(outputChromosome chromosome){
	outputChromosome.gene = inputChromosome.gene ^ maxUint
	for key, bit := range checkMap{
		// check if the portion should be overlapped
		if bit & overlapPortions == 0{
			continue
		}

		// if overlap, apply not with the masked section
		outputChromosome.gene ^= masks[key]
	}
	return
}

func TestMapping(t *testing.T){

	var totalMask uint32 = 0
	for _, mask := range masks{
		totalMask ^= mask
	} 


	if totalMask != maxUint{
		var bits string = strconv.FormatUint(uint64(totalMask), 2)
		t.Error("Masks do not xor to max value, Bit represenation: ", bits)
	}

	for key1, mask1 := range masks{
		for key2, mask2 := range masks{
			if key1 == key2{
				continue
			}

			if mask1 & mask2 > 0{
				overlap := strconv.FormatUint(uint64(mask1 & mask2), 2)
				t.Error("Mask overlap between ", key1, " and ", key2, " bit representation: ", overlap)
			}
			
		}
	}
}