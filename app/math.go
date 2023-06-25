package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"wechat/common"
	"wechat/service"
)

func GetMathItemList(c *gin.Context) {
	uaText := c.Request.Header.Get("User-Agent")
	isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	if !isFlag {
		common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
		return
	}
	count, _ := strconv.Atoi(c.Query("count"))
	max, _ := strconv.Atoi(c.Query("max"))
	//op 1加法，2减法，3混合
	op, _ := strconv.Atoi(c.Query("op"))

	var service service.MathService

	list := service.GetMathItemList(op, count, max, 1)

	common.ReturnResponse(common.SUCCESS, list, common.SUCCESS_MSG, c)
}

func GetMathLottoList(c *gin.Context) {
	//uaText := c.Request.Header.Get("User-Agent")
	//isFlag := strings.Contains(strings.ToLower(uaText), "micromessenger")
	//if !isFlag {
	//	common.ReturnResponse(common.FORBID, map[string]interface{}{}, common.FORBID_MSG, c)
	//	return
	//}

	//1 ssq 2dlt
	op, _ := strconv.Atoi(c.Query("type"))
	if op < 1 {
		op = 1
	}

	red, _ := strconv.Atoi(c.Query("red"))
	blue, _ := strconv.Atoi(c.Query("blue"))
	if op == 1 {
		if red < 1 {
			red = 6
		}
		if blue < 1 {
			blue = 1
		}
	}else{
		if red < 1 {
			red = 5
		}
		if blue < 1 {
			blue = 2
		}
	}

	redList := c.Query("red_list")
	blueList := c.Query("blue_list")

	fmt.Println("类型:" , op)
	fmt.Println("红球:" , red)
	fmt.Println("红球集合:" , redList)
	fmt.Println("篮球:" , blue)
	fmt.Println("篮球集合:" , blueList)
	var service service.MathService
	str := service.GetMathLottoList(op, red, blue, redList, blueList)

	common.ReturnResponse(common.SUCCESS, str, common.SUCCESS_MSG, c)
}
