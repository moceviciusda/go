package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	diagram  string
	children []string
}

func (g *Garden) GetRows() []string {
	return strings.Split(g.diagram, "\n")[1:]
}

var PlantMap = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func ValidateDiagram(diagram string) error {
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return fmt.Errorf("diagram must gave 3 rows. diagram rows: %v", len(rows))
	}
	if len(rows[1]) != len(rows[2]) {
		return fmt.Errorf("mismatched garden row length, row 1: %v, row 2: %v", rows[1], rows[2])
	}
	if len(rows[1])%2 != 0 {
		return fmt.Errorf("invalid number of cups: %v", len(rows[1]))
	}

	for ri, row := range rows[1:] {
		for _, r := range row {
			_, ok := PlantMap[r]
			if !ok {
				return fmt.Errorf("invalid cup code, row: %v, code: %v", ri, r)
			}
		}
	}

	return nil
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	err := ValidateDiagram(diagram)
	if err != nil {
		return nil, err
	}

	// should add the map to the struct and use it to retrieve childs plants
	childMap := make(map[string]bool, 0)
	for _, child := range children {
		_, ok := childMap[child]
		if ok {
			return nil, fmt.Errorf("duplicate child name: %v", child)
		}
		childMap[child] = true
	}

	return &Garden{diagram, children}, nil
}

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	sortedChildren := make([]string, len(g.children))
	copy(sortedChildren, g.children)
	sort.Slice(sortedChildren, func(i, j int) bool {
		return sortedChildren[i] < sortedChildren[j]
	})

	for ci, c := range sortedChildren {
		if c == child {
			rows := g.GetRows()
			for _, p := range rows[0][ci*2 : ci*2+2] {
				var decoded string
				decoded, ok = PlantMap[p]
				plants = append(plants, decoded)
			}
			for _, p := range rows[1][ci*2 : ci*2+2] {
				var decoded string
				decoded, ok = PlantMap[p]
				plants = append(plants, decoded)
			}
		}
	}

	return
}
