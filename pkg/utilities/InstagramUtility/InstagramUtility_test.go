package InstagramUtility_test

import (
	"dep-ex/pkg/utilities/InstagramUtility"
	"log"
	"testing"
)

func TestGetTag(t *testing.T) {
	//TODO: write test
	// log.Println("handle:", t)
	instaData := InstagramUtility.GetTag("test")
	if instaData.Data.Hashtag.Name != "test" {
		log.Println("DATA WAS BAD!")
		t.Errorf("Value of Instagram response was invalid.")
	}

}
