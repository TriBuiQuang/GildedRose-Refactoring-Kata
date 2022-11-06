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

	if item.Name != _agedBrie && item.Name != _backStage {
		if item.Quality > 0 {

			if item.Name != _sulFurans {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == _backStage {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.Name != _sulFurans {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != _agedBrie {
			if item.Name != _backStage {
				if item.Quality > 0 {
					if item.Name != _sulFurans {
						item.Quality = item.Quality - 1
					}
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			adjustQuality(item)
		}
	}
}

func adjustQuality(item *Item) {
	if item.Quality < 50 && item.Quality >= 0 {
		item.Quality = item.Quality + 1
	}
}
