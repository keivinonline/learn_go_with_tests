package integers

func Calculate(operator string, input_a int, input_b int) (result int) {
	switch operator {
	case "-":
		result = input_a - input_b
	case "*":
		result = input_a * input_b
	default: // "+"
		result = input_a + input_b
	}
	return
}
