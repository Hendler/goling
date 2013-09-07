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
    _"os"
    "log"
) 

func Hello() {
    fmt.Printf("Hello, world.\n")
}

func FreelingTest(){
    conn, err := net.Dial("tcp", "localhost:50005")
    if err != nil {
        // handle error
        fmt.Printf("ERROR CONNECTING\n\n")
        log.Fatal(err)
    }

    fmt.Printf("Connected\n\n")
    //fmt.Fprintf(conn, "RESET_STATS\n\n")

    //http://stackoverflow.com/questions/12359777/how-can-i-convert-a-null-terminated-string-in-a-byte-buffer-to-a-string-in-go

    //var zb = []byte{0}

    fmt.Fprintf(conn, "RESET_STATS%c", '\x00')
   
    //fmt.Fprintf(conn, zb)
    status, err := bufio.NewReader(conn).ReadString('\x00')
    if err != nil {
        // handle error
        fmt.Printf("ERROR READING\n\n")
    }

    fmt.Printf("Results: %s\n\n", status)
     
}


