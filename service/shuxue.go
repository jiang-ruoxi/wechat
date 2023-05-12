package service

import (
	"fmt"
	"math/rand"
	"time"
)

type ShuxueService struct {
	
}

func (ss *ShuxueService) GenerateAdditionList(op string, count, max int) (list []map[string]interface{}) {
	for i := 0; i < count; i++ {
		var all, fh string
		var first, second, third int
		if op == "3" {
			all, first, second, third, fh = ss.MixGenerateAddition(max)
		} else {
			all, first, second, third, fh = ss.GenerateAddition(max, op)
		}

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
		if a >= b {
			c = a - b
		} else {
			c = b - a
			var temp int = a
			a = b
			b = temp
		}

	}
	if symbol == "3" {
		symbol = "x"
		c = a * b
	}

	problem := fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
	return problem, a, b, c, symbol
}

func (ss *ShuxueService) MixGenerateAddition(max int) (string, int, int, int, string) {
	rand.Seed(time.Now().UnixNano())

	var a, b, c int
	var op string
	for {
		// 随机生成两个 0 到 10 之间的数字和一个运算符
		a = rand.Intn(max)
		b = rand.Intn(max)
		if a < b {
			a, b = b, a // 保证 a >= b
		}

		switch rand.Intn(2) {
		case 0:
			c = a + b
			op = "+"
		case 1:
			c = a - b
			op = "-"
		}

		// 如果结果大于等于 0，则跳出循环
		if c >= 0 {
			break
		}
	}

	result := fmt.Sprintf("%d %s %d = %d \n", a, op, b, c)
	return result, a, b, c, op
}
