/**

goling Go Freeling client

Assumes the Freeling server has been started already. 
analyze -f en.cfg  --server --port 50005 &


Version: 0.1
Author: Jonathan Hendler
Contact: jonathan dot hendler at gmail dot com
License: MIT
Repository: https://github.com/Hendler/goling

*/

package freeling

import (
    "fmt"
    "net"
    "bufio"
    "log"
    "strings"
) 


func Init( host string , port string )(net.Conn){
    server := [] string{ host, port}
    conn, err := net.Dial( "tcp" , strings.Join( server , ":") )
    if err != nil {
        // handle error
        log.Fatal(err)
    }
    return conn
}

/**

*/
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

/**
    Close the connection
*/
func Close( conn net.Conn ) {
    err := conn.Close( )
    if err != nil {
        // handle error
        log.Fatal(err)
    }
}

 


