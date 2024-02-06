package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// error user sudah ada
var (
	ErrUserNotFound   = errors.New("user tidak di temukan")
	ErrIdUserInvalid  = errors.New("ID user tidak valid")
	ErrIdUserNotFound = errors.New("ID user tidak di temukan")
)

type User struct {
	Name  string
	Age   string
	Email string
}

// data user
type Application struct {
	Data  []User
	name  string
	age   string
	email string
}

func (app *Application) CommandCreate() (message string) {
	// ambil nama user
	fmt.Println("Nama Anda: ?")
	Input(&app.name)

	// ambil umur user
	fmt.Println("Umur Anda: ?")
	Input(&app.age)

	// ambil email
	fmt.Println("Email Anda: ?")
	Input(&app.email)

	newUser := User{
		Name:  app.name,
		Age:   app.age,
		Email: app.email,
	}

	temp := app.Data
	app.Data = append(temp, newUser)

	// return
	message = fmt.Sprintf("user baru berhasil di tambahkan (%+v)", newUser)
	return
}

func (app *Application) CommandDisplay() {
	result := app.Data
	if len(result) > 0 {
		for i, v := range result {
			fmt.Printf("(%d). Nama: %s Umur: %s Email: %s \n", i+1, v.Name, v.Age, v.Email)
		}
	} else {
		fmt.Println("---- data kosong ----")
	}
}

func (app *Application) CommandEdit() (message string, err error) {
	var getInput string

	fmt.Println("ID Target")

	Input(&getInput)

	id, err := strconv.Atoi(getInput)
	if err != nil {
		message = ""
		err = ErrIdUserInvalid
		return
	}

	// name
	fmt.Println("Nama Anda: ?")
	Input(&app.name)

	fmt.Println("Umur Anda: ?")
	Input(&app.age)

	fmt.Println("Email Anda: ?")
	Input(&app.email)

	// validasi ID
	if id > len(app.Data) || id < 1 {
		message = ""
		err = ErrIdUserNotFound
		return
	}

	newUser := User{
		Name:  app.name,
		Age:   app.age,
		Email: app.email,
	}

	app.Data[id-1] = newUser

	message = fmt.Sprintf("user dengan ID %d berhasil di perbaharui. (%+v)", id, newUser)
	err = nil
	return
}

func (app *Application) CommandDelete() (message string, err error) {
	var data []User = app.Data
	var getInput string

	fmt.Println("ID target")
	Input(&getInput)
	id, err := strconv.Atoi(getInput)

	if err != nil {
		message = ""
		err = ErrIdUserInvalid
		return

	} else {
		if id > len(data) || id < 1 {
			message = ""
			err = ErrUserNotFound
			return
		}
		cutA := data[:id-1]
		cutB := data[id:]
		newData := append(cutA, cutB...)
		app.Data = newData

		message = fmt.Sprintf("user dengan ID %d berhasil di hapus", id)
		err = nil
		return
	}
}

func NewApplication() *Application {
	return &Application{}
}

func Input(name *string) {
	in := bufio.NewReader(os.Stdin)
	read := bufio.NewScanner(in)
	read.Scan()

	*name = read.Text()
}
