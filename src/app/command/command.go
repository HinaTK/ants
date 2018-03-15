package command

/*命令分2种， 100开始 100以下为系统
1,无ON开始的为发送
2,回答的为ON开始
*/

const COMMAND_BEGIN_INT = 100

//cmds
const (
	_ int = 300 + iota
	CLIENT_LOGON
	CLIENT_LOGOUT
	SERVER_LOGON_RESULT
	SERVER_KICK_PLAYER
)

//world
const (
	_ int = 400 + iota
	SERVER_WORLD_REMOVE_PLAYER
	SERVER_WORLD_ADD_PLAYER
	SERVER_WORLD_NOTICE_PLAYERS
	//暗渡陈仓
	SERVER_WORLD_KICK_PLAYER
	SERVER_WORLD_GET_ONLINE_NUM
	SERVER_WORLD_NOTICE_TEST
)

//chat
const (
	_ int = 500 + iota
	CLIENT_JOIN_CHANNEL
	CLIENT_QUIT_CHANNEL
	CLIENT_NOTICE_CHANNEL
	//服务器和客户端都能使用
	SERVER_BUILD_CHANNEL
	SERVER_REMOVE_CHANNEL
)

//game texas
const (
	_ int = 2000 + iota
	CLIENT_ENTER_TEXAS_ROOM
	CLIENT_LEAVE_TEXAS_ROOM
	CLIENT_TEXAS_SITDOWN
	CLIENT_TEXAS_STAND
	CLIENT_TEXAS_CHIP
	//服务器调度
	SERVER_CREATE_TABLES
	SERVER_REMOVE_TABLES
)

//RPG
const (
	_ int = 100100 + iota
	CLIENT_MOVE
	CLIENT_ACTION
	CLIENT_RESULT
)
