package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
)

// function that sets ("updates") quantity stored in dst
type updFunc func(dst *data.Slice)

// Output Handle for a quantity that is stored on the GPU.
type buffered struct {
	*data.Synced
	updFn updFunc
	autosave
}

func newBuffered(synced *data.Synced, name string, f updFunc) *buffered {
	b := new(buffered)
	b.Synced = synced
	b.name = name
	b.updFn = f
	return b
}

func (b *buffered) update(goodstep bool) {
	dst := b.Write()
	b.updFn(dst)
	b.WriteDone()
	b.touch(goodstep)
}

// notify the handle that it may need to be saved
func (b *buffered) touch(goodstep bool) {
	if goodstep && b.needSave() {
		goSave(b.fname(), b.Read(), Time, func() { b.ReadDone() })
		b.saved()
	}
}

// Memset with synchronization.
func (b *buffered) memset(val ...float32) {
	s := b.Write()
	cuda.Memset(s, val...)
	b.WriteDone()
}

// Normalize with synchronization.
func (b *buffered) normalize() {
	s := b.Write()
	cuda.Normalize(s)
	b.WriteDone()
}