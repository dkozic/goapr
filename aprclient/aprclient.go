package aprclient

import (
	"log"
	"strings"

	"github.com/sclevine/agouti"
)

type aprclient struct {
	url      string
	headless bool
}

func New(url string) aprclient {
	var c aprclient
	c.url = url
	c.headless = false
	return c
}

func (client aprclient) Headles() bool {
	return client.headless
}

func (client aprclient) SetHeadles(headless bool) {
	client.headless = headless
}

func parseLineByColon(in string) string {

	out := ""
	if colonPos := strings.Index(in, ":"); colonPos != -1 {
		out = in[colonPos+1:]
		out = strings.Trim(out, " ")
	}
	return out
}

func closeDriver(driver *agouti.WebDriver) {
	if err := driver.Stop(); err != nil {
		log.Printf("Failed to close driver: %v\n", err)
	}
}
