run:
	make -j 2 templ go

templ:
	@templ generate

go:
	 go run .

