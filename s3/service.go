package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/yun313350095/Noonde/api"
	"io"
	"strings"
	"time"
)

// Service ..
type Service struct {
	Conf     api.ConfService
	Log      api.LogService
	client   *s3.S3
	uploader *s3manager.Uploader
}

// Delete ..
func (s *Service) Delete(path string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.Conf.String("s3.bucket")),
		Key:    aws.String(path),
	}
	_, err := s.client.DeleteObject(input)
	if err != nil {
		s.Log.Error(err)
		return err
	}
	s.Log.Info("Delete from S3: " + path)

	return nil
}

//  Presign ..
func (s *Service) Presign(path string) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Conf.String("s3.bucket")),
		Key:    aws.String(path),
	})

	return req.Presign(1440 * time.Minute)
}

// Upload ..
func (s *Service) Upload(path string, body io.Reader) error {
	ctype := s.contentType(path)

	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.Conf.String("s3.bucket")),
		Key:         aws.String(path),
		ContentType: &ctype,
		Body:        body,
	})
	if err != nil {
		s.Log.Error(err)
		return err
	}

	s.Log.Info("Uploaded to S3: " + path)

	return nil
}

//--------------------------------------------------------------------------------
// Load ..
func (s *Service) Load() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(
			s.Conf.String("aws.access_key_id"),
			s.Conf.String("aws.secret_access_key"),
			"",
		),
	})
	if err != nil {
		s.Log.Error(err)
		return err
	}
	s.client = s3.New(sess)
	s.uploader = s3manager.NewUploader(sess)
	return nil
}

//--------------------------------------------------------------------------------

func (s *Service) contentType(path string) string {
	if strings.Contains(path, ".png") {
		return "image/png"
	}

	return "image/png"
}
