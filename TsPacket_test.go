package main

import (
	"testing"
)

// example TS packet which includes an adaptation field
var sample = []byte{
	0x47, 0x00, 0x30, 0x35, 0x07, 0x10, 0xFC, 0xDA, 0xDE, 0xFB, 0x7E, 0x14,
	0x88, 0x78, 0x23, 0x00, 0x21, 0x04, 0x13, 0x40, 0x0C, 0xD2, 0x43, 0x4E,
	0xAC, 0x8C, 0xEA, 0x69, 0x6C, 0x00, 0x00, 0x01, 0x10, 0x13, 0xFC, 0xC8,
	0xF9, 0x08, 0xB4, 0x4C, 0x92, 0x02, 0x4F, 0xFE, 0xF9, 0x80, 0x0C, 0x54,
	0x92, 0xDB, 0x45, 0x01, 0x97, 0xD2, 0x43, 0x53, 0x9A, 0x4F, 0xA2, 0x6D,
	0x04, 0xBF, 0xFC, 0xF3, 0x00, 0x18, 0xA9, 0x25, 0xAB, 0x21, 0x37, 0x8E,
	0xDF, 0x64, 0x4D, 0x52, 0x4B, 0x7D, 0x01, 0x04, 0xCB, 0x58, 0x51, 0xCD,
	0x9A, 0xC2, 0x7D, 0x46, 0xD8, 0x4D, 0x46, 0xA5, 0xF3, 0x91, 0x1A, 0x9B,
	0xCD, 0xCC, 0x9D, 0x15, 0xE0, 0xD7, 0x3F, 0x00, 0x0F, 0x44, 0x1B, 0xB0,
	0x86, 0xDA, 0x04, 0xC3, 0x38, 0xB4, 0x96, 0xDD, 0x5F, 0x7D, 0x80, 0x0E,
	0xC4, 0xCC, 0xC9, 0xE9, 0x2B, 0x3B, 0x7C, 0xAC, 0x56, 0x83, 0x9B, 0xE0,
	0x91, 0x7A, 0xDC, 0x98, 0xA8, 0xD5, 0x09, 0x15, 0x56, 0x1C, 0xAD, 0xEE,
	0x33, 0x9B, 0x28, 0x14, 0xD9, 0x63, 0xEF, 0x05, 0x16, 0x42, 0x68, 0x64,
	0xDF, 0x0E, 0x3A, 0x91, 0x76, 0x21, 0x0D, 0xB7, 0xF1, 0xD5, 0xF8, 0xE8,
	0x3E, 0xB6, 0xB7, 0x73, 0x7B, 0x8D, 0x73, 0xB4, 0xB0, 0x29, 0xB9, 0xF8,
	0x88, 0x3D, 0xB6, 0xDA, 0x41, 0x91, 0x1B, 0xF2}

func TestNewTsPacketHeader(t *testing.T) {
	hdr := NewTsPacketHeader(sample)
	if hdr.SyncByte != 0x47 {
		t.Error("SyncByte != 0x47!")
	}
	if hdr.Tei != false {
		t.Error("TEI should be false, got true!")
	}
	if hdr.Pusi != false {
		t.Error("PUSI should be false, got true!")
	}
	if hdr.Tp != false {
		t.Error("TP should be false, got true!")
	}
	if hdr.Pid != 0x30 {
		t.Errorf("Expected PID to be 0x30, got %x!", hdr.Pid)
	}
	if hdr.Tsc != 0x00 {
		t.Errorf("Expected TSC to be 0x00, got %x!", hdr.Tsc)
	}
	if hdr.Afc != 0x03 {
		t.Errorf("Expected AFC to be 0x03, got %x!", hdr.Afc)
	}
	if hdr.Cc != 0x05 {
		t.Errorf("Expected CC to be 0x05, got %x!", hdr.Cc)
	}

}

func TestNewAdaptationField(t *testing.T) {
	// TODO find a TS packet example with all fields present
	// only testing what is present in the sample TS packet
	af := NewAdaptationField(sample)

	// length
	if af.Length != 7 {
		t.Errorf("AF length is incorrect, expected 7 got %d", af.Length)
	}

	// indicators
	if af.Di != false {
		t.Errorf("DI should be false, got %t", af.Di)
	}
	if af.Rai != false {
		t.Errorf("Rai should be false, got %t", af.Rai)
	}
	if af.Espi != false {
		t.Errorf("Espi should be false, got %t", af.Espi)
	}

	// option flags
	if af.Pcrf != true {
		t.Errorf("Pcrf should be true, got %t", af.Pcrf)
	}
	if af.Opcrf != false {
		t.Errorf("Opcrf should be false, got %t", af.Opcrf)
	}
	if af.Spf != false {
		t.Errorf("Spf should be false, got %t", af.Spf)
	}
	if af.Tpdf != false {
		t.Errorf("Tpdf should be false, got %t", af.Tpdf)
	}
	if af.Afef != false {
		t.Errorf("Afef should be false, got %t", af.Afef)
	}
	if af.Afef != false {
		t.Errorf("Afef should be false, got %t", af.Afef)
	}

	// optional fields (controled by flags)
	if af.Pcrb != 0x01F9B5BDF6 {
		t.Errorf("Pcrb should be 0x01F9B5BDF6, got %x", af.Pcrb)
	}
	if af.Pcre != 0x14 {
		t.Errorf("Pcre should be 0x14, got %x", af.Pcre)
	}
}

// this test checks whether we handle zero-length adaptation layer
// it is a common error to no handle this correctly
func TestZlal(t *testing.T) {
	zlalSample := sample
	zlalSample[4] = 0

	af := NewAdaptationField(zlalSample)

	// length
	if af.Length != 0 {
		t.Errorf("AF length is incorrect, expected 0 got %d", af.Length)
	}

	// indicators
	if af.Di != false {
		t.Errorf("DI should be false, got %t", af.Di)
	}
	if af.Rai != false {
		t.Errorf("Rai should be false, got %t", af.Rai)
	}
	if af.Espi != false {
		t.Errorf("Espi should be false, got %t", af.Espi)
	}

	// option flags
	if af.Pcrf != false {
		t.Errorf("Pcrf should be false, got %t", af.Pcrf)
	}
	if af.Opcrf != false {
		t.Errorf("Opcrf should be false, got %t", af.Opcrf)
	}
	if af.Spf != false {
		t.Errorf("Spf should be false, got %t", af.Spf)
	}
	if af.Tpdf != false {
		t.Errorf("Tpdf should be false, got %t", af.Tpdf)
	}
	if af.Afef != false {
		t.Errorf("Afef should be false, got %t", af.Afef)
	}
	if af.Afef != false {
		t.Errorf("Afef should be false, got %t", af.Afef)
	}
}

func BenchmarkNewTsPacketHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewTsPacketHeader(sample)
	}
}

func BenchmarkNewAdaptationField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewAdaptationField(sample)
	}
}
