//channel
//宣言,操作

func main() {
	var ch1 chan int

	ch1 = make(chan int)

	fmt.Println(cap(ch1))

	var ch2 chan bool
	ch2 = make(chan bool)

	fmt.Println(cap(ch2))

	// バッファサイズ3のチャネルを作成する例
	ch3 := make(chan int, 3)

	ch3 <- 1
	fmt.Println(len(ch3))
	i := <-ch3
	fmt.Println(i)
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	fmt.Println(len(ch3))

	fmt.Println(<-ch3)
	fmt.Println(len(ch3))

	fmt.Println(<-ch3)
	fmt.Println(len(ch3))

	fmt.Println(<-ch3)
	fmt.Println(len(ch3))
}

//バッファサイズ
//デットロック

