package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	bucket   string
	key      string
	region   string
	filename string
)

func main() {
	flag.StringVar(&bucket, "bucket", "", "The bucket where the file is stored")
	flag.StringVar(&region, "region", "us-east-1", "The region of the Lambda function")
	flag.StringVar(&key, "key", "", "The key of the file that should be retrieved")
	flag.StringVar(&filename, "filename", "", "Optional: The name under which the file should be stored")
	flag.Parse()

	svc := s3.New(session.New(), &aws.Config{Region: aws.String(region)})

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	resp, err := svc.GetObject(params)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	if filename != "" {
		ioutil.WriteFile(filename, bytes, 0644)
	} else {
		fmt.Print(string(bytes))
	}
}
