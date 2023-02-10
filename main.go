package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) { // ... 은 인자를 여러개 받을 때 사용
	fmt.Println(words)
}

func lenAndUpper2(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") //defer return 다음 실행
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return //naked return : return 변수를 미리 정의하면 return만 적으면 됨
}

func superAdd(numbers ...int) int {
	for index, number := range numbers { //for index, number in enumerate(numbers)같은 기능
		fmt.Println(number)
		fmt.Println(index, number)
	}
	return 1
}

func main() {
	const name string = "choi"
	fmt.Println(name)
	var varname string = "sanghyun"
	fmt.Println(varname)
	name2 := "tom" //:= 문법 사용시 변수 타입 자동으로 설정됨.
	fmt.Println(name2)
	fmt.Println(multiply(2, 2))
	totalLenght, upperName := lenAndUpper("nico")
	fmt.Println(totalLenght, upperName)
	repeatMe("choi", "sang", "hyun")
	fmt.Println(lenAndUpper2("choi"))
	superAdd(1, 2, 3, 4, 5)

}
