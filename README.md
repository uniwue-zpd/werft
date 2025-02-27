# werft

This repository contains the containerized hosting architecture and maintenance scripts for Semantic MediaWiki instances developed or hosted at the [ZPD](https://github.com/uniwue-zpd). The goal of this project is to provide a production-ready, opinionated, but extensible environment for hosting a variety of isolated Semantic MediaWiki instances, while minimizing boilerplate and maintenance work. This is achieved through a hierarchy of Docker images, starting from a bare-bones Semantic MediaWiki environment and extending to an optional, highly customized image. Each image builds upon the previous one, streamlining updates and reducing the workload for maintaining multiple instances.

## Getting Started

These instructions will help you set up a Semantic MediaWiki instance on your local machine for development and testing. While the deployment process may vary depending on your infrastructure, the underlying steps remain the same.

### Prerequisites

* Install [Docker](https://docs.docker.com/get-started/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
* Create the scaffolding for your own Semantic MediaWiki instance by using the [werft-cli](https://github.com/uniwue-zpd/werft-cli) **or** download a [release of the custom image](https://github.com/uniwue-zpd/werft/releases)
* In case you choose the latter installation option, keep in mind that you have to manually rename some of the setup files to make things work (the CLI does this automatically for you). The required replacements are
    * `Dockerfile.template` ➡️ `Dockerfile`
    * `docker-compose.template.yaml` ➡️ `docker-compose.yaml`
    * `template.env` ➡️ `.env`

### Configuration
#### Dockerfile
Adapt the Dockerfile to your own likings by adding installation instructions for additional skins or extensions.

Alternatively – and recommended – you can also do this by adding the skins or extensions to the `composer.local.json`.
To find out how to do this and how to add custom repositories you can have a look [here](https://github.com/uniwue-zpd/werft/blob/main/images/core/core.composer.local.json).

#### PHP
To overwrite the initial PHP configuration or to add additional configurations just add them to the `php.ini` file.

#### Maintenance tasks and backups
Maintenance scripts are running as cron jobs in the main Semantic MediaWiki container.
To configure the intervals just edit the corresponding line in the `crontab` file or add additional tasks to it.
These will then be automatically be scheduled at container startup. Make sure to rebuild so that your changes take effect.

The default tasks are: 
- [`runJobs`](https://www.mediawiki.org/wiki/Manual:RunJobs.php)
- [`rebuildData`](https://www.semantic-mediawiki.org/wiki/Help:Maintenance_script_rebuildData.php)
- [`dumpBackup`](https://www.mediawiki.org/wiki/Manual:DumpBackup.php)

The provided setup also adds a database backup container by default. To configure the database backup intervals edit the corresponding environment variables.

#### Environment variables
The setup depends on certain environment variables being set for building the images and running the containers.
Adjust the `.env` to match your desired configuration (if you want to use a different filename you have to specify this in the docker-compose call with the `--env-file` argument)
##### Semantic MediaWiki
- `SMW_IMAGE_NAME`: Desired name for the custom image (**required**)
- `SMW_IMAGE_TAG`: Specifies the core image tag that should get used for building (default: `1.39.7`) 
- `SMW_PORT`: Port through which the MediaWiki should be available (default: `8080`)
##### Database
- `DB_NAME`: Name of the database that gets used by (Semantic) MediaWiki (**required**)
- `DB_USER`: Username of the database that gets used by (Semantic) MediaWiki (**required**)
- `DB_PASSWORD`: Password that's required to access the created database (**required**)
> [!TIP]
> The values that are set here are also the ones that should be used during the MediaWiki web setup and/or in your `LocalSettings.php`
##### Database Backup
- `DB_BACKUP_CRON`: cron schedule expression which controls how often a database dump is created (default: `00 23 * * *`)

#### Docker Compose
Adjust the volumes in the `docker-compose.yaml` to match your local storage locations or add further custom volumes (e. g. for custom layouts when using the `chameleon` skin) 

If you don't want to start an instance from scratch but from the backup of an existing (Semantic) MediaWiki instance make sure to uncomment the following line in the `database` service and point the volume declaration to your backup file
```
# - ./backup.sql:/docker-entrypoint-initdb.d/datadump.sql
```

### Installation
If you have configured everything accordingly run the following commands while inside the `custom` directory:

#### Build your custom image
```
docker-compose build
```

#### Start all images
```
docker-compose up -d
```

You should now be able to access the instance through the configured port on your machine.

> [!IMPORTANT]  
> If you're creating a new (Semantic) MediaWiki instance make sure that you stop the setup with `docker-compose down` after setting everything up through the web interface and uncomment the follwing line in the `smw` service and point the volume to the created `LocalSettings.php`
> 
> `# - ./LocalSettings.php:/var/www/html/LocalSettings.php:ro`
>
> Afterwards start everything with `docker-compose up -d` again 
