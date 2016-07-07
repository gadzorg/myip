myip
======

Simple HTTP docker service that prints client ip

    $ docker run -d -p 8000:8000 --name myip -t gadzorg/myip
    
    $ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
    
https://hub.docker.com/r/gadzorg/myip/
