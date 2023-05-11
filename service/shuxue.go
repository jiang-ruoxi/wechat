package service

import (
	"fmt"
	"math/rand"
	"time"
)

type ShuxueService struct {
}

func (ss *ShuxueService) GenerateAdditionList(symbol string, count, max int) (list []map[string]interface{}) {
	for i := 0; i < count; i++ {
		all, first, second, third, fh := ss.GenerateAddition(max, symbol)
		info := make(map[string]interface{})
		info["all"] = all
		info["first"] = first
		info["second"] = second
		info["third"] = third
		info["fh"] = fh
		list = append(list, info)
	}
	return list
}

func (ss *ShuxueService) GenerateAddition(n int, symbol string) (string, int, int, int, string) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(n)
	b := rand.Intn(n - a)
	var c int
	if symbol == "1" {
		symbol = "+"
		c = a + b
	}
	if symbol == "2" {
		symbol = "-"
		c = a - b
	}
	if symbol == "3" {
		symbol = "x"
		c = a * b
	}

	problem := fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
	return problem, a, b, c, symbol
}
