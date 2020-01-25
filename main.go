package main

import (
	"fmt"
	"strings"
)

var str string = "faathi"
var chaine string = "fe@thiii#"
var system string = "windows"
var category string ="junior"

func main() {
	str = strings.Map(func(r rune) rune {
		switch r {
		case 'a':
			return '1'
		case 'e':
			return '1'
		default:
			return r
		}
	}, str)
	fmt.Println("hello", str)

	chaine = strings.Map(func(el rune) rune {
		if (el > 'a' && el < 'z') || (el > 'A' && el < 'Z') {
			return el
		} else {
			return rune(0)
		}

	}, chaine)
	fmt.Println(chaine)
	if strings.EqualFold(system, "windows") {
		//if (system=="windows")
		fmt.Println("windows")
	} else if strings.EqualFold(system, "linux") {
		fmt.Println("linux")
	} else if strings.EqualFold(system, "ios") {
		fmt.Println("ios")
	} else {
		fmt.Println("other system")
	}
	switch category {
	case "j1","j2","j3":
		fmt.Println("junior")
	case "c1","c2","c3":
		fmt.Println("confirmed")
	case "s1","s2","s3":
		fmt.Println("senior")
	case "m1","m2","m3":
		fmt.Println("manager")


	default:
		fmt.Println(category)

	}

}
