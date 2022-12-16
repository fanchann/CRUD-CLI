package setting

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"crud_mysqli_cli/repository"
)

func InsertData() {
	var name, email string
	ctx := context.Background()
	resultRepo := repository.NewUserRepository(GetConnection())

	fmt.Printf("Masukan namamu :")
	scanInput := bufio.NewScanner(os.Stdin)
	if scanInput.Scan() {
		name = scanInput.Text()
	}

	fmt.Printf("Masukan emailmu :")
	fmt.Scanln(&email)

	has, err := resultRepo.Insert(ctx, name, email)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Success : %v\n", has)
}

func DeleteData() {
}

func ShowsData() {
	resultRepo := repository.NewUserRepository(GetConnection())

	has, err := resultRepo.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for _, c := range has {
		fmt.Printf("id :%v | nama :%v | email :%v\n", strconv.Itoa(int(c.Id)), c.Name, c.Email)
	}

}
