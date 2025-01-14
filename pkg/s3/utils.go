package s3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
)

func DefaultAWSConfig() (aws.Config, error) {
	// Set a default IMDC TTL of 1 hour if not set to avoid hitting the metadata service too often, which can slow down the
	// node's startup time.
	if _, ok := os.LookupEnv("AWS_EC2_METADATA_TTL"); !ok {
		err := os.Setenv("AWS_EC2_METADATA_TTL", "3600")
		if err != nil {
			return aws.Config{}, err
		}
	}
	var optFns []func(*config.LoadOptions) error
	return config.LoadDefaultConfig(context.Background(), optFns...)
}

// HasValidCredentials returns true if the AWS config has valid credentials.
func HasValidCredentials(config aws.Config) bool {
	credentials, err := config.Credentials.Retrieve(context.Background())
	if err != nil {
		log.Debug().Err(err).Msg("Failed to check if we have valid AWS credentials")
		return false
	}
	return credentials.HasKeys()
}
