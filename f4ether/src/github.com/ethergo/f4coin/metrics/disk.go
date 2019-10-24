// Copyright 2015 The f4coin Authors
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

package metrics

// DiskStats is the per process disk io stats.
type DiskStats struct {
	ReadCount  int64 // Number of read operations executed
	ReadBytes  int64 // Total number of bytes read
	WriteCount int64 // Number of write operations executed
	WriteBytes int64 // Total number of byte written
}
