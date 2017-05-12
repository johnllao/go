package main

import "encoding/csv"
import "io"
import "log"
import "net/http"
import "strconv"

// Company represents the company details
type Company struct {
	Symbol 		 string
	Name		 string
	LastSale	 float32
	MarketCap    float64
	ADRTSO       string
	IPOYear      string
	Sector       string
	Industry     string	
	SummaryQuote string
}

// CompanyList retrieves the company list
type CompanyList struct {
	url string
}

// NewCompanyList returns new instance of the CompanyList
func NewCompanyList() CompanyList {
	return CompanyList{
		url: "http://www.nasdaq.com/screening/companies-by-industry.aspx?exchange=NASDAQ&render=download",
	}
}

// Download company list
func (c CompanyList) Download() (<-chan Company) {

	companyChan := make(chan Company)

	go func() {
		// columns
		// Symbol, Name, LastSale, MarketCap, ADR TSO, IPOyear, Sector, Industry, Summary Quote
		log.Printf("downloading company list. URL : %s", c.url)
		
		req, err := http.NewRequest("GET", c.url, nil)
		if err != nil {
			close(companyChan)
			return
		}

		cli := http.Client{}
		res, err := cli.Do(req)
		if err != nil {
			close(companyChan)
			return
		}
		
		body := res.Body
		defer body.Close()

		rdr := csv.NewReader(body)
		for {
			rec, err := rdr.Read()
			if err == io.EOF {
				break
			}
			lastSale, err := strconv.ParseFloat(rec[2], 32)
			if err != nil {
				lastSale = 0
			}
			mktCap, err := strconv.ParseFloat(rec[3], 64)
			if err != nil {
				mktCap = 0
			}
			companyChan <- Company{
				Symbol: rec[0],
				Name: rec[1],
				LastSale: float32(lastSale),
				MarketCap: mktCap,
				ADRTSO: rec[4],
				IPOYear: rec[5],
				Sector: rec[6],
				Industry: rec[7],
			}
			
		}

		close(companyChan)
	
	}()
	return companyChan
}