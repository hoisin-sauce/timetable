package main

import (
	"math/rand/v2"
)

type chromosome struct {
	gene    uint32
	fitness float64
	id      int
}

// get a bit mask of length length shifted by shift
func getBitMask(length int, shift int) (mask uint32) {
	mask = ((1 << length) - 1) << shift
	return
}

// conditions controlling the layout of genes and the necessary comparisons that are made between them
var constraints map[string]uint32
var startingMotion map[string]uint32 = map[string]uint32{"timeSlot": 15, "teacher": 7, "class": 28, "classroom": 0, "subject": 21}
var masks map[string]uint32 = map[string]uint32{"timeSlot": getBitMask(6, 15), "teacher": getBitMask(8, 7), "class": getBitMask(4, 28), "classroom": getBitMask(7, 0), "subject": getBitMask(7, 21)}
var checkMap map[string]uint32 = map[string]uint32{"timeSlot": 1, "teacher": 2, "class": 4, "classroom": 8, "subject": 16}
var comparisonChecks [3]uint32 = [3]uint32{checkMap["teacher"] | checkMap["timeSlot"], checkMap["subject"] | checkMap["class"] | checkMap["timeSlot"], checkMap["classroom"] | checkMap["timeSlot"]}
var protectedGene uint32 = masks["class"] | masks["subject"]
var monteCarloConstant float64 = 0.5
var chromoCount int = 0
var lessonCount int = 0

const maxUint uint32 = ^uint32(0)

// map subject to unique identifier
// map subject to number of classes
func GenerateClassMaps(subjects []string, classes []int) (subjectClassCount map[string]int, subjectClassCodes map[string]int) {
	subjectClassCount = make(map[string]int)
	subjectClassCodes = make(map[string]int)

	for i := 0; i < len(subjects); i++ {
		subjectClassCount[subjects[i]] = classes[i]
		subjectClassCodes[subjects[i]] = i
	}

	return
}

// assign unique identifier to teachers
func GenerateTeacherMap(teachers []string) (teacherMap map[int]string) {
	teacherMap = make(map[int]string)
	for i := 0; i < len(teachers); i++ {
		teacherMap[i] = teachers[i]
	}
	return
}

// set the constraints that all genes should abide by
func GetConstraints(days int, lessonsPerDay int, classroomCount int, teacherMap map[int]string, classCodes map[string]int, classCounts map[string]int) (_constraints map[string]uint32) {
	_constraints = make(map[string]uint32)

	_constraints["teacher"] = uint32(GetMaxKey(teacherMap) + 1)

	_constraints["classroom"] = uint32(classroomCount + 1)

	_constraints["subject"] = uint32(GetMaxValue(classCodes) + 1)

	_constraints["timeSlot"] = uint32(days * lessonsPerDay)

	_constraints["class"] = uint32(GetMaxValue(classCounts) + 1)
	return
}

// set environment
func SetConstraints(days int, lessonsPerDay int, classroomCount int, teacherMap map[int]string, classCodes map[string]int, classCounts map[string]int) {
	constraints = GetConstraints(days, lessonsPerDay, classroomCount, teacherMap, classCodes, classCounts)
}

// invert map
func InverseSubjectMap(subjectClassCodes map[string]int) (subjectClassDecodes map[int]string) {
	subjectClassDecodes = make(map[int]string)
	for k, v := range subjectClassCodes {
		subjectClassDecodes[v] = k
	}
	return
}

// get the maximum key
func GetMaxKey(m map[int]string) (max int) {
	for key, _ := range m {
		if key > max {
			max = key
		}
	}
	return
}

// get the maximum value
func GetMaxValue(m map[string]int) (max int) {
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	return
}

func InitialisePopulation(subjectCodes map[string]int, classCounts map[string]int, lessonsPerClass int) (population []chromosome) {
	var (
		chromo chromosome
	)

	// go through all subjects

	for subject, code := range subjectCodes {

		// iterate through each set for each subject
		for set := 0; set < classCounts[subject]; set++ {

			for lesson := 0; lesson < lessonsPerClass; lesson++ {
				chromo = GenerateChromosome(code, set)
				population = append(population, chromo)
			}
		}
	}
	return
}

//create a new chromosome for a subject and class
func GenerateChromosome(subjectCode int, classSet int) (chromo chromosome) {
	var gene uint32 = 0

	for section, _ := range constraints {
		gene = UpdateMaskedPortion(gene, section, rand.Uint32N(constraints[section]))
	}

	gene = UpdateMaskedPortion(gene, "subject", uint32(subjectCode))
	gene = UpdateMaskedPortion(gene, "class", uint32(classSet))

	chromo = chromosome{gene: gene, fitness: 0, id: chromoCount}
	chromoCount++
	return
}


//wizardry using bitmasks
func UpdateMaskedPortion(original uint32, mask string, newValue uint32) uint32 {
	original = original & (^masks[mask]) //wipe area of original
	original |= newValue << startingMotion[mask] //XOR to combine bits
	return original
}
