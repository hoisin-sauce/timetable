package timetable

import "testing"

func TestInitialisePopulation(t *testing.T){

	teachers := getTestTeachers("a", 10)
	subjects := []string{"maths", "english", "science", "history", "it"}
	classes  := []int{10, 10, 10, 5, 5}

	classCounts, classCodes := generateClassMaps(subjects, classes)
	teacherCodeMap := generateTeacherMap(teachers)
	setConstraints(5, 6, 127, teacherCodeMap, classCodes, classCounts)
	population := initialisePopulation(classCodes, classCounts, 4)

	if len(population) == 0{
		t.Error("Population not returned")
	}
	for i, c := range population{
		for region, _ := range masks{
			if OutsideRegion(c.gene, region){
				valueInRegion := (c.gene & masks[region]) >> startingMotion[region] 
				t.Error("Gene ", i, " is outside of ", region, " bounds upon initialisation. ", region, " is ", valueInRegion,", should be under ", constraints[region])
			}
		}
	}
}

func getTestTeachers(name string, amount int)(teachers []string){
	for i := 0; i < amount; i++{
		teachers = append(teachers, name)
	}
	return
}