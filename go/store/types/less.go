// Copyright 2019 Liquidata, Inc.
//
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
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"github.com/dolthub/dolt/go/store/hash"
)

type kindAndHash interface {
	Kind() NomsKind
	Hash(*NomsBinFormat) (hash.Hash, error)
}

func valueLess(nbf *NomsBinFormat, v1, v2 kindAndHash) (bool, error) {
	switch v2.Kind() {
	case UnknownKind:
		return false, ErrUnknownType

	case BoolKind, FloatKind, StringKind:
		return false, nil

	default:
		h1, err := v1.Hash(nbf)

		if err != nil {
			return false, err
		}

		h2, err := v2.Hash(nbf)

		if err != nil {
			return false, err
		}

		return h1.Less(h2), nil
	}
}
