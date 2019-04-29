package main

import (
	"FirstProject/factory"
	"FirstProject/predictor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

func fucntion1(input *[]map[string]string, predictor *predictor.Predictor, wg *sync.WaitGroup){
	defer (*wg).Done()
	t2 := time.Now()
	for _, i := range *input {
		(*predictor).Predict(&i)
	}
	fmt.Println((time.Now().Sub(t2)))
}



func main()  {
	t := time.Now()
	var wg sync.WaitGroup;
	f,_ := os.Open("/home/datashare/modelStr.json")
	f2,_ := os.Open("/home/datashare/input2.json")
	predictor := (&factory.LightGBM{}).Load(f)
	byteValue, _ := ioutil.ReadAll(f2)
	var input []map[string]string
	json.Unmarshal(byteValue, &input)

	wg.Add(9)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	go fucntion1(&input, &predictor, &wg)
	wg.Wait()
	fmt.Println((time.Now().Sub(t)))
}



