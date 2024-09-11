# werft

This repository contains the containerized hosting architecture and all maintenance scripts for Semantic MediaWiki instances developed and / or hosted at the [ZPD](https://github.com/uniwue-zpd).
The goal of this project is providing a production ready, opinionated but extensible environment for hosting a wide array of isolated Semantic MediaWiki instances while also keeping the necessary boilerplate and maintenance work at a minimum.
This is achieved by leveraging a hierarchy of Docker images which gradually range from a bare-bones Semantic MediaWiki environment at the base level to an optional, highly customized Docker image, all the while building on top of each other and therefore reducing the workload for keeping multiple instances up-to-date.

## Getting Started

These instructions will get you your own Semantic MediaWiki instance up and running on your local machine for development and testing purposes. The deployment process can differ based on your available infrastructure but the underyling processes stay the same.

### Prerequisites

* Install Docker and Docker Compose
* Clone the whole repository 

### Configuration
#### Dockerfile
Copy the `Dockerfile.template` to a file called `Dockerfile` in the same directory.
Adapt the Dockerfile to your own likings by adding installation instructions for additional skins or extensions.

Alternatively – and recommended – you can also do this by adding the skins or extensions to the `composer.local.json`. To find out how to do this and how to add custom repositories you can have a look [here](https://github.com/uniwue-zpd/werft/blob/main/images/core/core.composer.local.json).

#### PHP
To overwrite the initial PHP configuration or to add additional configurations just add them to the `php.ini` file.

#### Maintannce tasks and backups
The setup will start an additional helper container for running (Semantic) MediaWiki maintenance tasks via cron. 
To configure the intervals just edit the corresponding line in the `crontab` file or add additional tasks to it.
These will then be automatically be scheduled at container startup. Make sure to rebuild so that your changes take effect.

The default tasks are: 
- [`runJobs`](https://www.mediawiki.org/wiki/Manual:RunJobs.php)
- [`rebuildData`](https://www.semantic-mediawiki.org/wiki/Help:Maintenance_script_rebuildData.php)
- [`dumpBackup`](https://www.mediawiki.org/wiki/Manual:DumpBackup.php)

The provided setup also adds a database backup container by default. To configure the database backup invervals edit the corresponding environment variables.

#### Environment variables
The setup depends on certain environment variables being set for building the images and running the containers.

Move the `template.env` file to a `.env` file in the same directory and set the following environment variables:
##### Semantic MediaWiki
- `SMW_IMAGE_NAME`: Desired name for the custom image (**required**)
- `SMW_IMAGE_TAG`: Desired tag for the custom image (**required**)
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
Copy the `docker-compose.template.yaml` to a file called `docker-compose.yaml` in the same directory.
Adjust the volumes if you want to.

If you don't want to start an instance from scratch but from the backup of an existing (Semantic) MediaWiki instance make sure to uncomment the following line in the `database` service and point the volume declaration to your backup file
```
# - ./backup.sql:/docker-entrypoint-initdb.d/datadump.sql
```

### Installation
If you have configured everything accordingly run the following commands while inside the `custom` directoy:

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
