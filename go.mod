module github.com/alexdyukov/go-example-adserver

go 1.18

require (
	github.com/caarlos0/env/v6 v6.10.0
	github.com/mailru/easyjson v0.7.7
)

require github.com/josharian/intern v1.0.0 // indirect

replace github.com/alexdyukov/go-example-adserver/internal/creativehandler => ./internal/creativehandler
replace github.com/alexdyukov/go-example-adserver/internal/cliparams => ./internal/cliparams

