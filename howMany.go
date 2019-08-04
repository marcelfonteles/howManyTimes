package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	utils "howMany/pkg"
	"os"
	"strings"
	"time"
)

func init () {
	utils.BeautifulPrint("#")
	fmt.Println("Welcome to 'How many times my mother called me'!")
	utils.BeautifulPrint("#")
}
func main() {
	// Variável para controlar o programa
	var options string
	// Inicializando o leitor de entrado do usuário
	reader := bufio.NewReader(os.Stdin)
	// Inicializando a variável de controle
	options = "init"
	// String que é escrita no arquivo
	var texto [][]string
	// Preparando o Arquivo para receber os dados
	//file, _ := os.Open("/home/marcelvieira/go/bin/result.csv")
	//file, _ :=    os.Create("/home/marcelvieira/go/bin/result.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	leitor := csv.NewReader(file)
	message, err := leitor.Read()
	for err == nil {
		fmt.Println(message)
		message, err = leitor.Read()
	}
	defer writer.Flush()
	// Texto que será escrito no arquivo
	var dt time.Time
	var dtString string
	var last int // zero value of int is 0
	// Loop em que o programa está rodando
	for options != "exit" {
		fmt.Printf("Calling (stat for statistics): ")
		options, _ = reader.ReadString('\n')
		options = strings.TrimSuffix(options, "\n")
		if options != "stat" {
			dt = time.Now()
			dtString = dt.Format("02-Jan-2006 15:04:05")
			var appiendar = []string{"Data: ", dtString}
			texto = append(texto, appiendar)
			writer.Write(texto[last])
			last += 1
		}
		if options == "stat" {
			utils.ClearConsole()
			stat(texto)
			fmt.Println("Press any key to exit")
			options, _ = reader.ReadString('\n')
			utils.ClearConsole()
		}
		if options == "exit" {
			utils.BeautifulPrint("-")
		}

	}
}
// Função que mostra o log de chamamentos
func stat(texto [][]string) {
	utils.BeautifulPrint("-")
	for _, value := range texto {
		fmt.Println(value)
	}
	utils.BeautifulPrint("-")
}

