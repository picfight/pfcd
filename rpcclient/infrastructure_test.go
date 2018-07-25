// Copyright (c) 2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import "testing"

func TestClientStringer(t *testing.T) {
	type test struct {
		url      string
		host     string
		endpoint string
		post     bool
	}
	tests := []test{
		{"https://localhost:9709", "localhost:9709", "", true},
		{"wss://localhost:9709/ws", "localhost:9709", "ws", false},
	}
	for _, test := range tests {
		cfg := &ConnConfig{
			Host:                test.host,
			Endpoint:            test.endpoint,
			HTTPPostMode:        test.post,
			DisableTLS:          false,
			DisableConnectOnNew: true,
		}
		c, err := New(cfg, nil)
		if err != nil {
			t.Errorf("%v rpcclient.New: %v", test.url, err)
			continue
		}
		s := c.String()
		if s != test.url {
			t.Errorf("Expected %q, got %q", test.url, s)
		}
	}
}
