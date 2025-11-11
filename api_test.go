package rin

import (
	"net/http"
	"testing"
)

func TestAPICall(t *testing.T){
	api := NewApi("http://httpbin.org")
	router:=NewRouter()
	router.RegisterFunc(200, func(resp *http.Response,  _ interface{}) error {
		return nil
	})
	res := NewRessource("/get", "GET", router)
	api.AddRessource("get", res)
	if err := api.Call("get", nil, nil); err!=nil{
		t.Fail()
	}
		resources := api.RessourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}

}
func TestAPIAuth(t *testing.T) {
	api := NewApi("https://github.com/api")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	res := NewRessource("/basic-auth/{{.user}}/{{.pass}}", "GET", router)
	api.AddRessource("basicauth", res)
	api.SetAuth(&AuthBasic{
		Username: "user",
		Password: "passw0rd",
	})
	if err := api.Call("basicauth", map[string]string{
		"user": "user",
		"pass": "passw0rd",
	}); err != nil {
		t.Fail()
	}
}
