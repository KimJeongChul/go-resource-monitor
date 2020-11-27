package error

import "strconv"

type ErrCode int

const (
	// OS_OPEN_ERR cannot open file os.Open() error
	OS_OPEN_ERR = iota

	// GOPSUTIL_CPU_ERR Cannot get cpu percent gopsutil.cpu.Percent()
	GOPSUTIL_CPU_ERR

	// GOPSUTIL_MEM_ERR Cannot get memory usage gopsutil.mem.VirtualMemory()
	GOPSUTIL_MEM_ERR

	// GOPSUTIL_DISK_ERR Cannot get disk usage gopsutil.disk.Usage()
	GOPSUTIL_DISK_ERR

	// GOPSUTIL_NETWORK_ERR Cannot get network IO counter gopsutil.net.IOCounters()
	GOPSUTIL_NETWORK_ERR
)

type CError struct {
	Code    ErrCode
	Message string
}

func NewCError(code ErrCode, msg string) *CError {
	return &CError{
		Code:    code,
		Message: msg,
	}
}

func (e *CError) Error() string {
	return "CODE:" + strconv.Itoa(int(e.Code)) + ", MSG:" + e.Message
}
