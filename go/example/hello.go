package main

import "fmt"
import "net/http"

func main() {
	links := []string {
		"https://google.com",
		"https://instagram.com",
		"https://qq.com",
		"https://tiktok.com",
	}

	c := make(chan string)

	for _, link := range links {
		go check(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf(link, "connect err..")
		c <- "Oh no connect err.."
		return;
	}

	fmt.Printf(link, "connected!")
	c <- "Oh yes connected!"
}