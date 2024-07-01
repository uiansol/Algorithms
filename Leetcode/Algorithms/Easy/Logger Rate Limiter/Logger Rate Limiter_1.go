type Logger struct {
	table map[string]int
}

func Constructor() Logger {
	return Logger{
		table: make(map[string]int),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	val, ok := this.table[message]
	if !ok || timestamp >= val+10 {
		this.table[message] = timestamp
		return true
	} else {
		return false
	}
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */