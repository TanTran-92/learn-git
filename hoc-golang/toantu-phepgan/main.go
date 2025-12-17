package main

import "fmt"

func main() {
	s1 := 15
	s2 := 15	

	phepsosanh := s1 >= s2
	fmt.Printf("Ket qua phep so sanh: %t \n", phepsosanh)

	phepluanly := !((5<4) && (5>3) && (5>2))
	fmt.Printf("Ket qua phep luan ly: %t \n", phepluanly)

	// tong := s1 + s2
	// hieu := s1 - s2
	// tich := s1 * s2
	// thuong := float32(s1) / float32(s2)
	// chialaydu := s1 % s2
	

	// fmt.Printf("tong cua %d va %d la %d \n", s1, s2, tong)
	// fmt.Printf("hieu cua %d va %d la %d \n", s1, s2, hieu)
	// fmt.Printf("tich cua %d va %d la %d \n", s1, s2, tich)
	// fmt.Printf("thuong cua %d va %d la %.2f \n", s1, s2, thuong)
	// fmt.Printf("thuong cua %d va %d la %.d \n", s1, s2, chialaydu)
	// s3 := 15
	// s3 %= 4
	// fmt.Printf("s3 %.d \n", s3)
}
