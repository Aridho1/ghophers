package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	WAIT_PRIMARY   = 1 * time.Second
	WAIT_SECONDARY = 300 * time.Millisecond
)

var scanner = bufio.NewScanner(os.Stdin)

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()

	// fmt.Print("\033[2J")
}

func PressEnterToContinue() {
	fmt.Print("\nPress Enter To Continue...")
	if scanner.Scan() {
		fmt.Print("\n")
		return
	}

	os.Exit(1)
}

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
			// fmt.Print("\n\n\n\n")
			time.Sleep(WAIT_PRIMARY)
			PressEnterToContinue()
		}

		ClearScreen()
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
func (a *App) LogListProduct() {
	idW, nameW, priceW, stockW := a.productTableWidth()

	fmt.Printf("%*s | %-*s | %*s | %*s\n",
		idW, "ID",
		nameW, "Nama",
		priceW, "Harga",
		stockW, "Stok",
	)

	fmt.Println(strings.Repeat("-", idW+nameW+priceW+stockW+9))

	for _, p := range a.Products {
		time.Sleep(WAIT_SECONDARY)
		fmt.Printf("%*d | %-*s | %*s | %*d\n",
			idW, p.ID,
			nameW, p.Name,
			priceW, "Rp"+NumFormat(p.Price),
			stockW, p.Stock,
		)
	}
}

func (a *App) EnsureProductLen() bool {
	if len(a.Products) == 0 {
		fmt.Println("[SYSTEM]: Tidak ada barang.")
		return false
	}
	return true
}

func NewApp() *App {
	var app *App

	app = &App{
		Title: "Waroeng Serba Ada",
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
					app.LogListProduct()

					time.Sleep(WAIT_SECONDARY)
					fmt.Println("\n=========\nTotal Barang:", len(app.Products))

					return nil
				},
			}, {
				Name: "Transaksi",
				Handler: func() error {
					if !app.EnsureProductLen() {
						return nil
					}

					app.LogListProduct()

					var id, qty int
					var product *Product

					for {
						fmt.Print("\n> Masukkan ID Barang: ")
						if !scanner.Scan() {
							return nil
						}
						id, _ = strconv.Atoi(scanner.Text())

						for i := range app.Products {
							if app.Products[i].ID == id {
								product = &app.Products[i]
								break
							}
						}

						if product == nil {
							fmt.Println("[SYSTEM]: Barang tidak ditemukan.")
							continue
						}

						fmt.Printf("> Jumlah beli (%s): ", product.Name)
						if !scanner.Scan() {
							return nil
						}
						qty, _ = strconv.Atoi(scanner.Text())

						if qty <= 0 {
							fmt.Println("[SYSTEM]: Jumlah tidak valid.")
							continue
						}

						if qty > product.Stock {
							fmt.Println("[SYSTEM]: Stok tidak mencukupi.")
							continue
						}

						break

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
				Name: "Case: Flash Sale",
				Handler: func() error {
					if !app.EnsureProductLen() {
						return nil
					}

					app.LogListProduct()

					const FLASH_SALE_QUANTITY_PEOPLE = 6767
					var id int
					var product *Product

					for {
						fmt.Print("\n> Masukkan ID Barang untuk FLASH SALE 12.12: ")
						if !scanner.Scan() {
							return nil
						}
						id, _ = strconv.Atoi(scanner.Text())

						for i := range app.Products {
							if app.Products[i].ID == id {
								product = &app.Products[i]
								break
							}
						}

						if product == nil {
							fmt.Println("[SYSTEM]: Barang tidak ditemukan.")
							continue
						}

						break
					}

					fmt.Printf("\n>>> Serbuan 12.12 untuk \"%s\"! Stok %d, penyerbu %d orang\n\n", product.Name, product.Stock, FLASH_SALE_QUANTITY_PEOPLE)

					start := time.Now()
					var wg sync.WaitGroup

					prevStock := product.Stock

					for range FLASH_SALE_QUANTITY_PEOPLE {
						if product.Stock < 1 {
							break
						}

						wg.Add(1)
						go func() {
							defer wg.Done()
							product.Stock--
						}()
					}

					wg.Wait()
					elapsed := time.Since(start)

					const _INDENT = 12

					fmt.Printf("====== FLASH SALE 12.12 SELESAI  ======\n")
					fmt.Printf("%-*s: %dms\n%-*s: %d\n%-*s: %d\n\n", _INDENT, "Waktu Proses", elapsed.Milliseconds(), _INDENT, "Unit Terjual", prevStock-product.Stock, _INDENT, "Sisa stok", product.Stock)

					if product.Stock < 0 {
						fmt.Printf("\n>> BAHAYA! Stok \"%s\" MINUS %d -- Toko menjual barang yang tidak ada!\n", product.Name, product.Stock*-1)
					}

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

	app.Products = append(app.Products, Product{ID: app.generateProductID(), Name: "Minyak Goreng 2L", Stock: 100, Price: 30000})

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
