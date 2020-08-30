package locsave

import (
	"encoding/xml"
	"errors"
	"github.com/lozovoya/gohomework10_1/pkg/cbrf"
	"log"
)

var ErrNoIncomingData = errors.New("Import data is empty")
var ErrLocalStorageEmpty = errors.New("Local storage is empty")

type Currency struct {
	Code  string
	Name  string
	Value float64
}

func SaveData(extract []byte) (storage []Currency, err error) {

	var data *cbrf.Curses
	err = xml.Unmarshal(extract, &data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(data.Currencies) == 0 {
		return nil, ErrNoIncomingData
	}

	buf := Currency{}
	for _, j := range data.Currencies {
		buf.Code = j.CharCode
		buf.Name = j.Name
		buf.Value = j.Value

		storage = append(storage, buf)
	}
	if len(storage) == 0 {
		return nil, ErrLocalStorageEmpty
	}

	return storage, nil
}
