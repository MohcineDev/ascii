module output

go 1.22.2

replace example/handleArgs => ../handleArgs

require getLines v0.0.0-00010101000000-000000000000

require example/handleArgs v0.0.0-00010101000000-000000000000 // indirect

replace handleFlag => ../handleFlag

replace getLines => ../getLines
