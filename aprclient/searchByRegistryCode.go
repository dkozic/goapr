package aprclient

import (
	"fmt"
	"strings"
)

// SearchByRegistryCode searches by register code of the company
func (client AprClient) SearchByRegistryCode(registryCode string) (SearchByRegistryCodeResult, error) {

	var srs SearchByRegistryCodeResult

	driver, err := client.createAndStartDriver()
	if err != nil {
		return srs, err
	}

	defer closeDriver(driver)

	page, err := driver.NewPage()
	if err != nil {
		return srs, fmt.Errorf("failed to create page: %v", err)
	}

	if err := page.SetImplicitWait(10000); err != nil {
		return srs, fmt.Errorf("unable to set implicit wait timeout: %v", err)
	}

	if err := page.Navigate(client.url); err != nil {
		return srs, fmt.Errorf("failed to open page: %v", err)
	}

	actualURL, err := page.URL()
	if err != nil {
		return srs, fmt.Errorf("failed to get page URL: %v", err)
	}

	expectedURL := client.url
	if actualURL != expectedURL {
		return srs, fmt.Errorf("expected URL to be %s but got %s", expectedURL, actualURL)
	}

	title, err := page.Title()
	if err != nil {
		return srs, fmt.Errorf("failed to get title: %v", err)
	}
	if title != "Претрага правних лица и предузетника" {
		return srs, fmt.Errorf("wrong title: %s", title)
	}

	forms := page.AllByXPath("//html/body/form[@action='/ObjedinjenePretrage/Search/SearchResult']")
	form := forms.At(1)

	inputTypeRadio := form.FirstByXPath(".//input[@type='radio' and @id='rdbtnSelectInputType' and @value='mbr']")
	if err := inputTypeRadio.Click(); err != nil {
		return srs, fmt.Errorf("failed to click in inputTypeRadio: %v", err)
	}

	registryCodeTxt := form.FirstByName("SearchByRegistryCodeString")
	if err := registryCodeTxt.Click(); err != nil {
		return srs, fmt.Errorf("failed to click in registryCode: %v", err)
	}
	if err := registryCodeTxt.Fill(registryCode); err != nil {
		return srs, fmt.Errorf("failed to fill in registryCode: %v", err)
	}

	if err := form.Submit(); err != nil {
		return srs, fmt.Errorf("failed to submit: %v", err)
	}

	table := page.FirstByClass("ContentTable")
	link := table.FirstByXPath(".//a[contains(text(), 'детаљније')]")
	if err := link.Click(); err != nil {
		text1 := page.FirstByClass("field-validation-error")
		f1, err := text1.Text()
		if err != nil {
			text2 := page.FirstByClass("Message")
			f2, err := text2.Text()
			if err != nil {
				return srs, fmt.Errorf("failed to click on details")
			}
			return srs, fmt.Errorf("message exists: %s", f2)
		}
		return srs, fmt.Errorf("validation error exists: %s", f1)
	}

	// main data
	{
		mainDataP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Основни подаци')]]/div[@class='GroupContent']/p")
		t, err := mainDataP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get mainData text: %v", err)
		}
		srs.MainData = parseMainData(t)
	}

	// businessNames
	{
		businessNameP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Пословно име')]]/div[@class='GroupContent']")
		t, err := businessNameP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get business name text: %v", err)
		}
		srs.BusinessNames = parseBusinessNames(t)
	}

	//legal representatives
	{
		legalRepresentativesLink := page.FirstByLink("Законски заступници")
		if err := legalRepresentativesLink.Click(); err != nil {
			return srs, fmt.Errorf("failed to get link legalRepresentative: %v", err)
		}
		legalRepresentativePersonP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Физичка лица')]]/div[@class='GroupContent']")
		t, err := legalRepresentativePersonP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get legalRepresentative text: %v", err)
		}
		srs.LegalRepresentatives.Persons = parseLegalRepresentativePersons(t)
	}

	// members
	{
		membersLink := page.FirstByLink("Чланови")
		if err := membersLink.Click(); err != nil {
			return srs, fmt.Errorf("failed to get link members: %v", err)
		}
		membersP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Чланови')]]/div[@class='GroupContent']")
		t, err := membersP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get members text: %v", err)
		}
		srs.Members = parseMembers(t)
	}

	// address
	{
		adressesLink := page.FirstByLink("Подаци о адресама")
		if err := adressesLink.Click(); err != nil {
			return srs, fmt.Errorf("failed to get link adresses: %v", err)
		}
		hqAddressP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Адреса седишта')]]/div[@class='GroupContent']")
		t, err := hqAddressP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get headquarter address text: %v", err)
		}
		srs.Addresses.HeadquarterAddress = parseAddress(t)
	}

	{
		pAddressP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Адреса за пријем поште')]]/div[@class='GroupContent']")
		t, err := pAddressP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get postal address text: %v", err)
		}
		srs.Addresses.PostalAddress = parseAddress(t)
	}

	{
		eAddressP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Адреса за пријем електронске поште')]]/div[@class='GroupContent']")
		t, err := eAddressP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get email address text: %v", err)
		}
		srs.Addresses.EmailAddress = parseEmailAddress(t)
	}

	// businessData
	buinessDataLink := page.FirstByLink("Пословни подаци")
	if err := buinessDataLink.Click(); err != nil {
		return srs, fmt.Errorf("failed to get link business data: %v", err)
	}

	{
		establishementP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Подаци оснивања')]]/div[@class='GroupContent']")
		t, err := establishementP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get establishment data text: %v", err)
		}
		srs.BusinessData.Establishment = parseEstablishment(t)
	}

	{
		durationP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Време трајања')]]/div[@class='GroupContent']")
		t, err := durationP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get duration data text: %v", err)
		}
		srs.BusinessData.Duration = parseDuration(t)
	}

	{
		activityP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Претежна делатност')]]/div[@class='GroupContent']")
		t, err := activityP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get activity data text: %v", err)
		}
		srs.BusinessData.MainActivity = parseActivity(t)
	}

	{
		identP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Остали идентификациони подаци')]]/div[@class='GroupContent']")
		t, err := identP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get ident data text: %v", err)
		}
		srs.BusinessData.Ident = parseIdent(t)
	}

	{
		contactP := page.FirstByXPath("//div[@class='Group' and ./div[@class='GroupHeader' and contains(text(), 'Контакт подаци')]]/div[@class='GroupContent']")
		t, err := contactP.Text()
		if err != nil {
			return srs, fmt.Errorf("failed to get contact data text: %v", err)
		}
		srs.BusinessData.Contact = parseContact(t)
	}

	return srs, nil

}

