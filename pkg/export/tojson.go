package export

import (
	"encoding/json"
	"github.com/lozovoya/gohomework10_1/pkg/locsave"
	"io/ioutil"
	"log"
)

type Currency struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Currencies struct {
	Currency []Currency
}

func ExportJSON(storage []locsave.Currency, filename string) error {

	encoded, err := json.Marshal(storage)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(filename, encoded, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
