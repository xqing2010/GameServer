package main

import (
    "net"
    "GameServer/Common/Util"
    "GameServer/Common/Network"
    "GameServer/GateServer/Handler"
)

var sessionMap map[int]*Network.Session;

var serverLog *Util.Logger

func init()  {
    serverLog, _ = Util.NewLogger("ServerLog")
}

func main() {
    
    listner, err := net.Listen("tcp", ":9999");
    if nil != err {
        serverLog.FatalLog(err);
    }
    count := 0;
    sessionMap = make(map[int]*Network.Session);
    
    Network.SetProtocolHandler(Handler.ProtocolHandler)
    
    synCh := make(chan *Network.Session)
    
    for {
        conn, err := listner.Accept();
        if nil != err {
            serverLog.Println(err);
            continue;
        }
        count++
        session, _ := Network.NewSession(conn, nil)
        go func ()  {
            serverLog.Printf("new connection %s, connections = %d", conn.RemoteAddr().String(), count)

            synCh <- session
            session.Run();
        }()
        
        select{
            case session := <-synCh: {
                sessionMap[session.ID] = session
            }
            default:{
            }           
        }

    }
}

