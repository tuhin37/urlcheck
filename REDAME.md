# URLCHECK

Urlcheck - is a small but powerful utility (written in Golang) that helps you bulk check status of URLs present in a CSV file. 


### Installation
Clone the repo
```shell
git clone https://github.com/tuhin37/urlcheck.git
```

CD in to the folder`urlcheck`
```shell
cd urlcheck
```

To make the binaries 
```shell
cd urlcheck
```

This will build binaries for different platform inside the `./bin` folder. Please note this step will require 
- go 1.22 runtime
- make and build-essentials 

To install the binary 
```shell
sudo make install
```
This will put your OS specific binary in `/usr/local/bin` 


To check if the installation is successful, execute the following command
```shell
urlcheck
```

You should get this Output
```shell
Error: Exactly one input file must be provided.

Usage: urlcheck [OPTIONS] file.csv
Options:
  --output-format csv/table Print format (default: csv)
  --workers n               Number of workers for URL checking (default: 1)
  --verbose                 Enable verbose output (default: disabled)
  --silent                  Suppress all output (default: disabled)
  --write-csv filename.csv  Write output to a CSV file (default: none)
  --read-offset n           Offset for reading records (default: 0 no-offset)
  --read-limit n            Number of records from the Offset (default: 0 no-limit)
  -h, --help                Display this help message

```


### Status

| Status Code        | Status                                        |
| ------------------ | --------------------------------------------- |
| URL_MALFORMED      | String is not a valid URL                     |
| FAIL_CONNECT       | TCP connection failed                         |
| FAIL_RESOLVE       | DNS resolution failed                         |
| FAIL_TIMEOUT       | Response did not come within provided timeout |
| TOO_MANY_REDIRECTS | Infinite redirection                          |
| SSL_ERROR          | SSL error                                     |
| FAIL_RESPONSE      | Response code not 2xx                         |


### Run 
To run the program follow these steps

CD into ./data folder
```shell
cd data
```

Run this command
```shell
urlcheck test.csv 
```

Output
```shell
<http://google.com/>, OK
<https://linkedin.com/>, OK
<https://facebook.com/>, OK
<http://example.com/>, OK
<bla-bla>, URL_MALFORMED
<http://xinhuanet.com/>, OK
<http://httpstat.us/300/>, FAIL_CONNECT
<https://www.github.com/>, OK
<http://nonexistentdomain1234567.com/>, FAIL_RESOLVE
<https://www.thiswebsiteneverexists.com/>, FAIL_RESOLVE
<http://slowwebsite.com/>, FAIL_RESOLVE
<https://delayedresponse.com/>, FAIL_CONNECT
<http://invalidurl.com/>, FAIL_RESOLVE
<https://expired-cert.com/>, FAIL_RESOLVE
<http://httpbin.org/status/404>, FAIL_RESPONSE
<https://httpbin.org/status/500>, FAIL_RESPONSE
<http://httpbin.org/status/403>, FAIL_RESPONSE
<https://httpbin.org/status/301>, TOO_MANY_REDIRECTS
<https://redirected.com/>, OK
<http://example.org/>, OK
<https://www.stackoverflow.com/>, OK
<http://httpbin.org/status/400>, FAIL_CONNECT
<https://httpbin.org/status/401>, FAIL_CONNECT
<http://unreachablewebsite.com/>, FAIL_RESOLVE
<https://unresolveddns.com/>, FAIL_RESOLVE
<http://slowserver.com/>, FAIL_RESPONSE
<https://timeouttest.com/>, FAIL_TIMEOUT
<https://highlatency.com/>, OK
<http://success-but-slow.com/>, FAIL_RESOLVE
<http://redirect-loop.com/>, FAIL_RESOLVE
<https://test-ssl-fail.com/>, FAIL_RESOLVE
<https://unknown-domain.xyz/>, FAIL_RESOLVE
<http://localhost:9999/>, FAIL_RESOLVE
<http://127.0.0.1:9999/>, FAIL_RESOLVE
<https://httpbin.org/delay/3>, FAIL_CONNECT
<http://httpbin.org/delay/2>, FAIL_CONNECT
<https://unresolvablehost.invalid/>, FAIL_RESOLVE
<http://httpbin.org/status/418>, FAIL_CONNECT
<http://httpbin.org/status/429>, FAIL_RESPONSE
<http://httpbin.org/status/502>, FAIL_RESPONSE
<http://httpbin.org/status/503>, FAIL_RESPONSE
<http://httpbin.org/status/504>, FAIL_RESPONSE
<https://www.example.com/>, FAIL_CONNECT
```




