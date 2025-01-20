package main

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicInfo() string {
	return "Authorization: Basic " + a.username + ":" + a.password
}
