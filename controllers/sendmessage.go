package controllers

type WriteToProController struct{
	MainController
}

func (this *WriteToProController) Get(){
	this.sendMessage("sendmessge")
}

func(this *WriteToProController) Post(){
	

}