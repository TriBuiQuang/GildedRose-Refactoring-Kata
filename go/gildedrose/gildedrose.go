package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	agedBrie  = "Aged Brie"
	backStage = "Backstage passes to a TAFKAL80ETC concert"
	sulFurans = "Sulfuras, Hand of Ragnaros"
	conjured  = "Conjured Mana Cake"
)

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateItemQuality(item)
	}
}

func updateItemQuality(item *Item) {
	_isDegrade := item.Name != agedBrie && item.Name != backStage && item.Name != sulFurans
	_isExpired := item.SellIn < 1
	_degradeRate := determineDegradeRate(item)

	if _isExpired {
		_degradeRate *= 2
	}

	if _isDegrade {
		adjustQuality(item, _degradeRate)
	}

	if item.Name == agedBrie || item.Name == backStage {
		adjustQuality(item, 1)
	}

	if item.Name == backStage {
		if item.SellIn < 11 {
			adjustQuality(item, 1)
		}

		if item.SellIn < 6 {
			adjustQuality(item, 1)
		}
	}

	if item.Name != sulFurans {
		item.SellIn = item.SellIn - 1
	}

	if _isExpired {

		if item.Name != agedBrie {
			if item.Name == backStage || item.Name == sulFurans {
				item.Quality = 0
			}
		} else {
			adjustQuality(item, 1)
		}
	}
}

func determineDegradeRate(item *Item) int {
	degradeRate := -1
	if item.Name == conjured {
		degradeRate = -2
	}
	return degradeRate
}

func adjustQuality(item *Item, adjustment int) {
	if item.Quality < 50 && item.Quality >= 0 {
		item.Quality = item.Quality + adjustment
	}
}
