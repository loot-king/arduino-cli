// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package builder

import (
	sk "github.com/arduino/arduino-cli/arduino/sketch"
	"github.com/arduino/arduino-cli/legacy/builder/types"
	"github.com/pkg/errors"
)

type SketchLoader struct{}

func (s *SketchLoader) Run(ctx *types.Context) error {
	if ctx.SketchLocation == nil {
		return nil
	}

	sketchLocation := ctx.SketchLocation

	sketchLocation, err := sketchLocation.Abs()
	if err != nil {
		return errors.WithStack(err)
	}
	mainSketchStat, err := sketchLocation.Stat()
	if err != nil {
		return errors.WithStack(err)
	}
	if mainSketchStat.IsDir() {
		sketchLocation = sketchLocation.Join(mainSketchStat.Name() + ".ino")
	}

	ctx.SketchLocation = sketchLocation

	sketch, err := sk.New(sketchLocation)
	if err != nil {
		return errors.WithStack(err)
	}
	ctx.SketchLocation = sketchLocation
	ctx.Sketch = sketch

	return nil
}