### Advanced Options 

Here are examples of some advanced options 

Example-1
```shell
urlcheck --verbose --workers 10 --output-format table test.csv 
```
Explain:
- Verbose output is on
- Uses 10 parallel workers for faster results
- The output is printed in a tabular format


Output
```shell
=================== command parsed ===================
Parsed arguments:
  Input-file:      test.csv
  Read-offset:     0
  Read-limit:      0
  Verbose:         true
  Silent:          false
  Workers:         10
  Output-format:   table
  Output-file:     


=================== input data ===================
URL: http://google.com/, Timeout: 1000 ms, Status: pending
URL: https://linkedin.com/, Timeout: 1000 ms, Status: pending
URL: https://facebook.com/, Timeout: 1500 ms, Status: pending
URL: http://example.com/, Timeout: 2000 ms, Status: pending
URL: bla-bla, Timeout: 200 ms, Status: pending
URL: http://xinhuanet.com/, Timeout: 2000 ms, Status: pending
URL: http://httpstat.us/300/, Timeout: 1000 ms, Status: pending
URL: https://www.github.com/, Timeout: 1000 ms, Status: pending
URL: http://nonexistentdomain1234567.com/, Timeout: 2000 ms, Status: pending
URL: https://www.thiswebsiteneverexists.com/, Timeout: 1500 ms, Status: pending
URL: http://slowwebsite.com/, Timeout: 500 ms, Status: pending
URL: https://delayedresponse.com/, Timeout: 2500 ms, Status: pending
URL: http://invalidurl.com/, Timeout: 1000 ms, Status: pending
URL: https://expired-cert.com/, Timeout: 1500 ms, Status: pending
URL: http://httpbin.org/status/404, Timeout: 1000 ms, Status: pending
URL: https://httpbin.org/status/500, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/403, Timeout: 1000 ms, Status: pending
URL: https://httpbin.org/status/301, Timeout: 1000 ms, Status: pending
URL: https://redirected.com/, Timeout: 2000 ms, Status: pending
URL: http://example.org/, Timeout: 1000 ms, Status: pending
URL: https://www.stackoverflow.com/, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/400, Timeout: 1000 ms, Status: pending
URL: https://httpbin.org/status/401, Timeout: 1000 ms, Status: pending
URL: http://unreachablewebsite.com/, Timeout: 3000 ms, Status: pending
URL: https://unresolveddns.com/, Timeout: 2000 ms, Status: pending
URL: http://slowserver.com/, Timeout: 5000 ms, Status: pending
URL: https://timeouttest.com/, Timeout: 800 ms, Status: pending
URL: https://highlatency.com/, Timeout: 3000 ms, Status: pending
URL: http://success-but-slow.com/, Timeout: 1500 ms, Status: pending
URL: http://redirect-loop.com/, Timeout: 1000 ms, Status: pending
URL: https://test-ssl-fail.com/, Timeout: 1500 ms, Status: pending
URL: https://unknown-domain.xyz/, Timeout: 2000 ms, Status: pending
URL: http://localhost:9999/, Timeout: 1000 ms, Status: pending
URL: http://127.0.0.1:9999/, Timeout: 1000 ms, Status: pending
URL: https://httpbin.org/delay/3, Timeout: 500 ms, Status: pending
URL: http://httpbin.org/delay/2, Timeout: 2000 ms, Status: pending
URL: https://unresolvablehost.invalid/, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/418, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/429, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/502, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/503, Timeout: 1000 ms, Status: pending
URL: http://httpbin.org/status/504, Timeout: 1000 ms, Status: pending
URL: https://www.example.com/, Timeout: 1000 ms, Status: pending


=================== checking URLs ===================
URL                                      Timeout(ms)  Status            
---------------------------------------  -----------  ------------------
http://google.com/                       1000         OK                
https://linkedin.com/                    1000         OK                
https://facebook.com/                    1500         OK                
http://example.com/                      2000         OK                
bla-bla                                  200          URL_MALFORMED     
http://xinhuanet.com/                    2000         OK                
http://httpstat.us/300/                  1000         FAIL_RESPONSE     
https://www.github.com/                  1000         FAIL_TIMEOUT      
http://nonexistentdomain1234567.com/     2000         FAIL_RESOLVE      
https://www.thiswebsiteneverexists.com/  1500         FAIL_RESOLVE      
http://slowwebsite.com/                  500          FAIL_RESOLVE      
https://delayedresponse.com/             2500         FAIL_CONNECT      
http://invalidurl.com/                   1000         FAIL_RESOLVE      
https://expired-cert.com/                1500         FAIL_RESOLVE      
http://httpbin.org/status/404            1000         FAIL_CONNECT      
https://httpbin.org/status/500           1000         FAIL_CONNECT      
http://httpbin.org/status/403            1000         FAIL_RESPONSE     
https://httpbin.org/status/301           1000         TOO_MANY_REDIRECTS
https://redirected.com/                  2000         OK                
http://example.org/                      1000         OK                
https://www.stackoverflow.com/           1000         OK                
http://httpbin.org/status/400            1000         FAIL_RESPONSE     
https://httpbin.org/status/401           1000         FAIL_CONNECT      
http://unreachablewebsite.com/           3000         FAIL_RESOLVE      
https://unresolveddns.com/               2000         FAIL_RESOLVE      
http://slowserver.com/                   5000         FAIL_RESPONSE     
https://timeouttest.com/                 800          FAIL_TIMEOUT      
https://highlatency.com/                 3000         OK                
http://success-but-slow.com/             1500         FAIL_RESOLVE      
http://redirect-loop.com/                1000         FAIL_RESOLVE      
https://test-ssl-fail.com/               1500         FAIL_RESOLVE      
https://unknown-domain.xyz/              2000         FAIL_RESOLVE      
http://localhost:9999/                   1000         FAIL_RESOLVE      
http://127.0.0.1:9999/                   1000         FAIL_RESOLVE      
https://httpbin.org/delay/3              500          FAIL_CONNECT      
http://httpbin.org/delay/2               2000         FAIL_CONNECT      
https://unresolvablehost.invalid/        1000         FAIL_RESOLVE      
http://httpbin.org/status/418            1000         FAIL_CONNECT      
http://httpbin.org/status/429            1000         FAIL_CONNECT      
http://httpbin.org/status/502            1000         FAIL_RESPONSE     
http://httpbin.org/status/503            1000         FAIL_CONNECT      
http://httpbin.org/status/504            1000         FAIL_CONNECT      
https://www.example.com/                 1000         FAIL_CONNECT      
```

