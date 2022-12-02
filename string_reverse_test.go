package main

import (
	"fmt"
	"testing"
)

func TestStringReverse(t *testing.T) {
	s0 := ""
	fmt.Println(reverseStrRecursion(s0))

	s1a := "ㄅ"
	fmt.Println(reverseStrRecursion(s1a))

	s1b := "b"
	fmt.Println(reverseStrRecursion(s1b))

	s1c := "?"
	fmt.Println(reverseStrRecursion(s1c))

	s2a := "Hello"
	fmt.Println(reverseStrRecursion(s2a))

	s2b := "?????????????"
	fmt.Println(reverseStrRecursion(s2b))

	s3 := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println(reverseStrRecursion(s3))

	s4 := "明里つむぎ 佳苗るか 岬ななみ"
	fmt.Println(reverseStrRecursion(s4))

}
