package setting

import (
	"fmt"
	"time"
)

var (
	menu int
)

func Menus() {
	fmt.Printf(`[01] Insert Data
[02] Show Data
[03] Update Data
[04] Delete data
[05] Show Data by Id
`)
	Choose()
}

func Choose() {

	fmt.Printf("\n[*] Choose :")
	fmt.Scanln(&menu)

	switch menu {
	case 1 | 01:
		InsertData()
	case 2 | 02:
		ShowsData()
	case 3 | 03:
		UpdateData()
	case 4 | 04:
		DeleteData()
	case 5 | 05:
		ShowDataBId()
	default:
		fmt.Println("yang bener lu!")
		time.Sleep(2 * time.Second)
		fmt.Println("\033[2J")
		Menus()
	}
}
