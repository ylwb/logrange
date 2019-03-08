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

package runner

import (
	"context"
	"github.com/jrivets/log4g"
	"github.com/logrange/logrange/forwarder"
	"github.com/logrange/logrange/forwarder/scanner"
	"github.com/logrange/logrange/forwarder/sink"
	"github.com/logrange/logrange/pkg/storage"
)

var (
	logger = log4g.GetLogger("forwarder")
)

// Run executes the forwarder process using ctx context provided and the forwarder configuration
func Run(ctx context.Context, cfg *forwarder.Config) error {
	logger.Info("Running.")
	sink, err := createSink(cfg)
	if err != nil {
		return err
	}

	strg, err := storage.NewStorage(cfg.Storage)
	if err != nil {
		return err
	}

	scnr := scanner.NewRpcScanner(cfg, strg, sink)
	err = scnr.Run(ctx)
	if err != ctx.Err() {
		logger.Error("could not start forwarder, err=", err)
	}

	logger.Info("Shutdown.")
	return nil
}

func createSink(cfg *forwarder.Config) (forwarder.Sink, error) {
	return sink.NewStdSkink(), nil
}
