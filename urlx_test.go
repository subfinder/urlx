package urlx

import (
	"strings"
	"testing"
)

func TestExtractSubdomains(t *testing.T) {
	tests := [...]string{
		"http://aa.bb.cc",
		"aa.bb.cc",
		"ftp://aa.bb.cc",
		"aa.bb.cc",
		"aa-bb.cc.dd.cc",
		"1212-aa.bb.345-d.cc",
		"-www.aa.bb.cc",
		"aa.bb.cc.dd/dd.aa.cc",
		"-www.dd.aa.cc",
		"-site:www.dd.aa.cc -site:www.dd.aa.cc",
	}

	for _, test := range tests {
		t.Logf("Current test: %s", test)
		results := ExtractSubdomains(test, "cc")
		if len(results) == 0 {
			t.Errorf("No result found for test: %s", test)
		}
		for _, result := range results {
			t.Logf("Current result: %s", result)
			if !strings.Contains(test, result) {
				t.Errorf("Domain was incorrect, got: %s, want it contained in: %s.", result, test)
			}
		}
	}
}

func BenchmarkExtractSubdomainsWithCheck(b *testing.B) {
	tests := [...]string{
		"http://aa.bb.cc",
		"aa.bb.cc",
		"ftp://aa.bb.cc",
		"aa.bb.cc",
		"aa-bb.cc.dd.cc",
		"1212-aa.bb.345-d.cc",
		"-www.aa.bb.cc",
		"aa.bb.cc.dd/dd.aa.cc",
		"-www.dd.aa.cc",
		"-site:www.dd.aa.cc -site:www.dd.aa.cc",
	}

	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			results := ExtractSubdomains(test, "cc")
			for _, result := range results {
				if !strings.Contains(test, result) {
					b.Errorf("Expected to find a 'cc' domain in '%v'", test)
				}
			}
		}
	}
}

func BenchmarkExtractSubdomains(b *testing.B) {
	tests := [...]string{
		"http://aa.bb.cc",
		"aa.bb.cc",
		"ftp://aa.bb.cc",
		"aa.bb.cc",
		"aa-bb.cc.dd.cc",
		"1212-aa.bb.345-d.cc",
		"-www.aa.bb.cc",
		"aa.bb.cc.dd/dd.aa.cc",
		"-www.dd.aa.cc",
		"-site:www.dd.aa.cc -site:www.dd.aa.cc",
	}

	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			results := ExtractSubdomains(test, "cc")
			if !(len(results) > 1) {
				b.Errorf("expected to be more than one result, got '%v'", len(results))
			}
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel1(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc
		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel2(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc
		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 2x
	for i := 1; i <= 2; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel3(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 3x
	for i := 1; i <= 3; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel4(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 4x
	for i := 1; i <= 4; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel5(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 5x
	for i := 1; i <= 5; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel6(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 6x
	for i := 1; i <= 6; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel7(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 7x
	for i := 1; i <= 7; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel8(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 8x
	for i := 1; i <= 8; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}

func BenchmarkExtractSubdomainsFromBlobOfTextLevel9(b *testing.B) {
	text := `
		http://aa.bb.cc
		aa.bb.cc
		ftp://aa.bb.cc
		aa.bb.cc
		aa-bb.cc.dd.cc
		1212-aa.bb.345-d.cc

		-www.aa.bb.cc
		aa.bb.cc.dd/dd.aa.cc
		-www.dd.aa.cc
		-site:www.dd.aa.cc -site:www.dd.aa.cc
	`
	// grow text 9x
	for i := 1; i <= 9; i++ {
		text += text
	}

	for n := 0; n < b.N; n++ {
		results := ExtractSubdomains(text, "cc")
		if !(len(results) > 1) {
			b.Errorf("expected to be more than one result, got '%v'", len(results))
		}
	}
}
