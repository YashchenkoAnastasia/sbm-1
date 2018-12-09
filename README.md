# Semantic Brain Map Blog

## Prerequisites

You will need the following things properly installed on your computer. 

* [Docker](https://www.docker.com/) `brew install docker`
* [docker-compose] `brew install docker-compose `
* [Git](https://git-scm.com/)
* [Node.js](https://nodejs.org/) (with npm)
* [Ember CLI](https://ember-cli.com/) `brew install ember-cli`
* [Go](https://golang.org/) `brew install golang`
* [Google Chrome](https://google.com/chrome/) -- Recommended

## Installation

* `git clone <repository-url>` this repository
* `cd sbm/frontend
* `npm install`

## Running

* `cd sbm/frontend`
* `ember build`
* `docker build -t amir_frontend .` 
* `cd ../backend`
* `docker-compose up --build`

Visit the app at http://127.0.0.1.
