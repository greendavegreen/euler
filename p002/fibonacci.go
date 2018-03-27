package main

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func evenFibs(limit int) int {
	f := fib()
	sum := 0
	for i := f(); i < limit; i = f(){
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}
