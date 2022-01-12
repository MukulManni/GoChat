package main

type message struct {
	User  string `json:"user"`
	Msg   string `json:"msg"`
	Color string `json:"color"`
}

var globalmsgList = []message{{"Mod", "Welcome to the global chat.", "red"}}

func getAllMsgs() []message {
	return globalmsgList
}

func addMsg(data message) {
	globalmsgList = append(globalmsgList, data)
}
