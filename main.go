package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/truonglocbinh/packer/numbers"
	"github.com/truonglocbinh/packer/strings"
	"github.com/truonglocbinh/packer/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("truonglocbinh"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}