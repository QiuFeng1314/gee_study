module mygee

go 1.20

require gee v1.0.0

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	util v1.0.0 // indirect
)

replace (
	gee => ./gee
	util => ./util
)
