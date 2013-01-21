package count

// Count up from 0 to n.
func Count(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= n; i++ {
				ch <- i
		}
	}()

	return ch
}

// Send by step.
func StepBy(src <-chan int, step int) <-chan int {

	dest := make(chan int)
	go func() {
		defer close(dest)
		for i := 0;; i++ {
			if v, ok := <-src; ok {
				if i % step == 0 {
					dest<-v
				}
			} else {
				break
			}
		}
	}()

	return dest
}
