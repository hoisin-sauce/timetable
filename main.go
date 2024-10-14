package timetable

import "fmt"

func main(){
	allmask := masks["timeSlot"] | masks["teacher"] | masks["classroom"] | masks["class"] | masks["subject"]
	fmt.Print(allmask)
}