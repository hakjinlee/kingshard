package main

import "regexp"

func main() {

	re, err := regexp.Compile("list_.*_subscriber")
	if err != nil {
		panic(err)
	}

	if re.MatchString("list_123_subscriber") {
		println("Matched")
	} else {
		println("Not matched")
	}
}
