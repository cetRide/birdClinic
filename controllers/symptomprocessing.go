package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type SymptomController struct {
	MainController
}


type MgnErrorJson struct {
	Response string
	Results  string
}

var myresponsejson MgnErrorJson

var Diseases orm.ParamsList
var Symptoms string

func (this *SymptomController) Get() {
	this.Symptom_Checker()
}

func (this *SymptomController) Post() {

	fmt.Println(string(this.Ctx.Input.RequestBody))
	var dataform map[string][]string
	json.Unmarshal(this.Ctx.Input.RequestBody, &dataform)
	symptoms := dataform["data"]
	fmt.Println(symptoms)
	var choice []string
	for i, _ := range symptoms {
		fmt.Println(symptoms[i])
		choice = append(choice, symptoms[i])
	}

	var list orm.ParamsList
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("symptom").Filter("name__in", choice).ValuesFlat(&list, "disease_id")

	if err == nil {
		fmt.Println("My disease id are", list)
		var lists orm.ParamsList
		o.QueryTable("disease").Filter("disease_id__in", list).ValuesFlat(&lists, "name")
		Diseases = lists
		fmt.Println("My disease names are", Diseases)

		myresponsejson.Response = "/potential_risks"
		obj, _ := json.Marshal(myresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)

	} else {
		if err != nil {
			myresponsejson.Results = "invalid"
		}
		myresponsejson.Response = "/"
		obj, _ := json.Marshal(myresponsejson)
		this.Ctx.Output.Header("Content-Type", "application/json")
		this.Ctx.Output.Body(obj)
	}
}
