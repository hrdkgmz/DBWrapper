package main

import (
	"fmt"
	"github.com/hrdkgmz/dbWrapper/cache"
	"sync"
	"time"
)
func main(){

	cache.SetRedisParas("172.28.62.228:6379","",1,100,100)
	//c:=cache.GetInstance()
	val,err:= cache.Get
	if err!=nil{
		fmt.Errorf("%s",err)
	}
	fmt.Println(val)


	list:=[]string{"4","5","6"}
	val,err:= cache.GetAllChanDetail()
	if err!=nil{
		fmt.Errorf("%s",err)
	}
	fmt.Println(val)

	var wg sync.WaitGroup
	fmt.Println("开始：",time.Now())
	for i:=0;i<100;i++{
		if i%2==0{
			go func(i int) {
				wg.Add(1)
				val, err := cache.GetAllChanDetail()
				if err != nil {
					fmt.Errorf("%s", err)
				}
				fmt.Println(i,val)
				wg.Done()
				//fmt.Println("线程")
			}(i)
		}else{
			go func(i int64) {
				wg.Add(1)
				err := cache.SetTaskDataList(i,list,1800)
				if err != nil {
					fmt.Errorf("%s", err)
				}

				fmt.Println("插入成功：",i)
				wg.Done()
			}(int64(i))
		}

	}
	wg.Wait()
	fmt.Println("结束",time.Now())
	//err:= cache.SetTaskDataList(123,list,300)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//
	//val,err:= cache.GetTaskDataListByRange(123,0,1)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(val)

	//len,err:= cache.GetTaskDataListLen(123)
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(len)
}