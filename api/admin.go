// Copyright 2018-2019 The logrange Authors
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

package api

import "context"

type (
	// Admin interface allows to perform some administrative operations. The
	// Execute function accepts a LQL statement and it returns the server response
	// in the result. For 'SELECT' statement please use Querier interface instead
	// of the Execute function.
	Admin interface {
		// Execute runs the lql query on the server and provides the execution result
		// in text form or an error, if the query could not be run. Not all LQL queries
		// could be support
		Execute(ctx context.Context, req ExecRequest) (ExecResult, error)
	}

	// ExecResult struct is used as a response of the request to execute a LQL query
	ExecResult struct {
		// Output contains formatted output result of the command execution
		Output string

		// Err contains the operation error, if any
		Err error `json:"-"`
	}

	// ExecRequest struct must be provided as an input for Admin.Execute request
	ExecRequest struct {
		// Query contains a LQL statement to be executed
		Query string
	}
)
