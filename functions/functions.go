package functions

func func_main() {

	var f func(int) int
	f = func(i int) int { //Recursive func call
		if i == 0 {
			return 1
		}
		return i * f(i-1)
	}
	f(1)
}

func callTheSum() {
	var numbers = []float64{12.1, 22.3, 11.0}

	_, _ = sum(numbers...)
}

func sum(values ...float64) (float64, error) { ///"values..." says that is array of something in this case float64
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total, nil

}

//if first letter is capital, this means method is public
func Sum_Deneme() func(values ...float64) (float64, error) { //Return func from func.
	q := sum
	return q
}
