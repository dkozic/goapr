package aprclient

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sclevine/agouti"
)

// AprClient is WebDriver based client to www.apr.rs
type AprClient struct {
	url      string
	headless bool
}

// NewAprClient creates new AprClient
func NewAprClient(url string, headless bool) AprClient {
	var c AprClient
	c.url = url
	c.headless = headless
	return c
}

func (client AprClient) createAndStartDriver() (*agouti.WebDriver, error) {
	var driver *agouti.WebDriver

	options := []agouti.Option{}

	options = append(options, agouti.Debug)

	chromeBin := os.Getenv("GOOGLE_CHROME_SHIM")
	if chromeBin != "" {
		o := agouti.ChromeOptions("binary", chromeBin)
		options = append(options, o)
	}

	if client.headless {
		o := agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		})
		options = append(options, o)
	}

	driver = agouti.ChromeDriver(options...)
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
