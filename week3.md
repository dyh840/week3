## week3

#### task1

切片会越界，设置足够长切片

多个goroutine修改consumeMSG，用复制的副本修改

#### task2

```
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(20)

	ch := make(chan struct {
		a int
		b int
	}, 20)

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wait.Done()
			c := rand.Intn(100)
			ch <- struct {
				a int
				b int
			}{i, c}
		}(i)
	}

	go func() {
		wait.Wait()
		close(ch)
	}()

	Map := make(map[int]int)
	for i := range ch {
		Map[i.a] = i.b
	}

	fmt.Println("未排序：")
	for i, b := range Map {
		fmt.Printf("编号：%d 随机数：%d\n", i, b)
	}

	fmt.Println("\n排序后：")
	for i := 0; i < 20; i++ {
		fmt.Printf("编号：%d 随机数：%d\n", i, Map[i])
	}
}

```



#### task3

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := "0123456789"
	ch := make(chan string)
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		defer wait.Done()
		la, lb := 0, 0
		for la < len(a) || lb < len(b) {
			if la+2 <= len(a) {
				ch <- a[la : la+2]
				la += 2
			} else if la < len(a) {
				ch <- a[la:]
				la = len(a)
			}

			if lb+2 <= len(b) {
				ch <- b[lb : lb+2]
				lb += 2
			} else if lb < len(b) {
				ch <- b[lb:]
				lb = len(b)
			}
		}
		close(ch)
	}()
	wait.Add(1)
	go func() {
		defer wait.Done()
		for v := range ch {
			fmt.Print(v)
		}
	}()
	wait.Wait()
}

```

