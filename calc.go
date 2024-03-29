package calc

type calculatorOperations struct {
	addition       string
	subtraction    string
	multiplication string
	division       string
}

func OpenCalculator() *calculatorOperations {
	return &calculatorOperations{}
}

func (c *calculatorOperations) Calculate(num1, num2 float64, operation string) float64 {
	c.addition = "+"
	c.subtraction = "-"
	c.multiplication = "*"
	c.division = "/"

	switch operation {
	case c.addition:
		return c.add(num1, num2)
	case c.subtraction:
		return c.sub(num1, num2)
	case c.multiplication:
		return c.mul(num1, num2)
	case c.division:
		return c.div(num1, num2)
	default:
		return 0
	}
}

func (c *calculatorOperations) add(num1, num2 float64) float64 {
	return num1 + num2
}

func (c *calculatorOperations) sub(num1, num2 float64) float64 {
	return num1 - num2
}

func (c *calculatorOperations) mul(num1, num2 float64) float64 {
	return num1 * num2
}

func (c *calculatorOperations) div(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}