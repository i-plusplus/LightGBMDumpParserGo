package factory

import (
	"FirstProject/booster"
	"strconv"
	"strings"
)

type Pair struct{
	left int
	right int
}
type BuildTree struct {
	numLeaves int
	numCat int
	featureName []string
	splitFeature []int
	decisionType []int
	childList []Pair
	leftChild []int
	rightChild []int
	leafValues []float64
	catBoundry []int
	catThreshold []int64
	featureInfo [][]int
	pandasCategories [][]string
	threshold []float64
}





func build(bt BuildTree) booster.Booster{
	var boosters  = make(map[int]*booster.Booster,0)
	for i:=0;i< len(bt.decisionType);i++{
		var condition booster.Condition
		if bt.decisionType[i] == 5 || bt.decisionType[i] == 1 {
			var values []int64
			k := int64(bt.threshold[i])
			for d := bt.catBoundry[k];d<bt.catBoundry[k+1];d++{
				values = append(values, bt.catThreshold[d])
			}
			var v , f = booster.GetFeatures(values, bt.pandasCategories[bt.splitFeature[i]])
			condition = booster.CategoricalCondition{ bt.featureName[bt.splitFeature[i]], f,  v}
		}
		if bt.decisionType[i] == 2{
			condition = booster.NumericalCondition{bt.featureName[bt.splitFeature[i]], bt.threshold[i]}
		}
		boosters[i] = &booster.Booster{}
		b := boosters[i]
		b.SetCondition(condition)
	}

	for i:=0; i< len(bt.decisionType);i++{
		leftBooster := boosters[bt.childList[i].left]

		if bt.childList[i].left < 0 {
			leftBooster = &booster.Booster{}
			leftBooster.SetValue(bt.leafValues[bt.childList[i].left*-1 - 1])
			leftBooster.SetIsLeaf(true)
			leftBooster.SetIndex(bt.childList[i].left)
		}

		rightBooster := boosters[bt.childList[i].right]

		if bt.childList[i].right < 0 {
			rightBooster = &booster.Booster{}
			rightBooster.SetValue(bt.leafValues[bt.childList[i].right*-1 - 1])
			rightBooster.SetIsLeaf(true)
			rightBooster.SetIndex(bt.childList[i].right)
		}
		b := boosters[i]
		b.SetLeft(leftBooster)
		b.SetRight(rightBooster)

	}

	return *boosters[0]

}


func (bt BuildTree) set(lines []string, featureName []string, featureInfo [][]int, pandasCategories [][]string) booster.Booster{
	bt.featureInfo = featureInfo
	bt.featureName = featureName
	bt.pandasCategories = pandasCategories

	for _, line := range lines {
		k := strings.Split(line, "=")
		switch k[0] {
		case "num_leaves":
			bt.numLeaves, _ = strconv.Atoi(k[1])
			break
		case "num_cat":
			bt.numCat, _ = strconv.Atoi(k[1])
			break
		case "split_feature":
			var sf []int
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.Atoi(s)
				sf = append(sf, i)
			}
			bt.splitFeature = sf
			break
		case "decision_type":
			var sf []int
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.Atoi(s)
				sf = append(sf, i)
			}
			bt.decisionType = sf
			break
		case "left_child":
			var sf []int
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.Atoi(s)
				sf = append(sf, i)
			}
			bt.leftChild = sf
			break
		case "right_child":
			var sf []int
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.Atoi(s)
				sf = append(sf, i)
			}
			bt.rightChild = sf
			break
		case "leaf_value":
			var sf []float64
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.ParseFloat(s, 64)
				sf = append(sf, i)
			}
			bt.leafValues = sf
			break
		case "threshold":
			var sf []float64
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.ParseFloat(s, 64)
				sf = append(sf, i)
			}
			bt.threshold = sf
			break
		case "cat_boundaries":
			var sf []int
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.Atoi(s)
				sf = append(sf, i)
			}
			bt.catBoundry = sf
			break
		case "cat_threshold":
			var sf []int64
			for _, s := range strings.Split(k[1], " ") {
				i, _ := strconv.ParseInt(s, 10, 64)
				sf = append(sf, i)
			}
			bt.catThreshold = sf
			break
		}
	}

		for i := 0 ; i<len(bt.leftChild); i++ {
			bt.childList = append(bt.childList, Pair{bt.leftChild[i], bt.rightChild[i]})
		}
		return build(bt)
}




