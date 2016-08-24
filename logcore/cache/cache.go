package cache

var LogCache LogContent

type LogContent struct {
	Content []LogInfo
}

type LogInfo struct {
	Time            string
	Type            string
	Condition       string
	StackTrace      string
	OperatingSystem string
	Device          string
	DeviceIP        string
	DeviceModel     string
}
