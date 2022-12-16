package setting

import (
	"fmt"
	"time"
)

func Menus() {
	var menu int
	fmt.Printf(`
[01] Insert Data
[02] Show Data
[03] Update Data
[04] Delete data
[05] Show Data by Id
`)

	fmt.Printf("[*] Choose :")
	fmt.Scanln(&menu)

	switch menu {
	case 1 | 01:
		InsertData()
	case 2 | 02:
		ShowsData()
	default:
		fmt.Println("yang bener lu!")
		time.Sleep(2 * time.Second)
		fmt.Println("\033[2J")
		Menus()
	}
}
