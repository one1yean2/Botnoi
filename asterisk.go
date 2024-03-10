package main

import (
	"fmt"
)

func main() {
	fmt.Println("Type a number")

	var n int
	fmt.Scanf("%d", &n)

	//วนลูปจำนวน n-1 ครั้ง เพื่อแสดงครึ่งบน เพิ่มขึ้นเรื่อยๆทีละบรรทัด
	for i := 0; i < n; i++ {
		//แสดง * ตามค่าของ i ในลูป
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		//ขึ้นบรรทัดใหม่
		fmt.Println()
	}
	//วนลูปจำนวน n ครั้ง เพื่อแสดง ตรงกลาง และ ครึ่งล่าง ลดลงเรื่อยๆทีละบรรทัด
	for i := n; i > 0; i-- {
		//แสดง * ตามค่าของ i ในลูป
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		//ขึ้นบรรทัดใหม่
		fmt.Println()
	}

	//ใช้คําสั่ง strings.Repeat เพื่อแสดง * ตามค่าของ i
	// for i := 0; i < n; i++ {
	// 	fmt.Println(strings.Repeat("*", i))
	// }
	// for i := n; i > 0; i-- {
	// 	fmt.Println(strings.Repeat("*", i))
	// }
}
