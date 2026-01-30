package consts

const (
	BarkLevelCritical      = "critical"      // 重要警告, 在静音模式下也会响铃
	BarkLevelActive        = "active"        // 默认值，系统会立即亮屏显示通知
	BarkLevelTimeSensitive = "timeSensitive" // 时效性通知，可在专注状态下显示通知
	BarkLevelPassive       = "passive"       // 仅将通知添加到通知列表，不会亮屏提醒
)
