package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type config struct {
	Marketplace   string `json:"marketplace_url"`
	Proxy         string `json:"proxy_url"`
	App           string `json:"appplication_id"`
	Username      string `json:"marketplace_username"`
	Token         string `json:"marketplace_oauth2_token"`
	Brand         string `json:"brand"`
	FiwareService string `json:"fiware_service"`
}

type dataSource struct {
	EntityType string
	EntityID   string
	Query      string
	Name       string
}

type productData struct {
	Config     config
	DataSource dataSource
}

// UploadAPI : Asset Upload API endpoint
const UploadAPI = "/charging/api/assetManagement/assets/uploadJob"

// ValidateAPI : Asset Validate API endpoint
const ValidateAPI = "/charging/api/assetManagement/assets/validateJob"

// ProductAPI : Product Specification API endpoint
const ProductAPI = "/DSProductCatalog/api/catalogManagement/v2/productSpecification"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filename given")
		os.Exit(0)
	}

	fileName := os.Args[1]
	printBanner()

	// Loading configuration
	dat, err := ioutil.ReadFile("config.json")
	check(err)
	var c config
	err = json.Unmarshal(dat, &c)

	// Loading data sources file
	dat, err = ioutil.ReadFile(fileName)
	check(err)

	// Parsing data
	ds, err := parseData(c, dat)
	check(err)

	// Showing data sources to import
	fmt.Println()
	fmt.Println("The following data sources will be imported:")
	for _, el := range ds {
		fmt.Println(el.Name)
	}
	// Ask for confirmation
	fmt.Println("WARNING: Are you sure? (yes/no)")
	if !askForConfirmation() {
		os.Exit(0)
	}

	// init
	aTmpl := template.Must(template.ParseFiles("templates/asset"))
	pTmpl := template.Must(template.ParseFiles("templates/product"))
	var b bytes.Buffer

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{}

	for _, el := range ds {
		pd := productData{c, el}
		fmt.Println()
		fmt.Println("Response for:", el.Name)

		b.Reset()
		aTmpl.Execute(&b, pd)
		fmt.Println("Asset Upload - ", sendPost(client, c.Marketplace+UploadAPI, c.Token, b.Bytes()))

		b.Reset()
		pTmpl.ExecuteTemplate(&b, "productValidation", pd)
		fmt.Println("POST Validation - ", sendPost(client, c.Marketplace+ValidateAPI, c.Token, b.Bytes()))

		b.Reset()
		pTmpl.ExecuteTemplate(&b, "product", pd)
		fmt.Println("Data Source Specification - ", sendPost(client, c.Marketplace+ProductAPI, c.Token, b.Bytes()))
	}
}

func sendPost(c *http.Client, url string, t string, b []byte) string {
	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+t)
	// Save a copy of this request for debugging.
	// requestDump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(requestDump))
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return (resp.Status + " " + string(body))
}

func parseData(c config, d []byte) ([]dataSource, error) {
	ds := []dataSource{}
	var err error

	lines := strings.Split(string(d), "\n")
	for _, line := range lines {
		if (line != "") && (line != "\n") {
			element := dataSource{}
			t := strings.Fields(line)
			if len(t) > 0 && t[0] != "" {
				element.Query = "type=" + t[0]
				element.EntityType = t[0]
				element.Name = element.EntityType
			} else {
				err = errors.New("Error parsing file")
				break
			}
			if len(t) > 1 && t[1] != "" {
				element.Query = element.Query + "&id=" + t[1]
				element.EntityID = t[1]
				element.Name = element.Name + " : " + element.EntityID
			}
			ds = append(ds, element)
		}
	}
	return ds, err
}
