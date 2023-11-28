# Github Copilot Review

This repository contains a review of Github's Copilot on 3 different development tasks :
1. Text document editing.
2. JSON configuration file updates.
3. Code completion of an http.Handler.

The folders docs, handler, resource contain a beginning source file and additional source files for each exercise, _ex1, followed by a corresponding Copilot source file, _ex1_copilot.

The text document and JSON configuration Copilot examples did not provide much value. Copilot text generation was erratic, probably due to the small amount of text for training. JSON configuration provided better value although Copilot continued to replace JSON markup that was being deleted, and in one case generated a large amount of JSON markup that seemed to be plagiarized. 

The code generation worked better but again Copilot would generate code that didn't exist, with Intellisense highlighting it as invalid. Copilot would also not recognize pointers. It would generate a function that had the correct syntax, except for the return value not being a pointer. I did not try creating a comment for a function, detailing the requirements for that function, and then having Copilot generate that function. Having full function requirements before development is not realistic in practice. 

The best value add for Copilot in Go is that a large amount of common modules/packages could be created and then utilized by Copilot for code generation. So in essence, Copilot could be trained on large enterprise Go code bases, and provide code generation for all common development tasks. 

