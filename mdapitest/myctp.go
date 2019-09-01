package main

/*
#include "ThostFtdcUserApiStruct.h"
extern void OnMktgo(struct CThostFtdcDepthMarketDataField* pMkt);
*/
import "C"
import (
	"CTP/mdapi"
	"fmt"
	"os"
	"os/signal"
	"time"
)

//export  OnMktgo
func OnMktgo(pMkt *C.struct_CThostFtdcDepthMarketDataField) {
	fmt.Println("行情回调xxx:", float64(pMkt.HighestPrice))
}
func main() {
	//"net/http"
	//_ "net/http/pprof"
	//go func() {
	//		http.ListenAndServe("localhost:6060", nil)
	//}()
	md := mdapi.NewMdapi()
	//defer md.FreeMdapi()
	//md.ConnectMdBgo("\\", "tcp://180.168.146.187:10110")
	md.ConnectMdgo("\\", "tcp://180.168.146.187:10131")
	md.InitMdgo()
	//md.JoinMdgo()
	//fmt.Println("step1", md.GetVersionMdMdgo())
	time.Sleep(time.Duration(2) * time.Second)

	loginfield := mdapi.CThostFtdcReqUserLoginFieldgo{BrokerID: "9999", UserID: "xxx", Password: "xxx"}
	md.LoginMdgo(loginfield, 0)
	time.Sleep(time.Duration(2) * time.Second)
	
	syb := []string{"rb1910", "ni1910"}
	md.RegMarketDataMdgo(mdapi.OnMkt(C.OnMktgo))
	md.SubscribeMarketDataMdgo(syb)
	

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("按Ctrl+C可退出程序")
	<-quit
	fmt.Println("主程序退出")
}
