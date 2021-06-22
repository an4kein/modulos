package main
// package verifica365

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// this code will be modified and transformation in an function that return a value for add on project.
var banner = `
__ __ ________ __ ______  ___ ___          
|| ||||   || \\||||   || //  // \\     
\\ //||== ||_//||||== ||((   ||=||
 \V/ ||___|| \\||||   || \\__|| ||

 [+] by @anakein
 [+] https://github.com/an4kein					
																	   
`

func Verifica365(domain string) {
	resp, err := http.Get(fmt.Sprintf("https://login.microsoftonline.com/getuserrealm.srf?login=username@%s&xml=1", domain))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var LoginMSonline struct {
		NameSpaceType string `xml:"NameSpaceType"`
		IsFederatedNS string `xml:"IsFederatedNS"`
	}

	if err := xml.Unmarshal([]byte(body), &LoginMSonline); err != nil {
		log.Fatal(err)
	}

	tipo := LoginMSonline.NameSpaceType
	valor := LoginMSonline.IsFederatedNS

	switch {
	case tipo == "Unknown":
		fmt.Println("[-] Instância não pôde ser encontrada na plataforma (Office 365)")
	case tipo == "Managed" && valor == "false":
		fmt.Println("[+] Instância válida na plataforma (Office 365)\nO valor retornado indica uma solução LDAP local, que provavelmente foi usada para autenticar credenciais.")
	default:
		fmt.Println("[!] Nada encontrado")
	}
}

func main() {
	fmt.Println(banner)
	domain := flag.String("domain", "example.com", "Digite seu dominio!")
	flag.Parse()

	if *domain == "example.com" {
		flag.PrintDefaults()
	} else {
		Verifica365(*domain)
	}
}
