package aprclient

import (
	"fmt"
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
	c.headless = true
	return c
}

func (client aprclient) Headles() bool {
	return client.headless
}

func (client aprclient) SetHeadles(headless bool) {
	client.headless = headless
}

func (client aprclient) createAndStartDriver() (*agouti.WebDriver, error) {
	var driver *agouti.WebDriver
	if client.headless {
		cdo1 := agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
		})
		driver = agouti.ChromeDriver(cdo1)
	} else {
		driver = agouti.ChromeDriver()
	}
	if err := driver.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Chrome driver: %v", err)
	}
	return driver, nil
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
