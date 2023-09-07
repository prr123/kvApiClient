# apic

clients to test the apicservers

## apic
A most basic server to interact with an ap server.  

## apic2

A client to test the cli for interaction with the azulkv.  
Format is:  

 - http:[addr]/db/cmd  

 - http:[addr]/db/cmd?key=keyval]  

 - http:[addr]/db/cmd?key=keyval&val=valuestr  

cmd is one of the following:  

 - add 
 - upd
 - del
 - get
 - list
 - entries
 - info

More documentation will be provided soon.  

## apiclogin

A client designed to test a login procedure.  
The client transmits the information either on the command line and/or as json object in the body.  

Next version will return a jwt token to the client.  

