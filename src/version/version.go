// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// version.go [created: Tue, 27 Feb 2018]


package version

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit = "unset"
	// Release is a semantic version of current build
	Release = "unset"
)
