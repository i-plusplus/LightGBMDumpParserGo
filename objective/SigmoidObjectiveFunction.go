package objective

import "math"

type SigmoidObjectiveFunction struct {

}

func (SigmoidObjectiveFunction) Apply(d float64) float64  {
	return math.Exp(d)/(1+math.Exp(d))
}