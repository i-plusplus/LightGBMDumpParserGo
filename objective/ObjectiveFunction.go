package objective

type ObjectiveFunction interface {
	Apply(d float64) float64
}
