package onelog

import "sync"

const (
	// INFO is the numeric code for INFO log level
	INFO = 0x1
	// DEBUG is the numeric code for DEBUG log level
	DEBUG = 0x2
	// WARN is the numeric code for WARN log level
	WARN = 0x4
	// ERROR is the numeric code for WARN log level
	ERROR = 0x8
	// FATAL is the numeric code for WARN log level
	FATAL = 0x10
)

// ALL is a shortcut to INFO | DEBUG | WARN | ERROR | FATAL to enable all logging levels
var ALL = uint8(INFO | DEBUG | WARN | ERROR | FATAL)

// Levels is the mapping between int log levels and their string value
var Levels = make([]string, 256)
var levelsJSON = make([][]byte, 256)
var levelsStreamJSON = make([][]byte, 256)

var levelKey = "level"

var mux = sync.Mutex{}

func init() {
	Levels[INFO] = "info"
	Levels[DEBUG] = "debug"
	Levels[WARN] = "warn"
	Levels[ERROR] = "error"
	Levels[FATAL] = "fatal"
	genLevelSlices()
	genStreamLevelSlices()

}

func genLevelSlices() {
	levelsJSON[INFO] = []byte(`{"` + levelKey + `":"` + Levels[INFO] + `","` + msgKey + `":`)
	levelsJSON[DEBUG] = []byte(`{"` + levelKey + `":"` + Levels[DEBUG] + `","` + msgKey + `":`)
	levelsJSON[WARN] = []byte(`{"` + levelKey + `":"` + Levels[WARN] + `","` + msgKey + `":`)
	levelsJSON[ERROR] = []byte(`{"` + levelKey + `":"` + Levels[ERROR] + `","` + msgKey + `":`)
	levelsJSON[FATAL] = []byte(`{"` + levelKey + `":"` + Levels[FATAL] + `","` + msgKey + `":`)
}

func genStreamLevelSlices() {
	levelsStreamJSON[INFO] = []byte(`"` + levelKey + `":"` + Levels[INFO] + `","` + msgKey + `":`)
	levelsStreamJSON[DEBUG] = []byte(`"` + levelKey + `":"` + Levels[DEBUG] + `","` + msgKey + `":`)
	levelsStreamJSON[WARN] = []byte(`"` + levelKey + `":"` + Levels[WARN] + `","` + msgKey + `":`)
	levelsStreamJSON[ERROR] = []byte(`"` + levelKey + `":"` + Levels[ERROR] + `","` + msgKey + `":`)
	levelsStreamJSON[FATAL] = []byte(`"` + levelKey + `":"` + Levels[FATAL] + `","` + msgKey + `":`)
}