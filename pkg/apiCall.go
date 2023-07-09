package scoreplay

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ApiCall[T ScoreplayResponseType] (route string, apiKey string) (*T, error) {
	var response T

	route = route + ".json?api_key=" + apiKey

	res, err := http.Get(route); if err != nil {
		log.Println(fmt.Sprintf("[ApiCall] GET failure for route %s", route), err)
		return &response, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body); if err != nil {
		log.Println(fmt.Sprintf("[ApiCall] Read Body failure for route %s", route), err)
		return &response, err
	}

	err = json.Unmarshal(body, &response); if err != nil {
		log.Println(fmt.Sprintf("[ApiCall] Unmarshal failure for route %s", route), err)
		return &response, err
	}

	return &response, err
}
