package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 寫一猜拳遊戲:
// 函式一(from_pc): 電腦出拳 - 隨機產生0~2之間數字, 0=剪刀, 1=石頭, 2=布，並回傳
// 函式二(from_player): 玩家出拳 - 接收使用者輸入，, 0=剪刀, 1=石頭, 2=布，並外傳，除外的數字會出現「錯誤，請重新輸入」
// 函式三(who_win): 輸贏判斷 - 將前面兩個函式的結果傳入，並進行輸贏判斷，如果電腦贏，回傳0, 平手回傳1，玩家贏回傳2
// main函式: 根據回傳值顯示結果：0=您輸了，1=平手，2=您贏了
//
// 備註:
//   1. 程式中請勿使用魔術數字
const SCISSORS = 0
const STONE = 1
const PAPER = 2

const RESULT_PC_WIN = 0
const RESULT_DRAW = 1
const RESULT_PLAYER_WIN = 2

var SHOW []string = []string{"剪刀", "石頭", "布"}

func count_array(a []int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adder2(x int) int {
	sum := 0
	sum += x
	return sum
}

func main() {
	/*
		fmt.Println("===========")
		fmt.Println("= 猜拳遊戲 =")
		fmt.Println("===========")

		pc := from_pc()         // 電腦出拳
		player := from_player() // 玩家出拳
		result := who_win(pc, player)

		fmt.Printf("電腦出: %s\n", SHOW[pc])
		fmt.Printf("您出了: %s\n", SHOW[player])

		switch result {
		case RESULT_PC_WIN:
			fmt.Println("遊戲結果: 電腦贏了.")
		case RESULT_PLAYER_WIN:
			fmt.Println("遊戲結果: 您贏了.")
		default:
			fmt.Println("遊戲結果: 平手.")
		}

		fmt.Println("")
	*/

	/*
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		sum := count_array(arr)
		fmt.Printf("總和為: %d\n", sum)
	*/
	/*
		sum_func := func(a int, b int) int {
			return a + b
		}

		sum := sum_func(1, 3)
		fmt.Println(sum)

		a := sum_func

		sum = a(5, 6)
		fmt.Println(sum)
	*/

	a := adder()
	b := adder()

	fmt.Printf("%d, %d\n", a(1), b(1))
	fmt.Printf("%d, %d\n", a(1), b(1))

	fmt.Printf("%d, %d\n", adder2(1), adder2(1))
	fmt.Printf("%d, %d\n", adder2(1), adder2(1))
}

func from_pc() int {
	rand.Seed(int64(time.Now().UnixNano()))
	pc := rand.Intn(2)

	return pc
}

func from_player() int {
	var i int

	for {
		fmt.Printf("請猜拳(0=%s、1=%s、2=%s): ", SHOW[0], SHOW[1], SHOW[2])
		fmt.Scanf("%d", &i)
		if i < 0 || i > 2 {
			fmt.Println("輸入錯誤，請重新輸入!")
		} else {
			break
		}
	}

	return i
}

func who_win(pc int, player int) int {
	if pc == SCISSORS && player == STONE {
		return RESULT_PLAYER_WIN
	} else if pc == SCISSORS && player == PAPER {
		return RESULT_PC_WIN
	} else if pc == STONE && player == SCISSORS {
		return RESULT_PC_WIN
	} else if pc == STONE && player == PAPER {
		return RESULT_PLAYER_WIN
	} else if pc == PAPER && player == SCISSORS {
		return RESULT_PLAYER_WIN
	} else if pc == PAPER && player == STONE {
		return RESULT_PC_WIN
	}

	return RESULT_DRAW
}

func range_ex() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}
}

// slice
func slice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(a[:])   // 全部
	fmt.Println(a[1:4]) // [2, 3, 4]
	fmt.Println(a[3:])  // [4, 5, 6, 7, 8, 9, 0]
	fmt.Println(a[:4])  // [1, 2, 3, 4]
	fmt.Printf("%v(len=%d, cap=%d)\n", a, len(a), cap(a))

	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := make([]int, 0, 50) // len(b)=0, cap(b)=5
	fmt.Printf("%v(len=%d, cap=%d)\n", b, len(b), cap(b))

	b = b[:50]
	fmt.Printf("%v(len=%d, cap=%d)\n", b, len(b), cap(b))
}

// 開發
func show_weekday() {
	weekday_str := []string{"日", "一", "二", "三", "四", "五", "六"}

	weekday := int(time.Now().Weekday())

	fmt.Println(weekday)
	fmt.Println("星期" + weekday_str[weekday])
}
