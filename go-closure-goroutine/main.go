package main

import "fmt"

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)

	values := []string{"a", "b", "c"}

	for _, v := range values {
		// normal loop
		fmt.Println(v)
	}

	for _, v := range values {
		go func() {
			// this will log out c c c instead of a b c
			fmt.Println(v)
			done1 <- true
		}()
	}

	for _ = range values {
		<-done1
	}

	for _, v := range values {
		go func(u string) {
			// this will log out c a b
			fmt.Println(u)
			done2 <- true
		}(v)
	}

	for _ = range values {
		<-done2
	}

	for _, v := range values {
		v := v
		go func() {
			// this will log out c a b too
			fmt.Println(v)
			done3 <- true
		}()
	}

	for _ = range values {
		<-done3
	}
}
