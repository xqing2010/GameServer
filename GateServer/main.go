package main

import (
    "net"
    "log"
    "GameServer/Common/Network"
    "GameServer/GateServer/Handler"
)

var sessionMap map[int]*Network.Session;

func main() {
    
    listner, err := net.Listen("tcp", ":9999");
    if nil != err {
        log.Fatalln(err);
    }
    count := 0;
    sessionMap = make(map[int]*Network.Session);
    
    Network.SetProtocolHandler(Handler.ProtocolHandler)
    
    for {
        conn, err := listner.Accept();
        if nil != err {
            log.Println(err);
            continue;
        }
        count++
        go func ()  {
            log.Printf("new connection %s, connections = %d", conn.RemoteAddr().String(), count)

            session, _ := Network.NewSession(conn, nil)
            session.Run();
        }()
        
        //sessionMap[session.ID] = session;
        
    }
}

