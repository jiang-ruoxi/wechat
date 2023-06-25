package service

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MathService struct {
}

type Set struct {
	m map[string]bool
}

func (ms *MathService) GetMathItemList(op, num, max, defaultLine int) (list []map[string]interface{}) {
	var count int
	var sets Set = NewSet()
	for {
		tem1 := ms.getRandomWithAll(1, max)
		tem2 := ms.getRandomWithAll(1, max)
		temStr := string(tem1) + string(tem2)
		if sets.Contains(temStr) {
			continue
		} else {
			sets.Add(temStr)
		}
		if tem1 > tem2 {
			oper := ms.getRandomWithAll(0, 1)
			if op == 3 {
				if oper == 0 {
					if tem1-tem2 <= int64(max) {
						fmt.Printf("%d - %d = %d\t\t", tem1, tem2, tem1-tem2)
						info := make(map[string]interface{})
						info["first"] = tem1
						info["second"] = tem2
						info["third"] = tem1 - tem2
						info["all"] = fmt.Sprintf("%d - %d = %d\t\t", tem1, tem2, tem1-tem2)
						info["o"] = fmt.Sprintf("%d - %d = ______\t\t", tem1, tem2)
						info["t"] = fmt.Sprintf("%d - %d\t\t", tem1, tem2)
						list = append(list, info)
						count = ms.calCount(count, defaultLine)
					}
				} else {
					if tem1+tem2 <= int64(max) {
						fmt.Printf("%d + %d = %d \t\t", tem1, tem2, tem1+tem2)
						info := make(map[string]interface{})
						info["first"] = tem1
						info["second"] = tem2
						info["third"] = tem1 + tem2
						info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2, tem1+tem2)
						info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
						info["t"] = fmt.Sprintf("%d + %d\t\t", tem1, tem2)
						list = append(list, info)
						count++
						if count%defaultLine == 0 {
							fmt.Println()
						}
					}
				}
			} else if op == 2 {
				if tem1-tem2 > int64(max) {
					continue
				}
				if tem1-tem2 <= int64(max) {
					fmt.Printf("%d - %d = %d \t\t", tem1, tem2, tem1-tem2)
					info := make(map[string]interface{})
					info["first"] = tem1
					info["second"] = tem2
					info["third"] = tem1 - tem2
					info["all"] = fmt.Sprintf("%d - %d = %d\t\t", tem1, tem2, tem1-tem2)
					info["o"] = fmt.Sprintf("%d - %d = ______\t\t", tem1, tem2)
					info["t"] = fmt.Sprintf("%d - %d\t\t", tem1, tem2)
					list = append(list, info)
					count = ms.calCount(count, defaultLine)
				}
			} else {
				if tem1+tem2 < int64(max) {
					continue
				}
				if tem1+tem2 <= int64(max) {
					fmt.Printf("%d + %d = %d \t\t", tem1, tem2, tem1+tem2)
					info := make(map[string]interface{})
					info["first"] = tem1
					info["second"] = tem2
					info["third"] = tem1 + tem2
					info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2, tem1+tem2)
					info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
					info["t"] = fmt.Sprintf("%d + %d\t\t", tem1, tem2)
					list = append(list, info)
					count++
					if count%defaultLine == 0 {
						fmt.Println()
					}
				}
			}

		} else if tem1 < tem2 {
			if op == 1 {
				if tem1+tem2 <= int64(max) {
					fmt.Printf("%d + %d = %d\t\t", tem1, tem2, tem1+tem2)
					info := make(map[string]interface{})
					info["first"] = tem1
					info["second"] = tem2
					info["third"] = tem1 + tem2
					info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2, tem1+tem2)
					info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
					info["t"] = fmt.Sprintf("%d + %d\t\t", tem1, tem2)
					list = append(list, info)
					count++
					if count%defaultLine == 0 {
						fmt.Println()
					}
				}
			}
		}
		if count >= num {
			break
		}
	}
	return list
}

func (ms *MathService) calCount(count int, num4 int) int {
	count++
	if count%num4 == 0 {
		fmt.Println()
	}
	return count
}

func (ms *MathService) getRandomWithAll(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}

func NewSet() Set {
	m := make(map[string]bool)
	return Set{m: m}
}

func (s *Set) Contains(val string) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set) Add(val string) {
	s.m[val] = true
}

func (s *Set) Remove(val string) {
	delete(s.m, val)
}

func (ms *MathService) GetMathLottoList(op int, indexRed, indexBlue int, redList string, blueList string) (all string) {
	//1 ssq 2dlt
	redListArr := strings.Split(redList, ",")
	blueListArr := strings.Split(blueList, ",")

	var redArray, blueArray []int
	for _, item := range redListArr {
		intItem, _ := strconv.Atoi(item)
		redArray = append(redArray, intItem)
	}
	for _, item := range blueListArr {
		intItem, _ := strconv.Atoi(item)
		blueArray = append(blueArray, intItem)
	}
	//
	if op == 1 {
		if len(redArray) <= 1 {
			redArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33}
		}
		if len(blueArray) <= 1 {
			blueArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		}
	} else {
		if len(redArray) <= 1 {
			redArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35}
		}
		if len(blueArray) <= 1 {
			blueArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		}
	}
	fmt.Println(redArray, blueArray)
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
	// 随机生成 n 个不重复的下标
	redListR := rand.Perm(len(redArray))[:indexRed]
	sort.Slice(redListR, func(i, j int) bool {
		return redListR[i] < redListR[j]
	})

	rand.Seed(time.Now().UnixNano())
	// 随机生成 n 个不重复的下标
	blueListR := rand.Perm(len(blueArray))[:indexBlue]
	sort.Slice(blueListR, func(i, j int) bool {
		return blueListR[i] < blueListR[j]
	})

	fmt.Println(redListR, blueListR)
	redListStr := strings.Trim(strings.Replace(fmt.Sprint(redListR), " ", " ", -1), "[]")
	blueListStr := strings.Trim(strings.Replace(fmt.Sprint(blueListR), " ", "、", -1), "[]")
	all = fmt.Sprintf("%s + %s", redListStr, blueListStr)

	return all
}