package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RedList6 [6]int
type ListMap struct {
	Data  RedList6
	Count int
}

func main() {
	var sortListMap []ListMap
	map6 := make(map[[6]int]int)
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < 999999; n++ {
		fmt.Println(n)
		var list6 RedList6
		for i := 0; i < 6; i++ {
			x := rand.Intn(33) + 1
			for k := range list6 {
				if x == list6[k] {
					i = i - 1
					break
				}
			}
			list6[i] = x
		}
		//排序
		for i := 0; i < len(list6); i++ {
			for j := i + 1; j < len(list6); j++ {
				if list6[i] > list6[j] {
					tmp := list6[i]
					list6[i] = list6[j]
					list6[j] = tmp
				}
			}
		}
		if _, ok := map6[list6]; ok {
			map6[list6] = map6[list6] + 1
		} else {
			map6[list6] = 1
		}
	}
	for k, v := range map6 {
		if v >= 9 {
			var tmpListMap ListMap
			tmpListMap.Data = k
			tmpListMap.Count = v
			sortListMap = append(sortListMap, tmpListMap)
		}
	}
	for i := 0; i < len(sortListMap); i++ {
		for j := i + 1; j < len(sortListMap); j++ {
			if sortListMap[i].Count > sortListMap[j].Count {
				tmp := sortListMap[i]
				sortListMap[i] = sortListMap[j]
				sortListMap[j] = tmp
			}
		}
	}
	for k := range sortListMap {
		fmt.Println(sortListMap[k])
	}

}

// type ListBlue struct {
// 	Data []List2 `json:"data"`
// }

// type List2 [2]int

// func main() {
// 	tmpMap := make(map[int]int)
// 	tmpBlueMap := make(map[int]int)
// 	var tmpBlueList ListBlue
// 	var tmpRedList ListBlue
// 	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
// 	for i := 0; i < 99999; i++ {
// 		x := rand.Intn(33) + 1
// 		//fmt.Println(x) //返回[0,100)的随机整数
// 		if _, ok := tmpMap[x]; ok {
// 			tmpMap[x] = tmpMap[x] + 1
// 		} else {
// 			tmpMap[x] = 1
// 		}
// 	}
// 	for i := 0; i < 99999; i++ {
// 		x := rand.Intn(16) + 1
// 		//fmt.Println(x) //返回[0,100)的随机整数
// 		if _, ok := tmpBlueMap[x]; ok {
// 			tmpBlueMap[x] = tmpBlueMap[x] + 1
// 		} else {
// 			tmpBlueMap[x] = 1
// 		}
// 	}
// 	for k, v := range tmpMap {
// 		fmt.Println(k, v)
// 		var tmp List2
// 		tmp[0] = k
// 		tmp[1] = v
// 		tmpRedList.Data = append(tmpRedList.Data, tmp)
// 	}
// 	for i := 0; i < len(tmpRedList.Data); i++ {
// 		for j := i + 1; j < len(tmpRedList.Data); j++ {
// 			if tmpRedList.Data[i][1] < tmpRedList.Data[j][1] {
// 				tmp := tmpRedList.Data[i][1]
// 				tmpRedList.Data[i][1] = tmpRedList.Data[j][1]
// 				tmpRedList.Data[j][1] = tmp
// 			}
// 		}
// 	}
// 	for k, v := range tmpRedList.Data {
// 		fmt.Println(k, v)
// 	}
// 	fmt.Println("--------------------------------------------")

// 	for k, v := range tmpBlueMap {
// 		fmt.Println(k, v)
// 		var tmp List2
// 		tmp[0] = k
// 		tmp[1] = v
// 		tmpBlueList.Data = append(tmpBlueList.Data, tmp)
// 	}
// 	for i := 0; i < len(tmpBlueList.Data); i++ {
// 		for j := i + 1; j < len(tmpBlueList.Data); j++ {
// 			if tmpBlueList.Data[i][1] < tmpBlueList.Data[j][1] {
// 				tmp := tmpBlueList.Data[i][1]
// 				tmpBlueList.Data[i][1] = tmpBlueList.Data[j][1]
// 				tmpBlueList.Data[j][1] = tmp
// 			}
// 		}
// 	}
// 	for k, v := range tmpBlueList.Data {
// 		fmt.Println(k, v)
// 	}
// 	var sortList []int
// 	for i := 0; i < 6; i++ {
// 		// fmt.Println("red ", i+1, ":", tmpRedList.Data[i][0])
// 		sortList = append(sortList, tmpRedList.Data[i][0])
// 	}
// 	for i := 0; i < len(sortList); i++ {
// 		for j := i + 1; j < len(sortList); j++ {
// 			if sortList[i] > sortList[j] {
// 				tmp := sortList[i]
// 				sortList[i] = sortList[j]
// 				sortList[j] = tmp
// 			}
// 		}
// 	}
// 	for k := range sortList {
// 		fmt.Println("red ", k+1, ":", sortList[k])
// 	}
// 	fmt.Println("Blue:", tmpBlueList.Data[0][0])
// 	// time.Sleep(10 * time.Second)
// }
