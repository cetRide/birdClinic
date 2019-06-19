package controllers

import (
	// "encoding/json"
	"fmt"
	// "poultry/models"
	"github.com/astaxie/beego/orm"
)

type FindProController struct {
	MainController
}

type Pro struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

var Pros []Pro

var Search string
var Len int

func (this *FindProController) FindPro() {

	this.find_pro("findPro")
	if this.Ctx.Input.Method() == "POST" {
		Pros = nil
		search := this.GetString("search")
		Search = search
		//query the results
		o := orm.NewOrm()
		o.Using("default")

		_, err := o.Raw("select first_name, last_name, email, phone from pro where username=?", search).QueryRows(&Pros)
		if err != nil {
			// No result
			fmt.Printf("Not row found")
			return
		}
		if err == nil && len(Pros) > 0 {
			Len = len(Pros)
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return

		}

		_, err = o.Raw("select first_name, last_name, email, phone from pro where email=?", search).QueryRows(&Pros)
		if err != nil {
			// No result
			fmt.Printf("Not row found")
			return
		}
		if err == nil && len(Pros) > 0 {
			Len = len(Pros)
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return

		}
		_, err = o.Raw("select first_name, last_name, email, phone from pro where district=?", search).QueryRows(&Pros)
		if err != nil {
			// No result
			fmt.Printf("Not row found")
			return
		}
		if err == nil && len(Pros) > 0 {
			Len = len(Pros)
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return

		}

		_, err = o.Raw("select first_name, last_name, email, phone from pro where county=?", search).QueryRows(&Pros)
		if err != nil {
			// No result
			fmt.Printf("Not row found")
			return
		}
		if err == nil && len(Pros) > 0 {
			Len = len(Pros)
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return

		}

		_, err = o.Raw("select first_name, last_name, email, phone from pro where country=?", search).QueryRows(&Pros)
		if err != nil {
			// No result
			fmt.Printf("Not row found")
			return
		}
		if err == nil && len(Pros) > 0 {
			Len = len(Pros)
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return

		}

		if len(Pros) == 0 {
			Len = 0
			fmt.Println("Results", Pros)
			this.Redirect("/search-results", 302)
			return
		}

	}
	Len = len(Pros)
	fmt.Println("The length", Len)
}
