package data

import (
	"cmd/proto"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

func FakeClient() *proto.CustomerRequest {
	return &proto.CustomerRequest{
		Id:    int32(rand.Int()),
		Name:  faker.Name(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
		Addresses: []*proto.CustomerRequest_Address{
			{
				Street:            faker.FirstName(),
				City:              faker.Word(),
				State:             faker.Word(),
				Zip:               faker.UUIDDigit(),
				IsShippingAddress: rand.Int()%2 == 0,
			},
		},
	}
}
