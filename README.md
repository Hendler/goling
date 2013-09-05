goling
======

Go client for Freeling NLP - http://nlp.lsi.upc.edu/freeling/

Socket client based on the PHP client. 

export GOPATH=$GOPATH:$HOME/goling


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