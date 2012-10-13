package main

import (
    "time"
    "fmt"
)

// 入力を受け付ける
func input() <-chan string {
    done := make(chan string)

    go func() {
        var str string
        fmt.Print(">")
        fmt.Scanf("%s", &str)
        done <- str
    }()

    time.Sleep(2 * time.Second)

    return done
}

func timeout(d time.Duration) <-chan bool {
    done := make(chan bool)

    go func() {
        time.Sleep(d)
        done <- true
    }()

    return done
}

func main() {
    select {
    case str := <-input():
        fmt.Println("input is", str)
    case <-timeout(5 * time.Second):
        fmt.Println()
        fmt.Println("time out!")
    }
}
