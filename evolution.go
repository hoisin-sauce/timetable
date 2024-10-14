package timetable

func SelectElites(n int, population []chromosome)(elites []chromosome){
	elites = population[:n]
	elites = Quicksort(elites, 0, n-1) // ascending quicksort on fitness
	for n=n; n < len(population); n++{
		if population[n].fitness > elites[0].fitness{
			elites = Insert(elites, population[n]) // insertion sort removing the 1st item
		}
	}

	return
}

func Insert(elites []chromosome, newElite chromosome)(Nelites[] chromosome){
	Nelites = elites
	Nelites[0] = newElite
	for i:= 0; i < len(elites) - 1; i++{
		if elites[i].fitness > elites[i + 1].fitness{
			elites[i], elites[i + 1] = elites[i + 1], elites[i]
		}
	}
	return
}

func Quicksort(elites []chromosome, min int, max int)(outputElites []chromosome){
	var (
		left int
		right int
		pivotIndex int
		pivot float64
	)

	left, right = min, max
	pivotIndex = (min + max) / 2 // no need for div operator with type int
	pivot = elites[pivotIndex].fitness

	for left <= right{
		for elites[left].fitness < pivot && left < max{
			left++
		}

		for elites[right].fitness > pivot && right > min{
			right--
		}

		if left <= right{
			elites[right], elites[left] = elites[left], elites[right]
			left++
			right--
		}
	}

	if min < right{
		elites = Quicksort(elites, min, right)
	}
	if max > left{
		elites = Quicksort(elites, left, max)
	}

	outputElites = elites
	return
}