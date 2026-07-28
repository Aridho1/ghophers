package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	WAIT_PRIMARY   = 1 * time.Second
	WAIT_SECONDARY = 300 * time.Millisecond
)

var scanner = bufio.NewScanner(os.Stdin)

func NumFormat(num int) string {
	str := strconv.Itoa(num)
	strLen := len(str)
	if strLen <= 3 {
		return str
	}

	var result []byte
	for i := range strLen {
		if i > 0 && i%3 == 0 {
			result = append([]byte{'.'}, result...)
		}
		result = append([]byte{str[strLen-1-i]}, result...)
	}

	return string(result)
}

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
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

func (a *App) generateProductID() int {
	if a == nil || a.Products == nil {
		return 1
	}

	return len(a.Products) + 1
}

func (a *App) LogTitle() {
	fmt.Printf("=== %s | INVENTORY MANAGER ===\n\n", a.Title)
}

func (a *App) LogListMenu(menuLen int) {

	fmt.Printf("List Menu\n\n")

	for i, menu := range a.Menus {
		num := i + 1

		if i == menuLen-1 {
			num = 0
		}

		fmt.Printf("%d. %s\n", num, menu.Name)
	}
}

func (a *App) Run() {
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

		fmt.Printf("\n> Pilih Menu [0-%d]: ", menuLen-1)

		if !scanner.Scan() {
			os.Exit(1)
		}

		inputMenu, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Input harus berupa angka!")
			continue
		}

		fmt.Printf("\n")
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
		Menus: []Menu{
			{
				Name: "Tambah Barang",
				Handler: func() error {
					var inputItemName string
					var inputPrice, inputStock int

					fmt.Print("> Masukan Nama Barang: ")
					if !scanner.Scan() {
						os.Exit(1)
					}
					inputItemName = scanner.Text()

					fmt.Print("> Masukan Harga Barang: ")
					if scanner.Scan() {
						inputPrice, _ = strconv.Atoi(scanner.Text())
					}

					fmt.Print("> Masukan Stock Barang: ")
					if scanner.Scan() {
						inputStock, _ = strconv.Atoi(scanner.Text())
					}

					product := Product{
						ID:    app.generateProductID(),
						Name:  inputItemName,
						Price: inputPrice,
						Stock: inputStock,
					}

					app.Products = append(app.Products, product)

					fmt.Println("\n[SYSTEM]: Barang berhasil ditambahkan ke gudang!")

					return nil
				},
			}, {
				Name: "Lihat Semua Stock",
				Handler: func() error {
					fmt.Printf("=== Daftar Stock Barang ===\n\n%-5s | %-20s | %-16s | %-5s\n", "ID", "Nama", "Harga", "Stok")

					for _, product := range app.Products {
						time.Sleep(WAIT_SECONDARY)
						fmt.Printf("%-5d | %-20s | Rp. %-12s | %-5d\n", product.ID, product.Name, NumFormat(product.Price), product.Stock)
					}

					time.Sleep(WAIT_SECONDARY)
					fmt.Println("\n=========\nTotal Barang:", len(app.Products))

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

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Minyak Goreng 2L",
		Stock: 100,
		Price: 30000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Beras 200gr",
		Stock: 400,
		Price: 3000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Susu HU TAO RILL",
		Stock: 69,
		Price: 67000,
	})

	return app
}

func main() {
	app := NewApp()
	app.Run()
}
