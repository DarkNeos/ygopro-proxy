package util

type Context struct {
	ContextId         int64
	WsConnected       bool
	TcpConnected      bool
	InfaReadBufferLen int
}

func NewContext() Context {
	return Context{
		ContextId:    0, // todo
		WsConnected:  false,
		TcpConnected: false,
	}
}
