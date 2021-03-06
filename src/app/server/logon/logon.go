package logon

import "fmt"
import "app/command"

//非网关模块通用
import "ants/conf"
import "ants/gnet"
import "ants/actor"

//弱连接服务器，不用管心跳
func ServerLaunch(port int) {
	//数据服务器链接
	init_dber()
	//模块调度
	var refLogic actor.IActorRef = actor.NewRefRunning(new(LogicActor))
	//服务器快速启动
	gnet.ListenAndRunServer(port, func(session gnet.IBaseProxy) {
		session.SetHandle(func(b []byte) {
			refLogic.Router(gnet.NewPackBytes(b))
		})
	})
}

//逻辑块
type LogicActor struct {
	actor.BaseActor
}

func (this *LogicActor) OnReady(ref actor.IActorRef) {
	ref.SetMqNum(5000)
	ref.SetThreadNum(1000)
	ref.Open()
}

func (this *LogicActor) OnClose() {

}

func (this *LogicActor) OnMessage(args ...interface{}) {
	pack := args[0].(gnet.ISocketPacket)
	switch pack.Cmd() {
	case command.CLIENT_LOGON:
		this.on_logon(pack)
	default:
		println("login no handle:", pack.Cmd())
	}
}

//message
func (this *LogicActor) on_logon(pack gnet.ISocketPacket) {
	//header
	UserID, PassWord, SerID, SessionID := pack.ReadInt(), pack.ReadString(), pack.ReadInt(), pack.ReadUInt64()
	//other
	fmt.Println(fmt.Sprintf("Logon Info# uid=%d, session=%v gateid=%d", UserID, SessionID, SerID))
	err_code := check_user(UserID, PassWord)
	fmt.Println("Seach Result Code:", err_code, UserID, PassWord, SerID, SessionID)
	var body []byte = []byte{}
	if err_code == 0 {
		body = get_user_info(int(UserID))
	}
	//错误直接返回
	if err_code != 0 {
		actor.Main.Send(conf.TOPIC_WORLD, pack_logon_result(int16(err_code), UserID, SerID, SessionID, body))
	} else {
		actor.Main.Send(conf.TOPIC_WORLD, pack_logon_result(int16(err_code), UserID, SerID, SessionID, body))
	}
}

//send world(通知登录结果)
func pack_logon_result(code int16, uid int32, gateid int32, session uint64, body []byte) gnet.IBytes {
	return gnet.NewPackArgs(command.SERVER_WORLD_ADD_PLAYER, code, uid, gateid, session, body)
}
