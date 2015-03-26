package labs27

import "bytes"
import "runtime"

func goid() string {
	buf := make([]byte, 15)
	buf = buf[:runtime.Stack(buf, false)]
	return string(bytes.Split(buf, []byte(" "))[1])
}
