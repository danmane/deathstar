package implgame

import "testing"

type testpair struct {
	segment Segment
	values  []Hex
}

var segTests = []testpair{
	{
		Segment{
			base:        Hex{0, 0},
			Length:      1,
			player:      NullPlayer,
			orientation: NullDirection,
		},
		[]Hex{{0, 0}}},
	{
		Segment{
			base:        Hex{0, 0},
			Length:      3,
			player:      NullPlayer,
			orientation: MidRight,
		},
		[]Hex{{0, 0}, {1, 0}, {2, 0}},
	},
}

func TestSegPieces(t *testing.T) {
	for _, pair := range segTests {
		v := pair.segment.segPieces()
		if !hexesEq(pair.values, v) {
			t.Error("For", pair.segment, "expected", pair.values, "got", v)
		}
	}
}

var simple_state = State{
	Board: Board{
		WhitePositions: slice2HexIndexArray([]Hex{
			{-1, 2}, {0, 2}, {-1, 3}, {-1, 4},
		}),
		BlackPositions: slice2HexIndexArray([]Hex{}),
		EdgeLength:     5,
	},
	NextPlayer:     White,
	MovesRemaining: 0,
	MarblesPerMove: 3,
	LossThreshold:  100,
}

func Test_segments(t *testing.T) {
	if len(simple_state.Segments(Black)) != 0 {
		t.Error("wrong number segs in 0 case")
	}

	if len(simple_state.Segments(White)) != 9 {
		t.Error("wrong number segs in simple case. Expected", 9, "got", len(simple_state.Segments(White)))
	}

	segs := Standard.Segments(White)
	numSegs := len(segs)
	if numSegs != 55 {
		for i1, v1 := range segs {
			if v1.Length > 3 || v1.Length < 1 {
				t.Error("wrong size seg")
			}
			if v1.player != White {
				t.Error("wrong player seg")
			}
			pieces := v1.segPieces()
			for _, hex := range pieces {
				if !standardWhitePositions.has(hex) {
					t.Error("nonexistant hex detected")
				}
			}
			for i2, v2 := range segs {
				pieces2 := v2.segPieces()
				p1s := slice2HexSet(pieces)
				p2s := slice2HexSet(pieces2)
				if i1 != i2 && p1s.eq(p2s) {
					t.Error("dup seg")
				}
			}
			t.Logf("pos: %v, length: %v, player: %v, orientation: %v\n", v1.base, v1.Length, v1.player, v1.orientation)
		}
		t.Error("number segments in standard game: expected 55, got", numSegs)
	}
}
