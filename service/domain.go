package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/crowdsourcing/Demo/sdkInit"
)

type ServiceSetup struct{

    Environment *map[string]*sdkInit.Environ
	ChaincodeID string
	Info *sdkInit.InitInfo
}

func registerEvent(client *channel.Client, ChaincodeID, eventID string)(fab.Registration, <-chan *fab.CCEvent){
	reg, notifier, err := client.RegisterChaincodeEvent(ChaincodeID, eventID)
	if err!=nil{
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error{
	select{
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件 : %v\n", ccEvent)
	case <- time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件（%s)", eventID)
	}
	return nil
}