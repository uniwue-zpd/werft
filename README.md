# SMWD – Semantic MediaWiki with Docker

This repository contains the Docker based hosting architecture and all maintenance scripts for Semantic MediaWiki instances developed and / or hosted at the [ZPD](https://github.com/uniwue-zpd).
The goal of this project is providing a production ready, opinionated but extensible environment for hosting a wide array of isolated Semantic MediaWiki instances while also keeping the necessary boilerplate and maintenance work at a minimum.
This is achieved by leveraging a hierarchy of Docker images which gradually range from a bare-bones Semantic MediaWiki environment at the base level to an optional, highly customized Docker image, all the while building on top of each other and therefore reducing the workload for keeping multiple instances up-to-date.

![Untitled presentation](https://user-images.githubusercontent.com/33344081/173696532-2d3dc9b2-c4ff-40d5-9ef0-958e7883846e.jpg)

## Getting started
* Pull the latest base image by running `docker pull uniwuezpd/smw:base` or building it manually by running `docker-compose build -f images/base/docker-compose.yaml`
* Pull the latest core image by running `docker pull uniwuezpd/smw:core` or building it manually by running `docker-compose build -f images/core/docker-compose.yaml`
* Configure the custom image according to your requirements and start the instance.

## Configuration
Configuration is done through environment variables. An example file with all available variables can be found [here](images/custom/template.env)-
It's advised to create a `.env` file in the directory containing the Dockerfile of your custom image. By default the `.env` file is present in the `.gitignore` of this repository. Uploading this file anywhere is highly discouraged for security reasons. 

```bash
cp docker-compose.example.yml docker-compose.override.yml
$EDITOR docker-compose.override.yml
```
## Startup

```bash
docker-compose up -d
```

## Maintenance
We recommend using the maintenance scripts provided in this repository to ensure running battle-tested and automatic maintenance routines for all currently running Semantic MediaWiki instances on your system.
Just start `maintenance/runner.py` with the correct parameters (see `python maintenance/runner.py --help` for a list of all available parameters.)

## Import data from an existing Semantic MediaWiki instance

You can import the data form an existing SMW instance by importing a SQL file containing a `mysqldump` in `./data.sql` and then run the following commands.

```bash
cat docker-composer.override.yml <<<EOD
mysql:
  volumes:
    - ./data.sql:/tmp/data.sql
EOD
docker-compose up -d
docker-compose exec mysql mysql -uroot -p$MYSQL_ROOT_PASSWORD mediawiki -e "source /tmp/data.sql"
docker-compose exec wiki php maintenance/update.php --skip-external-dependencies --quick
```

If you used `<code><pre>` instead of `<syntaxhighlight>` you can do this:

```bash
docker-compose exec wiki php extensions/ReplaceText/maintenance/replaceAll.php --nsall '<code><pre>' '<syntaxhighlight lang="bash">'
docker-compose exec wiki php extensions/ReplaceText/maintenance/replaceAll.php --nsall '</pre></code>' '</syntaxhighlight>'
```

Now is also the time to do some final cleanup and optimization after having migrated your wiki to this stack.

```bash
docker-compose exec wiki php maintenance/runJobs.php
```

If you had to wait a long time for the jobs to run you might want to run the updater again so it can optimize the db tables after the jobs have run.

```bash
docker-compose exec wiki php maintenance/update.php --skip-external-dependencies --quick
```

