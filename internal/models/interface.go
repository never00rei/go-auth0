package models

/*
Whilst this may not seem like a lot of code, I thought it prudent
to abstract away the "interface" here so that I can concentrate
on building out the models modelling the data in their own files.

This means that each model _must_ contain a function that has an
"UnmarshalToModel" function in order for the interface to be
utilised correctly in the response handler.

Insert /giphy shrug here... \0/
*/

type Model interface {
	UnmarshalToModel()
}
