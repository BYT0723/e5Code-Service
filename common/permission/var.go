package permission

const (
	None = iota
	Exec
	Read
	ReadAndExec
	Write
	WriteAndExec
	ReadAndWrite
	ALL
)

const (
	Up   = 1
	Down = -1
)
