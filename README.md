# web-infra

This repository defines the infrastructure used by my site, [web.sh](https://github.com/scott-the-programmer/web.sh)

Running this project _will_ charge you money on your provided linode / cloudflare account

## running the project

### prerequisites

* pulumi
* cloudflare account
* linode account
* golang >1.7

### running stack

You'll need to set the following env var variables to continue

* WEB_INFRA_PUB_KEY: your pub key to for ssh
* WEB_INFRA_DNS: the dns name you wish to used
* CLOUDFLARE_ZONE_ID: the zone id corresponding to the above DNS record
* CLOUDFLARE_EMAIL: email used for your cloudflare account
* CLOUDFLARE_API_KEY: api key used to interact with your cloudflare account

#### install dependencies

``` sh
make
```

#### create infra

``` sh
pulumi up
```

#### test infra

``` sh
make test
```

#### destroy infra

``` sh
pulumi destroy
```
