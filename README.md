[![Build Status](https://travis-ci.com/oleg-balunenko/instadiff-cli.svg?branch=master)](https://travis-ci.com/oleg-balunenko/instadiff-cli)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=instadiff-cli&metric=alert_status)](https://sonarcloud.io/dashboard?id=instadiff-cli)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e1b08a94c9cb45f4ac86391ef936166e)](https://www.codacy.com/manual/oleg.balunenko/instadiff-cli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=oleg-balunenko/instadiff-cli&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/oleg-balunenko/instadiff-cli)](https://goreportcard.com/report/github.com/oleg-balunenko/instadiff-cli)
[![Coverage Status](https://coveralls.io/repos/github/oleg-balunenko/instadiff-cli/badge.svg?branch=master)](https://coveralls.io/github/oleg-balunenko/instadiff-cli?branch=master)
[![Latest release artifacts](https://img.shields.io/badge/artifacts-download-blue.svg)](https://github.com/oleg-balunenko/instadiff-cli/releases/latest)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Foleg-balunenko%2Finstadiff-cli.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Foleg-balunenko%2Finstadiff-cli?ref=badge_shield)

# instadiff-cli

<p align="center">
  <img src="https://github.com/oleg-balunenko/instadiff-cli/blob/master/.assets/gopher.png" alt="" width="300">
  <br>
</p>

instadiff-cli - a command line tool for managing instagram account followers and followings

## Usage

```shell script
instadiff-cli help
```

```text

NAME:
   instadiff-cli - a command line tool for managing instagram account followers and followings

USAGE:
   instadiff-cli [global options] command [command options] [arguments...]

COMMANDS:
   followers               List your followers
   followings              List your followings
   clean-followers, clean  Un follow not mutual followings, except of whitelisted
   unmutual                List all not mutual followings
   bots                    List all bots or business accounts
   help, h                 Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config_path value  Path to the config file (default: ".config.json")
   --log_level value    Level of output logs (default: "info")
   --debug              Debug mode, where actions has no real effect
   --help, -h           show help
   --version, -v        print the version

```

To get help for any supported command:

``` shell script
instadiff-cli help [command]
```

Example of config file:

```json
{  
  "instagram":{  
  "user":{  
  "username":"user",  
  "password":"pass"  
  },  
  "whitelist":[  
  "user1",  
  "user2",  
  "user3"  
  ],  
  "limits":{  
  "unfollow":100  
  }  
 },  
  "storage": {  
  "local": true,  
  "mongo": {  
  "url": "mongoURL:test",  
  "db": "testing",  
  "collection": "users"  
  }  
 },  
  "debug": "false"  
}
```

* instagram: it is an config for instagram
	* user: user credentials to login into user's account
		* username: login
		* password: password
	* whitelist: list of followings that will be not unfollowed even if they are unmutual.
	* limits: limits per one run.
		* unfollow: number of users that could be unfollowed in one run (be careful will big numbers - account could be bunned)
* debug: if true, all opereattions will be influendce on account (e.g. unfollow will just list users and not really unfollow users)
* storage: its a config for database storage. 
	* local: if true, memory cache will be used and connection to mongo will be not set.
	* mongo: is a config for mongo database
	  - url: url of mongo DB to connect
	  - db: name of Database
	  - collection: collection name in database where models will be stored

 

Create a json file with configuration and pass the path to it via flag `--config_path`

```shell script
instadiff-cli --config_path ".config.json" [command]
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Foleg-balunenko%2Finstadiff-cli.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Foleg-balunenko%2Finstadiff-cli?ref=badge_large)