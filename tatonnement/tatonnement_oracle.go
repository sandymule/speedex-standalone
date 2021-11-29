package tatonnement

import (
	"fmt"

	"github.com/sandymule/speedex-standalone/orderbook"
)

type TatonnementOracle struct {
	Params            ControlParams
	MOrderbookManager orderbook.OrderbookManager
}

func (to *TatonnementOracle) ComputePrices(params ControlParams, prices map[orderbook.Asset]float64) {
	to.Params = params
	fmt.Println(to)
	baselineDemand := to.MOrderbookManager.DemandQuery(prices, to.Params.MSmoothMult)
	baselineObjective := baselineDemand.GetObjective()
	fmt.Println(baselineDemand)
	fmt.Println(baselineObjective)

	stepSize := to.Params.KStartingStepSize

	for !to.Params.Done() {
		to.Params.IncrementRound()
		fmt.Println(prices)
		trialPrices := to.Params.SetTrialPrices(prices, *baselineDemand, stepSize)
		trialDemand := to.MOrderbookManager.DemandQuery(trialPrices, to.Params.MSmoothMult)
		prices = trialPrices
		baselineDemand = trialDemand
	}
}
