package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	dstBucket = "media-conv-dst"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	if err != nil {
		panic(err)
	}

	s3svc := s3.New(sess)

	/* Get bucket policy
	getBucketPolicyOutput, err := s3svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
		Bucket: aws.String("media-conv-dst"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(getBucketPolicyOutput.GoString())
	*/

	getBucketAclOutput, err := s3svc.GetBucketAcl(&s3.GetBucketAclInput{
		Bucket: aws.String(dstBucket),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(getBucketAclOutput.GoString())

	// Get object list
	listObjectV2Output, err := s3svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(dstBucket),
		Prefix: aws.String("uploads/video/file/c7d4e30d-eaa8-4d76-8415-6156044fef20"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(listObjectV2Output.GoString())
	/* Put acl on multiple objects
	for _, object := range listObjectV2Output.Contents {
		//fmt.Println(*object.Key)
		_, err := s3svc.PutObjectAcl(&s3.PutObjectAclInput{
			ACL:    aws.String("public-read"),
			Bucket: aws.String(dstBucket),
			Key:    object.Key,
		})
		if err != nil {
			panic(err)
		}
		//fmt.Println(putObjectAclOutput.GoString())
	}
	*/

	/* Put acl on single object
	putObjectAclOutput, err := s3svc.PutObjectAcl(&s3.PutObjectAclInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(dstBucket),
		Key:    aws.String("upload/video/file/04d08eab-b309-4fbc-bf5b-21321d1e77e9/*"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(putObjectAclOutput.GoString())
	*/

	/* Put bucket acl
	putBucketAclOutput, err := s3svc.PutBucketAcl(&s3.PutBucketAclInput{
		Bucket: aws.String("media-conv-dst"),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(putBucketAclOutput.GoString())
	*/
}
