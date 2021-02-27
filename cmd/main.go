package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func listS3Objects(bucket string, region string, profile string) {
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: profile,

		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	if err != nil {
		log.Fatalf("Couldn't initialize aws session with profile [%s], and region [%s]\n", profile, region)
	}
	s3Svc := s3.New(sess)
	fmt.Println(s3Svc)
}
