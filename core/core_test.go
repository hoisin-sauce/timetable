package timetable

import "testing"

func TestGenerateTimetable(t *testing.T){
    GenerateTimetable(4, 5, 5, 20, 50, []string{"Maths", "English", "Science"}, []int{5, 5, 5}, []string{"James", "Janet", "Julia", "Jules", "Jennifer", "Jyle", "Jyke", "Aaron", "Aardvark", "Avagadro"})
}