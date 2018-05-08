package trade

/*
trade执行器支持trade的创建和交易，

主要提供操作有以下几种：
1）挂单出售；
2）购买指定的卖单；
3）撤销卖单；
*/

import (
	log "github.com/inconshreveable/log15"

	"gitlab.33.cn/chain33/chain33/common"
	"gitlab.33.cn/chain33/chain33/executor/drivers"
	"gitlab.33.cn/chain33/chain33/executor/drivers/token"
	"gitlab.33.cn/chain33/chain33/types"
)

var tradelog = log.New("module", "execs.trade")

func init() {
	t := newTrade()
	drivers.Register(t.GetName(), t, types.ForkV2AddToken)
}

type trade struct {
	drivers.DriverBase
}

func newTrade() *trade {
	t := &trade{}
	t.SetChild(t)
	return t
}

func (t *trade) GetName() string {
	return "trade"
}

func (t *trade) Clone() drivers.Driver {
	clone := &trade{}
	clone.DriverBase = *(t.DriverBase.Clone().(*drivers.DriverBase))
	clone.SetChild(clone)
	return clone
}

func (t *trade) Exec(tx *types.Transaction, index int) (*types.Receipt, error) {
	var trade types.Trade
	err := types.Decode(tx.Payload, &trade)
	if err != nil {
		return nil, err
	}
	tradelog.Info("exec trade tx=", "tx hash", common.Bytes2Hex(tx.Hash()), "Ty", trade.GetTy())

	action := newTradeAction(t, tx)
	switch trade.GetTy() {
	case types.TradeSellLimit:
		return action.tradeSell(trade.GetTokensell())

	case types.TradeBuyMarket:
		return action.tradeBuy(trade.GetTokenbuy())

	case types.TradeRevokeSell:
		return action.tradeRevokeSell(trade.GetTokenrevokesell())

	case types.TradeBuyLimit:
		if t.GetHeight() < types.ForTradeBuyLimit {
			return nil, types.ErrActionNotSupport
		}
		return action.tradeBuyLimit(trade.GetTokenbuylimit())

	case types.TradeSellMarket:
		if t.GetHeight() < types.ForTradeBuyLimit {
			return nil, types.ErrActionNotSupport
		}
		return action.tradeSellMarket(trade.GetTokensellmarket())

	case types.TradeRevokeBuy:
		if t.GetHeight() < types.ForTradeBuyLimit {
			return nil, types.ErrActionNotSupport
		}
		return action.tradeRevokeBuyLimit(trade.GetTokenrevokebuy())

	default:
		return nil, types.ErrActionNotSupport
	}
}

func (t *trade) ExecLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := t.DriverBase.ExecLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	for i := 0; i < len(receipt.Logs); i++ {
		item := receipt.Logs[i]
		if item.Ty == types.TyLogTradeSell || item.Ty == types.TyLogTradeRevoke {
			var receipt types.ReceiptTradeSell
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveSell([]byte(receipt.Base.Sellid), item.Ty)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == types.TyLogTradeBuy {
			var receipt types.ReceiptBuyBase
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveBuy(&receipt)
			set.KV = append(set.KV, kv...)

			// 添加个人资产列表
			kv = token.AddTokenToAssets(receipt.Owner, t.GetLocalDB(), receipt.TokenSymbol)
			if kv != nil {
				set.KV = append(set.KV, kv...)
			}
		} else if item.Ty == types.TyLogTradeBuyRevoke || item.Ty == types.TyLogTradeBuyLimit {
			var receipt types.ReceiptTradeBuyLimit
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}

			kv := t.saveBuyLimit([]byte(receipt.Base.Buyid), item.Ty)
			set.KV = append(set.KV, kv...)

			// 添加个人资产列表
			kv = token.AddTokenToAssets(receipt.Base.Owner, t.GetLocalDB(), receipt.Base.TokenSymbol)
			if kv != nil {
				set.KV = append(set.KV, kv...)
			}
		} else if item.Ty == types.TyLogTradeSellMarket {
			var receipt types.ReceiptTradeBase
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.saveSellMarket(&receipt)
			tradelog.Info("saveSellMarket", "kv", kv)
			set.KV = append(set.KV, kv...)
		}
	}

	return set, nil
}

func (t *trade) ExecDelLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := t.DriverBase.ExecDelLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}

	for i := 0; i < len(receipt.Logs); i++ {
		item := receipt.Logs[i]
		if item.Ty == types.TyLogTradeSell || item.Ty == types.TyLogTradeRevoke {
			var receipt types.ReceiptTradeSell
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.deleteSell([]byte(receipt.Base.Sellid), item.Ty)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == types.TyLogTradeBuy {
			var receipt types.ReceiptBuyBase
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.deleteBuy(&receipt)
			set.KV = append(set.KV, kv...)
		} else if item.Ty == types.TyLogTradeBuyRevoke || item.Ty == types.TyLogTradeBuyLimit {
			var receipt types.ReceiptTradeBuyLimit
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.deleteBuyLimit([]byte(receipt.Base.Buyid), item.Ty)
			set.KV = append(set.KV, kv...)

		} else if item.Ty == types.TyLogTradeSellMarket {
			var receipt types.ReceiptTradeBase
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				panic(err) //数据错误了，已经被修改了
			}
			kv := t.deleteSellMarket(&receipt)
			set.KV = append(set.KV, kv...)
		}
	}
	return set, nil
}

