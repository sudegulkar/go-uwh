package main

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			print("A")
		} else {
			print("B")
		}
	}
	println()
	for j := 1; j < y-1; j++ {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				print("B")
			} else {
				print(" ")
			}
		}
		println()
	}
	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				print("C")
			} else {
				print("B")
			}
		}
		println()
	}
}
func main() {
	QuadC(5, 3)
}
