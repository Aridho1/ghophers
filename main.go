package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (a *App) productTableWidth() (idW, nameW, priceW, stockW int) {
	idW = len("ID")
	nameW = len("Nama")
	priceW = len("Harga")
	stockW = len("Stok")

	for _, p := range a.Products {
		if w := len(strconv.Itoa(p.ID)); w > idW {
			idW = w
		}

		if w := len(p.Name); w > nameW {
			nameW = w
		}

		price := "Rp" + NumFormat(p.Price)
		if w := len(price); w > priceW {
			priceW = w
		}

		if w := len(strconv.Itoa(p.Stock)); w > stockW {
			stockW = w
		}
	}

	return
}

func NewApp() *App {
	var app *App

	app = &App{
		Title: "Waoreng Serba Ada",
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
					idW, nameW, priceW, stockW := app.productTableWidth()

					fmt.Printf("%*s | %-*s | %*s | %*s\n",
						idW, "ID",
						nameW, "Nama",
						priceW, "Harga",
						stockW, "Stok",
					)

					fmt.Println(strings.Repeat("-", idW+nameW+priceW+stockW+9))

					for _, p := range app.Products {
						time.Sleep(WAIT_SECONDARY)
						fmt.Printf("%*d | %-*s | %*s | %*d\n",
							idW, p.ID,
							nameW, p.Name,
							priceW, "Rp"+NumFormat(p.Price),
							stockW, p.Stock,
						)
					}

					time.Sleep(WAIT_SECONDARY)
					fmt.Println("\n=========\nTotal Barang:", len(app.Products))

					return nil
				},
			}, {
				Name: "Transaksi",
				Handler: func() error {
					if len(app.Products) == 0 {
						fmt.Println("[SYSTEM]: Tidak ada barang.")
						return nil
					}

					var id, qty int

					fmt.Print("> Masukkan ID Barang: ")
					if !scanner.Scan() {
						return nil
					}
					id, _ = strconv.Atoi(scanner.Text())

					var product *Product

					for i := range app.Products {
						if app.Products[i].ID == id {
							product = &app.Products[i]
							break
						}
					}

					if product == nil {
						fmt.Println("[SYSTEM]: Barang tidak ditemukan.")
						return nil
					}

					fmt.Printf("> Jumlah beli (%s): ", product.Name)
					if !scanner.Scan() {
						return nil
					}
					qty, _ = strconv.Atoi(scanner.Text())

					if qty <= 0 {
						fmt.Println("[SYSTEM]: Jumlah tidak valid.")
						return nil
					}

					if qty > product.Stock {
						fmt.Println("[SYSTEM]: Stok tidak mencukupi.")
						return nil
					}

					total := qty * product.Price
					product.Stock -= qty

					fmt.Println("\n===== STRUK =====")
					fmt.Println("Barang :", product.Name)
					fmt.Println("Harga  : Rp.", NumFormat(product.Price))
					fmt.Println("Qty    :", qty)
					fmt.Println("-------------------------")
					fmt.Println("Total  : Rp.", NumFormat(total))
					fmt.Println("=========================")
					fmt.Println("[SYSTEM]: Transaksi berhasil!")

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
		Name:  "Susu HU TAO RILL",
		Stock: 69,
		Price: 67000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Indomie Rasa Tanggal Tua",
		Stock: 404,
		Price: 13337,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Aqua Air Mata Skripsi",
		Stock: 420,
		Price: 6969,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Oreo Isi Janji Manis",
		Stock: 100,
		Price: 15000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Parfum Bau Deadline",
		Stock: 99,
		Price: 45000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Power Bank 1000% (Sisa 1%)",
		Stock: 1,
		Price: 99999,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Mie Instan Rasa Balikan",
		Stock: 0,
		Price: 14000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Keyboard RGB FPS +999",
		Stock: 88,
		Price: 250000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Es Teh Rasa Chat Dibales",
		Stock: 123,
		Price: 5000,
	})

	app.Products = append(app.Products, Product{
		ID:    app.generateProductID(),
		Name:  "Kopi Anti Ngantuk (Bohong)",
		Stock: 200,
		Price: 18000,
	})

	return app
}

func main() {
	app := NewApp()
	app.Run()
}
