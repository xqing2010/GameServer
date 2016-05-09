package Util

import (
    "log"
    "os"
    "time"
    "fmt"
)

type Logger struct {
    log *log.Logger
    file *os.File
    nameprex string
    creatTime time.Time     
}

func NewLogger(name string) (*Logger, error)  {
    gameLogger := new(Logger)
    gameLogger.creatTime = time.Now()
    gameLogger.nameprex = name
    str := DateFormat(&gameLogger.creatTime)

    filename := name + str
    file, err := os.OpenFile(filename, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0660)
    gameLogger.file = file
    if nil != err {
        log.Fatalln("cann't create log file " + filename);
        return nil, fmt.Errorf("create log %s error ", filename)
    }
    gameLogger.log = log.New(file, name, log.Ldate | log.Ltime | log.Lmicroseconds)
    
    return gameLogger, nil    
}

func (logger *Logger)FatalLog(v ...interface{})  {
    logger.log.Fatalln(v...);
}

func (logger *Logger)Println(v ...interface{}) {
    logger.log.Println(v...)
}

func (logger *Logger)Printf(format string, v ...interface{})  {
    logger.log.Printf(format + "\n", v...)
}

