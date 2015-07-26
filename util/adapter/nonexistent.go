// Copyright (c) 2012-2015 The upper.io/db authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package adapter

import (
	"fmt"

	"upper.io/db"
)

var (
	_ = db.Result(&NonExistentResult{})
	_ = db.Collection(&NonExistentCollection{})
)

func err(e error) error {
	if e == nil {
		return db.ErrCollectionDoesNotExist
	}
	return e
}

type NonExistentCollection struct {
	Err error
}

type NonExistentResult struct {
	Err error
}

func (c *NonExistentCollection) Append(interface{}) (interface{}, error) {
	return nil, err(c.Err)
}

func (c *NonExistentCollection) Exists() bool {
	return false
}

func (c *NonExistentCollection) Find(...interface{}) db.Result {
	if c.Err != nil {
		return &NonExistentResult{Err: c.Err}
	}
	return &NonExistentResult{Err: fmt.Errorf("Collection reported an error: %q", err(c.Err))}
}

func (c *NonExistentCollection) Truncate() error {
	return err(c.Err)
}

func (c *NonExistentCollection) Name() string {
	return ""
}

func (r *NonExistentResult) Limit(uint) db.Result {
	return r
}

func (r *NonExistentResult) Skip(uint) db.Result {
	return r
}

func (r *NonExistentResult) Sort(...interface{}) db.Result {
	return r
}

func (r *NonExistentResult) Select(...interface{}) db.Result {
	return r
}

func (r *NonExistentResult) Where(...interface{}) db.Result {
	return r
}

func (r *NonExistentResult) Group(...interface{}) db.Result {
	return r
}

func (r *NonExistentResult) Remove() error {
	return err(r.Err)
}

func (r *NonExistentResult) Update(interface{}) error {
	return err(r.Err)
}

func (r *NonExistentResult) Count() (uint64, error) {
	return 0, err(r.Err)
}

func (r *NonExistentResult) Next(interface{}) error {
	return err(r.Err)
}

func (r *NonExistentResult) One(interface{}) error {
	return err(r.Err)
}

func (r *NonExistentResult) All(interface{}) error {
	return err(r.Err)
}

func (r *NonExistentResult) Close() error {
	return err(r.Err)
}
