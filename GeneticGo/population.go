package main

type chromosome struct{
	gene uint32
	fitness float64
}

var constraints    map[string]uint32 = map[string]uint32{"timeSlot": 35, "teacher": 10, "class":1024, "classroom": 20, "subject": 1024}
var startingMotion map[string]uint32 = map[string]uint32{"timeSlot": 26, "teacher": 18, "class": 14,  "classroom":7,   "subject": 0}
var masks          map[string]uint32 = map[string]uint32{"timeSlot": 4227858432, "teacher": 133693440, "class": 491520, "classroom": 32512, "subject": 255}
var checkMap       map[string]uint32 = map[string]uint32{"timeSlot": 1,  "teacher": 2,  "class": 4,   "classroom": 8,  "subject": 16}
var comparisonChecks       [3]uint32 = [3]uint32{checkMap["teacher"] | checkMap["timeSlot"], checkMap["class"] | checkMap["timeSlot"], checkMap["classroom"] | checkMap["timeSlot"]}
