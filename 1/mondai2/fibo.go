package fibo

func Fibo() <-chan int {
	ch := make(chan int)
	go func() {
		fibo := []int{0, 1}
		ch <- fibo[0]
		ch <- fibo[1]
		for i := 2;; i++ {
			fibo = append(fibo, fibo[i-2] + fibo[i-1])
			ch <- fibo[i]
		}
	}()

	return ch
}
