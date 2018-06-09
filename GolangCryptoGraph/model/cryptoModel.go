package model

type CryptoData struct {
	MetaData   MetaDataData              `json:"Meta Data"`
	TimeSeries map[string]TimeSeriesData `json:"Time Series (Digital Currency Daily)"`
	CryptoType []CryptoType
}

type MetaDataData struct {
	Information         string `json:"1. Information"`
	DigitalCurrencyCode string `json:"2. Digital Currency Code"`
	DigitalCurrencyName string `json:"3. Digital Currency Name"`
	MarketCode          string `json:"4. Market Code"`
	MarketName          string `json:"5. Market Name"`
	//Interval            string `json:"6. Interval"`
	LastRefreshed string `json:"6. Last Refreshed"`
	TimeZone      string `json:"7. Time Zone"`
}

type TimeSeriesData struct {
	OpenPhyCur  string `json:"1a. open (USD)"`
	OpenDigCur  string `json:"1b. open (USD)"`
	HighPhyCur  string `json:"2a. high (USD)"`
	HighDigCur  string `json:"2b. high (USD)"`
	LowPhyCur   string `json:"3a. low (USD)"`
	LowDigCur   string `json:"3b. low (USD)"`
	ClosePhyCur string `json:"4a. close (USD)"`
	CloseDigCur string `json:"4b. close (USD)"`
	Volume      string `json:"5. volume"`
	MarketCap   string `json:"6. market cap (USD)"`
}
type CryptoType struct {
	CryptoSymbol string
	CryptoName   string
}
