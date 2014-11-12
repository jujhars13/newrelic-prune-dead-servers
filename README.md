#Newrelic Prune Dead Servers

Clear dead (non-reporting) servers out of your NewRelic account.

If you're like us here at Buto you'll have a lot of transient EC2 servers popping in and out of production.

This quick and dirty go script will clear them out.

to run in go (uncompiled)

`go run prune-dead-servers.go --api-key=<your NR api key>`
