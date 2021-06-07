package crtshgo

// Author: Eduardo Barbosa (@_anakein)
// Date: 07/06/2021
// anakein@protonmail.ch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Crtsh struct {
	CommonName string `json:"common_name"`
}

func GetJsonFromCrt(domain string) ([]string, error) {

	resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	sb := []byte(body)
	var subdomains []Crtsh
	err = json.Unmarshal(sb, &subdomains)
	if err != nil {
		fmt.Println("error:", err)
	}

	output := make([]string, 0)
	for _, subdomains := range subdomains {
		output = append(output, subdomains.CommonName)
	}

	return output, nil
}
