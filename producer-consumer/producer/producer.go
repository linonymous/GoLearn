package producer

func Produce(i int, buffer chan<- int) {
	for {
		buffer <- i
		i++
		if i == 100 {
			break
		}
	}
}
