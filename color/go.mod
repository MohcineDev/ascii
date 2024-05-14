module output

go 1.22.2

replace example.moh/handleArgs => ./handleArgs

replace example.moh/handleFlag => ./handleFlag

replace example.moh/getLines => ./getLines

require (
	example.moh/getLines v0.0.0-00010101000000-000000000000
	example.moh/handleFlag v0.0.0-00010101000000-000000000000
)

require example.moh/handleArgs v0.0.0-00010101000000-000000000000 // indirect
