// Copyright 2019 The gotfx Authors
// This file is part of the gotfx library.
//
// The gotfx library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gotfx library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gotfx library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"github.com/gitshock-labs/gotfx/core/rawdb"
	"github.com/gitshock-labs/gotfx/ethdb"
	"github.com/gitshock-labs/gotfx/trie/triedb/hashdb"
)

// newTestDatabase initializes the trie database with specified scheme.
func newTestDatabase(diskdb ethdb.Database, scheme string) *Database {
	db := prepare(diskdb, nil)
	if scheme == rawdb.HashScheme {
		db.backend = hashdb.New(diskdb, db.cleans, mptResolver{})
	}
	//} else {
	//	db.backend = snap.New(diskdb, db.cleans, nil)
	//}
	return db
}