Example-2
```shell
urlcheck --workers 10 --output-format table  --offset 3 --limit 6 --csv-out results.csv  test.csv	
```

This command will 
- use 10 workers to check all the urls
- output in tabular view 
- read from the 3rd item in the input list
- read a total of 6 items from the offset
- the output will be saved in a file called results.csv (this is csv and tabular)

Output
```shell
------------------------------------  -----------  -------------
http://example.com/                   2000         OK           
bla-bla                               200          URL_MALFORMED
http://xinhuanet.com/                 2000         OK           
http://httpstat.us/300/               1000         FAIL_RESPONSE
https://www.github.com/               1000         OK           
http://nonexistentdomain1234567.com/  2000         FAIL_RESOLVE 
```

Let's check the output file
```shell
cat results.csv
```

```shell title:results.csv
<http://example.com/>,OK
<bla-bla>,URL_MALFORMED
<http://xinhuanet.com/>,OK
<http://httpstat.us/300/>,FAIL_RESPONSE
<https://www.github.com/>,OK
<http://nonexistentdomain1234567.com/>,FAIL_RESOLVE
```


Example-3
```shell
urlcheck --workers 5 --silent  --offset 5 --limit 5 --csv-out out.csv  test.csv
```

- Any `STDOUT` will be subpressed
- Will use 5 workers to do the checking
- Read a total of 5 urls from the 5th line of the input `test.csv` file
- Output will be saved in `out.csv`

```shell
cat out.csv
```

```shell title:out.csv
<http://xinhuanet.com/>,OK
<http://httpstat.us/300/>,FAIL_RESPONSE
<https://www.github.com/>,OK
<http://nonexistentdomain1234567.com/>,FAIL_RESOLVE
<https://www.thiswebsiteneverexists.com/>,FAIL_RESOLVE
```