func parseMainData(in string) (out MainData) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Пословно Име:") == 0 {
			out.BusinessName = parseLineByColon(line)
		} else if strings.Index(line, "Статус:") == 0 {
			out.Status = parseLineByColon(line)
		} else if strings.Index(line, "Матични број:") == 0 {
			out.RegistryCode = parseLineByColon(line)
		} else if strings.Index(line, "Правна форма:") == 0 {
			out.LegalForm = parseLineByColon(line)
		} else if strings.Index(line, "Седиште:") == 0 {
			out.Headquarter = parseHeadquarter(parseLineByColon(line))
		} else if strings.Index(line, "Датум оснивања:") == 0 {
			out.EstablishmentDate = parseLineByColon(line)
		} else if strings.Index(line, "ПИБ:") == 0 {
			out.PIB = parseLineByColon(line)
		}
	}
	return out
}

func parseHeadquarter(in string) (out Headquarter) {

	lines := strings.Split(in, " | ")
	for _, line := range lines {
		if strings.Index(line, "Општина:") == 0 {
			out.Municipality = parseLineByColon(line)
		} else if strings.Index(line, "Место:") == 0 {
			out.City = parseLineByColon(line)
		} else if strings.Index(line, "Улица и број:") == 0 {
			out.StreetAndHouseNo = parseLineByColon(line)
		}
	}
	return out
}

