package main

import (
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

var products []Product

func generateProductID() int {
	return len(products) + 1
}

func Task__() {

	var inputMenu, inputItemName string
	var inputPrice, inputStock int

	products = append(products, Product{
		ID:    generateProductID(),
		Name:  "Minyak Goreng 2L",
		Stock: 100,
		Price: 30000,
	})

	fmt.Print("=== Toko Kelontong | INVENTORY MANAGER ===\n\n")

	menus := []string{"Tambah Barang", "Lihat Semua Stok", "Beli", "Keluar"}

	for endless := 0; endless <= 1; endless-- {

		if endless != 0 {
			fmt.Print("\n\n\n\n")
		}

		for index, menu := range menus {
			fmt.Println("[", index+1, "]", menu)
		}

		fmt.Print("\nPilih Menu (1-3): ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case "1":
			fmt.Print("\nMasukan Nama Barang: ")
			fmt.Scanln(&inputItemName)

			fmt.Print("Masukan Harga Barang: ")
			fmt.Scanln(&inputPrice)

			fmt.Print("Masukan Stock Barang: ")
			fmt.Scanln(&inputStock)

			product := Product{
				ID:    generateProductID(),
				Name:  inputItemName,
				Price: inputPrice,
				Stock: inputStock,
			}

			products = append(products, product)

			print("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang!")

		case "2":
			fmt.Printf("=== Daftar Stock Barang ===\n\n%-5s | %-20s | %-16s | %-5s\n", "ID", "Nama", "Harga", "Stok")
			for _, product := range products {
				fmt.Printf("%-5d | %-20s | Rp. %-12d | %-5d\n", product.ID, product.Name, product.Price, product.Stock)
			}

			fmt.Println("\n=========\nTotal Barang:", len(products))

		case "3":

		case "4":
			return
		default:
		}
	}

}

func main() {
	Task__()
}
