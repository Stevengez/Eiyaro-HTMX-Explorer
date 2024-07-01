package models

type RPCMapResponse map[string]any
type RPCResponse[T any] struct {
	Status string
	Data   T
}

// ** Access a nested map of the response map
func (d RPCMapResponse) M(k string) RPCMapResponse {
	return d[k].(map[string]any)
}

// ** Access a value of the rpcResponse map
func (d RPCMapResponse) V(k string) any {
	return d[k]
}
