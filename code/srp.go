package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) name() string {
	return "circle"
}

type square struct {
	length float64
}

func (c square) area() float64 {
	return c.length * c.length
}

func (c square) name() string {
	return "square"
}

type shape interface {
	name() string
	area() float64
}

type output struct{}

func (o output) Text(sh shape) string {
	return fmt.Sprintf("area of shape %s is %f", sh.name(), sh.area())
}

func (o output) JSON(sh shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {

	c := circle{
		radius: 3,
	}
	o := output{}
	fmt.Println(o.Text(c))
	fmt.Println(o.JSON(c))

	s := square{
		length: 3,
	}

	fmt.Println(o.Text(s))
	fmt.Println(o.JSON(s))

}
