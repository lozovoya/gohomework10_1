package cbrf

type Currency struct {
	XMLName  string  `xml:"Valute"`
	Numcode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

type Curses struct {
	XMLName    string     `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}
