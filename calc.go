/*package calc

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
}*/
// Пакет calc реализует простой калькулятор с основными арифметическими операциями.
package calc

// Структура calculator представляет калькулятор.
type calculator struct{}

// NewCalculator создает новый экземпляр структуры калькулятора.
func NewCalculator() *calculator {
	return &calculator{}
}

// Calculate выполняет арифметическую операцию над двумя числами в зависимости от переданного оператора.
// Возвращает результат операции.
func (c *calculator) Calculate(num1, num2 float64, operator string) float64 {
	switch operator {
	case "+":
		return c.add(num1, num2)
	case "-":
		return c.subtract(num1, num2)
	case "*":
		return c.multiply(num1, num2)
	case "/":
		return c.divide(num1, num2)
	default:
		return 0.0
	}
}

// add выполняет сложение двух чисел.
func (c *calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

// subtract выполняет вычитание одного числа из другого.
func (c *calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}

// multiply выполняет умножение двух чисел.
func (c *calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// divide выполняет деление одного числа на другое.
// В случае деления на ноль возвращает 0.
func (c *calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}
