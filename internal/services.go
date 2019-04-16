package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	injson "github.com/jmg-duarte/statuspage/internal/json"
)

type Services map[string]Service

func (s *Services) Add(serv map[string]string) {
	id := string(serv["id"])
	_, ok := (*s)[id]
	if !ok {
		(*s)[id] = NewService(serv)
	}
}

func (s Services) PollServices(brief bool) {
	for _, service := range s {
		resp, err := http.Get(service.Endpoint + injson.SummaryJson)
		if err != nil {
			log.Println(err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println(resp.Status)
			continue
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}

		var summary injson.Summary
		err = json.Unmarshal(b, &summary)
		if err != nil {
			log.Println(err)
			continue
		}
		if brief {
			log.Printf("%s: %s", service.Name, summary.BriefStatus())
		} else {
			log.Printf("%s:\n%s", service.Name, summary.FullStatus("\t", "\n"))
		}
	}
}

func (s Services) FetchServices(brief bool, interval time.Duration) {
	wait := interval / time.Second
	for {
		s.PollServices(brief)
		log.Printf("Waiting for %d seconds...", wait)
		time.Sleep(interval)
	}
}

func (s Services) String() string {
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
