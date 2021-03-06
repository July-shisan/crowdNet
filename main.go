//main.go

package main
import "C"

import (
	"fmt"
	"os"
	"time"

	"github.com/hyperledger/crowdsourcing/Demo/sdkInit"
	"github.com/hyperledger/crowdsourcing/Demo/service"
	// "github.com/hyperledger/crowdsourcing/Demo/web"
	// "github.com/hyperledger/crowdsourcing/Demo/web/controller"
)

const (
	// configFile  = "config.yaml"
	configFile  = "Demo/config.yaml"

	initialized = false
	SimpleCC    = "crowdchain1"
)

var environment map[string]*sdkInit.Environ

var initInfo *sdkInit.InitInfo

var serviceSetup service.ServiceSetup

//export Start
func Start(){
	initInfo = &sdkInit.InitInfo{
		ChannelID:     "mychannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/hyperledger/crowdsourcing/Demo/fixtures/channel-artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "OrgReq",
		OrdererOrgName: "orderer.crowd.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/hyperledger/crowdsourcing/Demo/chaincode/",
		UserName:        "User1",
	}
	environment = make(map[string]*sdkInit.Environ)
	environment[initInfo.ChannelID] = &sdkInit.Environ{}

    var err error

	//实例化SDK
	environment[initInfo.ChannelID].Sdk, err = sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}


	//资源客户管理端

	err = sdkInit.CreateSourceClient(environment[initInfo.ChannelID], initInfo)

	err = sdkInit.CreateChannel(environment[initInfo.ChannelID], initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = sdkInit.JoinChannel(environment[initInfo.ChannelID], initInfo)

	err = sdkInit.InstallAndInstantiateCC(environment[initInfo.ChannelID], initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

		serviceSetup = service.ServiceSetup{

		Environment: &environment,
		ChaincodeID: SimpleCC,
		Info:        initInfo,
	}
	sdkInit.Enroll(environment[initInfo.ChannelID], "admin", "adminpw")


	// Enroll(C.CString("admin"), C.CString("adminpw"))
	// Register(C.CString("userwang"), C.CString("yyt"))
    // Enroll(C.CString("userwang"), C.CString("yyt"))

}

//export Set
func Set(){
// =================SETINFO
	

    
    now := time.Now()
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	dd2 := dd1.Add(dd)
    

    // Name id type detail reward posttime receivetime deadline requirement[]

    msg, err := serviceSetup.RegisterChain("testuser1", "aaaaa", initInfo.ChannelID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}


	msg, err = serviceSetup.Recharge("2000", initInfo.ChannelID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
    
    msg, err = serviceSetup.Getusers(initInfo.ChannelID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}



	msg, err = serviceSetup.Posttask("task2", "bbbbb", "private", "this is detail", "30", "",initInfo.ChannelID, now, dd1, dd2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	// app := controller.Application{
	// 	Fabric: &serviceSetup,
	// }
	// web.WebStart(&app)
}

//export PostTask
func PostTask(name, taskid, tasktype, detail, reward, requirement *C.char) *C.char{

    now := time.Now()
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	dd2 := dd1.Add(dd)

	msg, err := serviceSetup.Posttask(C.GoString(name), C.GoString(taskid), C.GoString(tasktype), C.GoString(detail), C.GoString(reward), C.GoString(requirement), initInfo.ChannelID, now, dd1, dd2)
	if err != nil {
        return C.CString("0" + err.Error())
	} else {
		return C.CString("1"+msg)
	}
}

//export RecieveTask
func RecieveTask(taskid *C.char) *C.char{
	now := time.Now()
	msg, err := serviceSetup.Recievetask(C.GoString(taskid), initInfo.ChannelID, now)
	if err != nil {
        return C.CString("0" + err.Error())
	} else {
		return C.CString("1"+msg)
	}
}

//export CommitTask
func CommitTask(taskid, solution *C.char) *C.char{
	now := time.Now()
	msg, err := serviceSetup.Committask(C.GoString(taskid), C.GoString(solution), initInfo.ChannelID, now)
	if err != nil {
        return C.CString("0" + err.Error())
	} else {
		return C.CString("1"+msg)
	}
}


//export GetTask
func GetTask(taskid *C.char) *C.char{
	msg, err := serviceSetup.Gettask(C.GoString(taskid),initInfo.ChannelID)
	if err != nil {
		// fmt.Println(err)
		return C.CString(err.Error())
	} else {
		// fmt.Println(msg)
		return C.CString(msg)
	}

}

//export GetAllTasks
func GetAllTasks() *C.char{
	//GETINFO
	msg, err := serviceSetup.Gettasks(initInfo.ChannelID)
	if err != nil {
		// fmt.Println(err)
		return C.CString(err.Error())
	} else {
		// fmt.Println(msg)
		return C.CString(msg)
	}

} 

//export Register
func Register(name, password *C.char) *C.char{

	err:=sdkInit.Register(environment[initInfo.ChannelID], C.GoString(name), C.GoString(password))

	if err != nil {
		return C.CString("0" + err.Error())

	} else {
		return C.CString("1")
	}
}

//export Enroll
func Enroll(name, password, info *C.char) *C.char{

	err:=sdkInit.Enroll(environment[initInfo.ChannelID], C.GoString(name), C.GoString(password))

	if err != nil {
		return C.CString("0" + err.Error())

	} 
	err = sdkInit.CreateChannelClient(environment[initInfo.ChannelID],initInfo, C.GoString(name))

	if err != nil {
		return C.CString("0" + err.Error())

	} 

	msg, err := serviceSetup.RegisterChain(C.GoString(name), C.GoString(info), initInfo.ChannelID)
	if err!= nil{
		return C.CString("0" + err.Error())
	}	
	return C.CString("1"+"success"+msg)

}

//export Recharge
func Recharge(number *C.char) *C.char{
	msg, err := serviceSetup.Recharge(C.GoString(number), initInfo.ChannelID)
	if err != nil {
		return C.CString("0" + err.Error())
	} else {
		return C.CString("1" + msg)
	}
}

//export AddSkills
func AddSkills(skills *C.char) *C.char{
	msg, err := serviceSetup.Addskills(C.GoString(skills), initInfo.ChannelID)
	if err != nil {
		return C.CString("0" + err.Error())
	} else {
		return C.CString("1" + msg)
	}
}

//export GetUser
func GetUser() *C.char{
	msg, err := serviceSetup.Getuser(initInfo.ChannelID)
	if err != nil {
		// fmt.Println(err)
		return C.CString(err.Error())
	} else {
		// fmt.Println(msg)
		return C.CString(msg)
	}
}

//export GetAllUsers
func GetAllUsers() *C.char{
	msg, err := serviceSetup.Getusers(initInfo.ChannelID)
	if err != nil {
		// fmt.Println(err)
		return C.CString(err.Error())
	} else {
		// fmt.Println(msg)
		return C.CString(msg)
	}

}


func main() {
	// Start()

    // Register(C.CString("userwang"), C.CString("yyt"))
    // Enroll(C.CString("userwang"), C.CString("yyt"))
    // sdkInit.Geidd(environment[initInfo.ChannelID])
	for key := range environment {
		defer environment[key].Sdk.Close()
	}


}