func parseLegalRepresentativePersons(in string) (out []LegalRepresentativePerson) {

	lines := strings.Split(in, "\n")
	var person LegalRepresentativePerson
	for _, line := range lines {
		if strings.Index(line, "Име Презиме:") == 0 {
			person.Name = parseLineByColon(line)
		} else if strings.Index(line, "Јмбг/Лични број:") == 0 {
			person.JMBG = parseLineByColon(line)
		} else if strings.Index(line, "Функција:") == 0 {
			person.Role = parseLineByColon(line)
		} else if strings.Index(line, "Самостално заступа:") == 0 {
			person.StandsAlone = parseLineByColon(line)
		}
	}
	out = []LegalRepresentativePerson{person}
	return
}

func parseMembers(in string) []Member {

	out := []Member{}
	lines := strings.Split(in, "\n")
	var member *Member
	for _, line := range lines {
		if strings.Index(line, "Име Презиме:") == 0 || strings.Index(line, "Пословно име:") == 0 {
			if member != nil {
				out = append(out, *member)
			}
			member = new(Member)
			member.Name = parseLineByColon(line)
		} else if strings.Index(line, "Јмбг/Лични број:") == 0 {
			member.JMBG = parseLineByColon(line)
		} else if strings.Index(line, "Матични број:") == 0 {
			member.RegistrationCode = parseLineByColon(line)
		} else if strings.Index(line, "Земља:") == 0 {
			member.Country = parseLineByColon(line)
		} else if strings.Index(line, "Тип:") == 0 {
			member.Type = parseLineByColon(line)
		} else if strings.Index(line, "Удео:") == 0 {
			member.Share = parseLineByColon(line)
		}
	}
	if member != nil {
		out = append(out, *member)
	}
	return out
}

func parseAddress(in string) (out Address) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Општина:") == 0 {
			out.Municipality = parseLineByColon(line)
		} else if strings.Index(line, "Место:") == 0 {
			out.City = parseLineByColon(line)
		} else if strings.Index(line, "Улица:") == 0 {
			out.Street = parseLineByColon(line)
		} else if strings.Index(line, "Број:") == 0 {
			out.HouseNo = parseLineByColon(line)
		} else if strings.Index(line, "Спрат, број стана, слово:") == 0 {
			out.FloorFlatLetter = parseLineByColon(line)
		} else if strings.Index(line, "Додатни опис:") == 0 {
			out.AdditionalNote = parseLineByColon(line)
		}
	}
	return out
}

func parseEmailAddress(in string) (out string) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Е-пошта:") == 0 {
			out = parseLineByColon(line)
		}
	}
	return out
}

func parseEstablishment(in string) (out Establishment) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Датум регистрације:") == 0 {
			out.RegistrationDate = parseLineByColon(line)
		}
	}
	return out
}

func parseDuration(in string) (out Duration) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Трајање ограничено до:") == 0 {
			out.ValidUntil = parseLineByColon(line)
		}
	}
	return out
}

func parseActivity(in string) (out Activity) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Шифра делатности:") == 0 {
			out.Code = parseLineByColon(line)
		} else if strings.Index(line, "Назив делатности:") == 0 {
			out.Name = parseLineByColon(line)
		}
	}
	return out
}

func parseIdent(in string) (out Ident) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Порески идентификациони број ПИБ:") == 0 {
			out.PIB = parseLineByColon(line)
		} else if strings.Index(line, "РЗЗО број:") == 0 {
			out.RZZO = parseLineByColon(line)
		} else if strings.Index(line, "Пио број:") == 0 {
			out.PIO = parseLineByColon(line)
		}
	}
	return out
}

func parseContact(in string) (out Contact) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Телефон 1:") == 0 {
			out.Phone1 = parseLineByColon(line)
		} else if strings.Index(line, "Телефон 2:") == 0 {
			out.Phone2 = parseLineByColon(line)
		} else if strings.Index(line, "Факс:") == 0 {
			out.Fax = parseLineByColon(line)
		} else if strings.Index(line, "Интернет адреса:") == 0 {
			out.Web = parseLineByColon(line)
		}
	}
	return out
}

func parseBusinessNames(in string) (out BusinessNames) {

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.Index(line, "Пословно име:") == 0 {
			out.BusinessName = parseLineByColon(line)
		} else if strings.Index(line, "Скраћено пословно име:") == 0 {
			out.ShortBusinessName = parseLineByColon(line)
		}
	}
	return out
}
