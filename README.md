#Newrelic Prune Dead Servers

Clear dead (non-reporting) servers out of your NewRelic account.

If you're like us here at Buto you'll have a lot of transient EC2 servers popping in and out of production.

This quick and dirty go script will clear them out. I'm learning Go at the moment so this little task seemed perfect.

#Instructions To run in go (uncompiled)

`go run prune-dead-servers.go --api-key=<your NR api key>`

##Via Docker

As this is 2014 we use docker to get our stuff done, save you having to setup your env:

`docker run jujhars13/newrelic-prune-dead-servers --api-key=<your NR api key>`
