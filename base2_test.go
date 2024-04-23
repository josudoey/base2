package base2

import "testing"

func BenchmarkFmtToString(b *testing.B) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < b.N; i++ {
		FmtToString(data)
	}
}

func BenchmarkEncodeToString(b *testing.B) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < b.N; i++ {
		EncodeToString(data)
	}
}

func TestEncodeToString(t *testing.T) {
	bs := make([]byte, 256)
	for i, _ := range bs {
		bs[i] = byte(i)
	}

	if EncodeToString(bs) != FmtToString(bs) {
		t.Fatalf("EncodeToString(%v): got %v, not equal %v", bs, EncodeToString(bs), FmtToString(bs))
	}
}
