module main

go 1.22.2


replace example/handleArgs => ../handleArgs

require example/handleArgs v0.0.0-00010101000000-000000000000

require example.moh/handleFlag v0.0.0-00010101000000-000000000000 // indirect

replace example.moh/handleFlag => ../handleFlag

replace example.moh/handleArgs => ../handleArgs
