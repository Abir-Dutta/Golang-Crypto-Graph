package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"goApps/GolangCryptoGraph/model"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var CryptoDataMap = map[string]model.CryptoData{}
var CryptoDataMapFiltered = map[string]model.CryptoData{}

func getCryptoHistory(cryptoChannel chan model.CryptoData, cryptoSymb, cryptoName string) {
	defer wg.Done()
	var CryptoData model.CryptoData
	resp, geterr := http.Get("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=" + cryptoSymb + "&market=USD&apikey=E76NR0L8EXZI71B0")
	if geterr == nil {
		bytes, readerr := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if readerr == nil {
			err := json.Unmarshal(bytes, &CryptoData)
			if err == nil {
				cryptoChannel <- CryptoData
			}
		}
	}

}

func goGetHandler(w http.ResponseWriter, r *http.Request) {

	// var Btc model.CryptoData
	// cryptos := r.FormValue("cryptos")
	// fmt.Println(cryptos)
	// if r.Method == "GET" {
	// 	resp, _ := http.Get("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=BTC&market=USD&apikey=E76NR0L8EXZI71B0")
	// 	bytes, _ := ioutil.ReadAll(resp.Body)
	// 	json.Unmarshal(bytes, &Btc)
	// } else {
	// 	resp, _ := http.Get("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=" + cryptos + "&market=USD&apikey=E76NR0L8EXZI71B0")
	// 	bytes, _ := ioutil.ReadAll(resp.Body)
	// 	json.Unmarshal(bytes, &Btc)
	// }
	// t, _ := template.ParseFiles("./template/view/newView.gohtml")
	// fmt.Println(t.Execute(w, Btc))
	//CryptoDataMap := make(map[string]model.CryptoData)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./template/view/newView.gohtml")
		fmt.Println(t.Execute(w, CryptoDataMapFiltered["BTC"]))
	} else {
		cryptos := r.FormValue("cryptos")
		t, _ := template.ParseFiles("./template/view/newView.gohtml")
		fmt.Println(t.Execute(w, CryptoDataMapFiltered[cryptos]))
	}
}

func main() {
	var CryptoList []model.CryptoType

	b, err := ioutil.ReadFile("digital_currency_list.csv") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	r := csv.NewReader(strings.NewReader(str))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] != "currency code" {
			CryptoList = append(CryptoList, model.CryptoType{record[0], record[1]})
		}
	}
	cryptoChannel := make(chan model.CryptoData, 1000)
	for _, ele := range CryptoList {
		wg.Add(1)
		go getCryptoHistory(cryptoChannel, ele.CryptoSymbol, ele.CryptoName)
	}
	wg.Wait()
	close(cryptoChannel)
	CryptoListFilterd := make([]model.CryptoType, 0)
	for ele := range cryptoChannel {
		if len(ele.MetaData.DigitalCurrencyCode) > 0 {
			CryptoDataMap[ele.MetaData.DigitalCurrencyCode] = ele
			CryptoListFilterd = append(CryptoListFilterd, model.CryptoType{ele.MetaData.DigitalCurrencyCode, ele.MetaData.DigitalCurrencyName})
		}
	}
	for idx, ele := range CryptoDataMap {
		ele.CryptoType = CryptoListFilterd
		CryptoDataMapFiltered[idx] = ele
		fmt.Println(ele.MetaData.DigitalCurrencyCode + "," + ele.MetaData.DigitalCurrencyName)
	}
	http.HandleFunc("/", goGetHandler)
	appStartErr := http.ListenAndServe(":8000", nil)
	log.Println(appStartErr)
	fmt.Println("Server Started on 8000")
}
