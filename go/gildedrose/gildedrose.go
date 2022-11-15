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

	degradeRate := -1
	if item.Name == _conjured {
		degradeRate = -2
	}

	if item.Name != _agedBrie && item.Name != _backStage && item.Name != _sulFurans {
		adjustQuality(item, degradeRate)

	} else {
		adjustQuality(item, 1)

		if item.Name == _backStage {
			if item.SellIn < 11 {
				adjustQuality(item, 1)
			}

			if item.SellIn < 6 {
				adjustQuality(item, 1)
			}
		}

	}

	if item.Name != _sulFurans {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != _agedBrie {
			if item.Name != _backStage && item.Name != _sulFurans {
				adjustQuality(item, degradeRate)
			} else {
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
