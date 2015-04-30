package main

import "fmt"

type TransLog interface {
	Commit(*Database)
	Rollback(*Database)
}

type Database struct {
	transLogs  []TransLog
	playerItem map[int]*PlayerItem
}

type PlayerItem struct {
	Id     int
	ItemId int
	Num    int
}

func NewDatabase() *Database {
	return &Database{
		playerItem: make(map[int]*PlayerItem),
	}
}

func (db *Database) Transaction(trans func()) {
	defer func() {
		if err := recover(); err != nil {
			for i := len(db.transLogs) - 1; i >= 0; i-- {
				db.transLogs[i].Rollback(db)
			}
			panic(err)
		} else {
			for _, tl := range db.transLogs {
				tl.Commit(db)
			}
		}
		db.transLogs = db.transLogs[0:0]
	}()
	trans()
}

func (db *Database) LookupPlayerItem(id int) *PlayerItem {
	return db.playerItem[id]
}

func (db *Database) InsertPlayerItem(playerItem *PlayerItem) {
	db.playerItem[playerItem.Id] = playerItem
	db.transLogs = append(db.transLogs, &PlayerItemTransLog{
		Type: INSERT, New: playerItem,
	})
}

func (db *Database) DeletePlayerItem(playerItem *PlayerItem) {
	old := db.playerItem[playerItem.Id]
	delete(db.playerItem, playerItem.Id)
	db.transLogs = append(db.transLogs, &PlayerItemTransLog{
		Type: DELETE, Old: old,
	})
}

func (db *Database) UpdatePlayerItem(playerItem *PlayerItem) {
	old := db.playerItem[playerItem.Id]
	db.playerItem[playerItem.Id] = playerItem
	db.transLogs = append(db.transLogs, &PlayerItemTransLog{
		Type: UPDATE, Old: old, New: playerItem,
	})
}

type TransType int

const (
	INSERT TransType = iota
	DELETE
	UPDATE
)

type PlayerItemTransLog struct {
	Type TransType
	Old  *PlayerItem
	New  *PlayerItem
}

func (transLog *PlayerItemTransLog) Commit(db *Database) {
	switch transLog.Type {
	case INSERT:
		fmt.Printf(
			"INSERT INTO player_item (id, item_id, num) VALUES (%d, %d, %d)\n",
			transLog.New.Id, transLog.New.ItemId, transLog.New.Num,
		)
	case DELETE:
		fmt.Printf(
			"DELETE player_item WHERE id = %d\n",
			transLog.Old.Id,
		)
	case UPDATE:
		fmt.Printf(
			"UPDATE player_item SET id = %d, item_id = %d, num = %d\n",
			transLog.New.Id, transLog.New.ItemId, transLog.New.Num,
		)
	}
}

func (transLog *PlayerItemTransLog) Rollback(db *Database) {
	switch transLog.Type {
	case INSERT:
		delete(db.playerItem, transLog.New.Id)
	case DELETE:
		db.playerItem[transLog.Old.Id] = transLog.Old
	case UPDATE:
		db.playerItem[transLog.Old.Id] = transLog.Old
	}
}

func main() {
	db := NewDatabase()

	db.Transaction(func() {
		db.InsertPlayerItem(&PlayerItem{
			Id:     1,
			ItemId: 100,
			Num:    1,
		})
		db.InsertPlayerItem(&PlayerItem{
			Id:     2,
			ItemId: 100,
			Num:    1,
		})
	})

	fmt.Println(db.playerItem)
	fmt.Println()

	db.Transaction(func() {
		item := db.LookupPlayerItem(1)
		db.DeletePlayerItem(item)
	})

	fmt.Println(db.playerItem)
	fmt.Println()

	func() {
		defer func() {
			recover()
			fmt.Println("rollback")
		}()

		db.Transaction(func() {
			item := db.LookupPlayerItem(2)
			db.DeletePlayerItem(item)
			panic("error")
		})
	}()

	fmt.Println(db.playerItem)
	fmt.Println()
}
