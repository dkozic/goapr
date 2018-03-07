package aprclient

import (
	"encoding/json"
	"log"
)

const searchUrl = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"

func ExampleSearchByBusinessCode() {

	client := New(searchUrl)

	if sr, err := client.SearchByRegistryCode("21180408"); err != nil {
		log.Fatalf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Fatalf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}

func ExampleSearchByRegistryName() {

	client := New(searchUrl)

	if sr, err := client.SearchByBusinessName("asw"); err != nil {
		log.Fatalf("Unsuccesful search: %v", err)
	} else {
		if b, err := json.MarshalIndent(sr, "", "  "); err != nil {
			log.Fatalf("Error marshalling result: %+v", sr)
		} else {
			log.Printf("Search result: \n%s", string(b))
		}
	}

	//Output:
}
