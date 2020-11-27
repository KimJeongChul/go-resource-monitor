package resourceprofiler

import (
	"strconv"
	"time"

	cerror "github.com/KimJeongChul/go-resource-monitor/error"
	"github.com/KimJeongChul/go-resource-monitor/logger"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

const packageName = "resourceprofiler"

// ResourceProfiler
type ResourceProfiler struct {
	period int
}

// New Create ResourceProfiler
func New(period int) *ResourceProfiler {
	rp := &ResourceProfiler{
		period: period,
	}
	return rp
}

// Monitor Start profiling of computing resource
// cpuUsage, memoryUsage, diskUsage, rxUsage, txUsage
func (r ResourceProfiler) Monitor() {
	ticker := time.NewTicker(time.Duration(r.period) * time.Second)
	exRx, exTx, cErr := r.getNetwork()
	if cErr != nil {
		logger.LogE(packageName, "Monitor", cErr.Error())
	}
	for {
		select {
		case <-ticker.C:
			// CPU
			cpuUsage, cErr := r.getCPU()
			if cErr != nil {
				logger.LogE(packageName, "Monitor", cErr.Error())
			}
			// Memory
			memoryUsage, cErr := r.getMemory()
			if cErr != nil {
				logger.LogE(packageName, "Monitor", cErr.Error())
			}
			// Disk
			diskUsage, cErr := r.getDisk()
			if cErr != nil {
				logger.LogE(packageName, "Monitor", cErr.Error())
			}
			// Network
			curRx, curTx, cErr := r.getNetwork()
			if cErr != nil {
				logger.LogE(packageName, "Monitor", cErr.Error())
			}

			rx := (curRx - exRx) / uint64(r.period)
			rxUsage := strconv.Itoa(int(rx))
			tx := (curTx - exTx) / uint64(r.period)
			txUsage := strconv.Itoa(int(tx))

			logger.LogI(packageName, "Start", "cpu usage:"+cpuUsage+" memory usage:"+memoryUsage+" disk usage:"+diskUsage+" network rx:"+rxUsage+" network tx:"+txUsage)

			exRx = curRx
			exTx = curTx
		}
	}
}

// getCPU Get cpu usages
func (r ResourceProfiler) getCPU() (string, *cerror.CError) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		cErr := cerror.NewCError(cerror.GOPSUTIL_CPU_ERR, err.Error())
		return "", cErr
	}
	cpuUsage := strconv.FormatFloat(cpuPercent[0], 'f', 2, 64)
	return cpuUsage, nil
}

// getMemory Get memory usages
func (r ResourceProfiler) getMemory() (string, *cerror.CError) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		cErr := cerror.NewCError(cerror.GOPSUTIL_MEM_ERR, err.Error())
		return "", cErr
	}
	memoryUsage := strconv.FormatFloat(memInfo.UsedPercent, 'f', 2, 64)
	return memoryUsage, nil
}

// getDisk Get disk usages
func (r ResourceProfiler) getDisk() (string, *cerror.CError) {
	diskInfo, err := disk.Usage("/")
	if err != nil {
		cErr := cerror.NewCError(cerror.GOPSUTIL_DISK_ERR, err.Error())
		return "", cErr
	}
	diskUsage := strconv.FormatFloat(diskInfo.UsedPercent, 'f', 2, 64)
	return diskUsage, nil
}

// getNetwork Get Network tx, rx
func (r ResourceProfiler) getNetwork() (uint64, uint64, *cerror.CError) {
	netInfo, err := net.IOCounters(false)
	if err != nil {
		cErr := cerror.NewCError(cerror.GOPSUTIL_NETWORK_ERR, err.Error())
		return 0, 0, cErr
	}
	// receive
	rx := netInfo[0].BytesRecv * 8
	// transmit
	tx := netInfo[0].BytesSent * 8
	return rx, tx, nil
}
