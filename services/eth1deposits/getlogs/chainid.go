// Copyright © 2021 Weald Technology Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package getlogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type chainIDResponse struct {
	Result string `json:"result"`
}

// chainID fetches the current block number from an Ethereum 1 client.
func (s *Service) chainID(ctx context.Context) (uint64, error) {
	reference, err := url.Parse("/")
	if err != nil {
		return 0, errors.Wrap(err, "invalid endpoint")
	}
	url := s.base.ResolveReference(reference).String()

	body := bytes.NewBuffer([]byte(`{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1901}`))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return 0, errors.Wrap(err, "failed to setup request context")
	}
	req.Header.Set("Content-type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "failed to call POST endpoint")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, errors.Wrap(err, "failed to read failed POST response")
		}
		return 0, fmt.Errorf("POST failed with status %d: %s", resp.StatusCode, string(data))
	}

	var response chainIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, errors.Wrap(err, "failed to parse newFilter response")
	}

	chainID, err := strconv.ParseUint(strings.TrimPrefix(response.Result, "0x"), 16, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse result")
	}

	return chainID, nil
}