/**




Based on 
Version: 0.1
Author: Jonathan Hendler
Contact: jonathan dot hendler at gmail dot com
License: MIT
Repository: https://github.com/Hendler/goling

*/

package client

import (
    "fmt"
    "net"
    "bufio"
    "log"
    "strings"
) 


func Init()(net.Conn){
    conn, err := net.Dial("tcp", "localhost:50005")
    if err != nil {
        // handle error
        log.Fatal(err)
    }
    return conn
}

func Send( conn net.Conn , message string )(string){
    //null terminated string
    //var zb = []byte{0}
    //http://stackoverflow.com/questions/12359777/how-can-i-convert-a-null-terminated-string-in-a-byte-buffer-to-a-string-in-go
    fmt.Fprintf(conn, "RESET_STATS%c", '\x00')
    status, err := bufio.NewReader(conn).ReadString('\x00')
    if err != nil {
        // handle error
        log.Fatal(err)
    }

    if strings.EqualFold(status, "FL-SERVER-READY") {
        log.Fatal("FREELING SERVER NOT READY")
    }

    fmt.Fprintf(conn, "%s%c", message,  '\x00')
    status, err = bufio.NewReader(conn).ReadString('\x00')
    if err != nil {
        // handle error
        log.Fatal(err)
    }
    return status
}

func Close( conn net.Conn ) {
    err := conn.Close( )
    if err != nil {
        // handle error
        log.Fatal(err)
    }
}

/**
    Example usage 
*/
func FreelingTest(){
    conn := Init()
    fmt.Printf("Connected\n\n")
    status := Send( conn , "Every Good Boy Deserves Fudge.")
    fmt.Printf("\nResults: %s\n", status)
    Close(conn)
}


