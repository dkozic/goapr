package aprclient

import (
	"encoding/json"
	"log"
)

const searchURL = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"

func ExampleAprClient_SearchByRegistryCode() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByRegistryCode("21180408"); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByRegistryCode_second() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByRegistryCode(""); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByRegistryCode_third() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByRegistryCode("123"); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByRegistryCode_fourth() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByRegistryCode("99999999"); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByBusinessName() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByBusinessName("sedecom"); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByBusinessName_second() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByBusinessName(""); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleAprClient_SearchByBusinessName_third() {

	client := NewAprClient(searchURL, false)

	if sr, err := client.SearchByBusinessName("aa"); err != nil {
		log.Printf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Printf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}
