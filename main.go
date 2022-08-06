package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Message struct {
	Name string `json:"name"`
	Text string `json:"message"`
}
type Username struct {
	Name        string `json:"username"`
	Msg         string `json:"msg"`
	Pic         string `json:"dp"`
	No          string `json:"clientno"`
	CurrentUser string `json:"current_user_name"`
}
type GroupUser struct {
	Name        string `json:"username"`
	Msg         string `json:"msg"`
	Pic         string `json:"dp"`
	No          string `json:"clientno"`
	CurrentUser string `json:"current_user_name"`
	GroupId     string `json:"groupid"`
}
type Group struct {
	Room string `json:"room"`
	Name string `json:"name"`
}

var (
	fileInfo *os.FileInfo
	err      error
)

// send email
func sendemail(tomail string, textfile string, c *gosocketio.Channel) {
	log.Println("email Started")
	var htmlBody = `
	<html>
	<head>
	   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	   <title>Chat Conversation</title>
	</head>
	<body>
	   <p>PFA</p>
	</body>
	`
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "golangchatapplication@gmail.com"
	server.Password = "golangchatapplication123"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		c.Emit("email-fail", "fail")
		log.Println(err)
	}

	// Create email host
	email := mail.NewMSG()
	email.SetFrom("From Me <gochat@host.com>")
	email.AddTo(tomail)
	email.AddCc("golangchatapplication@gmail.com")
	email.SetSubject("GO chat convo")

	email.SetBody(mail.TextHTML, htmlBody)

	if _, err := os.Stat(textfile); err == nil {
		email.AddAttachment(textfile)
	}

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		c.Emit("email-fail", "fail")
		log.Println(err)
	}
	log.Println("email sent")
}

// logger
func logging() {
	// log
	file, err := os.OpenFile("chatLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Logs Started")
}

// mux router serve static files
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/views/chat"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}

