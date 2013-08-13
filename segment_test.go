package gox12

import (
	"testing"
)

// AAA
func TestSegmentParseSegmentId(t *testing.T) {
	str2 := "TST&AA!1!1&BB!5"
	seg := NewSegment(str2, '&', '!', '^')
	expectedSegId := "TST"
	if seg.SegmentId != expectedSegId {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedSegId, seg.SegmentId)
	}
}

func TestSegmentParseComposites(t *testing.T) {
	var segtests = []struct {
		refdes   string
		expected string
	}{
		{"TST01-1", "AA"},
		{"TST01-2", "1"},
		{"TST01-3", "5"},
		{"TST02-1", "BB"},
	}
	segmentStr := "TST&AA!1!5&BB!5"
	seg := NewSegment(segmentStr, '&', '!', '^')
	for _, tt := range segtests {
		actual, found, err := seg.GetValue(tt.refdes)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if !found {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if actual != tt.expected {
			t.Errorf("Didn't get expected result [%s] for path [%s], instead got [%s]", tt.expected, tt.refdes, actual)
		}
	}
}

func TestSegmentIndexNotFound(t *testing.T) {
	var segtests = []struct {
		refdes   string
		expected string
	}{
		{"TST01-5", ""},
		{"TST06", ""},
		{"TST07", ""},
		{"TST05-2", ""},
	}
	segmentStr := "TST&AA!1!5&BB!5&&X"
	seg := NewSegment(segmentStr, '&', '!', '^')
	for _, tt := range segtests {
		actual, found, err := seg.GetValue(tt.refdes)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if found {
			t.Errorf("Found should be false for [%s]", tt.refdes)
		}
		if actual != tt.expected {
			t.Errorf("Didn't get expected result [%s] for path [%s], instead got [%s]", tt.expected, tt.refdes, actual)
		}
	}
}

func TestSegmentIdentity(t *testing.T) {
	var segtests = []struct {
		rawseg string
	}{
		{"TST*AA:1:1*BB:5*ZZ"},
		{"ISA*00*          *00*          *ZZ*ZZ000          *ZZ*ZZ001          *030828*1128*U*00401*000010121*0*T*:\n"},
	}
	for _, tt := range segtests {
		seg := NewSegment(tt.rawseg, '*', ':', '^')
		actual := seg.String()
		if actual != tt.rawseg {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.rawseg, actual)
		}
	}
}

func BenchmarkSegmentParse(b *testing.B) {
	str2 := "TST&AA!1!1&BB!5"
	for i := 0; i < b.N; i++ {
		_ = NewSegment(str2, '&', '!', '^')
	}
}

func BenchmarkSegmentString(b *testing.B) {
	rawseg := "TST&AA!1!1&BBbbbbbbbbb!5&&B!FjhhealkjF&&J&HJY&IU"
	s := NewSegment(rawseg, '&', '!', '^')
	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}
func BenchmarkSegmentFormat(b *testing.B) {
	rawseg := "TST&AA!1!1&BBbbbbbbbbb!5&&B!FjhhealkjF&&J&HJY&IU"
	s := NewSegment(rawseg, '&', '!', '^')
	for i := 0; i < b.N; i++ {
		_ = s.Format('*', ':', '^')
	}
}
