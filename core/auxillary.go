package main


// sum function for laziness
func sum(values map[string]int) (total int) {
	total = 0
	for _, v := range values {
		total += v
	}
	return
}
