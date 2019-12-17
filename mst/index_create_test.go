package mst

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"testing"
)

func TestIndexCreate(T *testing.T) {
	db_invert, _ := leveldb.OpenFile("path/to/db_invert", nil)
	defer db_invert.Close()
	db_mst, _ := leveldb.OpenFile("path/to/db_mst", nil)
	defer db_mst.Close()
	index := Inverted_list{Db: db_invert}
	index.RenewList()
	fmt.Println(index)
	mst := MST{RootHash: [32]byte{}, Db:db_mst}
	mst.ReNewMst()
	mst.printMst()
	CreateIndex(File{Name: "tom",Keys: []string{"block","chain"}},88,&index,&mst)
	mst.printMst()
	fmt.Println(mst.search([]string{"篮球"}))
}