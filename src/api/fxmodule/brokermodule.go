package fxmodule

import (
	"go.uber.org/fx"
	"siigo.com/kubgo/src/api/config"
	"siigo.com/kubgo/src/domain/kubgo"

	"dev.azure.com/SiigoDevOps/Siigo/_git/go-slim.git/kafka"
	"dev.azure.com/SiigoDevOps/Siigo/_git/go-slim.git/slim"

	"siigo.com/kubgo/src/application/consumer"
)

// BrokerModule Create FX Broker Kafka Module
var BrokerModule = fx.Options(
	fx.Provide(
		NewMessageBus,
	),
	fx.Invoke(),
)

func NewMessageBus(config *config.Configuration) *slim.MessageBusBuilder {
	return slim.
		NewMessageBusBuilder().
		WithProviderKafka(NewKafkaConfig(config)).
		WithConsumer("Kubgo.Created", consumer.TestConsumer).
		WithProduce("Kubgo.Created", &kubgo.CreatedEvent{}).
		WithSsl().
		Build()
}

// NewKafkaConfig Create KafkaConfig Instance
func NewKafkaConfig(config *config.Configuration) *kafka.KafkaConfig {

	kafkaConfig := kafka.NewKafkaConfig()

	if config.KeyVault != nil && config.KeyVault.Kafka != nil && config.KeyVault.Kafka.Enabled {
		azureVaultConfig := &kafka.AzureVaultConfig{
			AzureClientId:     *config.KeyVault.AzureClientId,
			AzureClientSecret: *config.KeyVault.AzureClientSecret,
			AzureTenantId:     *config.KeyVault.AzureTenantId,
			VaultName:         config.KeyVault.VaultName,

			Kafka: config.KeyVault.Kafka,
		}

		kafkaConfig.AzureVaultConfig = azureVaultConfig

	}
	delete(config.Kafka, "brokerUrl")
	for key, value := range config.Kafka {
		_ = kafkaConfig.ConfigMap.SetKey(key, value)
	}

	return kafkaConfig
}
