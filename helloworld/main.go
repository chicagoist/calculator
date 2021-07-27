package main

import calculator "github.com/chicagoist/calculator/study_package_golang"

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
}
