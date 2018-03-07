package aprclient

type SearchByRegistryCodeResult struct {
	MainData             MainData
	BusinessNames        BusinessNames
	LegalRepresentatives LegalRepresentatives
	Members              []Member
	Addresses            Addresses
	BusinessData         BusinessData
}

type BusinessNames struct {
	BusinessName              string
	ShortBusinessName         string
	BusinessNameTranslations  []string
	ShortBusinessTranslations []string
}

type MainData struct {
	BusinessName      string
	Status            string
	RegistryCode      string
	LegalForm         string
	Headquarter       Headquarter
	EstablishmentDate string
	PIB               string
}

type Headquarter struct {
	Municipality     string
	City             string
	StreetAndHouseNo string
}

type LegalRepresentatives struct {
	Persons   []LegalRepresentativePerson
	Companies []LegalRepresentativeCompany
}

type LegalRepresentativePerson struct {
	Name        string
	JMBG        string
	Role        string
	StandsAlone string
}

type LegalRepresentativeCompany struct {
	Name string
}

type Member struct {
	Name              string
	RegistrationCode  string
	JMBG              string
	Country           string
	Type              string
	Share             string
	SubscribedCapital string
	PaidCapital       string
}

type Addresses struct {
	HeadquarterAddress Address
	PostalAddress      Address
	EmailAddress       string
}

type Address struct {
	Municipality    string
	City            string
	Street          string
	HouseNo         string
	FloorFlatLetter string
	AdditionalNote  string
}

type BusinessData struct {
	Establishment Establishment
	Duration      Duration
	MainActivity  Activity
	Ident         Ident
	Contact       Contact
}

type Establishment struct {
	RegistrationDate string
}

type Duration struct {
	ValidUntil string
}

type Activity struct {
	Code string
	Name string
}

type Ident struct {
	PIB  string
	RZZO string
	PIO  string
}

type Contact struct {
	Phone1 string
	Phone2 string
	Fax    string
	Web    string
}

type SearchByBusinessNameResult struct {
	LegalForm    string
	RegistryCode string
	BusinessName string
	Status       string
}
