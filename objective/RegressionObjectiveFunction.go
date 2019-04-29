package objective


type RegressionObjectiveFunction struct {

}

func (RegressionObjectiveFunction) Apply(d float64) float64{
	return d
}