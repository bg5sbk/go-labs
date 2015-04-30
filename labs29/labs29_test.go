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

type M map[string]interface{}

type Transaction struct {
	Time    int
	API     string
	Pid     int64
	Actions []Action
}

type Action struct {
	Type    string
	Table   string
	OldData M `json:",omitempty"`
	NewData M `json:",omitempty"`
}

func (t *Transaction) ByteSize() int {
	s := 0
	for i := 0; i < len(t.Actions); i++ {
		s += t.Actions[i].ByteSize()
	}
	return 8 + 4 + len(t.API) + 8 + 4 + len(t.Actions) + s
}

func (a *Action) ByteSize() int {
	return 4 + len(a.Type) + 4 + len(a.Table) + len(a.OldData)*4 + len(a.OldData)*4
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

const TestLines = 10000

func Test_ByteSize(t *testing.T) {
	t.Log(TestData.ByteSize()*TestLines/1000/1000, "MB")
}

func Test_Normal(t *testing.T) {
	var (
		f, _ = os.OpenFile("./json.normal", os.O_WRONLY|os.O_CREATE, 0777)
		b    = bufio.NewWriter(f)
		e    = json.NewEncoder(b)
	)
	for i := 0; i < TestLines; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	f.Close()
}

func Test_Gzip_Level1(t *testing.T) {
	var (
		f, _ = os.OpenFile("./json.gzip1", os.O_WRONLY|os.O_CREATE, 0777)
		b    = bufio.NewWriter(f)
		g, _ = gzip.NewWriterLevel(b, 1)
		e    = json.NewEncoder(g)
	)
	for i := 0; i < TestLines; i++ {
		e.Encode(TestData)
	}
	g.Flush()
	g.Close()
	b.Flush()
	f.Close()
}

func Test_Gzip_Level5(t *testing.T) {
	var (
		f, _ = os.OpenFile("./json.gzip5", os.O_WRONLY|os.O_CREATE, 0777)
		b    = bufio.NewWriter(f)
		g, _ = gzip.NewWriterLevel(b, 5)
		e    = json.NewEncoder(g)
	)
	for i := 0; i < TestLines; i++ {
		e.Encode(TestData)
	}
	g.Flush()
	g.Close()
	b.Flush()
	f.Close()
}

func Test_Gzip_Level9(t *testing.T) {
	var (
		f, _ = os.OpenFile("./json.gzip9", os.O_WRONLY|os.O_CREATE, 0777)
		b    = bufio.NewWriter(f)
		g, _ = gzip.NewWriterLevel(b, 9)
		e    = json.NewEncoder(g)
	)
	for i := 0; i < TestLines; i++ {
		e.Encode(TestData)
	}
	g.Flush()
	g.Close()
	b.Flush()
	f.Close()
}

func Test_Snappy(t *testing.T) {
	var (
		f, _ = os.OpenFile("./json.snappy", os.O_WRONLY|os.O_CREATE, 0777)
		g    = snappy.NewWriter(f)
		b    = bufio.NewWriter(g)
		e    = json.NewEncoder(b)
	)
	for i := 0; i < TestLines; i++ {
		e.Encode(TestData)
	}
	b.Flush()
	f.Close()
}

type CountWriter struct {
	c int
}

func (w *CountWriter) Write(p []byte) (n int, err error) {
	w.c += len(p)
	return len(p), nil
}

type CountProxyWriter struct {
	w io.Writer
	c int
}

func (w *CountProxyWriter) Write(p []byte) (n int, err error) {
	w.c += len(p)
	return w.w.Write(p)
}

func Test_Normal_10S(t *testing.T) {
	var (
		f = &CountWriter{}
		e = json.NewEncoder(f)
		s = time.Now()
	)
	for {
		e.Encode(TestData)
		if time.Since(s) >= time.Second*10 {
			break
		}
	}
	t.Log(f.c/10/1000/1000, "MB/S")
}

func Test_Gzip_10S_Level1(t *testing.T) {
	var (
		g, _ = gzip.NewWriterLevel(&CountWriter{}, 1)
		f    = &CountProxyWriter{w: g}
		e    = json.NewEncoder(f)
		s    = time.Now()
	)
	for {
		e.Encode(TestData)
		if time.Since(s) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	g.Close()
	t.Log(f.c/10/1000/1000, "MB/S")
}

func Test_Gzip_10S_Level5(t *testing.T) {
	var (
		g, _ = gzip.NewWriterLevel(&CountWriter{}, 5)
		f    = &CountProxyWriter{w: g}
		e    = json.NewEncoder(f)
		s    = time.Now()
	)
	for {
		e.Encode(TestData)
		if time.Since(s) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	g.Close()
	t.Log(f.c/10/1000/1000, "MB/S")
}

func Test_Gzip_10S_Level9(t *testing.T) {
	var (
		g, _ = gzip.NewWriterLevel(&CountWriter{}, 9)
		f    = &CountProxyWriter{w: g}
		e    = json.NewEncoder(f)
		s    = time.Now()
	)
	for {
		e.Encode(TestData)
		if time.Since(s) >= time.Second*10 {
			break
		}
	}
	g.Flush()
	t.Log(f.c/10/1000/1000, "MB/S")
}

func Test_Snappy2_10S(t *testing.T) {
	var (
		g = snappy.NewWriter(&CountWriter{})
		f = &CountProxyWriter{w: g}
		e = json.NewEncoder(f)
		s = time.Now()
	)
	for {
		e.Encode(TestData)
		if time.Since(s) >= time.Second*10 {
			break
		}
	}
	t.Log(f.c/10/1000/1000, "MB/S")
}
