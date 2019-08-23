API requester
===============
   Makes several requests to API url and show results.

Usage
----------
#### Build

    make apirequester
   
#### Start insolard

    ./scripts/insolard/launchnet.sh -g
   
#### Start apirequester

    ./bin/apirequester -k=.artifacts/launchnet/configs/

### Options

        -k path to members keys
                Path to dir with members keys.

        -u url
                API url for requests (default - http://localhost:19101/api).
