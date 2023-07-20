package JsonReader

import (
	"dep-ex/pkg/entities/InstaEntity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ParseJSON(file string) InstaEntity.InstagramResponse {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened JSON")
	fmt.Println(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result InstaEntity.InstagramResponse

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatalln("Error with data response:", err)
	}

	return result
}
