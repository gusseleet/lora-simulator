package logging
import (
	"github.com/op/go-logging"
	"os"
)
var log = logging.MustGetLogger("example")


var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init(){
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)
	logging.SetLevel(loggingLevel(), "example")


}

func Log() (*logging.Logger) {

	return log
}

func loggingLevel() (logging.Level){
	level := os.Getenv("DEBUG_LEVEL")

	switch level {
	case "ERROR":
		return logging.ERROR
	case "INFO":
		return logging.INFO
	case "DEBUG":
		return logging.DEBUG
	default:
		return logging.DEBUG

	}

}