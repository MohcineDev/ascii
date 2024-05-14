# output

/*
	###	if < 0 || > 2 +>>>> err msg
	###	if == 1 add func if the flag is valid  //--- go run . --output=result.  txt "HELLO"
			//y : save the result to the result.txt file using standard file format
			//no : display the text in the terminal using the standard format
	###	if == 2
			//check if it's a valid flag
				/// y : check if the second arg is a valid file format name
				 	/// y : save the result to the result.txt file using specified file format
					/// n : display err file not found
				///n : display err file not found
	*/
	/*
			   - or -- you decalre a flag
			   // the flag is not counted it's handled by the flag.parse()
			   //args are returned from the flag.Args() function
		go run . //done
		one arg
		two args
	*/