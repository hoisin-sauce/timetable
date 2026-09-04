package main

import timetable "timetable/core"

func main() {
	t := timetable.GenerateTimetableGo(4, 5, 5, 5, 60, []string{"Maths", "English", "Science"}, []int{5, 5, 5}, []string{"James", "Janet", "Julia", "Jules", "Jennifer", "Jyle", "Jyke"})
	timetable.GetClassContradictions(t, 60, 4)
}
