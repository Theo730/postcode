#!/bin/sh
curl  -H 'Authorization:Basic dXNlcjpwYXNzd2Q=' http://<ip>:<port>/api/v1/getAddressByPostCodes -d '{"indexes":["606559","641008","641134"]}'
