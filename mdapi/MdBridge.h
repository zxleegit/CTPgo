// 下列 ifdef 块是创建使从 DLL 导出更简单的
// 宏的标准方法。此 DLL 中的所有文件都是用命令行上定义的 MDBRIDGE_EXPORTS
// 符号编译的。在使用此 DLL 的
// 任何项目上不应定义此符号。这样，源文件中包含此文件的任何其他项目都会将
// MDBRIDGE_API 函数视为是从 DLL 导入的，而此 DLL 则将用此宏定义的
// 符号视为是被导出的。
#ifdef MDBRIDGE_EXPORTS
#define MDBRIDGE_API extern "C" __declspec(dllexport)
#else
#define MDBRIDGE_API __declspec(dllimport)
#endif
#pragma once
#include "ThostFtdcUserApiStruct.h"
#include "stdbool.h"

typedef void(*OnMarketDataMd)(struct CThostFtdcDepthMarketDataField *pMkt);
typedef void(*OnConnected)();
typedef void(*OnDisconnected)(int nReason);
typedef void(*OnHeartWarning)(int nTimeLapse);
typedef void(*OnLogin)(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnLogout)(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnError)(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnSubMarketData)(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct  CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnUnSubMarketData)(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct  CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnSubForQuoteRsp)(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct  CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnUnSubForQuoteRsp)(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct  CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
typedef void(*OnForQuoteRsp)(struct CThostFtdcForQuoteRspField *pForQuoteRsp);



typedef void* pMdBridge;
MDBRIDGE_API pMdBridge NewMdBridge();
MDBRIDGE_API void FreeMdBridge(pMdBridge x);

MDBRIDGE_API void ConnectMd(pMdBridge x, char *logpath, char *addr, OnConnected regConn);

MDBRIDGE_API void regDissConn(pMdBridge x, OnDisconnected regDissCon);
MDBRIDGE_API void regError(pMdBridge x, OnError regError);
MDBRIDGE_API void regHeartWarning(pMdBridge x, OnHeartWarning HeartWarning);

MDBRIDGE_API void InitMd(pMdBridge x);
MDBRIDGE_API const char *GetVersionMd(pMdBridge x);
MDBRIDGE_API void JoinMd(pMdBridge x);

MDBRIDGE_API int LoginMd(pMdBridge x,struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID, OnLogin regLogin);
MDBRIDGE_API const char *GetTradeDayMd(pMdBridge x);


MDBRIDGE_API int SubscribeMarketDataMd(pMdBridge x, char **ppInstrumentID, int count, OnSubMarketData Submkt);
MDBRIDGE_API int SubscribeForQuoteRspMd(pMdBridge x, char **ppInstrumentID, int count, OnSubForQuoteRsp SubQuote);

MDBRIDGE_API int UnSubscribeMarketDataMd(pMdBridge x, char **ppInstrumentID, int count, OnUnSubMarketData UnSubmkt);
MDBRIDGE_API int UnSubscribeForQuoteRspMd(pMdBridge x, char **ppInstrumentID, int count, OnUnSubForQuoteRsp UnSubQuote);

MDBRIDGE_API void regMarketDataMd(pMdBridge x, OnMarketDataMd mkt);
MDBRIDGE_API void regRTForQuoteRsp(pMdBridge x, OnForQuoteRsp rtQuoteRsp);

MDBRIDGE_API int LogoutMd(pMdBridge x,struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID, OnLogout regLogout);