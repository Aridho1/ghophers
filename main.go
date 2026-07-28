package main

import (
	"fmt"
	"os"
	"time"
)

const (
	WAIT_PRIMARY = 1 * time.Second
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

type Menu struct {
	Name    string
	Handler MenuHandler
}

type MenuHandler func() error

type App struct {
	Title    string
	Menus    []Menu
	Products []Product
}

func (a *App) LogTitle() {
	fmt.Printf("=== %s | INVENTORY MANAGER ===\n\n", a.Title)
}

func (a *App) LogListMenu(menuLen int) {
	for i, menu := range a.Menus {
		num := i + 1

		if i == menuLen {
			num = 0
		}

		fmt.Printf("%-2d. %s\n", num, menu.Name)
	}
}

func (a *App) Run() {
	var inputMenu int

	menuLen := len(a.Menus)
	loop := 0

	for {
		loop += 1
		if loop != 1 {
			fmt.Print("\n\n\n\n")
			time.Sleep(WAIT_PRIMARY)
		}

		a.LogTitle()
		a.LogListMenu(menuLen)

		fmt.Printf("\nPilih Menu [0-%d]: ", menuLen-1)
		fmt.Scan(&inputMenu)

		time.Sleep(WAIT_PRIMARY)

		if inputMenu < 0 || inputMenu > menuLen-1 {
			fmt.Printf("Menu Tidak ditemukan!\n")
		} else if inputMenu == 0 {
			a.Menus[menuLen-1].Handler()
		} else {
			a.Menus[inputMenu-1].Handler()
		}

	}
}

func NewApp() *App {
	var app *App

	app = &App{
		Title: "Toko Kelontong",
		Products: []Product{
			{
				ID:    generateProductID(),
				Name:  "Minyak Goreng 2L",
				Stock: 100,
				Price: 30000,
			},
		},
		Menus: []Menu{
			{
				Name: "Tambah Barang",
				Handler: func() error {
					var inputItemName string
					var inputPrice, inputStock int

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

					app.Products = append(app.Products, product)

					print("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang!")

					return nil
				},
			}, {
				Name: "Lihat Semua Stock",
				Handler: func() error {
					fmt.Printf("=== Daftar Stock Barang ===\n\n%-5s | %-20s | %-16s | %-5s\n", "ID", "Nama", "Harga", "Stok")
					for _, product := range products {
						fmt.Printf("%-5d | %-20s | Rp. %-12d | %-5d\n", product.ID, product.Name, product.Price, product.Stock)
					}

					fmt.Println("\n=========\nTotal Barang:", len(products))

					return nil
				},
			}, {
				Name: "Exit",
				Handler: func() error {
					os.Exit(0)
					return nil
				},
			},
		},
	}

	return app
}

// func Task__() {

// 	var inputMenu string

// 	fmt.Print("=== Toko Kelontong | INVENTORY MANAGER ===\n\n")

// 	// menus := []string{"Tambah Barang", "Lihat Semua Stok", "Beli", "Keluar"}

// 	for endless := 0; endless <= 1; endless-- {

// 		if endless != 0 {
// 			fmt.Print("\n\n\n\n")
// 		}

// 		// for index, menu := range menus {
// 		// 	fmt.Println("[", index+1, "]", menu)
// 		// }

// 		fmt.Print("\nPilih Menu (1-3): ")
// 		fmt.Scanln(&inputMenu)

// 		// switch inputMenu {
// 		// case "1":

// 		// case "2":

// 		// case "3":

// 		// case "4":
// 		// 	return
// 		// default:
// 		// }
// 	}

// }

func main() {
	// Task__()
	app := NewApp()
	app.Run()
}
