package main

// Author: Eduardo Barbosa (@anakein) 
// 25/06/2021 
// anakein@protonmail.ch

// This function check if username is valid using the method office.com
// Next, will altered for an funcion and will also be added on project
// suggest name for package is  "valida_email"


import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var banner = `
__ __ ________ __ ______  ___ ___          
|| ||||   || \\||||   || //  // \\     
\\ //||== ||_//||||== ||((   ||=||
 \V/ ||___|| \\||||   || \\__|| ||
 [+] by @anakein
 [+] https://github.com/an4kein					
																	   
`
var lista_emails []string

func ReadLine() []string {
	file, err := os.Open("/home/anakein/Desktop/email.txt") // Caminho do seu arquivo com emails, exemplo: emails.txt
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		email_list := scanner.Text()
		lista_emails = append(lista_emails, email_list)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lista_emails
}

func main() {
	fmt.Println(banner)
	emails := ReadLine()
	for _, usernames := range emails {
		authAuthenticatorUrl := "https://login.microsoftonline.com/common/GetCredentialType?mkt=en-US"
		password := "13371337" /* Senha padrao, voce pode modificar.
		Ainda vou incorporar uns ajustes para fazer password spray.
		Nao tenho uma conta valida na MS, so posso continuar os trabalhos, quando conseguir uma emprestada.*/
		values := map[string]string{"username": usernames, "password": password}
		jsonValue, _ := json.Marshal(values)
		resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var MSofficeResponse struct {
			Username       string `json:"Username"`
			IfExistsResult int    `json:"IfExistsResult"`
		}
		if err := json.Unmarshal([]byte(body), &MSofficeResponse); err != nil {
			panic(err)
		}
		if MSofficeResponse.IfExistsResult == 0 {
			fmt.Println("[+] VALID USERNAME", MSofficeResponse.Username)
		} else {
			fmt.Println("[-] INVALID USERNAME", MSofficeResponse.Username)
		}
	}
}
