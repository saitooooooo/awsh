package pkg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"awsh/libs"
)

func Cfg() aws.Config {
	aws_region := libs.ChooseValueFromPromptItems("Select aws region", []string{"ap-northeast-1", "other"})
	if aws_region == "other" {
		aws_region = libs.ChooseValueFromPromptItems("Select aws region", []string{
			"us-east-2",
			"us-east-1",
			"us-west-1",
			"us-west-2",
			"af-south-1",
			"ap-east-1",
			"ap-south-1",
			"ap-northeast-3",
			"ap-northeast-2",
			"ap-southeast-1",
			"ap-southeast-2",
			"ap-northeast-1",
			"ca-central-1",
			"cn-north-1",
			"cn-northwest-1",
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-south-1",
			"eu-west-3",
			"eu-north-1",
			"me-south-1",
			"sa-east-1",
		})
	}

	aws_profile := libs.ChooseValueFromPrompt("Please enter aws profile(If empty, default settings are loaded)", "")
	
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(aws_region), config.WithSharedConfigProfile(aws_profile))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}

