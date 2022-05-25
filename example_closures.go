package main

import "math"

func getmyfunction(functype string) func(int) int {
	agg := 1
	switch {
	case functype == "square":
		return func(gm int) int {
			agg1 := math.Pow(float64(gm), float64(agg))
			agg = int(agg1)
			return agg
		}
	case functype == "double":
		return func(gm int) int {
			agg = gm * 2
			return agg
		}
	case functype == "simple":
		return func(gm int) int {
			return agg
		}
	default:
		return func(gm int) int {
			return agg
		}
	}
}
