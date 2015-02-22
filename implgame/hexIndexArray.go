package implgame

import "encoding/json"

type HexIndexArray struct {
	arr    [61]bool
	length int
}

var idx_help [9]int = [9]int{0, 6, 13, 21, 30, 39, 47, 54, 60}
var idx_back [61]int = [61]int{-4, -4, -4, -4, -4, -3, -3, -3, -3, -3, -3, -2, -2, -2, -2, -2, -2, -2, -1, -1, -1, -1, -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4}

func (p *Hex) idx() int {
	if p.Dist2Origin() >= 10 {
		return -1
	}
	// r_idx := p.R + 4
	// if r_idx < 0 || r_idx > 8 {
	// 	return -1
	// }
	i := p.Q + idx_help[p.R+4]
	// if i < 0 || i > 60 {
	// 	return -1
	// }
	return i
}

func fromIdx(x int) *Hex {
	r := idx_back[x]
	off := idx_help[r+4]
	return &Hex{Q: x - off, R: r}
}

var standardWhitePositions = slice2HexIndexArray(
	[]Hex{
		{-4, 3}, {-4, 4}, {-3, 3}, {-3, 4}, {-2, 2},
		{-2, 3}, {-2, 4}, {-1, 2}, {-1, 3}, {-1, 4},
		{0, 2}, {0, 3}, {0, 4}, {1, 3},
	})

var standardBlackPositions = slice2HexIndexArray(
	[]Hex{
		{-1, -3}, {0, -4}, {0, -3}, {0, -2}, {1, -4},
		{1, -3}, {1, -2}, {2, -4}, {2, -3}, {2, -2},
		{3, -4}, {3, -3}, {4, -4}, {4, -3},
	},
)

func (e HexIndexArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToSlice())
}

func (s *HexIndexArray) UnmarshalJSON(data []byte) error {
	var in []Hex
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	s.addHexes(in)
	return nil
}

func (e *HexIndexArray) ToSlice() []Hex {
	out := make([]Hex, e.length)
	x := 0
	for i, v := range e.arr {
		if v {
			out[x] = *fromIdx(i)
			x++
		}
	}
	return out
}

func (e *HexIndexArray) copy() HexIndexArray {
	var out HexIndexArray
	for i := 0; i < 61; i++ {
		out.arr[i] = e.arr[i]
	}
	out.length = e.length
	return out
}

func (e *HexIndexArray) has(h Hex) bool {
	i := h.idx()
	if i < 0 {
		return false
	}
	return e.arr[i]
}

func slice2HexIndexArray(hexes []Hex) HexIndexArray {
	var out HexIndexArray
	for _, v := range hexes {
		out.arr[v.idx()] = true
	}
	out.length = len(hexes)
	return out
}

func (e *HexIndexArray) removeHexes(hexes []Hex) {
	for _, v := range hexes {
		idx := v.idx()
		if e.arr[idx] {
			e.length--
		}
		e.arr[v.idx()] = false
	}
}

func (e *HexIndexArray) addHexes(hexes []Hex) {
	for _, v := range hexes {
		idx := v.idx()
		if !e.arr[idx] {
			e.length++
		}
		e.arr[idx] = true
	}
}
