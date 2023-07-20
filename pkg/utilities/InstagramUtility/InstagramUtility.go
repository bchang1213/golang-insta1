package InstagramUtility

import (
	"dep-ex/pkg/entities/InstaEntity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetTag(tag string) InstaEntity.InstagramResponse {
	// todo: get JSON of {handle} from the Instagram public json route
	// example url: https://www.instagram.com/graphql/query/?query_hash=298b92c8d7cad703f7565aa892ede943&variables={"tag_name":"dogs","first":12}
	// parse json into an object that you define.
	// optional hint: https://mholt.github.io/json-to-go/
	// replace return type interface{} with this object
	// return the new object
	// NOTE: make sure to handle fails.
	type Payload struct {
		Tag_Name string
		First    int
	}

	payload := &Payload{tag, 12}

	out, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get("https://www.instagram.com/graphql/query/?query_hash=298b92c8d7cad703f7565aa892ede943&variables=" + strings.ToLower(string(out)))
	if err != nil {
		log.Fatalln("Error during request to Instagram:", err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error during response from Instagram:", err)
	}
	//Convert the body to type string
	sb := string(body)

	result := GetPictures(sb)
	return result
}

func GetPictures(responseBody string) InstaEntity.InstagramResponse {

	var result InstaEntity.InstagramResponse
	err := json.Unmarshal([]byte(responseBody), &result)
	if err != nil {
		log.Fatalln("Error with data response:", err)
	}

	return result
}
