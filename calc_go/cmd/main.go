package main

import (
	"github.com/MrM2025/rpforcalc/tree/master/calc_go/pkg/calculation"
	"fmt"
)

func main() {

	var calc calculation.IHistory

	calc.Init()
	fmt.Println(calc.Calc("-1+1*2.54+41+((3/3+10)/2-(-2.5-1+(-1))*10)-1"))
	fmt.Println(calc.Calc("1+1"))
	fmt.Println(calc.Calc("(2+2)*2"))
	fmt.Println(calc.Calc("2+2*2"))
	fmt.Println(calc.Calc("1/2"))
	fmt.Println(calc.Calc("1+1*"))
	fmt.Println(calc.Calc(""))
	fmt.Println(calc.Calc("1+((3/3+10)/2-2.5*10)"))
	fmt.Println(calc.Calc("-1+2+(-3)"))
	history := calc.GetCalcHistory()
	for key, _ := range history {
		fmt.Println(key, "result", history[key])
	}
	calc.RemoveHistory()
	fmt.Println("123", calc.GetCalcHistory())
}
