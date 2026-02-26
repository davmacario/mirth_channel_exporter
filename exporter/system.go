package exporter

/* System stats XML example:
 *
 * <com.mirth.connect.model.SystemStats>
 *   <timestamp>
 *     <time>1772032297737</time>
 *     <timezone>Etc/UTC</timezone>
 *   </timestamp>
 *   <cpuUsagePct>0.04252184090624086</cpuUsagePct>
 *   <allocatedMemoryBytes>200278016</allocatedMemoryBytes>
 *   <freeMemoryBytes>111327256</freeMemoryBytes>
 *   <maxMemoryBytes>268435456</maxMemoryBytes>
 *   <diskFreeBytes>18790436864</diskFreeBytes>
 *   <diskTotalBytes>31526391808</diskTotalBytes>
 * </com.mirth.connect.model.SystemStats>
 */
type SystemStats struct {
	CPUUsagePct          float64              `xml:"cpuUsagePct"`
	AllocatedMemoryBytes int64                `xml:"allocatedMemoryBytes"`
	FreeMemoryBytes      int64                `xml:"freeMemoryBytes"`
	MaxMemoryBytes       int64                `xml:"maxMemoryBytes"`
	DiskFreeBytes        int64                `xml:"diskFreeBytes"`
	DiskTotalBytes       int64                `xml:"diskTotalBytes"`
	Timestamp            SystemStatsTimestamp `xml:"timestamp"`
}

type SystemStatsTimestamp struct {
	Time     int64  `xml:"time"`
	Timezone string `xml:"timezone"`
}
