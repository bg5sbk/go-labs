package labs23

type Conn interface {
	Writer(i int) int
}

type NormalConn struct {
}

func (c *NormalConn) Writer(i int) int {
	return i + 1000
}

type BufferConn struct {
}

func (c *BufferConn) Writer(i int) int {
	return i + 1000
}

type MyObject struct {
	UseBufferConn bool
	Conn          Conn
}

func (obj *MyObject) UseInterface(i int) int {
	if conn, ok := obj.Conn.(*BufferConn); ok {
		return conn.Writer(i)
	}
	return obj.Conn.Writer(i)
}

func (obj *MyObject) UseBoolean(i int) int {
	if obj.UseBufferConn {
		return obj.Conn.(*BufferConn).Writer(i)
	}
	return obj.Conn.Writer(i)
}
