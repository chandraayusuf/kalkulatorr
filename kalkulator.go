package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64
	var num2 float64
	var operasi string
	var result float64
	var tanya string
	tanya = "y"

	fmt.Println("Jika Ingin mulai Menghitung Tekan Y")
	fmt.Println("Jika Tidak Ingin mulai Tekan T")

	for tanya == "Y" || tanya == "y" {
		fmt.Println("Apakah Anda Ingin Menghitung ?")
		fmt.Scanln(&tanya)
		if tanya == "y" || tanya == "Y" {
			fmt.Println("Masukkan Nilai Pertama")
			fmt.Scanln(&num1)

			fmt.Println("Masukkan operasi")
			fmt.Scanln(&operasi)

			fmt.Println("Masukkan Nilai Kedua")
			fmt.Scanln(&num2)

			if operasi == "^" {
				result = math.Pow(num1, num2)
			} else if operasi == "%" {
				result = math.Mod(num1, num2)
			} else if operasi == "+" {
				result = num1 + num2
			} else if operasi == "-" {
				result = num1 - num2
			} else if operasi == "*" {
				result = num1 * num2
			} else if operasi == "/" {
				result = num1 / num2
			}
			fmt.Println("Hasil Perhitungan", result)

		}
	}
}
