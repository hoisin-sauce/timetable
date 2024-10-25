package timetable

import "fmt"

func GenerateTimetable(lessonsPerClass int, days int, lessonsPerDay int, classroomCount int, eliteCount int,  subjects []string, classes []int, teachers []string)(population []chromosome){
    var (
        elites []chromosome
        classCounts map[string]int
        classCodes  map[string]int
        teacherCodeMap map[int]string
    )

    classCounts, classCodes = GenerateClassMaps(subjects, classes)
    teacherCodeMap = GenerateTeacherMap(teachers)
    SetConstraints(days, lessonsPerDay, classroomCount, teacherCodeMap, classCodes, classCounts)
    population = InitialisePopulation(classCodes, classCounts, lessonsPerClass)
    for GetMeanPopulationFitness(population) < 0.97{
        elites = SelectElites(eliteCount, population)
        CreateMutations(population, elites)
        i := GetMeanPopulationFitness(population)
        fmt.Println(i)
    }
    return
}