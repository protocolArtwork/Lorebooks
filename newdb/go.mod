module newdb

go 1.24.0

require gorm.io/driver/sqlite v1.5.7

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	gorm.io/gorm v1.25.7-0.20240204074919-46816ad31dde
	lbcli v0.0.0-00010101000000-000000000000
)

replace lbcli => ../lbcli
