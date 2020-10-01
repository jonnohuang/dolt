// Copyright 2020 Liquidata, Inc.
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


package nbs

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
)

type gcErrAccum map[string]error

var _ error = gcErrAccum{}

func (ea gcErrAccum) add(path string, err error) {
	ea[path] = err
}

func (ea gcErrAccum) isEmpty() bool {
	return len(ea) == 0
}

func (ea gcErrAccum) Error() string {
	var sb strings.Builder
	sb.WriteString("error garbage collecting the following files:")
	for filePath, err := range ea {
		sb.WriteString(fmt.Sprintf("\t%s: %s", filePath, err.Error()))
	}
	return sb.String()
}

type gcCopier struct {
	mt     *memTable
	mtSize uint64

	tables  tableSet
	ftp     *fsTablePersister
	tmpDir  string

	stats  *Stats
}

func newGarbageCollectionCopier(tmpDir string, mtSize uint64) gcCopier {
	// todo: minimize table index memory overhead
	indexCache := newIndexCache(defaultIndexCacheSize)
	// todo: are FD's valid after copy?
	fdCache := newFDCache(defaultMaxTables)
	ftp := &fsTablePersister{tmpDir, fdCache, indexCache}

	return gcCopier{
		mtSize:  mtSize,
		tables:  newTableSet(ftp),
		ftp:     ftp,
		tmpDir:  tmpDir,
		stats:   NewStats(),
	}
}

func (gcc gcCopier) addChunk(ctx context.Context, h addr, data []byte) bool {
	if gcc.mt == nil {
		gcc.mt = newMemTable(gcc.mtSize)
	}
	if !gcc.mt.addChunk(h, data) {
		gcc.tables = gcc.tables.Prepend(ctx, gcc.mt, gcc.stats)
		gcc.mt = newMemTable(gcc.mtSize)
		return gcc.mt.addChunk(h, data)
	}
	return true
}

func (gcc gcCopier) copyTablesToDir(destDir string) error {
	specs, err := gcc.tables.ToSpecs()

	if err != nil {
		return err
	}

	for _, spec := range specs {
		tmp := path.Join(gcc.tmpDir, spec.name.String())
		dest := path.Join(destDir, spec.name.String())

		// if copy does not complete, new files will be orphaned
		err = os.Rename(tmp, dest)

		if err != nil {
			return err
		}
	}

	return nil
}