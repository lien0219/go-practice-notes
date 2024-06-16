package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
问题：

编写一个函数；随机猜数游戏；随机生成一个1到100的整数，
有十次机会，如果第一次就猜中，提示“你真是个天才”，
如果2到3次猜中提示你很聪明，赶上我了，
如果4到9次猜中，提示一般般，
如果最后一次猜中提示可算猜对了，一次都没猜对提示再接再厉
*/

/*
思路：

使用rand.Seed来设置随机数生成器的种子，以确保每次运行程序时生成的随机数都不同。
接着，使用rand.Intn(100) + 1来生成一个1到100的随机数。
然后，使用一个for循环来允许用户最多猜测10次。在每次循环中，我们读取用户的输入，
检查它是否与随机数匹配，并相应地更新猜测次数。
根据用户的猜测次数和结果，我们使用switch语句来输出不同的提示信息。
如果用户在第一次就猜中了，就输出“你真是个天才！”；
如果在2到3次之间猜中，就输出“你很聪明，赶上我了！”；
如果在4到9次之间猜中，就输出“一般般。”；
如果在最后一次猜中，就输出“可算猜对了！”。
如果10次都没猜中，就输出“再接再厉！”。
*/
func guessNumberGame() {
	//	初始化随机数
	rand.Seed(time.Now().UnixNano())
	//	1-100随机数
	numberToGuess := rand.Intn(100) + 1
	//	猜测次数
	guesses := 0

	//	最多猜十次
	for guesses < 10 {
		fmt.Println("请输入你猜的数字（1-100）：")
		var guess int
		fmt.Scanln(&guess)
		guesses++
		if guess == numberToGuess {
			switch guesses {
			case 1:
				fmt.Println("你真是天才！")
			case 2, 3:
				fmt.Println("你很聪明！")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("一般般！")
			case 10:
				fmt.Println("再接再厉！")
			}
			return
		} else if guess < numberToGuess {
			fmt.Println("猜小了，再试一次")
		} else {
			fmt.Println("猜大了，再试一次")
		}
	}
	//	十次都不中
	fmt.Println("很遗憾，没机会了~")
}

func main() {
	guessNumberGame()
}
