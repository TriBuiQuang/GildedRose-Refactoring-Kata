package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func Test_Foo(t *testing.T) {
	t.Parallel()
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func Test_UpdateQuality(t *testing.T) {
	t.Parallel()
	type TestCases struct {
		name            string
		request, expect []*gildedrose.Item
	}

	testCases := []TestCases{
		{
			name: "sulfuras never changes",
			request: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 100, 100},
			},
			expect: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 100, 100},
			},
		},
		{
			name: "AgedBrie increasesInQuality",
			request: []*gildedrose.Item{
				{"Aged Brie", 2, 2},
				{"Aged Brie", 0, 2},
				{"Aged Brie", 2, 50},
			},
			expect: []*gildedrose.Item{
				{"Aged Brie", 1, 3},
				{"Aged Brie", -1, 4},
				{"Aged Brie", 1, 50},
			},
		},
		{
			name: "conjured increasesInQuality",
			request: []*gildedrose.Item{
				{"Conjured Mana Cake", 3, 6},
				{"Conjured Mana Cake", 0, 6},
			},
			expect: []*gildedrose.Item{
				{"Conjured Mana Cake", 2, 4},
				{"Conjured Mana Cake", -1, 2},
			},
		},
		{
			name: "BackStagePasses increasesInQuality",
			request: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 20, 2},
				{"Backstage passes to a TAFKAL80ETC concert", 10, 2},
				{"Backstage passes to a TAFKAL80ETC concert", 5, 2},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
			},
			expect: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 19, 3},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 4},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 5},
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			items := testCase.request

			gildedrose.UpdateQuality(items)

			assert.Equal(t, testCase.expect, items)
			for index, itemResult := range items {
				assert.Equal(t, testCase.expect[index].Name, itemResult.Name)
				assert.Equal(t, testCase.expect[index].Quality, itemResult.Quality)
				assert.Equal(t, testCase.expect[index].SellIn, itemResult.SellIn)
			}
		})
	}

	// var items = []*gildedrose.Item{
	// 	{"foo", 0, 0},
	// }

	// gildedrose.UpdateQuality(items)

	// if items[0].Name != "fixme" {
	// 	t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	// }
}
