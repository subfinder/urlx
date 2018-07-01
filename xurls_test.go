package xurls

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
