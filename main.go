package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// func main() {
//     fmt.Println(Red + "Error!" + Reset)
//     fmt.Println(Green + "Success!" + Reset)
// }

func main() {
	// fmt.Println("hello world")

	// // txt, err := fmt.Scanln("test: ")
	// var txt string

	// fmt.Printf("Test: ")
	// fmt.Scanln(&txt)

	// fmt.Printf("hasil: %s", txt)

	var inputPrice, inputMoney string
	// var price, money int
	// var err Error

	fmt.Printf("\n\n========================\n\n")
	fmt.Printf("Harga Barang: ")
	fmt.Scanln(&inputPrice)

	price, err := strconv.Atoi(inputPrice)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Printf("Error, inputnya harus int, %s itu bukan int :D", inputPrice)
		} else {
			fmt.Println("Error: ", err)
		}

		main()
		return
	}

	fmt.Printf("Uang yang dimiliki: ")
	fmt.Scanln(&inputMoney)
	money, err := strconv.Atoi(inputMoney)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Printf(Red+"Error, inputnya harus int, %s itu bukan int :D"+Reset, inputMoney)
		} else {
			fmt.Println("Error: ", err)
		}

		main()
		return
	}

	var sisa int = money - price

	if sisa < 0 {
		fmt.Println("[SYSTEM]: Transaksi ditolak! Uang kurang ", sisa*-1)
		main()
		return
	}

	fmt.Println("[SYSTEM]: Transaksi berhasil. Kembalian anda ", sisa)

	fmt.Printf("\n\n========================\n\n")
}
