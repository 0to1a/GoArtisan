package Controller

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

const (
	GoMainLess = `package main

import (
	"log"
	"math/rand"
	"time" 
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Print("Server: Starting")
 
	go func() {
		log.Print("Server: Service Running")
 	}()

	select {}
}
`
	GoMain = `package main

import (
	"log"
	"math/rand"
	"time"

	"./Model"
	"./Routing"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Print("Server: Starting")
	Model.Database = ConfigDatabase()

	go func() {
		log.Print("Server: Service Running")
		Routing.ConfigWebServer()
	}()

	select {}
}
`
	GoConfig = `package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

func ConfigDatabase() *goqu.Database {
	var (
		err      error
		database *sql.DB
		db       *goqu.Database
	)
	database, err = sql.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/DATABASE?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	db = goqu.New("mysql", database)
	return db
}
`
	GoDatabase = `package Model

import "gopkg.in/doug-martin/goqu.v5"

var Database *goqu.Database
`
	GoRoute = `package Routing

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigWebServer() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
`
	GoRandom = `package Tools

import (
	"math/rand"
)

var (
	letterRunes    = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letterRunesLow = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringLower(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunesLow[rand.Intn(len(letterRunesLow))]
	}
	return string(b)
}
`
)

// CreatePrintHelp to print command list for Create function.
func CreatePrintHelp() {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("%s\n", yellow("Description:"))
	fmt.Printf("  %s\n", "Create a new project with GoArtisan standart structure")
	fmt.Println()

	fmt.Printf("%s\n", yellow("Usage:"))
	fmt.Printf("  %s\n", "create [options] <name>")
	fmt.Println()

	fmt.Printf("%s\n", yellow("Arguments:"))
	fmt.Printf("  %s %s\n", green("name        "), "Name of folder project")
	fmt.Println()

	fmt.Printf("%s\n", yellow("Options:"))
	fmt.Printf("  %s %s\n", green("-h, --help  "), "Display this help message")
	fmt.Printf("  %s %s\n", green("-l, --less  "), "Less project structure based")
}

// CreateFolder is create project folder with some files in framework.
//
// args value contains some argument or option to run this command.
func CreateFolder(args []string) {
	var (
		data       []byte
		err        error
		isLess     bool = false
		folderpath string
	)
	folderpath = args[1]

	switch folderpath {
	case "-h", "--help":
		CreatePrintHelp()
		return
	case "-l", "--less":
		isLess = true
		folderpath = args[2]
	}

	if _, err = os.Stat(folderpath); os.IsNotExist(err) {
		os.MkdirAll(folderpath, os.ModePerm)
		os.MkdirAll(folderpath+"/Tools", os.ModePerm)
		data = []byte(GoConfig)
		err = ioutil.WriteFile(folderpath+"/config.go", data, 0755)
		CheckErr(err)
		data = []byte(GoRandom)
		err = ioutil.WriteFile(folderpath+"/Tools/RandomString.go", data, 0755)
		CheckErr(err)

		if !isLess {
			os.MkdirAll(folderpath+"/Model", os.ModePerm)
			os.MkdirAll(folderpath+"/Routing", os.ModePerm)
			data = []byte(GoDatabase)
			err = ioutil.WriteFile(folderpath+"/Model/database.go", data, 0755)
			CheckErr(err)
			data = []byte(GoRoute)
			err = ioutil.WriteFile(folderpath+"/Routing/route.go", data, 0755)
			CheckErr(err)
			data = []byte(GoMain)
			err = ioutil.WriteFile(folderpath+"/main.go", data, 0755)
			CheckErr(err)
		} else {
			data = []byte(GoMainLess)
			err = ioutil.WriteFile(folderpath+"/main.go", data, 0755)
			CheckErr(err)
		}

		if !IsQuiet {
			blue := color.New(color.FgBlue).SprintFunc()
			fmt.Printf("Project	'%s' is created\n", blue(folderpath))
		}
	} else {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("Error: Folder '%s' is exist.\n", red(folderpath))
	}
}
