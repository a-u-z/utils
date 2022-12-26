package main

import "log"

func sliceDetail() {
	// 當 range slice 的時候
	aa := []SliceDetail{
		{SA: "sasa"},
		{SA: "BBBB"},
	}

	// 這樣做會改不到 SB 的值，因為是對 v 這個值改動，而 v 又是從 aa 身上複製出來的
	for _, v := range aa {
		v.SB = "CCCC"
	}
	log.Printf("here is aa:%+v", aa)

	// 要對 aa 本身做修改才可以
	for i := range aa {
		aa[i].SB = "cccccccccc"
	}
	log.Printf("here is aa:%+v", aa)
}

type SliceDetail struct {
	SA string
	SB string
}
