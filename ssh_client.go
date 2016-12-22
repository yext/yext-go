package yext

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func NewSSHClient(username string, password string, host string, port string) (*sftp.Client, error) {
	c := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				// Just send the password back for all questions
				answers := make([]string, len(questions))
				for i, _ := range answers {
					answers[i] = password
				}

				return answers, nil
			}),
		},
	}

	hostAndPort := host + ":" + port
	connection, err := ssh.Dial("tcp", hostAndPort, c)
	if err != nil {
		return nil, err
	}

	server, err := sftp.NewClient(connection)
	if err != nil {
		return nil, err
	}

	return server, nil
}
