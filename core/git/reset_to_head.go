// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package git

import "github.com/google/gapid/core/log"

// ResetToHead performs a 'git clean -f -d' followed by 'git checkout .'.
func (g Git) ResetToHead(ctx log.Context) error {
	if _, _, err := g.run(ctx, "clean", "-f", "-d"); err != nil {
		return err
	}
	if _, _, err := g.run(ctx, "checkout", "."); err != nil {
		return err
	}
	return nil
}