func (t *trade) Query(funcName string, params []byte) (types.Message, error) {
	tradelog.Info("trade Query", "name", funcName)
	switch funcName {
	//查询某个特定用户的一个或者多个token的卖单,包括所有状态的卖单
	//TODO:后续可以考虑支持查询不同状态的卖单
	case "GetOnesSellOrder":
		var addrTokens types.ReqAddrTokens
		err := types.Decode(params, &addrTokens)
		if err != nil {
			return nil, err
		}
		return t.GetOnesSellOrder(&addrTokens)
	case "GetOnesBuyOrder":
		var addrTokens types.ReqAddrTokens
		err := types.Decode(params, &addrTokens)
		if err != nil {
			return nil, err
		}
		return t.GetOnesBuyOrder(&addrTokens)
		//查寻所有的可以进行交易的卖单
	case "GetAllSellOrdersWithStatus":
		var addrTokens types.ReqAddrTokens
		err := types.Decode(params, &addrTokens)
		if err != nil {
			return nil, err
		}
		return t.GetAllSellOrdersWithStatus(addrTokens.Status)
	case "GetTokenSellOrderByStatus": // 根据token 分页显示未完成成交卖单
		var req types.ReqTokenSellOrder
		err := types.Decode(params, &req)
		if err != nil {
			return nil, err
		}
		if req.Status == 0 {
			req.Status = types.TracdOrderStatusOnSale
		}
		return t.GetTokenByStatus(&req, req.Status)
	case "GetTokenBuyLimitOrderByStatus": // 根据token 分页显示未完成成交买单
		var req types.ReqTokenBuyLimitOrder
		err := types.Decode(params, &req)
		if err != nil {
			return nil, err
		}
		if req.Status == 0 {
			req.Status = types.TracdOrderStatusOnBuy
		}
		return t.GetTokenBuyLimitOrderByStatus(&req, req.Status)
	case "GetAllBuyOrdersWithStatus":
		var addrTokens types.ReqAddrTokens
		err := types.Decode(params, &addrTokens)
		if err != nil {
			return nil, err
		}
		return t.GetAllSellOrdersWithStatus(addrTokens.Status)

	default:
	}
	tradelog.Error("trade Query", "Query type not supprt with func name", funcName)
	return nil, types.ErrQueryNotSupport
}

func (t *trade) getSellOrderFromDb(sellid []byte) *types.SellOrder {
	value, err := t.GetStateDB().Get(sellid)
	if err != nil {
		panic(err)
	}
	var sellorder types.SellOrder
	types.Decode(value, &sellorder)
	return &sellorder
}

func genSaveSellKv(sellorder *types.SellOrder) []*types.KeyValue {
	status := sellorder.Status
	var kv []*types.KeyValue
	kv = saveSellOrderKeyValue(kv, sellorder, status)
	if types.TracdOrderStatusSoldOut == status || types.TracdOrderStatusRevoked == status {
		tradelog.Debug("trade saveSell ", "remove old status onsale to soldout or revoked with sellid", sellorder.Sellid)
		kv = deleteSellOrderKeyValue(kv, sellorder, types.TracdOrderStatusOnSale)
	}
	return kv
}

func (t *trade) saveSell(sellid []byte, ty int32) []*types.KeyValue {
	sellorder := t.getSellOrderFromDb(sellid)
	return genSaveSellKv(sellorder)
}

func deleteSellOrderKeyValue(kv []*types.KeyValue, sellorder *types.SellOrder, status int32) []*types.KeyValue {
	return genSellOrderKeyValue(kv, sellorder, status, nil)
}

func saveSellOrderKeyValue(kv []*types.KeyValue, sellorder *types.SellOrder, status int32) []*types.KeyValue {
	sellid := []byte(sellorder.Sellid)
	return genSellOrderKeyValue(kv, sellorder, status, sellid)
}

func genDeleteSellKv(sellorder *types.SellOrder) []*types.KeyValue {
	status := sellorder.Status
	var kv []*types.KeyValue
	kv = deleteSellOrderKeyValue(kv, sellorder, status)
	if types.TracdOrderStatusSoldOut == status || types.TracdOrderStatusRevoked == status {
		tradelog.Debug("trade saveSell ", "remove old status onsale to soldout or revoked with sellid", sellorder.Sellid)
		kv = saveSellOrderKeyValue(kv, sellorder, types.TracdOrderStatusOnSale)
	}
	return kv
}

func (t *trade) deleteSell(sellid []byte, ty int32) []*types.KeyValue {
	sellorder := t.getSellOrderFromDb(sellid)
	return genDeleteSellKv(sellorder)
}

