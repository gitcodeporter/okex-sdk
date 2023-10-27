package publicdata

import (
	"github.com/gitcodeporter/okex-sdk"
)

type (
	Instrument struct {
		InstID    string               `json:"instId"`
		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     string               `json:"ctVal,omitempty"`
		CtMult    string               `json:"ctMult,omitempty"`
		Stk       string               `json:"stk,omitempty"`
		TickSz    string               `json:"tickSz,omitempty"`
		LotSz     string               `json:"lotSz,omitempty"`
		MinSz     string               `json:"minSz,omitempty"`
		Lever     string               `json:"lever"`
		InstType  okex.InstrumentType  `json:"instType"`
		Category  okex.FeeCategory     `json:"category,string"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  string               `json:"listTime"`
		ExpTime   string               `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      int64                             `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     string                    `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       string              `json:"oi"`
		OiCcy    string              `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       int64               `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     string              `json:"fundingRate"`
		NextFundingRate string              `json:"NextFundingRate"`
		FundingTime     int64               `json:"fundingTime"`
		NextFundingTime int64               `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   string              `json:"buyLmt"`
		SellLmt  string              `json:"sellLmt"`
		TS       int64               `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx string              `json:"settlePx"`
		TS       int64               `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    string              `json:"delta"`
		Gamma    string              `json:"gamma"`
		Vega     string              `json:"vega"`
		Theta    string              `json:"theta"`
		DeltaBS  string              `json:"deltaBS"`
		GammaBS  string              `json:"gammaBS"`
		VegaBS   string              `json:"vegaBS"`
		ThetaBS  string              `json:"thetaBS"`
		Lever    string              `json:"lever"`
		MarkVol  string              `json:"markVol"`
		BidVol   string              `json:"bidVol"`
		AskVol   string              `json:"askVol"`
		RealVol  string              `json:"realVol"`
		TS       int64               `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string          `json:"ccy"`
		Amt          string          `json:"amt"`
		DiscountLv   string          `json:"discountLv"`
		DiscountInfo []*DiscountInfo `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate string `json:"discountRate"`
		MaxAmt       string `json:"maxAmt"`
		MinAmt       string `json:"minAmt"`
	}
	SystemTime struct {
		TS int64 `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss string                    `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string            `json:"ccy,omitempty"`
		Side    okex.OrderSide    `json:"side"`
		OosSide okex.PositionSide `json:"posSide"`
		BkPx    string            `json:"bkPx"`
		Sz      string            `json:"sz"`
		BkLoss  string            `json:"bkLoss"`
		TS      int64             `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   string              `json:"markPx"`
		TS       int64               `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         string              `json:"tier"`
		MinSz        string              `json:"minSz"`
		MaxSz        string              `json:"maxSz"`
		Mmr          string              `json:"mmr"`
		Imr          string              `json:"imr"`
		OptMgnFactor string              `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan string              `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  string              `json:"baseMaxLoan,omitempty"`
		MaxLever     string              `json:"maxLever"`
		TS           int64               `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string `json:"ccy"`
		Rate  string `json:"rate"`
		Quota string `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string `json:"level"`
		IrDiscount    string `json:"irDiscount"`
		LoanQuotaCoef int    `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string `json:"title"`
		State       string `json:"state"`
		Href        string `json:"href"`
		ServiceType string `json:"serviceType"`
		System      string `json:"system"`
		ScheDesc    string `json:"scheDesc"`
		Begin       int64  `json:"begin"`
		End         int64  `json:"end"`
	}
)
