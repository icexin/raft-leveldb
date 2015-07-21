package raftleveldb

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/raft/bench"
)

func LeveldbTestStore(b testing.TB) (string, *Store) {
	dir, err := ioutil.TempDir("", "raft")
	if err != nil {
		b.Fatal(err)
	}

	store, err := NewStore(dir)
	if err != nil {
		b.Fatal(err)
	}
	return dir, store
}

func BenchmarkStore_FirstIndex(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.FirstIndex(b, store)
}

func BenchmarkStore_LastIndex(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.LastIndex(b, store)
}

func BenchmarkStore_GetLog(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.GetLog(b, store)
}

func BenchmarkStore_StoreLog(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.StoreLog(b, store)
}

func BenchmarkStore_StoreLogs(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.StoreLogs(b, store)
}

func BenchmarkStore_DeleteRange(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.DeleteRange(b, store)
}

func BenchmarkStore_Set(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.Set(b, store)
}

func BenchmarkStore_Get(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.Get(b, store)
}

func BenchmarkStore_SetUint64(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.SetUint64(b, store)
}

func BenchmarkStore_GetUint64(b *testing.B) {
	dir, store := LeveldbTestStore(b)
	defer store.Close()
	defer os.RemoveAll(dir)

	raftbench.GetUint64(b, store)
}
