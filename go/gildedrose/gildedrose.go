package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateItemQuality(item)
	}
}

func updateItemQuality(item *Item) {
	_agedBrie := "Aged Brie"
	_backStage := "Backstage passes to a TAFKAL80ETC concert"
	_sulFurans := "Sulfuras, Hand of Ragnaros"
	_conjured := "Conjured Mana Cake"

	_isDegrade := item.Name != _agedBrie && item.Name != _backStage && item.Name != _sulFurans
	degradeRate := -1
	if item.Name == _conjured {
		degradeRate = -2
	}

	if _isDegrade {
		adjustQuality(item, degradeRate)
	}

	if item.Name == _agedBrie || item.Name == _backStage {
		adjustQuality(item, 1)
	}

	if item.Name == _backStage {
		if item.SellIn < 11 {
			adjustQuality(item, 1)
		}

		if item.SellIn < 6 {
			adjustQuality(item, 1)
		}
	}

	if item.Name != _sulFurans {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if _isDegrade {
			adjustQuality(item, degradeRate)
		}

		if item.Name != _agedBrie {
			if item.Name == _backStage || item.Name == _sulFurans {
				item.Quality = 0
			}
		} else {
			adjustQuality(item, 1)
		}
	}
}

func adjustQuality(item *Item, adjustment int) {
	if item.Quality < 50 && item.Quality >= 0 {
		item.Quality = item.Quality + adjustment
	}
}