func (t *trade) saveBuy(receiptTradeBuy *types.ReceiptBuyBase) []*types.KeyValue {
	var kv []*types.KeyValue
	return saveBuyMarketOrderKeyValue(kv, receiptTradeBuy, types.TracdOrderStatusBoughtOut, t.GetHeight())
}

func (t *trade) deleteBuy(receiptTradeBuy *types.ReceiptBuyBase) []*types.KeyValue {
	var kv []*types.KeyValue
	return deleteBuyMarketOrderKeyValue(kv, receiptTradeBuy, types.TracdOrderStatusBoughtOut, t.GetHeight())
}

// BuyLimit Local
func (t *trade) getBuyOrderFromDb(buyid []byte) *types.BuyLimitOrder {
	value, err := t.GetStateDB().Get(buyid)
	if err != nil {
		panic(err)
	}
	var buyOrder types.BuyLimitOrder
	types.Decode(value, &buyOrder)
	return &buyOrder
}

func genSaveBuyLimitKv(buyOrder *types.BuyLimitOrder) []*types.KeyValue {
	status := buyOrder.Status
	var kv []*types.KeyValue
	kv = saveBuyLimitOrderKeyValue(kv, buyOrder, status)
	if types.TracdOrderStatusBoughtOut == status || types.TracdOrderStatusBuyRevoked == status {
		tradelog.Debug("trade saveBuyLimit ", "remove old status with Buyid", buyOrder.Buyid)
		kv = deleteBuyLimitKeyValue(kv, buyOrder, types.TracdOrderStatusOnSale)
	}
	return kv
}

func (t *trade) saveBuyLimit(buyid []byte, ty int32) []*types.KeyValue {
	buyOrder := t.getBuyOrderFromDb(buyid)
	return genSaveBuyLimitKv(buyOrder)
}

func saveBuyLimitOrderKeyValue(kv []*types.KeyValue, buyOrder *types.BuyLimitOrder, status int32) []*types.KeyValue {
	buyid := []byte(buyOrder.Buyid)
	return genBuyLimitOrderKeyValue(kv, buyOrder, status, buyid)
}

func deleteBuyLimitKeyValue(kv []*types.KeyValue, buyOrder *types.BuyLimitOrder, status int32) []*types.KeyValue {
	return genBuyLimitOrderKeyValue(kv, buyOrder, status, nil)
}

func genDeleteBuyLimitKv(buyOrder *types.BuyLimitOrder) []*types.KeyValue {
	status := buyOrder.Status
	var kv []*types.KeyValue
	kv = deleteBuyLimitKeyValue(kv, buyOrder, status)
	if types.TracdOrderStatusBoughtOut == status || types.TracdOrderStatusBuyRevoked == status {
		tradelog.Debug("trade saveSell ", "remove old status onsale to soldout or revoked with sellid", buyOrder.Buyid)
		kv = saveBuyLimitOrderKeyValue(kv, buyOrder, types.TracdOrderStatusOnBuy)
	}
	return kv
}

func (t *trade) deleteBuyLimit(buyid []byte, ty int32) []*types.KeyValue {
	buyOrder := t.getBuyOrderFromDb(buyid)
	return genDeleteBuyLimitKv(buyOrder)
}

func (t *trade) saveSellMarket(receiptTradeBuy *types.ReceiptTradeBase) []*types.KeyValue {
	var kv []*types.KeyValue
	return saveSellMarketOrderKeyValue(kv, receiptTradeBuy, types.TracdOrderStatusSoldOut, t.GetHeight())
}

func (t *trade) deleteSellMarket(receiptTradeBuy *types.ReceiptTradeBase) []*types.KeyValue {
	var kv []*types.KeyValue
	return deleteSellMarketOrderKeyValue(kv, receiptTradeBuy, types.TracdOrderStatusSoldOut, t.GetHeight())
}

func saveSellMarketOrderKeyValue(kv []*types.KeyValue, receipt *types.ReceiptTradeBase, status int32, height int64) []*types.KeyValue {
	txhash := []byte(receipt.Txhash)
	return genSellMarketOrderKeyValue(kv, receipt, status, height, txhash)
}

func deleteSellMarketOrderKeyValue(kv []*types.KeyValue, receipt *types.ReceiptTradeBase, status int32, height int64) []*types.KeyValue {
	return genSellMarketOrderKeyValue(kv, receipt, status, height, nil)
}

func saveBuyMarketOrderKeyValue(kv []*types.KeyValue, receipt *types.ReceiptBuyBase, status int32, height int64) []*types.KeyValue {
	txhash := []byte(receipt.Txhash)
	return genBuyMarketOrderKeyValue(kv, receipt, status, height, txhash)
}

func deleteBuyMarketOrderKeyValue(kv []*types.KeyValue, receipt *types.ReceiptBuyBase, status int32, height int64) []*types.KeyValue {
	return genBuyMarketOrderKeyValue(kv, receipt, status, height, nil)
}
