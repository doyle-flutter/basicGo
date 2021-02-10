package main

import (
	"fmt"
	"strconv"

	mypack "./myPack"
)

func main() {

	// * customPack : 직접 패키지를 구성할 수 있습니다
	// * path : ./myPack/myPack.go
	fmt.Println(mypack.MyPack())

	// - 변수
	// var 변수명 타입
	var name string = "Name"
	fmt.Println(name)
	// 타입이 동일할 경우 묶어서 지정가능 var a,b,c int = 1,2,3
	var a, b, c int = 1, 2, 3
	fmt.Println(strconv.Itoa(a + b + c))
	// 타입이 다른 경우
	var (
		a1 int    = 1
		b2 string = "B22"
	)
	fmt.Println(strconv.Itoa(a1) + b2)

	// - 상수
	// const 상수명 타입
	const myName string = "MYNAME"
	fmt.Println(myName)
	// 타입이 동일할 경우 묶어서 지정가능(변수와 동일)
	const aNumber, bNumber, cNumber int = 1, 2, 3
	fmt.Println(strconv.Itoa(aNumber + bNumber + cNumber))

	// - 타입 : 사용하는 타입을 추론하지만 최대한 입력하는 것을 권장
	// 문자열 : string
	// 정수 : int
	// 소수 : float32, float64
	// 배열 []int : 배열 안에 요소타입 정수( []string : 배열 안에 요소타입 문자 )
	// struct : 구조
	// 참거짓 : bool

	// 포인터 : & 주소 / * 참조
	var (
		pointerAd = &name
		pointerRf = *pointerAd
	)
	fmt.Println(fmt.Sprint("PONTER : ", pointerRf))

	// - 조건 및 반복문 : 소괄호를 사용하지 않음
	if a > 0 {
		fmt.Println("a는 0보다 크다")
	} else {
		fmt.Println("a는 0보다 작다")
	}

	for i := 0; i < a; i++ {
		var data string = fmt.Sprint("i : ", i)
		fmt.Println(data)
	}

	// - 함수 실행
	fmt.Println(dataf())                                   // --- 함수
	fmt.Println(dataf2("MSG", 123))                        // --- 인자
	fmt.Println(dataf3("ABC", "DEF"))                      // --- 가변 인자 1
	fmt.Println(dataf3("ABC", "DEF", "!@#", "$%^"))        // --- 가변 인자 2
	var aa, bb string = dataf3("LKJ", "POI", "!@#", "$%^") // --- 가변 인자 3
	fmt.Println(fmt.Sprint(aa, " ", bb))
	var arg *Dataf4Arg = new(Dataf4Arg) // --- struct 인자 활용
	arg.id = 12345678
	arg.name = "NAME"
	fmt.Println(fmt.Sprint(strconv.Itoa(dataf4(arg))))

	// --- struct 와 함수 활용
	var dataf4ArgConstructorClass1 *Dataf4Arg = dataf4ArgConstructor(123, "제임쓰")
	var dataf4ArgConstructorClass2 *Dataf4Arg = dataf4ArgConstructor(123, "제임쓰")
	fmt.Println(fmt.Sprint("(1) Constructor func 1 : ", dataf4ArgConstructorClass1.id))
	fmt.Println(fmt.Sprint("(1) Constructor func 2 : ", dataf4ArgConstructorClass2.id))
	dataf4ArgConstructorClass1.id = 321321
	dataf4ArgConstructorClass2.id = 890890
	fmt.Println(fmt.Sprint("(2) Constructor func 1 : ", dataf4ArgConstructorClass1.id))
	fmt.Println(fmt.Sprint("(2) Constructor func 2 : ", dataf4ArgConstructorClass2.id))

	// - 함수와 포인터
	var pointerData string = "data"
	fmt.Println(fmt.Sprint("(0) POINTDATA ? ", pointerData))
	changeData(&pointerData)
	fmt.Println(fmt.Sprint("(1) POINTDATA ? ", pointerData))
	// -- 포인터를 사용하지 않는 경우
	// : 인자로 받은 변수 안의 값이 변경되지 않음
	changeDataNonePointer(pointerData)
	fmt.Println(fmt.Sprint("(4) POINTDATA ? ", pointerData))

	// - 익명 함수
	var noName func() string = func() string {
		return "NONAME"
	}
	fmt.Println(noName())

	// - 익명 함수와 일급 함수 / 콜백
	var callBack = func(data1 int, data2 int) int {
		return (data1 + data2)
	}
	var wentCallBack = func(cb func(int, int) int) string {
		var a, b int = 10, 20
		return fmt.Sprint("CallBack : ", strconv.Itoa(cb(a, b)))
	}
	fmt.Println(wentCallBack(callBack))

	// - 배열
	// var 변수명[크기]타입
	// var 변수명 = [크기]타입{요소}
	// var 변수명 = [...]타입{요소} // 크기 동적
	var arr1 = [...]int{1, 2, 3}
	fmt.Println(arr1)

}

// - 함수 선언
// func 함수명(){}
func dataf() string {
	return "Hi?"
}

// -- 함수 인자
func dataf2(msg string, num int) string {
	return fmt.Sprint(msg, " - ", num)
}

// -- 가변 인자
func dataf3(data ...string) (string, string) {
	return data[0], data[1]
}

// -- 생성자 constructor
func dataf4ArgConstructor(id int, name string) *Dataf4Arg {
	var class *Dataf4Arg = new(Dataf4Arg)
	class.id = id
	class.name = name
	return class
}

// -- 네이밍 옵션 대신 struct 인자
type Dataf4Arg struct {
	id   int
	name string
}

func dataf4(arg *Dataf4Arg) int {
	return arg.id
}

// --- 함수와 포인터
func changeData(data *string) {
	*data = "CHANGE DATA"
	fmt.Println(fmt.Sprint("(2) POINTDATA ? ", *data))
	return
}
func changeDataNonePointer(data string) {
	data = "NONE:Pointer CHANGE DATA"
	fmt.Println(fmt.Sprint("*(3) NONE:POINTDATA ? ", data))
	return
}
