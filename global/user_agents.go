package global

var GUserAgents *UserAgents

//Adding user agent because streaming service website will block access to webpage if it was not received
//from a valid browser. We will randomly choose user agent at runtime so that bots don't detect us
type UserAgents []string
