package sshx

import (
	"fmt"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Cli struct {
	user       string
	pwd        string
	addr       string
	client     *ssh.Client
	session    *ssh.Session
	LastResult string
}

func NewCli(addr, user, pwd string) *Cli {
	return &Cli{
		user: user,
		pwd:  pwd,
		addr: addr,
	}
}

func (c *Cli) Connect() error {
	config := &ssh.ClientConfig{}
	config.SetDefaults()

	config.User = c.user
	config.Auth = []ssh.AuthMethod{ssh.Password(c.pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	client, err := ssh.Dial("tcp", c.addr, config)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *Cli) Run(cmds ...string) (string, error) {
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	cmd := strings.Join(cmds, " && ")
	fmt.Printf("cmd: %v\n", cmd)

	buf, err := session.CombinedOutput(cmd)
	c.LastResult = string(buf)
	return c.LastResult, err
}
