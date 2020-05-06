package main

import (
	"github.com/hrdkgmz/dbWrapper/cache"
)
func main(){

	cache.SetRedisParas("172.28.62.228:6379","",1,100,100)
	//c:=cache.GetInstance()

	//list:=[]string{"4","5","6"}
	////val,err:= cache.GetAllCCDetail()
	//err:= cache.SetTaskDataList(123,list,300)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//
	//val,err:= cache.GetTaskDataListByRange(123,0,-1)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(val)
	//
	//len,err:= cache.GetTaskDataListLen(123)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(len)
}