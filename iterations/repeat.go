package iterations

func Repeat(text string, numberOfTimes int) string {
	var repeated string
	for i := 0; i < numberOfTimes; i++ {
		repeated += text
	}
	return repeated
}
