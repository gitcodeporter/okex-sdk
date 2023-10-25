package account

import (
	"github.com/gitcodeporter/okex-sdk"
)

type (
	Balance struct {
		TotalEq     string            `json:"totalEq"`
		IsoEq       string            `json:"isoEq"`
		AdjEq       string            `json:"adjEq,omitempty"`
		OrdFroz     string            `json:"ordFroz,omitempty"`
		Imr         string            `json:"imr,omitempty"`
		Mmr         string            `json:"mmr,omitempty"`
		MgnRatio    string            `json:"mgnRatio,omitempty"`
		NotionalUsd string            `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       int64             `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string `json:"ccy"`
		Eq            string `json:"eq"`
		CashBal       string `json:"cashBal"`
		IsoEq         string `json:"isoEq,omitempty"`
		AvailEq       string `json:"availEq,omitempty"`
		DisEq         string `json:"disEq"`
		AvailBal      string `json:"availBal"`
		FrozenBal     string `json:"frozenBal"`
		OrdFrozen     string `json:"ordFrozen"`
		Liab          string `json:"liab,omitempty"`
		Upl           string `json:"upl,omitempty"`
		UplLib        string `json:"uplLib,omitempty"`
		CrossLiab     string `json:"crossLiab,omitempty"`
		IsoLiab       string `json:"isoLiab,omitempty"`
		MgnRatio      string `json:"mgnRatio,omitempty"`
		Interest      string `json:"interest,omitempty"`
		Twap          string `json:"twap,omitempty"`
		MaxLoan       string `json:"maxLoan,omitempty"`
		EqUsd         string `json:"eqUsd"`
		NotionalLever string `json:"notionalLever,omitempty"`
		StgyEq        string `json:"stgyEq"`
		IsoUpl        string `json:"isoUpl,omitempty"`
		UTime         int64  `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosID       string              `json:"posId"`
		TradeID     string              `json:"tradeId"`
		Pos         string              `json:"pos"`
		AvailPos    string              `json:"availPos,omitempty"`
		AvgPx       string              `json:"avgPx"`
		Upl         string              `json:"upl"`
		UplRatio    string              `json:"uplRatio"`
		Lever       string              `json:"lever"`
		LiqPx       string              `json:"liqPx,omitempty"`
		Imr         string              `json:"imr,omitempty"`
		Margin      string              `json:"margin,omitempty"`
		MgnRatio    string              `json:"mgnRatio"`
		Mmr         string              `json:"mmr"`
		Liab        string              `json:"liab,omitempty"`
		Interest    string              `json:"interest"`
		NotionalUsd string              `json:"notionalUsd"`
		ADL         string              `json:"adl"`
		Last        string              `json:"last"`
		DeltaBS     string              `json:"deltaBS"`
		DeltaPA     string              `json:"deltaPA"`
		GammaBS     string              `json:"gammaBS"`
		GammaPA     string              `json:"gammaPA"`
		ThetaBS     string              `json:"thetaBS"`
		ThetaPA     string              `json:"thetaPA"`
		VegaBS      string              `json:"vegaBS"`
		VegaPA      string              `json:"vegaPA"`
		PosSide     okex.PositionSide   `json:"posSide"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		InstType    okex.InstrumentType `json:"instType"`
		CTime       int64               `json:"cTime"`
		UTime       int64               `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     int64             `json:"pTime"`
		UTime     int64             `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   string                               `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      int64                                `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string `json:"ccy"`
		Eq    string `json:"eq"`
		DisEq string `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy string              `json:"notionalCcy"`
		Pos         string              `json:"pos"`
		NotionalUsd string              `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    string              `json:"balChg"`
		PosBalChg string              `json:"posBalChg"`
		Bal       string              `json:"bal"`
		PosBal    string              `json:"posBal"`
		Sz        string              `json:"sz"`
		Pnl       string              `json:"pnl"`
		Fee       string              `json:"fee"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        int64               `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType okex.GreekType    `json:"greeksType"`
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string            `json:"instId"`
		Lever   string            `json:"lever"`
		MgnMode okex.MarginMode   `json:"mgnMode"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string `json:"instId"`
		Ccy     string `json:"ccy"`
		MaxBuy  string `json:"maxBuy"`
		MaxSell string `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string `json:"instId"`
		AvailBuy  string `json:"availBuy"`
		AvailSell string `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string            `json:"instId"`
		Amt     string            `json:"amt"`
		PosSide okex.PositionSide `json:"posSide,string"`
		Type    okex.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy"`
		Ccy     string          `json:"ccy"`
		MaxLoan string          `json:"maxLoan"`
		MgnMode okex.MarginMode `json:"mgnMode"`
		Side    okex.OrderSide  `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    string              `json:"taker"`
		Maker    string              `json:"maker"`
		Delivery string              `json:"delivery,omitempty"`
		Exercise string              `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       int64               `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string          `json:"instId"`
		Ccy          string          `json:"ccy"`
		Interest     string          `json:"interest"`
		InterestRate string          `json:"interestRate"`
		Liab         string          `json:"liab"`
		MgnMode      okex.MarginMode `json:"mgnMode"`
		TS           int64           `json:"ts"`
	}
	InterestRate struct {
		Ccy          string `json:"ccy"`
		InterestRate string `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string `json:"ccy"`
		MaxWd string `json:"maxWd"`
	}
)
