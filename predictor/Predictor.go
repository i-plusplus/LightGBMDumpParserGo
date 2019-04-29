package predictor

type Predictor interface {
	Predict(input *map[string]string) float64
}
