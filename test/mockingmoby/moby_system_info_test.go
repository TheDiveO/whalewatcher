// Copyright 2021 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockingmoby

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("informs", func() {

	It("returns mocked engine information", func() {
		mm := NewMockingMoby()
		defer mm.Close()
		info, err := mm.Info(context.Background())
		Expect(err).NotTo(HaveOccurred())
		Expect(info.ID).To(HaveLen(6*(4+1+4+1) - 1))
	})

})
