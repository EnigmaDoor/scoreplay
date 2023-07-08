package scoreplay

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ApiCall[T ScoreplayResponseType] (route string, apiKey string) (*T, error) {
	var response T

	route = route + ".json?api_key=" + apiKey
	fmt.Println("myroute = ", route)

	res, err := http.Get(route); if err != nil {
		// todo err handling
		return &response, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body); if err != nil {
		// todo err handling
		return &response, err
	}

	// fmt.Println("==============>", string(body[:]))

	json.Unmarshal(body, &response)

	// fmt.Println("==============>", response)
	return &response, err
}
