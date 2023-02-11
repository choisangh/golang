package main

import (
	"fmt"
	"strings"

	"github.com/choisangh/golang/banking/banking"
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

func superAdd(numbers ...int) int { //for 반복문
	for index, number := range numbers { //for index, number in enumerate(numbers)같은 기능
		fmt.Println(number)
		fmt.Println(index, number)
	}
	return 1
}

func canIDrink(age int) bool { //조건문
	if koreanAge := age + 2; koreanAge < 18 { //if 문 안에 변수 선언 가능(variable expression )
		return false
	}
	return true
}

func canIDrink2(age int) bool { //switch문 - if else문 간소화
	switch { //varialble expression 가능
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	return false
}

func pointers() {
	a := 2
	b := 5
	fmt.Println(&a, &b) //메모리 주소 보기
	c := &a
	fmt.Println(*c) //메모리 주소 참조 -> *c = a
}

func test_array() {
	names := [5]string{"choi", "sang", "hyun"}
	names[3] = "tom"
	names[4] = "tom"
	fmt.Println(names)
	slices := []string{"choi", "sang", "hyun"} //slices : 길이 가변적인 Array
	slices = append(slices, "hi")
	fmt.Println(slices)

}

func test_map() {
	choi := map[string]string{"name": "choi", "age": "29"} //[key type]valuetype
	// map의 key, value는 한 타입으로만 설정 가능
	fmt.Println(choi)
	for key, value := range choi {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func test_structs() {
	favFood := []string{"kimchi", "ramen"}
	choi := person{name: "choi", age: 29, favFood: favFood}
	fmt.Println(choi)
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
	pointers()
	test_array()
	test_structs()
	account := banking.Account{owner: "choi", balance: 100}

}
