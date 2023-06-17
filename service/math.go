package service

import (
	"fmt"
	"math/rand"
	"time"
)

type MathService struct {
}

type Set struct {
	m map[string]bool
}

func (ms *MathService) GetMathItemList(op, num, max, defaultLine int) (list []map[string]interface{}){
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
						fmt.Printf("%d - %d = %d\t\t", tem1, tem2,tem1-tem2)
						info := make(map[string]interface{})
						info["first"] = tem1
						info["second"] = tem2
						info["third"] = tem1 - tem2
						info["all"] = fmt.Sprintf("%d - %d = %d\t\t", tem1, tem2,tem1-tem2)
						info["o"] = fmt.Sprintf("%d - %d = ______\t\t", tem1, tem2)
						list = append(list, info)
						count = ms.calCount(count, defaultLine)
					}
				} else {
					if tem1+tem2 <= int64(max) {
						fmt.Printf("%d + %d = %d \t\t", tem1, tem2,tem1+tem2)
						info := make(map[string]interface{})
						info["first"] = tem1
						info["second"] = tem2
						info["third"] = tem1 + tem2
						info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2,tem1+tem2)
						info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
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
					info["all"] = fmt.Sprintf("%d - %d = %d\t\t", tem1, tem2,tem1-tem2)
					info["o"] = fmt.Sprintf("%d - %d = ______\t\t", tem1, tem2)
					list = append(list, info)
					count = ms.calCount(count, defaultLine)
				}
			}else{
				if tem1+tem2 < int64(max) {
					continue
				}
				if tem1+tem2 <= int64(max) {
					fmt.Printf("%d + %d = %d \t\t", tem1, tem2,tem1+tem2)
					info := make(map[string]interface{})
					info["first"] = tem1
					info["second"] = tem2
					info["third"] = tem1 + tem2
					info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2,tem1+tem2)
					info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
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
					fmt.Printf("%d + %d = %d\t\t", tem1, tem2,tem1+tem2)
					info := make(map[string]interface{})
					info["first"] = tem1
					info["second"] = tem2
					info["third"] = tem1 + tem2
					info["all"] = fmt.Sprintf("%d + %d = %d\t\t", tem1, tem2,tem1+tem2)
					info["o"] = fmt.Sprintf("%d + %d = ______\t\t", tem1, tem2)
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
