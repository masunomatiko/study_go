package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(c *Conn) {
	c.respond("220 Service ready for new user.")

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			c.cwd(args)
		case "LIST": // ls
			c.list(args)
		case "LPRT":
			c.lprt(args)
		case "USER":
			c.user(args)
		case "QUIT": // close: QUITコマンドに対するメッセージ送信
			c.respond("221 Service closing control connection.")
			return
		case "RETR": // get
			c.retr(args)
		default:
			c.respond("This command is not implemented")
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}
