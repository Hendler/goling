goling
======

Go client for Freeling NLP - http://nlp.lsi.upc.edu/freeling/

Socket client based on the PHP client. 

export GOPATH=$GOPATH:$HOME/goling

To test Freeling sockets from Bash, use socat, e.g.
    
    socat tcp-connect:localhost:50005 exec:'bash -i'


echo -ne "n\0m\0k"

# Appendix
 
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

 - more socat examples
   http://stuff.mit.edu/afs/sipb/machine/penguin-lust/src/socat-1.7.1.2/EXAMPLES
   http://www.geeklab.info/2011/12/playing-with-the-sockets-socat-and-netcat/