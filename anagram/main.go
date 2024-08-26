
// "abbad"  "slice" 100
// [2,2,0,1,...1,..]
// {"a" -> 1,"b"->2}
// [1,2]

// // 0 1 2 3
// // a b c d
// [1,0,0,1]

// [a,b,b,b] 100

package main

import (
	"fmt"
	"reflect"
)

func Anagram(a string, b string) bool {

	A := [26]int{}
	B := [26]int{}
	for _, char := range a {
		if char >= 'a' && char <= 'z' {
			A[char-'a']++
		}

	}
	for _, char := range b {
		if char >= 'a' && char <= 'z' {
			B[char-'a']++
		}
	}
	return reflect.DeepEqual(A, B)
	// for i := 0; i < 26; i++ {
	// 	if A[i] != B[i] {
	// 		return false
	// 	}
	// }
	// return true
}

func main() {

	fmt.Println(Anagram("hello", "heblo"))

}
