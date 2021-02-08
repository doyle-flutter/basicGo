// 직접 패키지를 구성하여 사용할 수 있습니다

package mypack

// MyPack 함수를 만들고 외부에서 사용할 수 있도록 합니다
// 내보낼 함수는 대문자로 사용해야합니다
func MyPack() string {
	return "제임쓰의 패키지!"
}

// -> 소문자로 만든 함수는 export 가 되지 않습니다
func myPack2() string {
	return "제임쓰의 패키지!22"
}
