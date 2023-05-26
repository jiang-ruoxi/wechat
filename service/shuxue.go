package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"wechat/model"
	"wechat/pkg/mysql"
	"wechat/pkg/redis"
)

type ShuxueService struct {
}
func (bs *ShuxueService) GetHBToken() string {
	//队列名称
	redisToken := "huiben_wechat_token"

	token, _ := redis.RedisClient.Get(context.Background(), redisToken).Result()

	return token
}

// GetSXList 数学列表
func (bs *ShuxueService) GetSXList(page, pageSize int) (list []model.ShuXue, total int64, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	// 创建db
	var sxList []model.ShuXue
	db := mysql.DB.Model(&model.ShuXue{}).Debug()
	err = db.Count(&total).Error
	db = db.Order("id desc")
	db = db.Limit(limit).Offset(offset).Find(&sxList)

	return sxList, total, err
}

func (ss *ShuxueService) GenerateAdditionList(op string, count, max int, et string) (list []map[string]interface{}) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		var all, fh string
		var first, second, third int

		if max == 20 {
			if op == "1" && et == "1" {
				//rand.Seed(time.Now().UnixNano())
				var a, b, c int
				var symbol, problem string
				a = rand.Intn(20)
				b = rand.Intn(20 - a)
				c = a + b
				symbol = "+"
				problem = fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
				all = problem
				first = a
				second = b
				third = c
				fh = "+"
			} else if op == "1" && et == "3" {
				var a, b, c int
				var symbol, problem string
				//// 设置随机数种子
				//rand.Seed(time.Now().Unix())

				// 首先生成两个数字
				a = rand.Intn(max)
				b = rand.Intn(max)

				// 处理进位
				if a%10+b%10 >= 10 {
					a += 10
				}
				c = a + b
				problem = fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
				all = problem
				first = a
				second = b
				third = c
				fh = "+"
			} else {
				all, first, second, third, fh = ss.TwGenerateAddition(op, max, et)
			}
		} else {
			//op == 3是混合的加减
			if op == "3" {
				all, first, second, third, fh = ss.MixGenerateAddition(max)
			} else {
				all, first, second, third, fh = ss.GenerateAddition(max, op)
			}
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

func (ss *ShuxueService) TwGenerateAddition(op string, max int, et string) (string, int, int, int, string) {

	//op 1加法，2减法，3混合
	//et 1加法不进位，2减法不退位 3加法进位 4减法退位

	if op == "1" && et == "1" {
		rand.Seed(time.Now().UnixNano())
		var a, b, c int
		var symbol, problem string
		a = rand.Intn(10)
		b = rand.Intn(10 - a)
		c = a + b
		symbol = "+"
		problem = fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
		return problem, a, b, c, symbol
	} else if op == "1" && et == "3" {
		var problem string
		var a, b, c int
		var symbol string
		//// 设置随机数种子
		rand.Seed(time.Now().Unix())

		// 首先生成两个数字
		a = rand.Intn(max)
		b = rand.Intn(max)

		// 处理进位
		if a%10+b%10 >= 10 {
			a += 10
		}
		c = a + b
		symbol = "+"
		problem = fmt.Sprintf("%d %s %d = %d", a, symbol, b, c)
		return problem, a, b, c, symbol
	} else if op == "2" && et == "2" {
		var problem string
		var a, b, c int
		var symbol string
		rand.Seed(time.Now().UnixNano())
		a = rand.Intn(max)
		b = rand.Intn(a + 1)
		c = a - b
		symbol = "-"
		problem = fmt.Sprintf("%d - %d = %d\n", a, b, c)
		return problem, a, b, c, symbol
	} else if op == "2" && et == "4" {
		var problem string
		var a, b, c int
		var symbol string
		rand.Seed(time.Now().UnixNano())
		a = rand.Intn(20)
		b = rand.Intn(a + 1)
		if a%10 < b%10 { // 发生借位
			a = (a/10 - 1) * 10 // 借位操作
			b = (b/10 + 1) * 10 // 进位操作
		}
		if b > a {
			c = b - a
			var temp int = a
			a = b
			b = temp
		} else {
			c = a - b
		}
		symbol = "-"
		problem = fmt.Sprintf("%d - %d = %d\n", a, b, c)
		return problem, a, b, c, symbol
	} else if op == "3" {
		var problem string
		var a, b, c int
		var symbol string
		problem, a, b, c, symbol = ss.MixGenerateAddition(max)
		return problem, a, b, c, symbol
	}
	return "", 0, 0, 0, ""
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
