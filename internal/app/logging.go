package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// LogRecord holds a log record line from the logging file
type LogRecord struct{ Time, Level, MSG string }

func logRecordsFromFile() []LogRecord {
	contents, _ := ioutil.ReadFile("leif.log")

	logRows := strings.Split(strings.ReplaceAll(string(contents), `"`, ""), "\n")

	m := regexp.MustCompile(`(\w+)=([^=]*.)(?:\s|$)`)

	var logRecords []LogRecord
	for _, logRow := range logRows {
		if logRow == "" {
			continue
		}
		args := m.FindAllStringSubmatch(logRow, -1)
		logRecord := LogRecord{Time: args[0][2], Level: args[1][2], MSG: args[2][2]}
		logRecords = append(logRecords, logRecord)
	}

	return logRecords
}

// GetLogRecords collects all logs for a specific log level
func GetLogRecords(level string) string {
	if _, err := logrus.ParseLevel(strings.ToLower(level)); err != nil {
		return fmt.Sprintf("Unsupported log level '%s' "+
			"(supported are: Trace, Debug, Info, Warning, Error, Fatal and Panic)", level)
	}

	logRecords := logRecordsFromFile()
	if len(logRecords) == 0 {
		return "Found no log records, sorry"
	}

	var answer string
	for _, record := range logRecords {
		if strings.ToUpper(record.Level) == strings.ToUpper(level) {
			answer = answer + fmt.Sprintf("```%s %s %s```\n", record.Time, record.Level, record.MSG)
		}
	}

	if answer != "" {
		return answer
	}

	return fmt.Sprintf("Found no %s logs", level)
}

// InitLogger creates a new logrus logger using a multiwriter and returns it
func InitLogger() *logrus.Logger {
	f, _ := os.OpenFile("leif.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	multiWriter := io.MultiWriter(os.Stdout, f)

	logger := logrus.New()

	Formatter := new(logrus.TextFormatter)
	logger.SetFormatter(Formatter)
	logger.SetOutput(multiWriter)

	return logger
}

func logEnabledPlugins(conf Config, logger *logrus.Logger) {
	var enabled []string
	if conf.GitLabToken != "" {
		enabled = append(enabled, "GitLab")
	}
	if conf.JIRAURL != "" && conf.JIRAUsername != "" && conf.JIRAToken != "" {
		enabled = append(enabled, "JIRA")
	}
	if conf.PostMordToken != "" {
		enabled = append(enabled, "PostMord")
	}

	if len(enabled) > 0 {
		logger.Infof("Enabled plugins: %s", strings.Join(enabled, ", "))
	}
}
