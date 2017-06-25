package quadratic

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestStringSolver(t *testing.T) {
	type EquationTest struct {
		Equation string
		Err      error
		Ans      []float64
	}

	tests := []EquationTest{
		{
			Equation: "2x^2 -1x -1 = 0",
			Err:      nil,
			Ans:      []float64{1.0, -0.5},
		},
		{
			Equation: "2x^2 -1x -1",
			Err:      nil,
			Ans:      []float64{1.0, -0.5},
		},
		{
			Equation: "-1x -1 +2x^2",
			Err:      nil,
			Ans:      []float64{1.0, -0.5},
		},
		{
			Equation: "-1x -1 = -2x^2",
			Err:      nil,
			Ans:      []float64{1.0, -0.5},
		},
		{
			Equation: "2x^2 +5x -3 = 0",
			Err:      nil,
			Ans:      []float64{-3.0, 0.5},
		},
		{
			Equation: "2x^2 +5x = 3",
			Err:      nil,
			Ans:      []float64{-3.0, 0.5},
		},
		{
			Equation: "2x^2 -2x = 0",
			Err:      nil,
			Ans:      []float64{1.0, 0.0},
		},
		{
			Equation: "2x^2  -2x = 0",
			Err:      nil,
			Ans:      []float64{1.0, 0.0},
		},
	}

	for _, test := range tests {
		t.Logf("Performing Test: %s", test.Equation)
		f, err := StringToFormula(test.Equation)

		if err != nil {
			t.Log("StringToFormula unexpectedly failed")
			t.Fail()
		}
		root1, root2, err := f.Solve()

		if err != test.Err {
			t.Log("Solve unexpectedly failed")
			t.Fail()
		}

		if root1 != test.Ans[0] && root1 != test.Ans[1] {
			t.Log(fmt.Sprintf("Solve gave the wrong answer Got:%f expected:%f or %f", root1, test.Ans[0], test.Ans[1]))
			t.Fail()
		}

		if root2 != test.Ans[0] && root2 != test.Ans[1] {
			t.Log(fmt.Sprintf("Solve gave the wrong answer Got:%f expected:%f or %f", root2, test.Ans[0], test.Ans[1]))
			t.Fail()
		}

	}

	//too many ='s case'
	_, err := StringToFormula("2x^2 -1x -1 = = 0")
	if err != ErrTooManyEquals {
		t.Log("StringToFormula gave an unknown error for too many equals case")
		t.Fail()
	}

	//too many ='s case'
	_, err = StringToFormula("2x^2 -1x -1 = = 0")
	if err != ErrTooManyEquals {
		t.Log("StringToFormula gave an unknown error for too many equals case")
		t.Fail()
	}

	//can't solve
	f, err := StringToFormula("70x^2 -2x +2 = 0")
	if err != nil {
		t.Log("StringToFormula gave an unknown error for can't solve case")
		t.Fail()
	}

	_, _, err = f.Solve()
	if err != ErrCantSolve {
		pretty.Println(err)
		t.Log("StringToFormula gave an unknown error for can't solve case")
		t.Fail()
	}
}

func TestCodeEquationSolver(t *testing.T) {
	//2x^2 -1x -1 = 0

	f := NewFormula()
	f.Add(NewVariable(2, 2))
	f.Minus(NewVariable(1, 1))
	f.Minus(NewConstant(1))
	f.Equals(NewConstant(0))

	root1, root2, err := f.Solve()
	if err != nil {
		t.Log("Expected no error")
		t.Fail()
	}

	if root1 != 1.0 {
		t.Log(fmt.Sprintf("Solve gave the wrong answer Got:%f expected:%f", root1, 1.0))
		t.Fail()
	}

	if root2 != -0.5 {
		t.Log(fmt.Sprintf("Solve gave the wrong answer Got:%f expected:%f", root2, -0.5))
		t.Fail()
	}
}
