/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"strings"
)

const (
	BrokerConfigMapKey           = "bootstrapServers"
	MaxIdleConnectionsKey        = "maxIdleConns"
	MaxIdleConnectionsPerHostKey = "maxIdleConnsPerHost"

	RocketmqChannelSeparator = "."

	knativeRocketmqTopicPrefix = "knative-messaging-rocketmq"

	DefaultMaxIdleConns        = 1000
	DefaultMaxIdleConnsPerHost = 100
	DefaultBrokers             = "127.0.0.1:9876"
)

type RocketmqConfig struct {
	Brokers             []string
	MaxIdleConns        int32
	MaxIdleConnsPerHost int32
}

// GetRocketmqConfig returns the details of the Rocketmq cluster.
func GetRocketmqConfig() (*RocketmqConfig, error) {
	return &RocketmqConfig{
		Brokers:             []string{DefaultBrokers},
		MaxIdleConns:        DefaultMaxIdleConns,
		MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
	}, nil
}

func TopicName(separator, namespace, name string) string {
	topic := []string{knativeRocketmqTopicPrefix, namespace, name}
	return strings.Join(topic, separator)
}
