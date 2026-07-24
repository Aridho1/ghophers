package main

import (
	"fmt"
	"strconv"
)

func learnArray() {
	// array
	// declare array
	// case 1
	var arr1 [4]string

	arr1[0] = "Budi"
	arr1[1] = "Joko"
	arr1[2] = "Bendi"
	arr1[3] = "Bima"

	fmt.Println("arr1", arr1)

	// case 2
	arr2 := [2]string{"Pen", "Book"}
	fmt.Println("arr2", arr2)

	// slice | array but have no strict length
	var arr3 []string
	arr3 = append(arr3, "hello")
	fmt.Println("arr3", arr3)

	// slice but another declare case
	arr4 := []string{"Table", "Chair"}
	fmt.Println("arr4", arr4)
}

// type Car {
// 	Color: string
// }

func learnLoop() {
	// using for
	for i := 1; i < 10; i++ {
		fmt.Println("Halo", i)
	}
}

func learnRangeLoop() {
	arr1 := []string{"Semangka", "Apel", "Sirsak", "Mangga"}

	// range
	// case 1, using traditional for loop
	for i := range arr1 {
		fmt.Println(i, arr1[i])
	}

	// case 2, using range loop
	for index, item := range arr1 {
		fmt.Println(index, item)
	}
}

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func Task__() {

	var inputMenu, inputItemName, inputPrice, inputStock string
	products := []Product{}

	fmt.Print("=== Toko Kelontong | INVENTORY MANAGER\n\n")

	menus := []string{"Tambah Barang", "Lihat Semua Stok", "Keluar"}

	for endless := 0; endless <= 1; endless-- {

		fmt.Print("\n\n\n\n")
		for index, menu := range menus {
			fmt.Println("[", index+1, "]", menu)
		}

		fmt.Print("Pilih Menu (1-3): ")
		fmt.Scanln(&inputMenu)

		if inputMenu == "1" {
			fmt.Print("\nMasukan Nama Barang: ")
			fmt.Scanln(&inputItemName)

			fmt.Print("Masukan Harga Barang: ")
			fmt.Scanln(&inputPrice)

			fmt.Print("Masukan Stock Barang: ")
			fmt.Scanln(&inputStock)

			price, _ := strconv.Atoi(inputPrice)
			stock, _ := strconv.Atoi(inputStock)

			product := Product{
				ID:    len(products) + 1,
				Name:  inputItemName,
				Price: price,
				Stock: stock,
			}

			products = append(products, product)

			print("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang!")

		} else if inputMenu == "2" {
			fmt.Print("=== Daftar Stock Barang ===\n\n")
			for _, product := range products {
				print("ID: ", product.ID, "|", "Nama:", product.Name, "|", "Harga: Rp", product.Price, "|", "stock:", product.Stock, "\n")
			}

			fmt.Println("\n=========\nTotal Barang:", len(products))
		} else if inputMenu == "3" {
			return
		} else {
		}
	}

}

func main() {
	// learnArray()
	// learnLoop()
	// learnRangeLoop()
	Task__()
}
