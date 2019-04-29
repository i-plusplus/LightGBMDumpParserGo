package booster

type Condition interface {
	IsLeft(input *map[string]string) (bool, error)
}
