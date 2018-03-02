package controllers

type Token struct {

	Token string `json:"token"`
}

type Row struct {
	Createtime string `json:"createtime"`
	Name string `json:"name"`
	Recordtype string `json:"recordtype"`
	Size string `json:"size"`
	Record_id string	`json:"record_id"`
	Fileextname string `json:"fileextname"`
 }

type Record struct {
	Total int `json:"total"`
	Rows  []Row `json:"rows"`
}