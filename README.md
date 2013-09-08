goling
======

Go client for Freeling C++ NLP Library - http://nlp.lsi.upc.edu/freeling/

## Usage 

    go get github.com/Hendler/goling/freeling

    import (
        "github.com/Hendler/goling/freeling"
    )

## TODO

 - Response parser. Strings come back in one line per word. 
 - cgo embedded version (interact with C/C++ directly and build into programe)
   - reasons for this are that Freeling has many configurable parsing levels  
 - proper test setup
 


## Appendix
 - installing Freeling on OSX http://nlp.lsi.upc.edu/freeling/index.php?option=com_simpleboard&Itemid=65&func=view&id=2485&catid=5

    aclocal
    sudo glibtoolize --force
    autoconf
    automake -a
    env LDFLAGS="-L/opt/local/lib -L/opt/local/lib/db46" CPPFLAGS="-I/opt/local/include -I/opt/local/include/boost -I/opt/local/include/db46" ./configure --enable-boost-locale 
    make 
    sudo make install


- launching Freeling in server mode
    analyze -f en.cfg  --server --port 50005 &

Tried test Freeling sockets from Bash, use socat, e.g. - needs null terminated strings...
    
    socat tcp-connect:localhost:50005 exec:'bash -i'
    echo -ne "n\0m\0k"

 - more socat examples
   http://stuff.mit.edu/afs/sipb/machine/penguin-lust/src/socat-1.7.1.2/EXAMPLES
   http://www.geeklab.info/2011/12/playing-with-the-sockets-socat-and-netcat/

