// Copyright 2012-2014 Aerospike, Inc.
//
// Portions may be licensed to Aerospike, Inc. under one or more contributor
// license agreements.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package aerospike

type Role string

// Pre-defined user roles.
const (
	// Manage users their roles.
	UserAdmin Role = "user-admin"

	// Manage indicies, user defined functions and server configuration.
	SysAdmin Role = "sys-admin"

	// Allow read and write transactions with the database.
	ReadWrite Role = "read-write"

	// Allow read transactions with the database.
	Read Role = "Read"
)
