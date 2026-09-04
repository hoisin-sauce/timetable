package main

import (
	"fmt"
)

//structure for lesson for easier return
type timetableLesson struct {
	Id        int
	Timeslot  uint32
	Classroom uint32
	Teacher   string
	Class     uint32
	Subject   string
	Students  []int
}

//generate timetable
func GenerateTimetableGo(lessonsPerClass int, days int, lessonsPerDay int, classroomCount int, eliteCount int, subjects []string, classes []int, teachers []string) (lessons []timetableLesson) {
	var (
		elites            []chromosome
		classCounts       map[string]int
		classCodes        map[string]int
		teacherCodeMap    map[int]string
		populationFitness float64
		eliteInGeneration int
		population        []chromosome
		fullElite         int
		iterations        uint
	)

    //generate encodings
	classCounts, classCodes = GenerateClassMaps(subjects, classes)
	teacherCodeMap = GenerateTeacherMap(teachers)

	//set problem constraints
	SetConstraints(days, lessonsPerDay, classroomCount, teacherCodeMap, classCodes, classCounts)

	//set up initial conditions
	population = InitialisePopulation(classCodes, classCounts, lessonsPerClass)
	populationFitness = 0
	iterations = 0
	fullElite = len(population) - 1
	for populationFitness < 1 {
	    //calculate the number of elites in the generation (unchanged members)
		eliteInGeneration = min(int(float64(eliteCount)*(float64(1)-(populationFitness*0.01))), fullElite)

        //select elite members of population
		elites = SelectElites(eliteInGeneration, population)

		//apply mutations
		CreateMutations(population, elites)

		//calculate overall fitness
		populationFitness = GetMeanPopulationFitness(population)

		//adjust randomness according to arbitrary bounds
		if populationFitness > 0.99 {
			monteCarloConstant = 0.25
		} else if populationFitness > 0.97 {
			monteCarloConstant = 0.1
		}

        //in case problem cannot be solved, break condition at 10000 iterations
		iterations++
		if iterations > 10000 {
			fmt.Println(populationFitness)
			break
		}
	}

    //convert from uint32 to timetableLesson
	lessons = make([]timetableLesson, 0)
	for _, member := range population {
		lesson := convertPopulationReadable(member, teacherCodeMap, classCodes)
		lessons = append(lessons, lesson)
	}
	return
}

//make the population human readable (helps for debugging and returning to bython)
func convertPopulationReadable(individual chromosome, teacherCodeMap map[int]string, classCodes map[string]int) (lesson timetableLesson) {
	subjectClassDecodes := InverseSubjectMap(classCodes)
	lesson = timetableLesson{
		Id:        lessonCount,
		Timeslot:  GetRegion(individual.gene, "timeSlot"),
		Classroom: GetRegion(individual.gene, "classroom"),
		Teacher:   teacherCodeMap[int(GetRegion(individual.gene, "teacher"))],
		Class:     GetRegion(individual.gene, "class"),
		Subject:   subjectClassDecodes[int(GetRegion(individual.gene, "subject"))],
	}
	lessonCount++
	return
}

//check that no duplicate IDs are introduced (used for testing)
func debugIDS(population []chromosome, arg string) {
	var seen []int = make([]int, 0)
	var duplicount int = 0
	for _, member := range population {
		for _, id := range seen {
			if member.id == id {
				duplicount++
			}
		}
		seen = append(seen, member.id)
	}
	fmt.Println("DUPLICOUNT", duplicount, arg)
}
