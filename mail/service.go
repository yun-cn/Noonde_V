package mail

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/yun313350095/Noonde/api"
	"html/template"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"regexp"
	"strings"
)

// Service ..
type Service struct {
	Conf   api.ConfService
	Log    api.LogService
	client *ses.SES
}

// Send ..
func (s *Service) Send(name string, data interface{}, locale string, dests []string, from *api.MailPerson, to *api.MailPerson) error {
	var err error

	// Prepare.
	path := s.Conf.String("gotmpl") + "/" + name
	_, err = os.Stat(path + "/" + locale + ".mail")
	if err != nil {
		s.Log.Error(err)
		return err
	}
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Parse mail template.
	var mailBuf bytes.Buffer
	tmpl := template.Must(template.ParseFiles(path + "/" + locale + ".mail"))
	err = tmpl.Execute(&mailBuf, data)
	if err != nil {
		s.Log.Error(err)
		return err
	}
	scanner := bufio.NewScanner(strings.NewReader(mailBuf.String()))
	var sub, text string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			sub = line
			i++
			continue
		}
		if i == 1 && line != "" {
			err = errors.New("There must be empty line after subject")
			s.Log.Error(err)
			return err
		}
		text += line + "\n"
		i++
	}
	sub = strings.TrimSpace(sub)
	text = strings.TrimSpace(text)

	// Build basic part.
	headBasic := make(textproto.MIMEHeader)
	headBasic.Set("From", mime.BEncoding.Encode("UTF-8", from.Name)+" <"+from.Addr+">")
	headBasic.Set("To", mime.BEncoding.Encode("UTF-8", to.Name)+" <"+to.Addr+">")
	headBasic.Set("Return-Path", from.Addr)
	headBasic.Set("Subject", mime.BEncoding.Encode("UTF-8", sub))
	headBasic.Set("Content-Type", "multipart/alternative; boundary=\""+writer.Boundary()+"\"")
	_, err = writer.CreatePart(headBasic)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	// Build text part.
	headText := make(textproto.MIMEHeader)
	headText.Set("Content-Transfer-Encoding", "base64")
	headText.Set("Content-Type", "text/plain; charset=UTF-8")
	part, err := writer.CreatePart(headText)
	if err != nil {
		s.Log.Error(err)
		return err
	}
	_, err = part.Write([]byte(base64.StdEncoding.EncodeToString([]byte(text))))
	if err != nil {
		s.Log.Error(err)
		return err
	}

	// Build html part.
	_, err = os.Stat(path + "/" + locale + ".html")
	if !os.IsNotExist(err) {
		var htmlBuf bytes.Buffer
		tmpl = template.Must(template.ParseFiles(path + "/" + locale + ".html"))
		if err := tmpl.Execute(&htmlBuf, data); err != nil {
			s.Log.Error(err)
			return err
		}
		html := htmlBuf.String()

		// Replace imgs.
		r, _ := regexp.Compile("<img src=\"(.*)\"")
		mm := r.FindAllStringSubmatch(html, -1)
		var files []string
		files = append(files, mm[0][1])
		files = append(files, mm[1][1])
		html = r.ReplaceAllString(html, "<img src=\"cid:${1}\"")
		html = strings.TrimSpace(html)

		// Build html part.
		headHTML := make(textproto.MIMEHeader)
		headHTML.Set("Content-Transfer-Encoding", "base64")
		headHTML.Set("Content-Type", "text/html; charset=UTF-8")
		part, err := writer.CreatePart(headHTML)
		if err != nil {
			s.Log.Error(err)
			return err
		}
		_, err = part.Write([]byte(base64.StdEncoding.EncodeToString([]byte(html))))
		if err != nil {
			s.Log.Error(err)
			return err
		}

		// Build image parts.
		for _, f := range files {
			data, err := ioutil.ReadFile(path + "/" + f) // TODO
			if err != nil {
				s.Log.Error(err)
				return err
			}
			ctype := http.DetectContentType(data)
			headImg := make(textproto.MIMEHeader)
			headImg.Set("Content-Disposition", "inline; filename="+f)
			headImg.Set("Content-Type", ctype+"; name=\""+f+"\"")
			headImg.Set("Content-Transfer-Encoding", "base64")
			headImg.Set("Content-ID", "<"+f+">")
			part, err := writer.CreatePart(headImg)
			if err != nil {
				s.Log.Error(err)
				return err
			}
			_, err = part.Write([]byte(base64.StdEncoding.EncodeToString(data)))
			if err != nil {
				s.Log.Error(err)
				return err
			}
		}
	}

	// Close.
	err = writer.Close()
	if err != nil {
		s.Log.Error(err)
		return err
	}
	str := buf.String()
	if strings.Count(str, "\n") < 2 {
		err = errors.New("Invalid E-mail content")
		s.Log.Error(err)
		return err
	}
	str = strings.SplitN(str, "\n", 2)[1]

	// Build raw input.
	raw := ses.RawMessage{
		Data: []byte(str),
	}
	var dd []*string
	for _, d := range dests {
		dd = append(dd, aws.String(d))
	}
	input := &ses.SendRawEmailInput{
		Destinations: dd,
		Source:       aws.String(from.Addr),
		RawMessage:   &raw,
	}

	// Debug
	// fmt.Println(str)

	// Send.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
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
	client := ses.New(sess)
	_, err = client.SendRawEmail(input)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	for _, d := range dests {
		s.Log.Info("Mail sent to: " + d)
	}

	return nil
}

//--------------------------------------------------------------------------------

// Load ..
func (s *Service) Load() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
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
	s.client = ses.New(sess)
	return nil
}
