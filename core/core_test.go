package main

import "testing"

func TestGenerateTimetableGo(t *testing.T) {
	GenerateTimetableGo(4, 5, 5, 20, 10, []string{"Maths", "English", "Science"}, []int{5, 5, 5}, []string{"James", "Janet", "Julia", "Jules", "Jennifer", "Jyle", "Jyke", "Aaron", "Aardvark", "Avagadro"})
}
