///
// Copyright 2019 Insolar Technologies GmbH
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
///

package executor

import (
	"context"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/pkg/errors"
)

//go:generate minimock -i github.com/insolar/insolar/ledger/heavy/executor.dropTruncater -o ./ -s _gen_mock.go
type dropTruncater interface {
	TruncateHead(ctx context.Context, lastPulse insolar.PulseNumber) error
}

type DBRollback struct {
	drops     dropTruncater
	jetKeeper JetKeeper
}

func NewDBRollback(drops dropTruncater, jetKeeper JetKeeper) *DBRollback {
	return &DBRollback{
		drops:     drops,
		jetKeeper: jetKeeper,
	}
}

func (d *DBRollback) Start(ctx context.Context) error {
	logger := inslogger.FromContext(ctx)
	pn := d.jetKeeper.TopSyncPulse()

	logger.Debug("[ DBRollback.Start ] last finalized pulse number: ", pn)
	if pn == insolar.GenesisPulse.PulseNumber {
		logger.Debug("[ DBRollback.Start ] No finalized data. Nothing done")
		return nil
	}

	err := d.drops.TruncateHead(ctx, pn)
	if err != nil {
		return errors.Wrapf(err, "can't truncate db to pulse: %d", pn)
	}

	return nil
}
