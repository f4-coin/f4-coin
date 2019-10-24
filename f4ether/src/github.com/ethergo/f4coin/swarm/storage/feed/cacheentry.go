// Copyright 2018 The f4coin Authors
// This file is part of the f4coin library.
//
// The go-f4coin library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-f4coin library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with The go-f4coin library. If not, see <http://www.gnu.org/licenses/>.

package feed

import (
	"bytes"
	"context"
	"time"

	"github.com/ethergo/f4coin/swarm/storage"
)

const (
	hasherCount            = 8
	feedsHashAlgorithm     = storage.SHA3Hash
	defaultRetrieveTimeout = 100 * time.Millisecond
)

// cacheEntry caches the last known update of a specific Swarm feed.
type cacheEntry struct {
	Update
	*bytes.Reader
	lastKey storage.Address
}

// implements storage.LazySectionReader
func (r *cacheEntry) Size(ctx context.Context, _ chan bool) (int64, error) {
	return int64(len(r.Update.data)), nil
}

//returns the feed's topic
func (r *cacheEntry) Topic() Topic {
	return r.Feed.Topic
}
