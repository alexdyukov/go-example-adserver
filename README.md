# go-example-adserver
Golang's example advert server
====
[![CI](https://github.com/alexdyukov/go-example-adserver/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/alexdyukov/go-example-adserver/actions/workflows/lint.yml?query=branch%3Amaster)

Golang's example advert server which provides 3 services: adserver, creative and site.

## Creative service

Have 1 handler:
`GET /?param1=1&param2=fake` - returns json like {"price":231} where `price` identifies put call for current user request, based on URL params

## Adserver service

Have 1 handler:
`GET /` - enumerate creative endpoints to get best match based on provided price. Simple logic:
1. calls creative services with grabbed user data like IP/lang/country/OS and etc.
2. wait max response time for any creative service
3. find out best match based on higher price in creatives response
4. get redirect URL by winner ID of creative service
4. return 307 http code with detected URL

## Site service

Have 4 handlers:
`POST /registration` - registrate new user
`POST /login` - auth user
`DELETE /advert/{id}` - disable ad company
`POST /advert` - create new ad company
`GET /advert?status=active` - get ad companies, filtered (or not, its optional parameter) by `status`

## Architecture

[EN](ARCH_en.md)
[RU](ARCH_ru.md)

## Example

See docker-compose.yaml to find out how to configure test environment

## License

MIT licensed. See the included LICENSE file for details.