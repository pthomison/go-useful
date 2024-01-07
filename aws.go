package utilkit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type NewConfigInput struct {
	Region string
}

func NewConfig(input NewConfigInput) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(input.Region),
	)
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}

type RequestParameterInput struct {
	Name   string
	Region string
}

func MustRequestParameter(input RequestParameterInput) string {
	cfg, err := NewConfig(NewConfigInput{
		Region: input.Region,
	})
	Check(err)

	ssmsvc := ssm.NewFromConfig(cfg)

	resp, err := ssmsvc.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(input.Name),
		WithDecryption: aws.Bool(true),
	})
	Check(err)

	return *resp.Parameter.Value
}
