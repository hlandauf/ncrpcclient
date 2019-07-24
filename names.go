// Copyright (c) 2014-2017 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package nmcrpcclient

import (
	"encoding/json"

	"github.com/btcsuite/btcd/rpcclient"

	"github.com/namecoin/nmcjson"
)

// *********************
// Name Lookup Functions
// *********************

// FutureNameShowResult is a future promise to deliver the result
// of a NameShowAsync RPC invocation (or an applicable error).
type FutureNameShowResult chan *rpcclient.Response

// Receive waits for the Response promised by the future and returns detailed
// information about a name.
func (r FutureNameShowResult) Receive() (*nmcjson.NameShowResult, error) {
	res, err := rpcclient.ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a name_show result object
	var nameShow nmcjson.NameShowResult
	err = json.Unmarshal(res, &nameShow)
	if err != nil {
		return nil, err
	}

	return &nameShow, nil
}

// NameShowAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See NameShow for the blocking version and more details.
func (c *Client) NameShowAsync(name string) FutureNameShowResult {
	cmd := nmcjson.NewNameShowCmd(name, nil)
	return c.SendCmd(cmd)
}

// NameShow returns detailed information about a name.
func (c *Client) NameShow(name string) (*nmcjson.NameShowResult, error) {
	return c.NameShowAsync(name).Receive()
} 