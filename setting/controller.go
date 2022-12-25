package setting

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"crud_mysqli_cli/repository"
)

var (
	Connection = repository.NewUserRepository(GetConnection())
	CtxB       = context.Background()
)

func InsertData() {
	var name, email string

	fmt.Printf("\nMasukan namamu :")
	scanInput := bufio.NewScanner(os.Stdin)
	if scanInput.Scan() {
		name = scanInput.Text()
	}

	fmt.Printf("Masukan emailmu :")
	fmt.Scanln(&email)

	has, err := Connection.Insert(CtxB, name, email)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success : %v\n", has)
}

func UpdateData() {
	var id int32
	var name, email string

	fmt.Printf("\nMasukan Id :")
	fmt.Scanln(&id)

	fmt.Printf("Masukan nama baru :")
	nameScan := bufio.NewScanner(os.Stdin)
	if nameScan.Scan() {
		name = nameScan.Text()
	}

	fmt.Printf("Masukan email baru :")
	fmt.Scanln(&email)

	_, err := Connection.Update(CtxB, name, email, id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSucces update id %d\nName -> %s\nEmail -> %s \n", id, name, email)
}

func DeleteData() {
	var id int32

	fmt.Printf("\nMasukan id yang akan dihapus :")
	fmt.Scanln(&id)

	has, err := Connection.Delete(CtxB, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf(has)
	fmt.Println()
	ShowsData()
}

func ShowsData() {

	has, err := Connection.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for _, c := range has {
		fmt.Printf("id : %v | nama : %v | email : %v\n", strconv.Itoa(int(c.Id)), c.Name, c.Email)
	}
	fmt.Println()
}

func ShowDataBId() {
	var id int32

	fmt.Printf("\nMasukan id yang dicari :")
	fmt.Scanln(&id)

	has, err := Connection.FindById(CtxB, id)
	if err != nil {
		panic(err)
	}

	if len(has.Email) == 0 {
		fmt.Println("kosong")
		return
	}
	fmt.Println("Id -> ", has.Id)
	fmt.Println("Name -> ", has.Name)
	fmt.Println("Email -> ", has.Email)

}
