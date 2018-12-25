package array1

func SumAll(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func FindMaxMin(a []int) (max, min int) {
	max = a[0]
	min = a[0]
	for _, v := range a {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

func IsPrime(a int) bool {
	if a == 2 {
		return true
	}
	if a%2 == 0 {
		return false
	}
	if a <= 1 {
		return false
	} else {
		count := 0
		for i := 3; i < a; i += 2 {
			if a%i == 0 {
				count += 1
			}
			if count > 1 {
				return false
			}
		}
	}
	return true
}

func PrimeCount(a []int) (count int) {
	for _, v := range a {
		if IsPrime(v) {
			// f.Println(v)
			count += 1
		}
	}
	return count
}

func ListPrime(a []int, m, n int) (b []int) {
	for _, v := range a {
		if v <= n && v >= m {
			if IsPrime(v) {
				b = append(b, v)
			}
		}
	}
	return
}

func HasPrime(a []int) bool {
	for _, v := range a {
		if IsPrime(v) {
			return true
		}
	}
	return false
}
