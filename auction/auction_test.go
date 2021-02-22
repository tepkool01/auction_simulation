package auction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lookout_interview/auction/model"
	"testing"
)

func TestKingOfHillStrategyIsOfBidWarType(t *testing.T) {
	t.Parallel()

	kingOfHillStrategy := &KingOfTheHillStrategy{}
	var s BidWarStrategy = kingOfHillStrategy
	assert.NotNil(t, s)
}

func TestItemBidIsOfItemType(t *testing.T) {
	t.Parallel()

	itemBid := &model.ItemBid{}
	var i model.Item = itemBid
	assert.NotNil(t, i)
}

func TestAuction_RunHappyPath(t *testing.T) {
	t.Parallel()

	kingOfTheHill := &KingOfTheHillStrategy{}

	auction, err := NewAuction("testdata/happyPath.json", 50)
	assert.Nil(t, err)

	auction.Run(kingOfTheHill)
	results := auction.GetListingStatuses()

	assert.NotNil(t, results)
	assert.Equal(t, 3, len(results))
}

func TestAuction_RunFileFailures(t *testing.T) {
	t.Parallel()

	t.Run("ValidationFail", func(t *testing.T) {
		t.Parallel()

		_, err := NewAuction("testdata/corruptedEntries.json", 50)
		assert.NotNil(t, err)
		fmt.Println(err)
	})

}
