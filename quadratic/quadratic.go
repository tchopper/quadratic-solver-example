package quadratic

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var ErrTooManyEquals = errors.New("Sorry only one = sign allowed")
var ErrCantSolve = errors.New("Couldn't find a solution")

type Variable struct {
	Coefficient float64
	Exponent    float64
}

type Constant float64

type Formula struct {
	Variables []Variable
	Constants []Constant
	Equal     Constant
}

func NewFormula() Formula {
	return Formula{}
}

func NewVariable(c float64, e float64) Variable {
	return Variable{Coefficient: c, Exponent: e}
}

func NewConstant(c float64) Constant {
	return Constant(c)
}

func (f *Formula) Add(v interface{}) {
	switch v.(type) {
	case Variable:
		f.Variables = append(f.Variables, v.(Variable))
	case Constant:
		f.Constants = append(f.Constants, v.(Constant))
	default:
		panic("wrong type")
	}
}

func (f *Formula) Minus(v interface{}) {
	switch v.(type) {
	case Variable:
		variable := v.(Variable)
		variable.Coefficient *= -1
		f.Variables = append(f.Variables, variable)
	case Constant:
		f.Constants = append(f.Constants, v.(Constant)*-1)
	default:
		panic("wrong type")
	}
}

func (f *Formula) Equals(v interface{}) {
	switch v.(type) {
	case Variable:
		f.Variables = append(f.Variables, v.(Variable))
	case Constant:
		f.Equal = v.(Constant)
	default:
		panic("wrong type")
	}
}

func (f *Formula) Solve() (float64, float64, error) {
	var a float64 = 0.0
	var b float64 = 0.0
	var c float64 = 0.0

	for _, v := range f.Variables {
		switch v.Exponent {
		case 0:
			c += v.Coefficient
		case 1:
			b += v.Coefficient
		case 2:
			a += v.Coefficient
		}
	}

	for _, constant := range f.Constants {
		c += float64(constant)
	}

	c -= float64(f.Equal)

	negB := float64(-b)
	twoA := float64(2 * a)
	bSquared := float64(b * b)
	fourAC := float64(4 * a * c)
	discrim := float64(bSquared - fourAC)

	if discrim < 0 {
		return 0, 0, ErrCantSolve
	}

	sq := math.Sqrt(discrim)
	xpos := (negB + sq) / twoA
	xneg := (negB - sq) / twoA

	return float64(xpos), float64(xneg), nil
}

func StringToFormula(formula string) (*Formula, error) {
	f := NewFormula()
	equalsSplit := strings.Split(formula, "=")

	if len(equalsSplit) > 2 {
		return nil, ErrTooManyEquals
	}

	if len(equalsSplit) == 2 {
		//
		err := f.stringToVars(equalsSplit[0], 1)
		if err != nil {
			return nil, err
		}
		err = f.stringToVars(equalsSplit[1], -1)
		if err != nil {
			return nil, err
		}
	} else {
		err := f.stringToVars(equalsSplit[0], 1)
		if err != nil {
			return nil, err
		}
	}

	return &f, nil
}

func (f *Formula) stringToVars(str string, multiplier float64) error {
	spaceSplit := strings.Split(str, " ")

	for _, i := range spaceSplit {
		if i == "" {
			continue
		}

		if strings.Contains(i, "x^2") {
			replaced := strings.Replace(i, "x^2", "", 1)
			if replaced == "" {
				f.Add(NewVariable(1*multiplier, 2))
				continue
			}
			a, err := strconv.ParseFloat(replaced, 64)
			if err != nil {
				return err
			}
			f.Add(NewVariable(a*multiplier, 2))
			continue
		}

		if strings.Contains(i, "x") {
			replaced := strings.Replace(i, "x", "", 1)
			if replaced == "" {
				f.Add(NewVariable(1*multiplier, 1))
				continue
			}
			a, err := strconv.ParseFloat(replaced, 64)
			if err != nil {
				return err
			}
			f.Add(NewVariable(a*multiplier, 1))
			continue
		}

		a, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return err
		}
		f.Add(NewConstant(a * multiplier))
	}

	return nil
}
