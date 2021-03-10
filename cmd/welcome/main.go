package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	code := flag.Arg(1)
	message := `Добро пожаловать, {name}! Ваш код доступа: {code}.`
	message = strings.ReplaceAll(message, "{name}",name)
	message = strings.ReplaceAll(message,"{code}",code)
	fmt.Println(message)
}
