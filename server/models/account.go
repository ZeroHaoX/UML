package models

import (
	
)

type Account struct{
	Profit float32	`json:"profit"`
	Turnover float32	`json:"turnover"`
}

func GetAccount(year int,month int)(account *Account,err error){
	
	return
}

func AccountList()(accounts []Account,err error){

	return
}