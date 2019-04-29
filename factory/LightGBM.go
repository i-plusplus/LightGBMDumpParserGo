package factory

import (
	"FirstProject/booster"
	"FirstProject/objective"
	"FirstProject/predictor"
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type LightGBM struct {
	objectiveFunction objective.ObjectiveFunction
	featureInfo       [][]int
	pandasCategories  [][]string
	featureName       []string
}


func (lg *LightGBM) loadMetaData(reader *bufio.Scanner){
	for reader.Scan() {

		line := reader.Text()
		k := strings.Split(line, "=")
		switch k[0] {
		case "feature_infos":
			{
				for _, sub := range strings.Split(k[1], " ") {
					sub = strings.Replace(sub, "none", "0", -1)
					if strings.Contains(sub, "[") {
						continue
					}
					subList := strings.Split(sub, ":")
					var fi []int
					for _, sl := range subList {
						i, _ := strconv.Atoi(sl)
						fi = append(fi, i)
					}
					lg.featureInfo = append(lg.featureInfo, fi)
				}
				break
			}
		case "feature_names":
			{
				lg.featureName = strings.Split(k[1], " ")
				break
			}
		case "objective":
			{
				if k[1] == "binary sigmoid:1" {
					lg.objectiveFunction = objective.SigmoidObjectiveFunction{}
				} else if k[1] == "regression" {
					lg.objectiveFunction = objective.RegressionObjectiveFunction{}
				}
			}
		}
		if strings.HasPrefix(line, "pandas_categorical") {
			line = strings.SplitN(line, ":", 2)[1]
			json.Unmarshal([]byte(line), &lg.pandasCategories)
		}
	}


}

func (lg *LightGBM) GetBoosters(reader *bufio.Scanner) []booster.Booster {
	var boosters []booster.Booster
	notBreak := true
	var les []string

	for reader.Scan(){
		line := reader.Text()
		if(notBreak && !strings.HasPrefix(line, "Tree=")){
			continue
		}
		notBreak = false


		if(len(les) > 0 && strings.HasPrefix(line, "Tree=")){
			boosters = append(boosters, BuildTree{}.set(les, lg.featureName, lg.featureInfo, lg.pandasCategories))
			les = nil
		}
		les = append(les, line)
	}
	if(len(les) > 0){
		boosters = append(boosters, BuildTree{}.set(les, lg.featureName, lg.featureInfo, lg.pandasCategories))
		les = nil
	}
	return boosters
}


func (lg *LightGBM) Load(modelFile *os.File) predictor.Predictor{
	reader := bufio.NewScanner(modelFile)
	reader.Buffer(make([]byte, 1024*1024*100),1024*1024*100)
	lg.loadMetaData(reader)
	reader = bufio.NewScanner(modelFile)
	reader.Buffer(make([]byte, 1024*1024*100),1024*1024*100)
	modelFile.Seek(0,0)
	return predictor.LightGBMPredictor{lg.GetBoosters(reader), lg.objectiveFunction}
}

