dwdtweetsgo
=========
CLI tool for downloading tweets given a specific query parameter

This project is a CLI tool for downloading tweets given a specific query parameter by making use of [twitterquerygo](https://github.com/MihaiBogdanEugen/twitterquerygo).

[![Build Status](https://travis-ci.org/MihaiBogdanEugen/dwdtweetsgo.svg?branch=master)](https://travis-ci.org/MihaiBogdanEugen/dwdtweetsgo)

Releases
-----
> UNTESTED, please run directly from source

Check the [Releases](https://github.com/MihaiBogdanEugen/dwdtweetsgo/releases/tag/1.0.0) page for a binary release for your current platform.

Compiling & running from source
-----
> Make sure you have GO installed and configured

Take a look at [run.sh](run.sh) use:
```shell
$ CONSUMER_KEY=hey CONSUMER_SECRET=secret QUERY=query OUTPUT_FOLDER=/full/folder/path go run main.go
```


Configuration
-----
This application uses [namsral/flag](github.com/namsral/flag) lib to configure flags or environment variable:

| Environment Variable | Flag | Type | Default Value | Description | 
| ------------- | ------------- | ------------- | ------------- | ------------- |
| CONSUMER_KEY | -consumer-key | String | - | Twitter API Consumer Key |
| CONSUMER_SECRET | -consumer-secret | String | - | Twitter API Consumer Secret |
| OUTPUT_FOLDER | -output-folder | String | - | Output folder path |
| QUERY | -query | String | - | Search query of 500 characters maximum, including operators |
| LANGUAGE | -language | String | en | Restricts tweets to the given language, given by an ISO 639-1 code |
| RESULT_TYPE | -result-type | String | recent | Specifies what type of search results you would prefer to receive. Accepted values are `recent`, `popular` or `mixed` |
| SINCE_ID | -since-id | Uint64 | - | Returns results with an ID greater than (that is, more recent than) the specified ID |
| JSON_LOGGING | -json-logging | Bool | true | Whether to log in JSON format or not |
| LOG_LEVEL | -log-level | String | debug | The log level (panic, fatal, error, warn, info, debug) |

