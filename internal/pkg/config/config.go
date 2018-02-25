/*
Load configuration from an SSM parameter

TODO: Documentation
*/
package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/jniedrauer/logs-to-elastic/internal/pkg/aws"
	"github.com/jniedrauer/logs-to-elastic/internal/pkg/env"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var defaultSsmParam string = "/LogsToElastic/config.yml"

type Configuration struct {
	LogstashEndpoint string     `yaml:"logstash"`
	LogLevel         string     `yaml:"log_level"`
	LogGroups        []LogGroup `yaml:"log_groups"`
}

type LogGroup struct {
	Name      string `yaml:"logGroup"`
	IndexName string `yaml:"indexname"`
}

func (c *Configuration) LoadConfig() {
	ssmParam := env.GetEnvOrDefault("SSM_CONFIG_PARAM", defaultSsmParam)

	sess := session.Must(aws.GetSession())
	svc := ssm.New(sess)

	response, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: &ssmParam},
	)
	if err != nil {
		log.Fatalf("Got SSM error: %v", err)
	}

	readConfig(c, response.Parameter.Value)

	log.Info("Initialized config %v", *c)
}

func readConfig(c *Configuration, data *string) {
	err := yaml.Unmarshal([]byte(*data), c)
	if err != nil {
		log.Fatalf("Cannot unmarshal data: %v", err)
	}
}
