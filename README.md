# Semantic Brain Map Blog
This blog is devoted to Semantic Brain Map project. One can find a description of Semantic Brain Map and related studies. 
## Content: 
 - Members
 - Information about Semantic Brain Map itself and Semantic Atlas
 - Review on Linking Techniques for NeuroimagingData and Computational Linguistics Models
 - A dataset of audio-stimulus with linguistic features
 
 # Semantic Brain Map project
 
 ## Members
 * Amir Bakarov
 * Artem Stepanov
 * Anastasia Yashchenko
 * Anastasia Lopukhina
 * Anna Medvedeva
 * Svetlana Malugina
 * Alexey Artemov 
 ## Idea
We are trying to develop an interface, providing information about linking text of audio-stimulus (annotated with different linguistic features) with neuroimaging data. 

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
