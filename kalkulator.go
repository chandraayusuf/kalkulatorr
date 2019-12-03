package main

import (
	"fmt"
)

func repeatHelper() {
	fmt.Println("Tekan huruf Y jika anda ingin memulai Perhitungan")
	fmt.Println("Tekan huruf T jika Anda ingin keluar dari apliaksi")
	fmt.Println("Apakah Anda ingin melakukan perhitungan??")
}
func outputHelper() {
	fmt.Println("Perhitungan dengan format num1 spasi oparasi spasi num2")
	fmt.Println("Operasi Tersedia = +,-,*,/,^,%")
	fmt.Println("Ket : ^ operasi pangkat sedangakan % operasi hasil bagi")
	fmt.Println("Contoh = 1 + 1")
}

func kalkulator() {
	var num1 float64
	var num2 float64
	var result float64
	var operasi string
	var tempArray[] string
	
	num1 = 0
	num2 = 0
	result = 0
	operasi = "+"
	
	
	
