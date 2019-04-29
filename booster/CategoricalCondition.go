package booster


type CategoricalCondition struct {
	Feature string
	Values  map[string]bool
	Match   []int32
}

func binaryOnes(values []int64) []int32{

	var l = make([]int32, 0)
	k := int32(0)
	for _, value := range values {
		for i := 0;i<32;i++ {
			if(value % 2 == 1) {
				l = append(l, k)
			}
			value = value>>1
			k++
		}
	}
	return l
}

func GetFeatures(values []int64, features []string) ([]int32, map[string]bool){
	featuresList := make(map[string]bool)
	l := binaryOnes(values)
	for _,l2:=range  l{
		featuresList[features[l2]] = true
	}
	return l, featuresList
}



func (c CategoricalCondition) IsLeft(input *map[string]string) (bool,error)  {
	return c.Values[(*input)[c.Feature]], nil
}