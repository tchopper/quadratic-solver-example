package main

import (
	"flag"
	"fmt"
	"os"
	"quadratic/quadratic"
)

func main() {
	//./quadratic -equation="x^2 +5x +18 = 7363094" -constraint_lower=0 -constraint_upper =10000
	eqPtr := flag.String("equation", "", " Equation to solve (Required)")
	constraintLowerPtr := flag.Int("constraint_lower", 0, "Constraint for the solution lower bound (Optional)")
	constraintUpperPtr := flag.Int("constraint_upper", 0, "Constraint for the solution upper bound (Optional)")
	flag.Parse()

	if *eqPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	lowerLimit := 0
	upperLimit := 0

	fmt.Printf("Solving equation: %s\n", *eqPtr)
	if *constraintLowerPtr != 0 || *constraintUpperPtr != 0 {
		if *constraintLowerPtr > *constraintUpperPtr {
			fmt.Printf("Upper limit must be greater than lower limit \n")
			os.Exit(1)
		}
		upperLimit = *constraintUpperPtr
		lowerLimit = *constraintLowerPtr
		fmt.Printf("Constraining answer to be between %d and %d\n", *constraintLowerPtr, *constraintUpperPtr)
	}

	f, err := quadratic.StringToFormula(*eqPtr)
	if err != nil {
		fmt.Printf("An error occurred parsing the equation %s \n", err.Error())
		os.Exit(1)
	}

	root1, root2, err := f.Solve()
	if err != nil {
		fmt.Printf("An error occurred solving the equation %s \n", err.Error())
		os.Exit(1)
	}

	if lowerLimit == 0 && upperLimit == 0 {
		fmt.Printf("Answers are %f and %f\n", root1, root2)
	} else {
		if root1 >= float64(lowerLimit) && root1 <= float64(upperLimit) {
			fmt.Printf("Answer is %f\n", root1)
		}

		if root2 >= float64(lowerLimit) && root2 <= float64(upperLimit) {
			fmt.Printf("Answer is %f\n", root2)
		}
	}

}
