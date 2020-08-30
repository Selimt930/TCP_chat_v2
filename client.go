package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	nick string
	room *room
	commands chan<-command
}


func (c *client)readInput(){
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil{
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg," ")
		cmd := strings.TrimSpace(args[0])

		var comMap = map[string]command {
			"/nick" : {
				id: NICK,
				client: c,
				args: args,
			},
			"/join":{
				id: JOIN,
				client: c,
				args: args,
			},
			"/rooms":{
				id: ROOMS,
				client: c,
			},
			"/msg":{
				id: MSG,
				client: c,
				args: args,
			},
			"/quit":{
				id:QUIT,
				client: c,
			},
			"/members":{
				id:MEMBERS,
				client: c,
			},

		}

		k, ok := comMap[cmd]
		if !ok{
			c.err(fmt.Errorf("unknown command: %s", cmd))
		} else{
			c.commands <- k
		}


/*
		switch cmd{
		case "/nick":
			c.commands <- command{
				id: NICK,
				client: c,
				args: args,
			}
		case "/join":
			c.commands <-command{
				id: JOIN,
				client: c,
				args: args,
			}
		case "/rooms":
			c.commands <- command{
				id: ROOMS,
				client: c,
			}
		case "/msg":
			c.commands <- command{
				id: MSG,
				client: c,
				args: args,
			}
		case "/quit":
			c.commands <- command{
				id:QUIT,
				client: c,
			}
		case "/members":
			c.commands <-command{
				id:MEMBERS,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))

		}

 */
	}
}

func (c *client) err (err error){
	c.conn.Write([]byte("err:" + err.Error() + "\n"))
}

func (c *client) msg (msg string){
	c.conn.Write([]byte(">" + msg + "\n"))
}

