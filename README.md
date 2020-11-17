# New Relic Prune Dead Servers

Clear dead (non-reporting) servers out of your [New Relic](http://newrelic.com) account.

If you're like us here at [Buto](http://get.buto.tv) you'll have a lot of transient EC2 instances popping in and out of production.

This quick and dirty go script will clear them out. I'm learning Go at the moment so this little task seemed perfect.

Thanks to the support team at New Relic for their inspiration for this script.

## Instructions To run in go (uncompiled)

`go run prune-dead-servers.go --api-key=<your NR api key>`

Grab your API key from Account Settings > Data Sharing

### Or via Docker

As this is 2014 we use docker to get our stuff done, save you having to setup your env:

`docker run jujhars13/newrelic-prune-dead-servers --api-key=<your NR api key>`
