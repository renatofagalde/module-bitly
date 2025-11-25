package bitly

import (
	"bytes"
	"testing"
)

func TestE_KnownValues(t *testing.T) {
	tests := []struct {
		in   uint64
		want string
	}{
		{0, "0"},
		{1, "1"},
		{9, "9"},
		{10, "A"},
		{35, "Z"},
		{36, "a"},
		{61, "z"},
		{62, "10"},
		{63, "11"},
		{124, "20"},
	}

	for _, tt := range tests {
		got := E(tt.in)
		if got != tt.want {
			t.Errorf("E(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestD_KnownValues(t *testing.T) {
	tests := []struct {
		in   string
		want uint64
	}{
		{"0", 0},
		{"1", 1},
		{"9", 9},
		{"A", 10},
		{"Z", 35},
		{"a", 36},
		{"z", 61},
		{"10", 62},
		{"11", 63},
		{"20", 124},
	}

	for _, tt := range tests {
		got, err := D(tt.in)
		if err != nil {
			t.Fatalf("D(%q) returned error: %v", tt.in, err)
		}
		if got != tt.want {
			t.Errorf("D(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestRoundTripUint64(t *testing.T) {
	// testa alguns valores, incluindo bordas
	values := []uint64{
		0,
		1,
		10,
		61,
		62,
		63,
		124,
		999,
		123456789,
		^uint64(0), // max uint64
	}

	for _, v := range values {
		enc := E(v)
		dec, err := D(enc)
		if err != nil {
			t.Fatalf("round-trip D(E(%d)) returned error: %v", v, err)
		}
		if dec != v {
			t.Errorf("round-trip failed: got %d, want %d (enc=%q)", dec, v, enc)
		}
	}
}

func TestD_InvalidChar(t *testing.T) {
	_, err := D("ABC!")
	if err == nil {
		t.Fatalf("expected error for invalid characters, got nil")
	}
}

func TestEncodeBytes_Empty(t *testing.T) {
	got := EncodeBytes(nil)
	if got != "" {
		t.Errorf("EncodeBytes(nil) = %q, want %q", got, "")
	}

	got = EncodeBytes([]byte{})
	if got != "" {
		t.Errorf("EncodeBytes([]byte{}) = %q, want %q", got, "")
	}
}

func TestEncodeBytes_Zero(t *testing.T) {
	got := EncodeBytes([]byte{0x00})
	if got != "0" {
		t.Errorf("EncodeBytes([]byte{0x00}) = %q, want %q", got, "0")
	}
}

func TestEncodeDecodeBytes_RoundTrip(t *testing.T) {
	tests := [][]byte{
		{0x01},
		{0x01, 0x02, 0x03},
		{0xFF, 0x00, 0xAB, 0xCD},
		{0x10, 0x20, 0x30, 0x40, 0x50, 0x60},
	}

	for _, original := range tests {
		enc := EncodeBytes(original)
		dec, err := DecodeBytes(enc)
		if err != nil {
			t.Fatalf("DecodeBytes(%q) returned error: %v", enc, err)
		}

		if !bytes.Equal(dec, original) {
			t.Errorf("round-trip EncodeBytes/DecodeBytes failed: got %v, want %v (enc=%q)",
				dec, original, enc)
		}
	}
}

func TestDecodeBytes_InvalidChar(t *testing.T) {
	_, err := DecodeBytes("abc!")
	if err == nil {
		t.Fatalf("expected error for invalid characters in DecodeBytes, got nil")
	}
}

func TestD_Overflow(t *testing.T) {
	s := "LygHa16AHYFz"

	_, err := D(s)
	if err == nil {
		t.Fatalf("expected overflow error for %q, got nil", s)
	}

	// se quiser ser mais específico:
	expected := "bitly: overflow ao decodificar uint64"
	if err.Error() != expected {
		t.Fatalf("expected error %q, got %q", expected, err.Error())
	}
}
