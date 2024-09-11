# werft

This repository contains the containerized hosting architecture and all maintenance scripts for Semantic MediaWiki instances developed and / or hosted at the [ZPD](https://github.com/uniwue-zpd).
The goal of this project is providing a production ready, opinionated but extensible environment for hosting a wide array of isolated Semantic MediaWiki instances while also keeping the necessary boilerplate and maintenance work at a minimum.
This is achieved by leveraging a hierarchy of Docker images which gradually range from a bare-bones Semantic MediaWiki environment at the base level to an optional, highly customized Docker image, all the while building on top of each other and therefore reducing the workload for keeping multiple instances up-to-date.

![Untitled presentation](https://user-images.githubusercontent.com/33344081/173696532-2d3dc9b2-c4ff-40d5-9ef0-958e7883846e.jpg)

## Documentation
The documentation is available here