// routing
func routing(server *gosocketio.Server) {
	router := NewRouter()
	router.Handle("/socket.io/", server)
	log.Println("server started at 80")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

// check username already exsists
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// get the index of room
func indexOf(clinetid string, cids map[string]string) int {
	count := 0
	for k, _ := range cids {
		if clinetid == k {
			return count
		}
		count = count + 1
	}
	return count
}

// Get the list of clients and check whetehr username is set

func getchannelinfo(c *gosocketio.Channel, cids map[string]string, cid_name map[string]string) (int, string) {
	channels := c.List(cids[c.Id()])
	isname := "set"
	for _, value := range channels {
		if _, ok := cid_name[value.Id()]; !ok {
			// the cid exists within the map
			log.Println("values--", value.Id(), "name not set")
			isname = "notset"
		}
	}
	roomclientAmount := c.Amount(cids[c.Id()])
	return roomclientAmount, isname
}

// driver code
func main() {

	logging()

	log.Println("create server instance ")
	//create server instance
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	// users array:
	var users []string
	// roomno array:
	var roomno []string
	// var cids [100]string
	count := 0

	var cids = make(map[string]string) // map[[clientid]->roomno,[clientid]->roomno]

	var cid_name = make(map[string]string) //map[[clientid]->name,[clientid]->name]

	// groups array:
	groups := [5]string{"group1", "group2"}

	// on connect
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel, args interface{}) {

		//client id is unique
		log.Println("New client connected, client id is ", c.Id())
		if len(roomno) > 0 {
			var roomcreate = "yes"
			for _, rooms := range roomno {

				if c.Amount(rooms) < 2 {
					// delete a msg file
					if c.Amount(rooms) < 1 {
						if _, err := os.Stat(rooms); err == nil {
							err := os.Remove(rooms)
							log.Println("deleteing file -- ", rooms)
							if err != nil {
								log.Fatal(err)
							}
						}
					}
					c.Join(rooms)
					cids[c.Id()] = rooms
					log.Println(c.Id(), "--Added in the room--", rooms)
					roomcreate = "no"
					log.Println(c.Amount(rooms))

					break
				}
			}
			if roomcreate == "yes" {
				log.Println("No rooms available creating new one")
				count = count + 1
				roomno = append(roomno, strconv.Itoa(count))
				cids[c.Id()] = strconv.Itoa(count)
				log.Println(strconv.Itoa(count))
				c.Join(strconv.Itoa(count))
				log.Println(c.Id(), "--Added in the room--", count)
			}
		} else {
			log.Println("creating the First room")
			count = count + 1
			roomno = append(roomno, strconv.Itoa(count))
			cids[c.Id()] = strconv.Itoa(count)
			log.Println(strconv.Itoa(count))
			c.Join(strconv.Itoa(count))
			log.Println(c.Id(), "--Added in the room--", count)
		}

		log.Println("total rooms---- ", len(roomno))
	})

	// on disconnect
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		_, isname := getchannelinfo(c, cids, cid_name)
		if isname == "set" {
			c.BroadcastTo(cids[c.Id()], "clientLeft", "success")
		}
		log.Println(c.Id(), "Disconnected")
		log.Println("before delting the cid from the map")
		log.Println(cids)
		delete(cids, c.Id())
		log.Println("After delete the cid from the map")
		log.Println(cids)
		c.Close()

	})
	server.On("setUsername", func(c *gosocketio.Channel, name string) {
		if contains(users, name) {
			log.Println("username taken already", name)
			users = append(users, name)
			username := Username{name, "user taken already!!!!", "", "", ""}
			c.Emit("userExists", username)
		} else {
			log.Println("username set ", name)
			users = append(users, name)
			username := Username{name, "success", "", "", ""}
			cid_name[c.Id()] = name
			c.Emit("userSet", username)
			if c.Amount(cids[c.Id()]) == 2 {
				// // both the client in the room should have a username
				_, isname := getchannelinfo(c, cids, cid_name)
				if isname == "set" {
					c.BroadcastTo(cids[c.Id()], "ready", "success")
				}
			}

		}
	})
	// join the users to the specified room
	server.On("joingroup", func(c *gosocketio.Channel, group *Group) {
		log.Println("groups----------", groups)
		log.Println(group.Room)
		c.Join(group.Room)
		log.Println("amount----------")
		log.Println(c.Amount(group.Room))
		c.BroadcastTo(group.Room, group.Room+"userjoined", group)
	})
	// remove the users to the specified room
	server.On("LeaveGroup", func(c *gosocketio.Channel, group *Group) {
		c.Leave(group.Room)
		log.Println("Leave groups----------", groups)
		log.Println("amount----------")
		log.Println(c.Amount(group.Room))
		log.Println(group.Room + "userleft")
		c.BroadcastTo(group.Room, group.Room+"userleft", group)
	})
	// send msg to the specified room
	server.On("/groupchat", func(c *gosocketio.Channel, groupuser *GroupUser) {
		log.Println("groupchat---------------")
		groupuser.Name = cid_name[c.Id()]
		log.Println(groupuser.GroupId)
		c.BroadcastTo(groupuser.GroupId, groupuser.GroupId, groupuser)
	})
	server.On("/chat", func(c *gosocketio.Channel, username *Username) {

		username.Name = cid_name[c.Id()]
		// check wether the room has 2 clients to start chat
		roomclientAmount, isname := getchannelinfo(c, cids, cid_name)
		// write the message in a textfile.
		fileName := cids[c.Id()]
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		datawriter := bufio.NewWriter(file)
		sampledata := username.Name + " : " + username.Msg
		for _, data := range sampledata {
			_, _ = datawriter.WriteString(string(data))
		}
		_, _ = datawriter.WriteString("\n")
		datawriter.Flush()
		file.Close()

		// send the message
		if roomclientAmount == 2 && isname == "set" {
			log.Println("cids")
			log.Println(cids)
			log.Println(c.Id(), "--sent a message")
			username.No = strconv.Itoa(indexOf(c.Id(), cids))
			log.Println(username.No)
			c.BroadcastTo(cids[c.Id()], "newmsg", username)
		} else {
			c.Emit("noclient", username)
		}

	})
	server.On("send_email", func(c *gosocketio.Channel, tomail string) {
		sendemail(tomail, cids[c.Id()], c)
		c.Emit("sent-email", "success")
	})

	// check prev messages
	server.On("check-prev-msg", func(c *gosocketio.Channel) {
		if _, err := os.Stat(cids[c.Id()]); err == nil {
			c.BroadcastTo(cids[c.Id()], "file-exists", "success")
		}

	})

	//error catching handler
	server.On(gosocketio.OnError, func(c *gosocketio.Channel) {
		log.Println("Error occurs")
	})

	routing(server)
}
