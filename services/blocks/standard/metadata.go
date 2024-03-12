// Copyright © 2020 Weald Technology Trading.
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

package standard

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// Metadata stores information about this service.
type Metadata struct {
	LatestSlot int64 `json:"latest_slot"`
}

// MetadataKey is the key for the metadata.
const MetadataKey = "blocks.standard"

// GetMetadata gets metadata for this service.
func (s *Service) GetMetadata(ctx context.Context) (*Metadata, error) {
	md := &Metadata{
		LatestSlot: -1,
	}
	mdJSON, err := s.chainDB.Metadata(ctx, MetadataKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch metadata")
	}
	if mdJSON == nil {
		return md, nil
	}
	if err := json.Unmarshal(mdJSON, md); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal metadata")
	}
	return md, nil
}

// SetMetadata sets metadata for this service.
func (s *Service) SetMetadata(ctx context.Context, md *Metadata) error {
	mdJSON, err := json.Marshal(md)
	if err != nil {
		return errors.Wrap(err, "failed to marshal metadata")
	}
	if err := s.chainDB.SetMetadata(ctx, MetadataKey, mdJSON); err != nil {
		return errors.Wrap(err, "failed to update metadata")
	}
	return nil
}

