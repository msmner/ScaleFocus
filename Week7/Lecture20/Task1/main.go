package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	bartender := CocktailBartender{}
	bartender.Start()
}

const (
	link = "https://www.thecocktaildb.com/api/json/v1/1/search.php"
)

type CocktailBartender struct{}

type Cocktail struct {
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	Instructions string `json:"strInstructions"`
}

func (d Drink) String() string {
	return fmt.Sprintf("%v", d.Instructions)
}

func (bartender *CocktailBartender) Start() {
	for {
		fmt.Println("What would you like to drink?")
		var input string
		fmt.Scanln(&input)
		if input == "nothing" {
			return
		}

		u, err := url.Parse(link)
		if err != nil {
			fmt.Printf("Error parsing link: %v", err)
		}

		q := u.Query()
		q.Add("s", input)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			fmt.Printf("Error building the request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Error getting the response: %v", err)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error reading the body: %v", err)
		}

		var result Cocktail
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("Error deserializing the body: %v", err)
		}

		str := fmt.Sprintf("%#v", result.Drinks[0].Instructions)
		lines := strings.Split(str, ".")

		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
