package ldb

import (
	"encoding/json"
	"fmt"
	"github.com/dnsdao/nft.pass/config"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strconv"
	"sync"
)

type Ldb struct {
	db *leveldb.DB
}

var (
	ldb     *Ldb
	ldbOnce sync.Once
)

func GetLdb() *Ldb {
	if ldb == nil {
		ldbOnce.Do(func() {
			opts := opt.Options{
				Strict:      opt.DefaultStrict,
				Compression: opt.NoCompression,
				Filter:      filter.NewBloomFilter(10),
			}
			db, err := leveldb.OpenFile(config.GetRConf().GetDBPath(), &opts)
			if err != nil {
				panic(err)
			}
			ldb = &Ldb{
				db: db,
			}

		})
	}

	return ldb
}

func (db *Ldb) CloseLdb() {
	if ldb != nil {
		db.db.Close()
	}
	ldb = nil
}

func (db *Ldb) GetLatestBlkNum(BlKnumberKey string) (blk uint64) {
	if v, err := db.db.Get([]byte(BlKnumberKey), nil); err != nil {
		return 0
	} else {
		if blk, err = strconv.ParseUint(string(v), 10, 64); err != nil {
			return 0
		}
	}
	return
}

func (db *Ldb) SaveLatestBlkNum(BlKnumberKey string, blk uint64) error {

	if err := db.db.Put([]byte(BlKnumberKey), []byte(strconv.FormatUint(blk, 10)), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}

	return nil
}

// get key
func (db *Ldb) GetKey(key string) (string, error) {
	if v, err := db.db.Get([]byte(key), nil); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}

func (db *Ldb) SaveMintCard(CardId string, val string) error {
	if err := db.db.Put([]byte(CardId), []byte(val), &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}

	return nil
}

func (db *Ldb) GetS3c() (*S3, error) {
	var ret *S3
	k := fmt.Sprintf("cS3")
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveS3c(ct *S3) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf("cS3")
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetS3m() (*S3, error) {
	var ret *S3
	k := fmt.Sprintf("mS3")
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveS3m(ct *S3) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf("mS3")
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

func (db *Ldb) GetIpfsc() (*Ipfs, error) {
	var ret *Ipfs
	k := fmt.Sprintf("cipfs")
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveIpfsc(ct *Ipfs) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf("cipfs")
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}
func (db *Ldb) GetIpfsm() (*Ipfs, error) {
	var ret *Ipfs
	k := fmt.Sprintf("mipfs")
	if v, err := db.db.Get([]byte(k), nil); err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(v, &ret)
		if err != nil {
			return nil, err
		}

		return ret, nil
	}
}

func (db *Ldb) SaveIpfsm(ct *Ipfs) error {
	v, _ := json.Marshal(ct)
	k := fmt.Sprintf("mipfs")
	if err := db.db.Put([]byte(k), v, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}
