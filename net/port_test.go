/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package net

import "testing"

func TestLocalPortString(t *testing.T) {
	testCases := []struct {
		description string
		ip          string
		family      IPFamily
		port        int
		protocol    string
		expectedStr string
		expectedErr bool
	}{
		{"IPv4 UDP", "1.2.3.4", "", 9999, "udp", `"IPv4 UDP" (1.2.3.4:9999/udp)`, false},
		{"IPv4 TCP", "5.6.7.8", "", 1053, "tcp", `"IPv4 TCP" (5.6.7.8:1053/tcp)`, false},
		{"IPv6 TCP", "2001:db8::1", "", 80, "tcp", `"IPv6 TCP" ([2001:db8::1]:80/tcp)`, false},
		{"IPv4 SCTP", "9.10.11.12", "", 7777, "sctp", `"IPv4 SCTP" (9.10.11.12:7777/sctp)`, false},
		{"IPv6 SCTP", "2001:db8::2", "", 80, "sctp", `"IPv6 SCTP" ([2001:db8::2]:80/sctp)`, false},
		{"IPv4 TCP, all addresses", "", IPv4, 1053, "tcp", `"IPv4 TCP, all addresses" (:1053/tcp4)`, false},
		{"IPv6 TCP, all addresses", "", IPv6, 80, "tcp", `"IPv6 TCP, all addresses" (:80/tcp6)`, false},
		{"No ip family TCP, all addresses", "", "", 80, "tcp", `"No ip family TCP, all addresses" (:80/tcp)`, false},
		{"IP family mismatch", "2001:db8::2", IPv4, 80, "sctp", "", true},
		{"IP family mismatch", "1.2.3.4", IPv6, 80, "sctp", "", true},
		{"Unsupported protocol", "2001:db8::2", "", 80, "http", "", true},
		{"Invalid IP", "300", "", 80, "tcp", "", true},
		{"Invalid ip family", "", "5", 80, "tcp", "", true},
	}

	for _, tc := range testCases {
		lp, err := NewLocalPort(
			tc.description,
			tc.ip,
			tc.family,
			tc.port,
			tc.protocol,
		)
		if tc.expectedErr {
			if err == nil {
				t.Errorf("Expected err when creating LocalPort %v", tc)
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected err when creating LocalPort %s", err)
			continue
		}
		str := lp.String()
		if str != tc.expectedStr {
			t.Errorf("Unexpected output for %s, expected: %s, got: %s", tc.description, tc.expectedStr, str)
		}
	}
}
