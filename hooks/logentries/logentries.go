package logrus_logentries

import (
	"fmt"
	"net"
	"os"
	"github.com/rfsbraz/logrus"
)

// Logentries to send logs via the Token-Based interface
type LogentriesHook struct {
	Token string
	TCPConn net.Conn
}

// NewLogEntriesHook creates a hook to be added to an instance of logger.
func NewLogentriesHook(token string) (*LogentriesHook, error) {
	host := "data.logentries.com"
	port := 10000

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))

	return &LogentriesHook{token, conn}, err
}

// Fire is called when a log event is fired.
func (hook *LogentriesHook) Fire(entry *logrus.Entry) error {
	msg, _ := entry.String()
	payload := fmt.Sprintf( hook.Token + "%s", msg)

	bytesWritten, err := hook.TCPConn.Write([]byte(payload))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to send log line to Papertrail via UDP. Wrote %d bytes before error: %v", bytesWritten, err)
		return err
	}

	return nil
}

// Levels returns the available logging levels.
func (hook *LogentriesHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
