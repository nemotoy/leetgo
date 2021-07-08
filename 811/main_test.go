package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

/*
	## summary
*/
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains)*1)
	for _, domain := range cpdomains {
		raw := strings.Split(domain, " ")
		visit, _ := strconv.Atoi(raw[0])
		domains := strings.Split(raw[1], ".")
		for i := range domains {
			// split by domain and add up the visits.
			m[strings.Join(domains[i:], ".")] += visit
		}
	}
	ret := make([]string, 0, len(m))
	for domain, count := range m {
		ret = append(ret, strconv.Itoa(count)+" "+domain)
	}
	return ret
}

func TestSubdomainVisits(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		// {
		// 	[]string{"9001 discuss.leetcode.com"},
		// 	[]string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		// },
		// {
		// 	[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		// 	[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		// },
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := subdomainVisits(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
