package Util

import (
    "time"
    "fmt"
)


func DateFormat(tm *time.Time) string {
    str := fmt.Sprintf("%d_%d_%d", tm.Year(), tm.Month(), tm.Day());
    return str;
}

func GetMilliSecond(tm *time.Time) int  {
    return tm.Nanosecond() / int(time.Millisecond) 
}
