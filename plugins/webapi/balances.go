package webapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"

	"github.com/iotaledger/iota.go/address"

	"github.com/gohornet/hornet/pkg/model/tangle"
)

func init() {
	addEndpoint("getBalances", getBalances, implementedAPIcalls)
}

func getBalances(i interface{}, c *gin.Context, _ <-chan struct{}) {
	e := ErrorReturn{}
	query := &GetBalances{}

	err := mapstructure.Decode(i, query)
	if err != nil {
		e.Error = "Internal error"
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	if len(query.Addresses) == 0 {
		e.Error = "No addresses provided"
		c.JSON(http.StatusBadRequest, e)
	}

	for _, addr := range query.Addresses {
		// Check if address is valid
		if err := address.ValidAddress(addr); err != nil {
			e.Error = "Invalid address: " + addr
			c.JSON(http.StatusBadRequest, e)
			return
		}
	}

	tangle.ReadLockLedger()
	defer tangle.ReadUnlockLedger()

	if !tangle.IsNodeSynced() {
		e.Error = "Node not synced"
		c.JSON(http.StatusBadRequest, e)
		return
	}

	cachedLatestSolidMs := tangle.GetMilestoneOrNil(tangle.GetSolidMilestoneIndex()) // bundle +1
	if cachedLatestSolidMs == nil {
		e.Error = "Ledger state invalid - Milestone not found"
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	defer cachedLatestSolidMs.Release(true) // bundle -1

	result := GetBalancesReturn{}

	for _, addr := range query.Addresses {

		balance, _, err := tangle.GetBalanceForAddressWithoutLocking(addr[:81])
		if err != nil {
			e.Error = "Ledger state invalid"
			c.JSON(http.StatusInternalServerError, e)
			return
		}

		// Address balance
		result.Balances = append(result.Balances, strconv.FormatUint(balance, 10))
	}

	// The index of the milestone that confirmed the most recent balance
	result.MilestoneIndex = uint32(cachedLatestSolidMs.GetBundle().GetMilestoneIndex())
	result.References = []string{cachedLatestSolidMs.GetBundle().GetMilestoneHash()}
	c.JSON(http.StatusOK, result)
}
