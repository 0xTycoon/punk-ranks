package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"math"
	"testing"
)

func TestScore(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	p := 2.0
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b := math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestScore2(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	p := 2.0
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b := math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 0.00044
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b = math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 1
	for i := range values {
		y += math.Pow(math.Log(values[i]), p)
	}
	b = math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestScore3(*testing.T) {
	values := []float64{24.59, 2.8600, 0.44}
	var y float64
	//	p := 2.0
	for i := range values {
		y += math.Log(values[i])
	}
	b := y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 0.00044
	for i := range values {
		y += math.Log(values[i])
	}
	b = y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)

	values[2] = 1
	for i := range values {
		y += math.Log(values[i])
	}
	b = y // math.Exp(math.Pow(y/float64(len(values)), 1.0/p))
	fmt.Println(b)
}

func TestUnescape(t *testing.T) {

	all, err := ioutil.ReadFile("../../ranks.md")
	if err != nil {
		t.Error(err)
		return
	}
	str := html.UnescapeString(string(all))

	err = ioutil.WriteFile("../../ranks-fixed.md", []byte(str), 0644)
	if err != nil {
		t.Error(err)
		return
	}

}
