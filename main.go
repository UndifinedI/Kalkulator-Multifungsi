package main

import (
	"fmt"
)

func Format() {
balik:
	var opf string
	var a int
	var b string

	for true {
		fmt.Print("\nMasukan Operator Format  \n")
		fmt.Scan(&opf)

		fmt.Print("Masukan Angka  \n")
		fmt.Scan(&a)

		switch opf {
		case "o":
			fmt.Printf("Hasil: %o", a)

		case "x":
			fmt.Printf("Hasil: %x", a)

		case "b":
			fmt.Printf("Hasil: %b", a)

		default:
			goto balik
		}

		fmt.Print("\nMulai Lagi? y/n  \n")
		fmt.Scan(&b)

		if b == "y" {
			Format()
		} else if b == "n" {
			main()
			break
		}
	}

}



func main() {
balik:
// 	variabel baru bro
	var tes string
	
	var a int
	var b int
	var op string

	for true {

		fmt.Print("\nMasukan Operator  \n")
		fmt.Scan(&op)

		if op == "end" {
			break
		} else if op == "f" {
			Format()
		}
		fmt.Print("Masukan Angka Pertama  \n")
		fmt.Scan(&a)

		fmt.Print("Masukan Angka Kedua  \n")
		fmt.Scan(&b)

		switch op {
		case "+":
			fmt.Print("Hasil: ", a+b, "\n")

		case "-":
			fmt.Print("Hasil: ", a-b, "\n")

		case "*":
			fmt.Print("Hasil: ", a*b, "\n")

		case "/":
			fmt.Print("Hasil: ", a/b, "\n")

		default:
			goto balik
		}
	}
}

//NO nyontek, progaymer tidak sombong
//made by panji and jundia
