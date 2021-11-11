# plnack-proto
protocol for plnack

**Plnack协议作为常规go服务和json-api服务的通信间协议**

**Make for Renj.io Apps**

### 使用

```bash
go get github.com/Landers1037/plnack-proto
```

### 引入

```go
import pp "github.com/Landers1037/plnack-proto"

type T1 struct {
	Name string
	Age  int
}

func main() {
    j := `{"name": "jiejie"}`
	var td T1
    pp.DecodeAnyJSON(&td, []byte(j))
}
```

