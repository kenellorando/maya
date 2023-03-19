// aws.go
// AWS API calls

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func describeInstances() (*ec2.DescribeInstancesOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ec2_client := ec2.New(sess)
	result, err := ec2_client.DescribeInstances(nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func describeInstanceStatus(id string) (*ec2.DescribeInstanceStatusOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ec2_client := ec2.New(sess)
	result, err := ec2_client.DescribeInstanceStatus(&ec2.DescribeInstanceStatusInput{
		InstanceIds: []*string{&id},
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func startInstance(id string) (*ec2.StartInstancesOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ec2_client := ec2.New(sess)
	result, err := ec2_client.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: []*string{&id},
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

func stopInstance(id string) (*ec2.StopInstancesOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ec2_client := ec2.New(sess)
	result, err := ec2_client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{&id},
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}
