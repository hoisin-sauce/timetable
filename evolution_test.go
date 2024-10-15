package timetable

import "testing"

func TestSelectElites(t *testing.T){
	classCounts, classCodes, _ := SetupEnvironment()
	population := InitialisePopulation(classCodes, classCounts, 4)
	for _, chromo := range population{
		chromo.fitness = GetFitness(chromo, population)
	}
	SelectElites(100, population)
}

func BenchmarkSelectElites(b *testing.B){
	classCounts, classCodes, _ := SetupEnvironment()
	population := InitialisePopulation(classCodes, classCounts, 4)
	i := 1
	limit := len(population)
	for range b.N{
		SelectElites(i, population)
		i++; i %= limit
		if i == 0{
			i = 1
		}
	}
}

func BenchmarkSelectElitesQsort(b *testing.B){
	classCounts, classCodes, _ := SetupEnvironment()
	population := InitialisePopulation(classCodes, classCounts, 4)
	i := 0
	limit := len(population)
	for range b.N{
		SelectElitesQsort(100, population)
		i++; i %= limit;
		if i == 0{
			i = 1
		}
	}
}
