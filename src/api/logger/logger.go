package logger

import (
	"dev.azure.com/SiigoDevOps/Siigo/_git/Siigo.Core.Logs.Golang.git/easy"
	"github.com/ic2hrmk/promtail"
	"github.com/sirupsen/logrus"
	"os"
	"siigo.com/kubgo/src/api/config"
)

const BusinessFieldKey = "business"
const BusinessFieldValue = "yes"

type BusinessHook struct {
	client promtail.Client
}

// NewLogrus Create a new instance of logrus with custom hooks
func NewLogrus(config *config.Configuration) *logrus.Logger {

	logger := logrus.New()
	logger.SetFormatter(&easy.Formatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	logger.AddHook(NewBusinessHook(config))

	return logger
}

func NewBusinessHook(config *config.Configuration) *BusinessHook {

	promtailClient, err := promtail.NewJSONv1Client(config.BusinessLogger.LokiUrl, config.BusinessLogger.DefaultLabels)

	if err != nil {
		panic(err)
	}

	return &BusinessHook{
		client: promtailClient,
	}

}

func (hook *BusinessHook) Fire(entry *logrus.Entry) error {

	if _, exist := entry.Data[BusinessFieldKey]; !exist {
		return nil
	}

	go hook.client.LogfWithLabels(promtail.Info, FieldsToMap(entry.Data), entry.Message)
	return nil
}

func (hook *BusinessHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func FieldsToMap(fields logrus.Fields) map[string]string {
	mapFields := make(map[string]string)
	for key, element := range fields {
		mapFields[key] = element.(string)
	}
	return mapFields
}

func WithBusinessFields(fields logrus.Fields) logrus.Fields {
	fields[BusinessFieldKey] = BusinessFieldValue
	return fields
}
