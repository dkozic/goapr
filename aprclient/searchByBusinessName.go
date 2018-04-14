package aprclient

import (
	"fmt"
)

// SearchByBusinessName searches by name of registered entity
func (client AprClient) SearchByBusinessName(businessName string) ([]SearchByBusinessNameResult, error) {

	driver, err := client.createAndStartDriver()
	if err != nil {
		return nil, err
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

	actualURL, err := page.URL()
	if err != nil {
		return nil, fmt.Errorf("failed to get page URL: %v", err)
	}

	expectedURL := client.url
	if actualURL != expectedURL {
		return nil, fmt.Errorf("expected URL to be %s but got %s", expectedURL, actualURL)
	}

	title, err := page.Title()
	if err != nil {
		return nil, fmt.Errorf("failed to get title: %v", err)
	}
	if title != "Претрага правних лица и предузетника" {
		return nil, fmt.Errorf("wrong title: %s", title)
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

	cnt, err := trs.Count()
	if err != nil {
		return nil, fmt.Errorf("failed to get count of tr elements: %v", err)
	}
	for i := 0; i < cnt; i++ {
		var r SearchByBusinessNameResult
		tr := trs.At(i)

		td1 := tr.FirstByXPath("./td[1]")
		t1, err := td1.Text()
		if err != nil {
			return nil, fmt.Errorf("failed to get text of td[1]: %v", err)
		}
		r.LegalForm = t1

		td2 := tr.FirstByXPath("./td[2]")
		t, err := td2.Text()
		if err != nil {
			return nil, fmt.Errorf("failed to get text of td[2]: %v", err)
		}
		r.RegistryCode = t

		td3 := tr.FirstByXPath("./td[3]")
		t3, err := td3.Text()
		if err != nil {
			return nil, fmt.Errorf("failed to get text of td[3]: %v", err)
		}
		r.BusinessName = t3

		td4 := tr.FirstByXPath("./td[4]")
		t4, err := td4.Text()
		if err != nil {
			return nil, fmt.Errorf("failed to get text of td[4]: %v", err)
		}
		r.Status = t4

		results = append(results, r)
	}

	return results, nil

}
