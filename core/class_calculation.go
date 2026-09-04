package main

import (
	"strings"
)

// calculate the class boundaries based on the number of students with each class
func GenerateClassBounds(lessonsPerClass int, subjectCounts map[string]int) (subjects []string, counts []int) {

	for key, value := range subjectCounts {
		subjects = append(subjects, key)
		counts = append(counts, (value/lessonsPerClass)+3)
	}
	return
}

// count the amount each subject has been chosen
func CountSubjects(studentSubjects []string) (subjectCounts map[string]int) {
	var subjectList []string
	subjectCounts = make(map[string]int)
	for _, subjects := range studentSubjects {
		subjectList = strings.Split(subjects, "-")
		for _, subject := range subjectList {
			subjectCounts[subject] += 1
		}
	}
	return
}

//calculate which classes cannot go with each other
func GetClassContradictions(classes []timetableLesson, classCount int, lessonsPerClass int) (lessonIdClashes map[int][]int) {
	var lessonGroups [][]timetableLesson
	lessonGroups = make([][]timetableLesson, 0)
	lessonIdClashes = make(map[int][]int)

	//group lessons into classes
	for _, class := range classes {
		flag := true
		for i, group := range lessonGroups {
			for _, lesson := range group {
				if lesson.Subject == class.Subject && lesson.Class == class.Class {
					lessonGroups[i] = InsertTimetable(lessonGroups[i], class)
					flag = false
				}
			}
		}

        //if first of the class, make a new entry to the map
		if flag {
			new_group := make([]timetableLesson, 0)
			new_group = append(new_group, class)
			lessonGroups = append(lessonGroups, new_group)
		}
	}

    //check groups for collisions then include mark contradictions when a clash occurs
	for i, group1 := range lessonGroups {
		for j, group2 := range lessonGroups {
			if i == j {
				continue
			}
			g1Collisions, _ := GroupTimeslotCollisions(group1, group2)
			for k := range lessonGroups[i] {
				lessonIdClashes[lessonGroups[i][k].Id] = append(lessonIdClashes[lessonGroups[i][k].Id], g1Collisions...)
			}
		}
	}
	return
}

//calculate collisions that occur between groups
func GroupTimeslotCollisions(g1 []timetableLesson, g2 []timetableLesson) ([]int, []int) {
	var g1Times []uint32 = make([]uint32, 0)
	var g1ID []int = make([]int, 0)
	var g2ID []int = make([]int, 0)
	var collision bool = false

	//load group 1 times and IDs
	for _, v := range g1 {
		g1Times = append(g1Times, v.Timeslot)
		g1ID = append(g1ID, v.Id)
	}

    //load group 2 IDs and check for a collision
	for _, v := range g2 {
		g2ID = append(g2ID, v.Id)
		if in(v.Timeslot, g1Times) {
			collision = true
		}
	}

    //if a collision occurs, return the opposite IDs to be added to an exclusion list
	if collision {
		return g2ID, g1ID
	}

    //otherwise no exclusions are necessary
	return make([]int, 0), make([]int, 0)
}

// Insert a timetableLesson into a slice if its not their (pet peeve for testing)
func InsertTimetable(times []timetableLesson, lesson timetableLesson) []timetableLesson {
	for _, v := range times {
		if v.Id == lesson.Id {
			return times
		}
	}
	times = append(times, lesson)
	return times
}

// Assign students to classes using a naive algorithm that skips classes when not immediately available
func AssignStudents(classes []timetableLesson, studentIDs []int, studentSubjects []string, contradictions map[int][]int, classSize int) []timetableLesson {
	var (
		id             int
		subjects       []string
		blockedClasses []int
		contradicts    bool
		classSet       uint32
		inClasses      []int
		preBlocked     []int
		iterCount      int
		subjectIndex   int
	)

	for i, v := range studentSubjects {
	    //initialise variables for student
		iterCount = 0
		blockedClasses = make([]int, 0)
		id = studentIDs[i]
		subjects = strings.Split(v, "-")
		inClasses = make([]int, 0)
		subjectIndex = 0

		for subjectIndex < len(subjects) {
		    //prepare in case reversion necessary
			copy(preBlocked, blockedClasses)
			classSet = maxUint - 1
			for classIndex := range classes {
			    //if the lesson has the correct subject and the lesson is in the class or class is not set
				if classes[classIndex].Subject == subjects[subjectIndex] && (classSet == maxUint-1 || classes[classIndex].Class == classSet) {
				    //check for contradictions
					contradicts = in(classes[classIndex].Id, blockedClasses)
					if !contradicts && len(classes[classIndex].Students) < classSize {
					    //add blockedClasses and set the class
						blockedClasses = append(contradictions[classes[classIndex].Id], blockedClasses...)
						inClasses = append(inClasses, classIndex)
						classSet = classes[classIndex].Class
					}
				}
			}

            //check for correctly adding students
			if len(inClasses) < subjectIndex-1 && iterCount < 100 {
				iterCount++
				subjectIndex -= 2
				//revert if necessary
				copy(blockedClasses, preBlocked)
				//block previous class
				blockedClasses = append(blockedClasses, inClasses[len(inClasses)-1])
				inClasses = inClasses[:len(inClasses)-1]
			}
			subjectIndex++
		}

        //add student to classes
		for _, index := range inClasses {
			classes[index].Students = append(classes[index].Students, id)
		}
	}
	return classes
}

//shortcut of in NO KEYWORD
func in[T comparable](value T, slice []T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

//check for students having overlap
func verifyStudents(timetables []timetableLesson) bool {
	var duplicount int
	var zeroCount int
	var studentsMap map[int][]int = map[int][]int{}
	for _, lesson := range timetables {
		for _, student := range lesson.Students {
			_, ok := studentsMap[student]
			if ok {
				if in(int(lesson.Timeslot), studentsMap[student]) {
					duplicount++
					if int(lesson.Timeslot) == 0 {
						zeroCount++
					}
				}
				studentsMap[student] = append(studentsMap[student], int(lesson.Timeslot))
			} else {
				studentsMap[student] = make([]int, int(lesson.Timeslot))
			}
		}
	}

	return duplicount == 0
}
