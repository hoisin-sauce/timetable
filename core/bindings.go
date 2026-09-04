package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"strconv"
	"unsafe"
)

func ProcessStrings(cStrings **C.char, length C.int) []string {
	// Convert C array to Go slice
	size := int(length)
	// Wizardry found online
	tmp := (*[1<<30 - 1]*C.char)(unsafe.Pointer(cStrings))[:size:size]

	goStrings := make([]string, size)
	for i, s := range tmp {
		goStrings[i] = C.GoString(s)
	}

	return goStrings
}

func ProcessInts(cArray *C.int, length C.int) []int {
	// Convert C array to Go slice
	goInts := (*[1 << 30]C.int)(unsafe.Pointer(cArray))[:int(length):int(length)]

	var returnInts []int

	// Calculate sum
	for _, num := range goInts {
		returnInts = append(returnInts, int(num))
	}

	return returnInts
}

func StringifyTechers(teachInts []int) (teachStrs []string) {
	for _, v := range teachInts {
		teachStrs = append(teachStrs, strconv.Itoa(v))
	}
	return
}

//export GenerateTimetable
func GenerateTimetable(studentIDsC *C.int, studentIDLengthC C.int,
	studentSubjectsC **C.char, studentSubjectLengthC C.int,
	teacherIDsC *C.int, teacherIDLengthC C.int,
	teacherSubjectsC **C.char, teacherSubjectLengthC C.int,
	lessonsPerClass C.int, days C.int, lessonsPerDay C.int,
	classroomCount C.int,
	studentsPerClass C.int) *C.char {

	var (
		studentIDs      []int
		studentSubjects []string
		teacherIDs      []int
		subjectCounts  map[string]int
		subjects       []string
		classes        []int
		teachers       []string
		timetable      []timetableLesson
		contradictions map[int][]int
		eliteCount     int
	)

    //convert to golang types
	studentIDs = ProcessInts(studentIDsC, studentIDLengthC)
	studentSubjects = ProcessStrings(studentSubjectsC, studentSubjectLengthC)
	teacherIDs = ProcessInts(teacherIDsC, teacherIDLengthC)
	teachers = StringifyTechers(teacherIDs)

	//prepare data and generate bounds
	subjectCounts = CountSubjects(studentSubjects)
	subjects, classes = GenerateClassBounds(int(studentsPerClass), subjectCounts)
	eliteCount = int(float64(sum(subjectCounts)) * float64(lessonsPerClass) * 0.5)

	//generate timetable
	timetable = GenerateTimetableGo(int(lessonsPerClass), int(days), int(lessonsPerDay), int(classroomCount), int(eliteCount), subjects, classes, teachers)

	//calculate incompatible classes and assign students
	contradictions = GetClassContradictions(timetable, sum(subjectCounts), int(lessonsPerDay))
	timetable = AssignStudents(timetable, studentIDs, studentSubjects, contradictions, int(studentsPerClass))

	//return data to python
	jsonData, err := json.Marshal(timetable)
	if err != nil {
		return C.CString("{}")
	}

	return C.CString(string(jsonData))
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {}
