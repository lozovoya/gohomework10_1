package cbrf

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Service struct {
	url     string
	timeout time.Duration
	client  *http.Client
}

func NewService(
	url string,
	timeout time.Duration,
	client *http.Client,
) *Service {
	return &Service{
		url: url, timeout: timeout * time.Minute, client: client,
	}
}

func (s *Service) Extract() (data []byte, err error) {
	ctx, _ := context.WithTimeout(context.Background(), s.timeout)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		s.url,
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err == nil {
				log.Println(err)
				err = cerr
			}
		}
	}()

	return respBody, err
}
