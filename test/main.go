package main

import (
	"fmt"
	"github.com/hrdkgmz/dbWrapper/cache"
)
func main(){

	cache.SetRedisParas("172.28.62.228:6379","",1,100,100)

	//orgKeys, _ := cache.GetInstance().KeysWitchDb(0, "b_org_info:hash:org_name:*")
	//fmt.Println(orgKeys)

	fmt.Println(cache.GetCCIDsByCCName("example3"))
	fmt.Println("\n")

	//fmt.Println(cache.GetAllCCNames())
	//fmt.Println("\n")
	//
	//org,_:= cache.GetCCDetailsByCCName("example3")
	//for _,v:=range org{
	//	fmt.Println(v)
	//}
	//
	//chann,_:=cache.GetChannelByPeer("peer0.org1.example.com")
	//fmt.Println(chann, "\n")
	//
	//
	//fmt.Println("\n")
	//peer,_:=cache.GetAllPeerCCDetail()
	//for _,v:=range peer{
	//	fmt.Println(v)
	//}
	//fmt.Println("\n")
	//ord,_:=cache.GetAllChannelCCDetail()
	//for _,v:=range ord{
	//	fmt.Println(v)
	//}
	//fmt.Println("\n")
	//
	//pp,_:=cache.GetPeerCCDetailsByCCID("99")
	//for _,v:=range pp{
	//	fmt.Println(v)
	//}
	//fmt.Println("\n")
	//
	//cc,_:=cache.GetChanCCDetailsByCCID("12345")
	//for _,v:=range cc{
	//	fmt.Println(v)
	//}
	//fmt.Println("\n")

	//c:=cache.GetInstance()
	//db.SetMysqlParas("172.28.62.228:3306", "bctest", "root","root", "utf8",2000,1000)
	//rand.Seed(time.Now().Unix())
	//var wg sync.WaitGroup
	//max:=10000
	//wg.Add(max)
	//for i:=0;i<max;i++{
	//	go func() {
	//		_,err:=db.GetInstance().Insert("INSERT INTO l_tx_data (tx_id, blk_no, tx_time, db_time, tx_type) VALUES (?,?,?,?,?)", strconv.Itoa(rand.Int()), rand.Intn(1000), "2020-05-19 11:22:33", "2020-05-19 11:22:33",4)
	//		if err!=nil{
	//			fmt.Println(err)
	//		}
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()

















	//dd,err:= cache.GetPeerJoinTime("mychannel","peer0.org1.example.com")
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(dd)
	//
	//
	//list:=[]string{"4","5","6"}
	//val,err:= cache.GetAllChanDetail()
	//if err!=nil{
	//	fmt.Errorf("%s",err)
	//}
	//fmt.Println(val)
	//
	//var wg sync.WaitGroup
	//fmt.Println("开始：",time.Now())
	//for i:=0;i<100;i++{
	//	if i%2==0{
	//		go func(i int) {
	//			wg.Add(1)
	//			val, err := cache.GetAllChanDetail()
	//			if err != nil {
	//				fmt.Errorf("%s", err)
	//			}
	//			fmt.Println(i,val)
	//			wg.Done()
	//			//fmt.Println("线程")
	//		}(i)
	//	}else{
	//		go func(i int64) {
	//			wg.Add(1)
	//			err := cache.SetTaskDataList(i,list,1800)
	//			if err != nil {
	//				fmt.Errorf("%s", err)
	//			}
	//
	//			fmt.Println("插入成功：",i)
	//			wg.Done()
	//		}(int64(i))
	//	}
	//
	//}
	//wg.Wait()
	//fmt.Println("结束",time.Now())
















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