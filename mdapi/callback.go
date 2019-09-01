package mdapi

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lMDBRIDGE
#cgo LDFLAGS: -L . -lthostmduserapi_se
#include "MdBridge.h"
#include "stdbool.h"
#include "ThostFtdcUserApiStruct.h"
#include "string.h"
typedef char TThostFtdcBrokerIDType[11];
void OnMarketDataMdgo_cgo(struct CThostFtdcDepthMarketDataField *pMktgo);

void OnConnectedMdgo_cgo();

void OnDisconnectedMdgo_cgo(int nReason);

void OnHeartWarningMdgo_cgo(int nTimeLapse);

void OnLoginMdgo_cgo(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnLogoutMdgo_cgo(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnErrorMdgo_cgo(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnSubMarketDataMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnUnSubMarketDataMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnSubForQuoteRspMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnUnSubForQuoteRspMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void OnForQuoteRspMdgo_cgo(struct CThostFtdcForQuoteRspField *pForQuoteRsp);
*/
import "C"
import "fmt"
import "unsafe"

type OnMkt C.OnMarketDataMd

type CThostFtdcReqUserLoginFieldgo struct {
	///交易日
	TradingDay string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///密码
	Password string
	///用户端产品信息
	UserProductInfo string
	///接口端产品信息
	InterfaceProductInfo string
	///协议信息
	ProtocolInfo string
	///Mac地址
	MacAddress string
	///动态密码
	OneTimePassword string
	///终端IP地址
	ClientIPAddress string
	///登录备注
	LoginRemark string
	///终端IP端口
	ClientIPPort int
}

//export OnMarketDataMdgo
func OnMarketDataMdgo(pMkt *C.struct_CThostFtdcDepthMarketDataField) {
	fmt.Println("行情回调:", float64(pMkt.HighestPrice))
}

type Mdapi struct {
	mdapi C.pMdBridge
}

func NewMdapi() *Mdapi {
	return &Mdapi{
		C.NewMdBridge(),
	}
}

func (x *Mdapi) FreeMdapi() {
	C.FreeMdBridge(x.mdapi)
}

func (x *Mdapi) ConnectMdgo(path string, addr string) {
	C.ConnectMd(x.mdapi, C.CString(path), C.CString(addr), (C.OnConnected)(unsafe.Pointer(C.OnConnectedMdgo_cgo)))
}

func (x *Mdapi) RegDissConnMdgo() {
	C.regDissConn(x.mdapi, (C.OnDisconnected)(unsafe.Pointer(C.OnDisconnectedMdgo_cgo)))
}

func (x *Mdapi) RegErrorMdgo() {
	C.regError(x.mdapi, (C.OnError)(unsafe.Pointer(C.OnErrorMdgo_cgo)))
}
func (x *Mdapi) RegHeartWarningMdBgo() {
	C.regHeartWarning(x.mdapi, (C.OnHeartWarning)(unsafe.Pointer(C.OnHeartWarningMdgo_cgo)))
}
func (x *Mdapi) InitMdgo() {
	C.InitMd(x.mdapi)
}
func (x *Mdapi) GetVersionMdMdgo() string {
	ver := C.GetVersionMd(x.mdapi)
	return C.GoString(ver)
}
func (x *Mdapi) JoinMdgo() {
	C.JoinMd(x.mdapi)
}

func (x *Mdapi) LoginMdgo(pLogin CThostFtdcReqUserLoginFieldgo, nRequestID int) int {
	var cLogin C.struct_CThostFtdcReqUserLoginField

	C.strcpy(&cLogin.TradingDay[0], C.CString(pLogin.TradingDay))
	C.strcpy(&cLogin.BrokerID[0], C.CString(pLogin.BrokerID))
	C.strcpy(&cLogin.UserID[0], C.CString(pLogin.UserID))
	C.strcpy(&cLogin.Password[0], C.CString(pLogin.Password))
	C.strcpy(&cLogin.UserProductInfo[0], C.CString(pLogin.UserProductInfo))
	C.strcpy(&cLogin.InterfaceProductInfo[0], C.CString(pLogin.InterfaceProductInfo))
	C.strcpy(&cLogin.ProtocolInfo[0], C.CString(pLogin.ProtocolInfo))
	C.strcpy(&cLogin.MacAddress[0], C.CString(pLogin.MacAddress))
	C.strcpy(&cLogin.OneTimePassword[0], C.CString(pLogin.OneTimePassword))
	C.strcpy(&cLogin.ClientIPAddress[0], C.CString(pLogin.ClientIPAddress))
	C.strcpy(&cLogin.LoginRemark[0], C.CString(pLogin.LoginRemark))
	cLogin.ClientIPPort = C.int(pLogin.ClientIPPort)

	//pReqUserLoginField := C.struct_CThostFtdcReqUserLoginField{}

	result := C.LoginMd(x.mdapi, &cLogin, C.int(nRequestID), (C.OnLogin)(unsafe.Pointer(C.OnLoginMdgo_cgo)))
	return int(result)
}

func (x *Mdapi) GetTradeDayMd() string {
	TradeDay := C.GetTradeDayMd(x.mdapi)
	return C.GoString(TradeDay)
}

func (x *Mdapi) SubscribeMarketDataMdgo(symbol []string) int {
	ppInstrumentID := []*C.char{}
	for _, syb := range symbol {
		ppInstrumentID = append(ppInstrumentID, C.CString(syb))
	}
	count := cap(ppInstrumentID)
	result := C.SubscribeMarketDataMd(x.mdapi, &ppInstrumentID[0], C.int(count), (C.OnSubMarketData)(unsafe.Pointer(C.OnSubMarketDataMdgo_cgo)))
	return int(result)
}

func (x *Mdapi) UnSubscribeMarketDataMdgo(symbol []string) int {
	ppInstrumentID := []*C.char{}
	for _, syb := range symbol {
		ppInstrumentID = append(ppInstrumentID, C.CString(syb))
	}
	count := cap(ppInstrumentID)
	result := C.UnSubscribeMarketDataMd(x.mdapi, &ppInstrumentID[0], C.int(count), (C.OnUnSubMarketData)(unsafe.Pointer(C.OnUnSubMarketDataMdgo_cgo)))
	return int(result)
}

func (x *Mdapi) SubscribeForQuoteRspMdgo(symbol []string) int {
	ppInstrumentID := []*C.char{}
	for _, syb := range symbol {
		ppInstrumentID = append(ppInstrumentID, C.CString(syb))
	}
	count := cap(ppInstrumentID)
	result := C.SubscribeForQuoteRspMd(x.mdapi, &ppInstrumentID[0], C.int(count), (C.OnSubForQuoteRsp)(unsafe.Pointer(C.OnSubForQuoteRspMdgo_cgo)))
	return int(result)
}

func (x *Mdapi) UnSubscribeForQuoteRspMdgo(symbol []string) int {
	ppInstrumentID := []*C.char{}
	for _, syb := range symbol {
		ppInstrumentID = append(ppInstrumentID, C.CString(syb))
	}
	count := cap(ppInstrumentID)
	result := C.UnSubscribeForQuoteRspMd(x.mdapi, &ppInstrumentID[0], C.int(count), (C.OnUnSubForQuoteRsp)(unsafe.Pointer(C.OnUnSubForQuoteRspMdgo_cgo)))
	return int(result)
}

func (x *Mdapi) RegMarketDataMdgo(mktgo OnMkt) {
	//C.regMarketDataMd(x.mdapi, (C.OnMarketDataMd)(unsafe.Pointer(C.OnMarketDataMdgo_cgo)))
	C.regMarketDataMd(x.mdapi, (C.OnMarketDataMd)(mktgo))
}

func (x *Mdapi) RegRTForQuoteRspMdgo() {
	C.regRTForQuoteRsp(x.mdapi, (C.OnForQuoteRsp)(unsafe.Pointer(C.OnForQuoteRspMdgo_cgo)))
}

func (x *Mdapi) LogoutMd(pUserLogout *C.struct_CThostFtdcUserLogoutField, nRequestID int) int {

	result := C.LogoutMd(x.mdapi, pUserLogout, C.int(nRequestID), (C.OnLogout)(unsafe.Pointer(C.OnLogoutMdgo_cgo)))
	return int(result)
}
