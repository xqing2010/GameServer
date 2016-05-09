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
    
    synCh := make(chan *Network.Session)
    
    for {
        conn, err := listner.Accept();
        if nil != err {
            log.Println(err);
            continue;
        }
        count++
        session, _ := Network.NewSession(conn, nil)
        go func ()  {
            log.Printf("new connection %s, connections = %d", conn.RemoteAddr().String(), count)

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

