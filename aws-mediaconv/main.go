package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

const (
	defaultRegion      = "ap-northeast-1"
	jobTempName        = "icook-media-conv"
	roleName           = "arn:aws:iam::635002287587:role/icook-media-conv"
	endpoint           = "https://4aycjjhd.mediaconvert.ap-northeast-1.amazonaws.com"
	hlsOutputGroupName = "hls-output-group"
)

// Global variable - create job input
var mcCreateJobInput *mediaconvert.CreateJobInput = &mediaconvert.CreateJobInput{
	JobTemplate: aws.String(jobTempName),
	Role:        aws.String(roleName),
	Settings: &mediaconvert.JobSettings{
		Inputs: []*mediaconvert.Input{
			&mediaconvert.Input{
				FileInput: aws.String("s3://media-conv-dst/uploads/video/file/7710aed3-356c-4102-907a-c012b86c23a3/d78cc69497f3dbafac81792a2f16fb78.mp4"),
			},
		},
		OutputGroups: []*mediaconvert.OutputGroup{
			&mediaconvert.OutputGroup{
				OutputGroupSettings: &mediaconvert.OutputGroupSettings{
					HlsGroupSettings: &mediaconvert.HlsGroupSettings{
						Destination: aws.String("s3://media-conv-dst/uploads/video/file/7710aed3-356c-4102-907a-c012b86c23a3/"),
					},
				},
			},
		},
	},
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(endpoint),
		Region:   aws.String(defaultRegion),
	})
	if err != nil {
		panic(err)
	}
	mcsvc := mediaconvert.New(sess)
	/*
		getJobTempOutput, err := mcsvc.GetJobTemplate(&mediaconvert.GetJobTemplateInput{
			Name: aws.String(jobTempName),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(getJobTempOutput.JobTemplate.String())

		createJobOutput, err := mcsvc.CreateJob(mcCreateJobInput)
		if err != nil {
			panic(err)
		}
		fmt.Println(createJobOutput.GoString())
	*/
	getJobResult, err := mcsvc.GetJob(&mediaconvert.GetJobInput{
		Id: aws.String("1524563564417-mw4zsc"),
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(*getJobResult.Job.Settings.OutputGroups[0].OutputGroupSettings.HlsGroupSettings.Destination)
	fmt.Println(getJobResult.GoString())
}
