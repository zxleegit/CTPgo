#include "_cgo_export.h"
#include "stdio.h"
#include "MdBridge.h"
#include "stdbool.h"
void OnMarketDataMdgo_cgo(struct CThostFtdcDepthMarketDataField *pMktgo)
{
	OnMarketDataMdgo(pMktgo);
	//OnMarketDataMdgo(struct CThostFtdcDepthMarketDataField *);
	printf("callback in CCC:%lf\n",pMktgo->LastPrice);
	printf("callback in CCC  pMktgo:%p\n",pMktgo);
}

void OnConnectedMdgo_cgo()
{
	printf("Connected");
}
void OnDisconnectedMdgo_cgo(int nReason)
{
	printf("dissconnect:%d",nReason);
}
void OnHeartWarningMdgo_cgo(int nTimeLapse)
{
	printf("OnHeartWarningMdgo:%d",nTimeLapse);
}

void OnLoginMdgo_cgo(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnLogoutMdgo_cgo(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnErrorMdgo_cgo(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnSubMarketDataMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnUnSubMarketDataMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnSubForQuoteRspMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnUnSubForQuoteRspMdgo_cgo(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
{
	printf("login:%s",pRspInfo->ErrorMsg);
}
void OnForQuoteRspMdgo_cgo(struct CThostFtdcForQuoteRspField *pForQuoteRsp)
{

}