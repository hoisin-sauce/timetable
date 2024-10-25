package main
import "timetable/core"

func main(){
	timetable.GenerateTimetable(4, 5, 5, 20, 5, []string{"Maths", "English", "Science"}, []int{5, 5, 5}, []string{"James", "Janet", "Julia", "Jules", "Jennifer", "Jyle", "Jyke"})
}