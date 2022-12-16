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

	fmt.Printf("Masukan namamu :")
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

func DeleteData() {
	var id int32

	fmt.Printf("Masukan id yang akan dihapus :")
	fmt.Scanln(&id)

	has, err := Connection.Delete(CtxB, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf(has)
	fmt.Println()

}

func ShowsData() {

	has, err := Connection.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for _, c := range has {
		fmt.Printf("id :%v | nama :%v | email :%v\n", strconv.Itoa(int(c.Id)), c.Name, c.Email)
	}

}
