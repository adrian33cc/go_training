package countries

import "fmt"

var myCountry string

var collection []string


func Add(country string) {
	collection = append(collection, country)
}

func SetMyCountry(country string)  {
	for _, c:= range collection {
		if c == country{
			myCountry = country
		}
	}
}	

func List() {

	for i, country := range collection {
		myCountryMsg := ""

		if country == myCountry{
			myCountryMsg = "[my country]"
		
		}

		fmt.Printf("%d. %s %s\n", i+1,country,myCountryMsg)
	}
}
