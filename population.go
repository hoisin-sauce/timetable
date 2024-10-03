package timetable

type chromosome struct{
	gene uint32
	fitness float64
}

func getBitMask(length int, shift int)(mask uint32){
	mask = ((1 << length) - 1) << shift
	return
}

var constraints    map[string]uint32 = map[string]uint32{"timeSlot": 35, "teacher": 10, "class":1024, "classroom": 20, "subject": 1024}
var startingMotion map[string]uint32 = map[string]uint32{"timeSlot": 26, "teacher": 18, "class": 14,  "classroom":7,   "subject": 0}
var masks          map[string]uint32 = map[string]uint32{"timeSlot": getBitMask(6, 26), "teacher": getBitMask(8, 18), "class": getBitMask(4, 14), "classroom": getBitMask(7, 7), "subject": getBitMask(7, 0)}
var checkMap       map[string]uint32 = map[string]uint32{"timeSlot": 1,  "teacher": 2,  "class": 4,   "classroom": 8,  "subject": 16}
var comparisonChecks       [3]uint32 = [3]uint32{checkMap["teacher"] | checkMap["timeSlot"], checkMap["class"] | checkMap["timeSlot"], checkMap["classroom"] | checkMap["timeSlot"]}
const maxUint                 uint32 = ^uint32(0)