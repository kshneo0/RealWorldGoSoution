package fuzz

func Fuzz(data []byte) int {
	amount := string(data)

	_, err := ConvertStringDollarsToPennies(amount)
	if err != nil {
		return -1
	}
	return 1
}
