package raftleveldb

import (
	"github.com/hashicorp/raft"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Store struct {
	db *leveldb.DB
}

func NewStore(dir string) (*Store, error) {
	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		return nil, err
	}
	return &Store{
		db: db,
	}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) FirstIndex() (uint64, error) {
	iter := s.db.NewIterator(nil, nil)
	defer iter.Release()
	ok := iter.First()
	if !ok {
		return 0, nil
	}

	return bytesToUint64(iter.Key()), nil
}

func (s *Store) LastIndex() (uint64, error) {
	iter := s.db.NewIterator(nil, nil)
	defer iter.Release()
	ok := iter.Last()
	if !ok {
		return 0, nil
	}

	return bytesToUint64(iter.Key()), nil
}

func (s *Store) GetLog(index uint64, log *raft.Log) error {
	key := uint64ToBytes(index)
	v, err := s.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return raft.ErrLogNotFound
	} else if err != nil {
		return err
	}
	return decodeMsgPack(v, log)
}

func (s *Store) StoreLog(log *raft.Log) error {
	return s.StoreLogs([]*raft.Log{log})
}

func (s *Store) StoreLogs(logs []*raft.Log) error {
	batch := new(leveldb.Batch)
	for _, log := range logs {
		key := uint64ToBytes(log.Index)
		val, err := encodeMsgPack(log)
		if err != nil {
			return err
		}
		batch.Put(key, val.Bytes())
	}
	return s.db.Write(batch, nil)
}

func (s *Store) DeleteRange(min, max uint64) error {
	Range := util.Range{
		Start: uint64ToBytes(min),
		Limit: uint64ToBytes(max),
	}
	batch := new(leveldb.Batch)

	iter := s.db.NewIterator(&Range, nil)
	defer iter.Release()

	for iter.Next() {
		batch.Delete(iter.Key())
	}
	// leveldb的range不包含Limit
	batch.Delete(uint64ToBytes(max))
	return s.db.Write(batch, nil)
}

func (s *Store) Set(key []byte, val []byte) error {
	return s.db.Put(key, val, nil)
}

func (s *Store) Get(key []byte) ([]byte, error) {
	return s.db.Get(key, nil)
}

func (s *Store) SetUint64(key []byte, val uint64) error {
	return s.Set(key, uint64ToBytes(val))
}

func (s *Store) GetUint64(key []byte) (uint64, error) {
	val, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	return bytesToUint64(val), nil
}
