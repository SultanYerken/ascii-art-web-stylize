# ASCII-ART-WEB

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of last project, ascii-art.

Authors: Sultanye

Usage: 
    1. Type in terminal from root folder: go run ./cmd/web
    2. Go to link http://127.0.0.1:1800
    3. On the web page, enter the text you want to see in Ascii-Art

Implementation details: 
   There are two handler functions in the code. First, main page and second page with output. The main page accepts data - text and banner. The output goes to the ascii-art page, the handler function of which processes the text with the banner through the ascii-art project and outputs the result