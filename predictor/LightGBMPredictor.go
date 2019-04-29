package predictor

import (
	"FirstProject/booster"
	"FirstProject/objective"
)

type LightGBMPredictor struct {
	Boosters    []booster.Booster
	ObjFunction objective.ObjectiveFunction
}

func (lgbm LightGBMPredictor) Predict(input *map[string]string)  float64 {
	d := 0.0;
	i := 0
	for _, booster := range lgbm.Boosters {
		l,_ := booster.GetValue(input)
		d+=l
		i++
	}
	return lgbm.ObjFunction.Apply(d);
}

