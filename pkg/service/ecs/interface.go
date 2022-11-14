package ecs

import (
	"awsh/pkg/api/ecs"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ECSServicer interface {
	StartEcs(aws.Config) error
	EcsExec(aws.Config) error
	StopEcsTask(aws.Config) error
}

type ECSService struct {
	Api ecs.ECSApi
}

// constructor関数
func NewECSService(s ecs.ECSApi) *ECSService {
	return &ECSService{Api: s}
}
