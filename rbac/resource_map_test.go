/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rbac_test

import (
	"github.com/go-chassis/cari/rbac"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResource(t *testing.T) {
	t.Run("given empty mapping,return empty resource", func(t *testing.T) {
		res := rbac.GetResource("/test")
		assert.Empty(t, res)
	})
	t.Run("add api and resource mapping,return resource", func(t *testing.T) {
		rbac.MapResource("/v1/order/1", "order")
		res := rbac.GetResource("/v1/order/1")
		assert.NotEmpty(t, res)
	})
	t.Run("add api and resource mapping,return empty resource", func(t *testing.T) {
		rbac.MapResource("/v1/item/1", "order")
		res := rbac.GetResource("/v1/a/1")
		assert.Empty(t, res)
	})
}
