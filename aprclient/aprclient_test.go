package aprclient

import (
	"encoding/json"
	"log"
)

const searchUrl = "http://pretraga2.apr.gov.rs/ObjedinjenePretrage/Search/Search"

func ExampleSearchByRegistryCode1() {

	client := New(searchUrl, false)

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

func ExampleSearchByRegistryCode2() {

	client := New(searchUrl, false)

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

func ExampleSearchByRegistryCode3() {

	client := New(searchUrl, false)

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

func ExampleSearchByRegistryCode4() {

	client := New(searchUrl, false)

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

func ExampleSearchByBusinessName1() {

	client := New(searchUrl, false)

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

func ExampleSearchByBusinessName2() {

	client := New(searchUrl, false)

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

func ExampleSearchByBusinessName3() {

	client := New(searchUrl, false)

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
