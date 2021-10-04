package main

import "strings"

// 482 https://leetcode-cn.com/problems/license-key-formatting/
func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(s, "-", "")
	if s == "" {
		return ""
	}
	firstCount := len(s) % k
	if firstCount == 0 {
		firstCount = k
	}
	var res = []string{s[:firstCount]}

	for i := 0; i < (len(s)-firstCount)/k; i++ {
		res = append(res, s[firstCount+i*k:firstCount+(i+1)*k])
	}
	return strings.ToUpper(strings.Join(res, "-"))
}
