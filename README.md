# go-example-adserver
Golang's example advert server
====
[![CI](https://github.com/alexdyukov/go-example-adserver/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/alexdyukov/go-example-adserver/actions/workflows/lint.yml?query=branch%3Amaster)

Golang's example advert server which provides 2 services: adserver and creative.

## Creative service

Provides simple auction server based on input parameters. For example:
`curl http://creative/?lang=ru` returns `{"id":0,"price":1000}` where:
- `id` counts as advert id response. Returns id is much better then returns iframe, cause it less burn out internal traffic.
- `price` put call for current user request

## Adserver service

Enumerate creative endpoints to get best match based on provided price. Simple logic:
1. calls creative services with grabbed user data like IP/lang/country/OS and etc.
2. wait max response time for any creative service
3. get advert ID with highest price
4. return 307 http code for winner ID

## Example

See docker-compose.yaml to find out how to configure test environment

## License

MIT licensed. See the included LICENSE file for details.