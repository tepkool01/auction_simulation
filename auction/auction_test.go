package auction

import (
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

	t.Run("LookoutProvidedJSON", func(t *testing.T) {
		t.Parallel()

		auction, err := NewAuction("testdata/happyPath.json", 50)
		assert.Nil(t, err)

		auction.Run(kingOfTheHill)
		results := auction.GetListingStatuses()
		auction.PrintResults()

		assert.NotNil(t, results)
		assert.Equal(t, 3, len(results))
	})

	t.Run("IdenticalBids", func(t *testing.T) {
		t.Parallel()

		auction, err := NewAuction("testdata/identicalBids.json", 50)
		assert.Nil(t, err)

		auction.Run(kingOfTheHill)
		results := auction.GetListingStatuses()

		assert.Equal(t, "Alice", results["a8cfcb76-7f24-4420-a5ba-d46dd77bdffd"].HighestBidder.GetName())
	})

	t.Run("SlowBidNotCounted", func(t *testing.T) {
		t.Parallel()

		auction, err := NewAuction("testdata/itemExpires.json", 1500)
		assert.Nil(t, err)

		auction.Run(kingOfTheHill)
		results := auction.GetListingStatuses()

		assert.Equal(t, "Alice", results["a8cfcb76-7f24-4420-a5ba-d46dd77bdffd"].HighestBidder.GetName())
	})
}

func TestAuction_RunFileFailures(t *testing.T) {
	t.Parallel()

	t.Run("ValidationFail", func(t *testing.T) {
		t.Parallel()

		_, err := NewAuction("testdata/corruptedEntries.json", 50)
		assert.NotNil(t, err)
	})

	t.Run("NonJSONFile", func(t *testing.T) {
		t.Parallel()

		_, err := NewAuction("testdata/nonJSONfile.txt", 50)
		assert.NotNil(t, err)
	})

	t.Run("FileDoesNotExist", func(t *testing.T) {
		t.Parallel()

		_, err := NewAuction("testdata/1241251235135dfsbvdfgsdfg2352tdfg.txt", 50)
		assert.NotNil(t, err)
	})
}

func TestListingStatusWithNil(t *testing.T) {
	t.Parallel()

	ls := model.ListingStatus{}
	assert.Equal(t, true, ls.IsClosed())
}
