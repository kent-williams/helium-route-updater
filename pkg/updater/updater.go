package updater

import (
	"context"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
)

type SignedRequest interface {
	proto.Message
	GetSignature() []byte
}

func Sign[T SignedRequest](req T) T {
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes([]byte{})

	//TODO sign
	//marshalled, _ := proto.Marshal(req)

	signature := []byte{}
	reflect.ValueOf(req).Elem().FieldByName("Signature").SetBytes(signature)

	return req
}

func Run(server string, routeId string, ch <-chan listener.DeviceEvent) {
	ctx := context.Background()
	conn, err := grpc.Dial(
		server,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := iot_config.NewRouteClient(conn)
	u, err := client.UpdateEuis(ctx)

	//TODO: auth
	for event := range ch {
		pair := &iot_config.EuiPairV1{
			RouteId: routeId,
			AppEui:  event.AppEui,
			DevEui:  event.DevEui,
		}

		if event.EventType == listener.EventType_Updated || event.EventType == listener.EventType_Deleted {
			err := u.Send(Sign(&iot_config.RouteUpdateEuisReqV1{
				Action:    iot_config.ActionV1_remove,
				EuiPair:   pair,
				Timestamp: 0,
				Signature: nil,
			}))
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		if event.EventType == listener.EventType_Updated || event.EventType == listener.EventType_Added {
			err := u.Send(Sign(&iot_config.RouteUpdateEuisReqV1{
				Action:    iot_config.ActionV1_add,
				EuiPair:   pair,
				Timestamp: 0,
				Signature: nil,
			}))
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}