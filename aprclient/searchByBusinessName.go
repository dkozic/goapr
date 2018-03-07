package aprclient

import (
	"fmt"

	"github.com/sclevine/agouti"
)

func (client aprclient) SearchByBusinessName(businessName string) ([]SearchByBusinessNameResult, error) {

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Chrome driver: %v", err)
	}

	defer closeDriver(driver)

	page, err := driver.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %v", err)
	}

	if err := page.SetImplicitWait(10000); err != nil {
		return nil, fmt.Errorf("unable to set implicit wait timeout: %v", err)
	}

	if err := page.Navigate(client.url); err != nil {
		return nil, fmt.Errorf("failed to open page: %v", err)
	}

	actualUrl, err := page.URL()
	if err != nil {
		return nil, fmt.Errorf("failed to get page URL: %v", err)
	}

	expectedUrl := client.url
	if actualUrl != expectedUrl {
		return nil, fmt.Errorf("expected URL to be %s but got %s", expectedUrl, actualUrl)
	}

	if title, err := page.Title(); err != nil {
		return nil, fmt.Errorf("failed to get title: %v", err)
	} else {
		if title != "Претрага правних лица и предузетника" {
			return nil, fmt.Errorf("wrong title: %s", title)
		}
	}

	forms := page.AllByXPath("//html/body/form[@action='/ObjedinjenePretrage/Search/SearchResult']")
	form := forms.At(1)

	inputTypeRadio := form.FirstByXPath(".//input[@type='radio' and @id='rdbtnSelectInputType' and @value='poslovnoIme']")
	if err := inputTypeRadio.Click(); err != nil {
		return nil, fmt.Errorf("failed to click in inputTypeRadio: %v", err)
	}

	nameTxt := form.FirstByName("SearchByNameString")
	if err := nameTxt.Click(); err != nil {
		return nil, fmt.Errorf("failed to click in name: %v", err)
	}
	if err := nameTxt.Fill(businessName); err != nil {
		return nil, fmt.Errorf("failed to fill in name: %v", err)
	}

	if err := form.Submit(); err != nil {
		return nil, fmt.Errorf("failed to submit: %v", err)
	}

	results := []SearchByBusinessNameResult{}

	table := page.FirstByClass("ContentTable")
	trs := table.AllByXPath(".//tr[@class='ContentTableRowData']")

	if cnt, err := trs.Count(); err != nil {
		return nil, fmt.Errorf("failed to get count of tr elements: %v", err)
	} else {
		for i := 0; i < cnt; i++ {
			var r SearchByBusinessNameResult
			tr := trs.At(i)

			td1 := tr.FirstByXPath("./td[1]")
			if t, err := td1.Text(); err != nil {
				return nil, fmt.Errorf("failed to get text of td[1]: %v", err)
			} else {
				r.LegalForm = t
			}

			td2 := tr.FirstByXPath("./td[2]")
			if t, err := td2.Text(); err != nil {
				return nil, fmt.Errorf("failed to get text of td[2]: %v", err)
			} else {
				r.RegistryCode = t
			}

			td3 := tr.FirstByXPath("./td[3]")
			if t, err := td3.Text(); err != nil {
				return nil, fmt.Errorf("failed to get text of td[3]: %v", err)
			} else {
				r.BusinessName = t
			}

			td4 := tr.FirstByXPath("./td[4]")
			if t, err := td4.Text(); err != nil {
				return nil, fmt.Errorf("failed to get text of td[4]: %v", err)
			} else {
				r.Status = t
			}

			results = append(results, r)
		}
	}

	return results, nil

}
