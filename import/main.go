package main

import "encoding/csv"
import "flag"
import "io"
import "log"
import "net/http"
import "strconv"

// Company represents the company details
type Company struct {
	Symbol 		string
	Name		string
	LastSale	float32
	MarketCap   float64
	Sector      string
	Industry    string	
}

var symbolArg = flag.String("sym", "", "symbol")
var symURL = "http://www.batstrading.com/market_data/symbol_listing/csv"
var companyListURL = "http://www.nasdaq.com/screening/companies-by-industry.aspx?exchange=NASDAQ&render=download"

func init() {
	flag.Parse()
}

func main() {

	syms, err := downloadCompanyList()
	if err != nil {
		log.Print(err.Error())
	}

	if comp, exists := syms[*symbolArg]; exists {
		log.Print(*symbolArg, comp)
	} else {
		log.Printf("invalid symbol. symbol : %s", *symbolArg)
	}
}

func downloadSymbols() (map[string]bool, error) {

	log.Printf("downloading symbols. URL : %s", symURL)
	
	req, err := http.NewRequest("GET", symURL, nil)
	if err != nil {
		return nil, err
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	
	body := res.Body
	defer body.Close()

	syms := make(map[string]bool)
	rdr := csv.NewReader(body)
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		syms[rec[0]] = true
	}

	return syms, nil
}

func downloadCompanyList() (map[string]Company, error) {

	// columns
	// Symbol, Name, LastSale, MarketCap, ADR TSO, IPOyear, Sector, Industry, Summary Quote
	log.Printf("downloading company list. URL : %s", companyListURL)
	
	req, err := http.NewRequest("GET", companyListURL, nil)
	if err != nil {
		return nil, err
	}

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	
	body := res.Body
	defer body.Close()

	syms := make(map[string]Company)
	rdr := csv.NewReader(body)
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		lastSale, _ := strconv.ParseFloat(rec[2], 32)
		mktCap, _ := strconv.ParseFloat(rec[3], 64)
		syms[rec[0]] = Company{
			Symbol: rec[0],
			Name: rec[1],
			LastSale: float32(lastSale),
			MarketCap: mktCap,
			Sector: rec[6],
			Industry: rec[7],
		}
		
	}

	return syms, nil
}