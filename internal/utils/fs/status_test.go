// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package fsutils_test

import (
	fsutils "github.com/TeaOSLab/EdgeNode/internal/utils/fs"
	"github.com/iwind/TeaGo/assert"
	"testing"
	"time"
)

func TestWrites(t *testing.T) {
	var a = assert.NewAssertion(t)

	for i := 0; i < int(fsutils.DiskMaxWrites); i++ {
		fsutils.WriteBegin()
	}
	a.IsFalse(fsutils.WriteReady())

	fsutils.WriteEnd()
	a.IsTrue(fsutils.WriteReady())
}

func TestWaitLoad(t *testing.T) {
	fsutils.WaitLoad(100, 5, 1*time.Minute)
}

func BenchmarkWrites(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fsutils.WriteReady()
			fsutils.WriteBegin()
			fsutils.WriteEnd()
		}
	})
}
