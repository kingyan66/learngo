package main

import "fmt"

type Item struct {
	ID    int
	Type  string
	Child *Item
}

type ItemClassifier interface {
	IsDoll() bool
}

func (item *Item) IsDoll() bool {
	if item.Type == "doll" {
		return true
	}
	return false
}

func main() {
	doll := Item{
		ID:   1,
		Type: "doll",
		Child: &Item{
			ID:   2,
			Type: "doll",
			Child: &Item{
				ID:   3,
				Type: "doll",
				Child: &Item{
					ID:    4,
					Type:  "diamond",
					Child: nil,
				},
			},
		},
	}
	diamond := findDiamond(doll)
	fmt.Printf("Item %d is diamond\n", diamond.ID)
}

func findDiamond(item Item) Item {
	if item.IsDoll() {
		return findDiamond(*item.Child)
	} else {
		return item
	}
}
