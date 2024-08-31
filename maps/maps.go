package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon web services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "Google")
}
