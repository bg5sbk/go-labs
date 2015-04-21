package labs29

import (
	"bufio"
	"code.google.com/p/snappy-go/snappy"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"testing"
)

import "time"

/*
{"Time":1429494711, "API":"make_item", "Pid":123, "Actions": [
	{"Type":"Update", "Table":"player_info",
		"OldData":{ "pid":123, "coins":2000 },
		"NewData":{ "pid":123, "coins":1000 }
	},
	{"Type":"Delete", "Table":"player_item",
		"OldData":{ "id":1, "pid":123, "item_id":100 }
	},
	{"Type":"Delete", "Table":"player_item",
		"OldData":{ "id":2, "pid":123, "item_id":200 }
	},
	{"Type":"Insert", "Table":"player_item",
		"NewData":{ "id":3, "pid":123, "item_id":300 }
	}
]}
*/

type Transaction struct {
	Time    int
	API     string
	Pid     int64
	Actions []Action
}

type M map[string]interface{}

type Action struct {
	Type    string
	Table   string
	OldData M `json:",omitempty"`
	NewData M `json:",omitempty"`
}

var TestData = Transaction{
	Time: time.Now().Nanosecond(), API: "make_item", Pid: 123, Actions: []Action{
		{Type: "Update", Table: "player_info",
			OldData: M{"pid": 123, "coins": 2000},
			NewData: M{"pid": 123, "coins": 1000},
		},
		{Type: "Delete", Table: "player_item",
			OldData: M{"id": 2, "pid": 123, "item_id": 100},
		},
		{Type: "Delete", Table: "player_item",
			OldData: M{"id": 2, "pid": 123, "item_id": 200},
		},
		{Type: "Insert", Table: "player_item",
			NewData: M{"id": 3, "pid": 123, "item_id": 300},
		},
	},
}

func Test_Normal(t *testing.T) {
	f, err := os.OpenFile("./json.normal", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fail()
	}
	b := bufio.NewWriter(f)
	e := json.NewEncoder(b)
	for i := 0; i < 10000; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	f.Close()
}

func Test_Gzip_Level1(t *testing.T) {
	f, err := os.OpenFile("./json.gzip1", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fail()
	}
	g, _ := gzip.NewWriterLevel(f, 1)
	b := bufio.NewWriter(g)
	e := json.NewEncoder(b)
	for i := 0; i < 10000; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	g.Flush()
	f.Close()
}

func Test_Gzip_Level5(t *testing.T) {
	f, err := os.OpenFile("./json.gzip5", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fail()
	}
	g, _ := gzip.NewWriterLevel(f, 5)
	b := bufio.NewWriter(g)
	e := json.NewEncoder(b)
	for i := 0; i < 10000; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	g.Flush()
	f.Close()
}

func Test_Gzip_Level9(t *testing.T) {
	f, err := os.OpenFile("./json.gzip9", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fail()
	}
	g, _ := gzip.NewWriterLevel(f, 9)
	b := bufio.NewWriter(g)
	e := json.NewEncoder(b)
	for i := 0; i < 10000; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	g.Flush()
	f.Close()
}

func Test_Snappy(t *testing.T) {
	f, err := os.OpenFile("./json.snappy", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		t.Fail()
	}
	g := snappy.NewWriter(f)
	b := bufio.NewWriter(g)
	e := json.NewEncoder(b)
	for i := 0; i < 10000; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	//g.Flush()
	f.Close()
}

type TestWriter1 struct {
	c int
}

func (w *TestWriter1) Write(p []byte) (n int, err error) {
	w.c += len(p)
	return len(p), nil
}

type TestWriter2 struct {
	w io.Writer
	c int
}

func (w *TestWriter2) Write(p []byte) (n int, err error) {
	w.c += len(p)
	return w.w.Write(p)
}

func Test_Normal_10S(t *testing.T) {
	f := &TestWriter1{}
	e := json.NewEncoder(f)
	t1 := time.Now()
	for {
		e.Encode(TestData)
		if time.Since(t1) >= time.Second*10 {
			break
		}
	}
	t.Log(f.c / 10 / 1024)
}

func Test_Gzip_10S_Level1(t *testing.T) {
	f1 := &TestWriter1{}
	g, _ := gzip.NewWriterLevel(f1, 1)
	f2 := &TestWriter2{w: g}
	e := json.NewEncoder(f2)
	t1 := time.Now()
	for {
		e.Encode(TestData)
		if time.Since(t1) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	t.Log(f2.c / 10 / 1024)
}

func Test_Gzip_10S_Level5(t *testing.T) {
	f1 := &TestWriter1{}
	g, _ := gzip.NewWriterLevel(f1, 5)
	f2 := &TestWriter2{w: g}
	e := json.NewEncoder(f2)
	t1 := time.Now()
	for {
		e.Encode(TestData)
		if time.Since(t1) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	t.Log(f2.c / 10 / 1024)
}

func Test_Gzip_10S_Level9(t *testing.T) {
	f1 := &TestWriter1{}
	g, _ := gzip.NewWriterLevel(f1, 9)
	f2 := &TestWriter2{w: g}
	e := json.NewEncoder(f2)
	t1 := time.Now()
	for {
		e.Encode(TestData)
		if time.Since(t1) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	t.Log(f2.c / 10 / 1024)
}

func Test_Snappy2_10S(t *testing.T) {
	f1 := &TestWriter1{}
	g := snappy.NewWriter(f1)
	f2 := &TestWriter2{w: g}
	e := json.NewEncoder(f2)
	t1 := time.Now()
	for {
		e.Encode(TestData)
		if time.Since(t1) >= time.Second*10 {
			break
		}
	}
	t.Log(f2.c / 10 / 1024)
}
