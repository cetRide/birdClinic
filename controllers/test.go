package controllers

import (
	"encoding/json"
	"fmt"
)

type TestController struct {
	MainController
}

type MgnErrorJson struct {
	Response string
	Name     string
}

var myresponsejson MgnErrorJson

func (this *TestController) Get() {
	this.Symptom_Checker() 
}

func (this *TestController) Post() {

	// fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string][]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	choices := dataform["data"]
	fmt.Println(choices)
	var choice []string
	for i , _ := range choices{
		fmt.Println(choices[i])
		choice = append(choice, choices[i])
	}
	fmt.Println(choice)
	fmt.Println(len(choice))
	myresponsejson.Response = "/" 
	obj, _ := json.Marshal(myresponsejson)
	this.Ctx.Output.Header("Content-Type", "application/json")
	this.Ctx.Output.Body(obj)

}
