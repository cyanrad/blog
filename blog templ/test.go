package main

func pthFactor(n, p int64) int64 {
	var counter int64 = 0
	var i int64 = 1
	for ; i <= n; i++ {
		if n%i == 0 {
			counter++
		} else if p == counter {
			return i
		}
	}

	return 0
}
