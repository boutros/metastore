/*
Copyright 2013 Tristan Wietsma

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metastore

import "testing"

func TestMetaStore(t *testing.T) {
	var M MetaStore
	M.Init(1000)

	h := M.GetHasher()
	idx := h([]byte("key123"))
	if idx != 334 {
		t.Errorf("Wrong bucket index. Got %d; should be 334.", idx)
		return
	}

	M.Bucket[idx].Set("key123", "value567")
	value, ok := M.Bucket[idx].Get("key123")
	if value != "value567" || !ok {
		t.Errorf("Set-Get failed.")
	}
}

func BenchmarkSet(b *testing.B) {
	var M MetaStore
	M.Init(1000)
	h := M.GetHasher()
	idx := h([]byte("key123"))
	M.Bucket[idx].Set("key123", "value567")
}

func TestMetaStoreFlushAll(t *testing.T) {
	var M MetaStore
	M.Init(1000)
	M.Set("key123", "value567")
	M.FlushAll()
	_, ok := M.Get("key123")
	if ok {
		t.Errorf("FlushAll failed")
	}
}
