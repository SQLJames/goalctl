// Copyright 2021 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !freebsd
// +build !freebsd

package sqlite // import "modernc.org/sqlite"

func setMaxOpenFiles(n int) error { return nil }
