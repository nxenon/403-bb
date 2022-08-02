# 403 BB
[![Tool Category](https://badgen.net/badge/Tool/Bypasser/black)](https://github.com/nxenon/403-bb)
[![APP Version](https://badgen.net/badge/Version/Beta/red)](https://github.com/nxenon/403-bb)
[![GO Version](https://badgen.net/badge/GO/1.18/blue)](https://www.python.org/download/releases/3.0/)
[![License](https://badgen.net/badge/License/GPLv2/purple)](https://github.com/nxenon/403-byebye/blob/master/LICENSE)

Hope to Bypass 403 Forbidden Errors :)
 - Still working on it.
 - I'll be glad if you wanna contribute

 # Installation

    go install github.com/nxenon/403-bb@latest

# Usage

    ./403-bb  -help
    ./403-bb  -url http://example.com/admin
    ./403-bb  -url http://example.com/admin -payload l27.0.0.1
    ./403-bb  -url http://example.com/admin -proxy http://127.0.0.1:8080 -timeout 0

# Help    
    _  _    ___ ____    ____  ____  
    | || |  / _ \___ \  |  _ \|  _ \  Version: beta
    | || |_| | | |__) | | |_) | |_) |
    |__   _| | | |__ <  |  _ <|  _ < 
       | | | |_| |__) | | |_) | |_) |
       |_|  \___/____/  |____/|____/
    
    Usage of ./main:
    ./main -url TARGET
      -payload string
            Bypass-Value(payload) for Replacing in Headers (default "127.0.0.1")
      -proxy string
            Send Request to Proxy [When Using Proxy set Timeout to 0 with -timeout] (Example: -proxy http://127.0.0.1:8080)
      -timeout float
            Timeout for Requests (default 3)
      -url string
            Target URL