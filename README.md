# go-sverigesradio

Just a simple Go API client for the Sveriges Radio open API: https://api.sr.se/api/documentation/v2/index.html

Most of the methods in the documentation are covered. Some endpoints doesn't seem to work and only returns empty objects, those are excluded.
There also seems to be some bugs in the API, for example sometimes the number of returned items varies depending on whether pagination is enabled or not.