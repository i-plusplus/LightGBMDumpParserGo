package booster

import "strconv"

type NumericalCondition struct {
	Feature string
	Value   float64
}

func (nc NumericalCondition) IsLeft(input *map[string]string) (bool, error){
	v, err := strconv.ParseFloat((*input)[nc.Feature], 64)
	if(err != nil){
		return false, err
	}
	return v <= nc.Value, nil
}