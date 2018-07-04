// Package urlx extracts urls from plain text using regular expressions.
package urlx

import (
	"html"
	"net/url"
	"strings"
	"unicode"
)

// ExtractSubdomains finds all subdomains from a given text
func ExtractSubdomains(text, domain string) (urls []string) {
	allUrls := findAllUrls(text)
	var finalUrls []string

	for _, u := range allUrls {
		finalUrls = append(finalUrls, handleURI(u)...)
	}

	// Filter by domains and remove duplicates
	finalUrls = filterByDomain(finalUrls, domain)

	return finalUrls
}

func findAllUrls(text string) (urls []string) {
	for i, r := range text {
		if r == '.' {
			bck := string(r)
			//Go back till first valid ascii or number
			for backIndex := i - 1; backIndex >= 0; backIndex-- {
				rr := rune(text[backIndex])
				if isValidRuneBack(rr) {
					bck = string(rr) + bck
				} else {
					break
				}
			}

			//Go forth till the last valid ascii or number
			for forwardIndex := i + 1; forwardIndex < len(text); forwardIndex++ {
				rr := rune(text[forwardIndex])
				if isValidRuneForward(rr) {
					bck = bck + string(rr)
				} else {
					break
				}
			}
			urls = append(urls, bck)
		}
	}

	return urls
}

func isValidRuneBack(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r) || r == ':' || r == '/' || r == '_' || r == '-' || r == '%'
}

func isValidRuneForward(r rune) bool {
	return isValidRuneBack(r) || r == '.'
}

func handleURI(u string) []string {
	var urls []string
	// Try to parse as normal URI
	if u, err := url.ParseRequestURI(u); err == nil {
		urls = append(urls, u.Host)
		return urls

	}

	// Html Unescape
	u = html.UnescapeString(u)

	// Query Unescape
	u, _ = url.QueryUnescape(u)

	replacer := strings.NewReplacer(
		"u003d", " ",
		"/", " ",
		"\\", " ",
		"-site:", " ",
		"-www", "www",
	)

	// Suppress bad chars
	u = replacer.Replace(u)

	// Suppress bad starting characters
	u = suppressLeftChar(u)

	// Split on spaces
	return strings.Split(u, " ")
}

func suppressLeftChar(s string) string {
	if strings.HasPrefix(s, "-www") {
		return s[1:]
	}

	if strings.HasPrefix(s, "-site:") {
		return s[6:]
	}

	for i, r := range s {
		if r == '/' {
			return s[i:]
		}
	}

	return s
}

func filterByDomain(urls []string, domain string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, u := range urls {
		if strings.HasSuffix(u, domain) {
			if _, ok := seen[u]; !ok {
				result = append(result, u)
				seen[u] = u
			}
		}
	}
	return result
}
