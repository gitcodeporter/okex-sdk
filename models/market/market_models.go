package market

import (
	"github.com/gitcodeporter/okex-sdk"
)

type (
	Ticker struct {
		InstID    string              `json:"instId"`
		Last      string              `json:"last"`
		LastSz    string              `json:"lastSz"`
		AskPx     string              `json:"askPx"`
		AskSz     string              `json:"askSz"`
		BidPx     string              `json:"bidPx"`
		BidSz     string              `json:"bidSz"`
		Open24h   string              `json:"open24h"`
		High24h   string              `json:"high24h"`
		Low24h    string              `json:"low24h"`
		VolCcy24h string              `json:"volCcy24h"`
		Vol24h    string              `json:"vol24h"`
		SodUtc0   string              `json:"sodUtc0"`
		SodUtc8   string              `json:"sodUtc8"`
		InstType  okex.InstrumentType `json:"instType"`
		TS        int64               `json:"ts"`
	}
	IndexTicker struct {
		InstID  string `json:"instId"`
		IdxPx   string `json:"idxPx"`
		High24h string `json:"high24h"`
		Low24h  string `json:"low24h"`
		Open24h string `json:"open24h"`
		SodUtc0 string `json:"sodUtc0"`
		SodUtc8 string `json:"sodUtc8"`
		TS      int64  `json:"ts"`
	}
	OrderBook struct {
		Asks []*OrderBookEntity `json:"asks"`
		Bids []*OrderBookEntity `json:"bids"`
		TS   int64              `json:"ts"`
	}
	OrderBookWs struct {
		Asks     []*OrderBookEntity `json:"asks"`
		Bids     []*OrderBookEntity `json:"bids"`
		Checksum int                `json:"checksum"`
		TS       int64              `json:"ts"`
	}
	OrderBookEntity struct {
		DepthPrice      string
		Size            string
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		Ts          int64
		O           string
		H           string
		L           string
		C           string
		Vol         string
		VolCcy      string
		VolCcyQuote string
		Confirm     string
	}
	IndexCandle struct {
		O  string
		H  string
		L  string
		C  string
		TS int64
	}
	Trade struct {
		InstID  string         `json:"instId"`
		TradeID string         `json:"tradeId"`
		Px      string         `json:"px"`
		Sz      string         `json:"sz"`
		Side    okex.TradeSide `json:"side,string"`
		TS      int64          `json:"ts"`
	}
	TotalVolume24H struct {
		VolUsd string `json:"volUsd"`
		VolCny string `json:"volCny"`
		TS     int64  `json:"ts"`
	}
	IndexComponent struct {
		Index      string       `json:"index"`
		Last       string       `json:"last"`
		Components []*Component `json:"components"`
		TS         int64        `json:"ts"`
	}
	Component struct {
		Exch   string `json:"exch"`
		Symbol string `json:"symbol"`
		SymPx  string `json:"symPx"`
		Wgt    string `json:"wgt"`
		CnvPx  string `json:"cnvPx"`
	}
)
