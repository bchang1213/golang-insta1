package index

import (
	"dep-ex/pkg/utilities/JsonReader"
	TemplateUtility "dep-ex/pkg/utilities/templateUtility"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	TemplateUtility.RenderTemplate("landing.gohtml", nil, w)
	return
}

func Instagram(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	// handle := r.Form.Get("tag")
	//TODO: call utility to get Instagram uploads from handler
	// log.Println("handle:", handle)

	//THIS LINE MUST BE COMMENTED BACK IN!!! I only removed it because the API doesnt work for me.
	// result := InstagramUtility.GetTag(handle)

	//Because the API doesnt work for me
	result := JsonReader.ParseJSON("../dep-ex/internal/handlers/index/example-response.json")
	TemplateUtility.RenderTemplate("instagram.gohtml", result, w)
	return
}
