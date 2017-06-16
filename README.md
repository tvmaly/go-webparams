# go-webparams 
a simple utility with functions to convert string based parameters passed to your web app to their respective types

This was created to reduce a common pattern that was observed in web apps where you needed to always convert from

string to some type.


Plain example

  intval, err := webparams.ExtractInt64("12345")

  if err != nil {
	// handle error
  }

  floatval, err := webparams.ExtractFloat64("12345.000")

  if err != nil {
	// handle error
  }


Example in a server using Gin framework with context variable named c


  intval, err := webparams.ExtractInt64( c.Params.ByName("id") )

  if err != nil {
	// handle error
  }


