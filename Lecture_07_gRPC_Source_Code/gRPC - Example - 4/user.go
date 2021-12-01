package data

import (
	userpb "github.com/dojinkimm/go-grpc-example/protos/v1/user"
)

var UserData = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "Meow",
		PhoneNumber: "01012345678",
		Age:         29,
	},

	{
		UserId:      "2",
		Name:        "MeowMeow",
		PhoneNumber: "01023456789",
		Age:         24,
	},

	{
		UserId:      "3",
		Name:        "MeowMeowMeow",
		PhoneNumber: "01034567890",
		Age:         30,
	},

	{
		UserId:      "4",
		Name:        "MeowMeowMeowMeow",
		PhoneNumber: "01045678901",
		Age:         25,
	},
}
