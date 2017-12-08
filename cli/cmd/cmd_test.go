package cmd

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println("=====> In TEST of Register")
	registerCmd.Flags().Set("username", "r0beRT")
	registerCmd.Flags().Set("password", "passw0rd")
	registerCmd.Flags().Set("email", "dr.paper@live.com")
	registerCmd.Flags().Set("phone", "17665310114")
	registerCmd.Run(registerCmd, nil)
}

func TestLogin(t *testing.T) {
	fmt.Println("=====> In TEST of Login")
	logoutCmd.Run(logoutCmd, nil)
	loginCmd.Flags().Set("username", "r0beRT")
	loginCmd.Flags().Set("password", "passw0rd")
	loginCmd.Run(loginCmd, nil)
}

func TestCreateNewMeeting(t *testing.T) {
	fmt.Println("=====> In TEST of Create a New Meeting")
	createmeetingCmd.Flags().Set("title", "testMeeting")
	createmeetingCmd.Flags().Set("members", "a,b,c")
	createmeetingCmd.Flags().Set("starttime", "2017/12/25/13:00")
	createmeetingCmd.Flags().Set("endtime", "2017/12/25/16:00")
	createmeetingCmd.Run(createmeetingCmd, nil)
}
