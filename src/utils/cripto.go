package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func EtheriumApi(ch chan<- interface{}, hash string) {

	base, err := url.Parse("https://api.etherscan.io/api")

	// Path params
	//base.Path += "this will get automatically encoded"

	// Query params
	params := url.Values{}
	params.Add("module", "account")
	params.Add("action", "txlist")
	params.Add("address", "0x08a92ad30602fd010f3aef93f419e88c06045f98")
	params.Add("startblock", "0")
	params.Add("endblock", "99999999")
	params.Add("page", "1")
	params.Add("offset", "1000")
	params.Add("sort", "asc")
	params.Add("apikey", "H47GQJQJJZUR7VSH21SA3Z5Z8C2UN8FD83")
	base.RawQuery = params.Encode()

	// Http request Json
	resp, err := http.Get(base.String())
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	//sb := string(body)
	byt := []byte(body)
	dat := make(map[string]interface{})
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	strs := dat["result"].([]interface{})
	var aux interface{}
	for k, v := range strs {

		if strs[k].(map[string]interface{})["hash"] == hash {
			aux = v
			break
		} else {
			aux = map[string]string{"message": "Hash incorrecto o no valido"}
		}

	}
	ch <- aux
}
