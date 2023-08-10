package main

import (
	"fyne.io/fyne/v2/app"
	//"log"
	//"os"
	//"strconv"
)

func main() {

	app := app.New()
	calculator := NewCalculator()

	calculator.Build(app)

	app.Run()

	/* scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			log.Fatal(err)
		}

		equation := scanner.Text()

		if len(equation) == 0 {
			break
		}
		op1, op2, operator := GetArgs(equation)

		op1AsInt, err := strconv.Atoi(op1)
		op2AsInt, err2 := strconv.Atoi(op2)

		if err != nil || err2 != nil {
			log.Fatal(err, err2)
		}

		Compute(op1AsInt, op2AsInt, operator)

	} */

}
