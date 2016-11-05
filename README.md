# NLPF - Nouveaux Langages de Programmation et Frameworks - New Programmation Languages and Frameworks

## What is this project ?
This project is a small experiment made for a course which is an introduction to non-standard languages (no PHP, Java, C++). This project was made with Roman Thiaw-Kine during only two days, which explains the rush and the code quality.
We had a free choice on the technologies, so we chose some challenge with the following stack:
- Go - REST API server
- CouchDB - NoSQL database
- AngularJS 1 - FrontEnd
- Docker - Containerisation of testing environment


The goal of the project was to realize a Minimum Viable Project of a small crowdfunding platform and to present it to a client in less than a week. After that realization, the point was to realize a feedback on our stack to the rest of the class

## Our Feedback
- AngularJS: Due to the lack of time, AngularJS has been only used as a 'shortcut' for Javascript, and we only used this framework at a fraction of its capacities

- CouchDB: The usage of CouchDB in this project was a first trial for this sort of technologies, and we haven't had the opportunity to use all the functionnalities offered by CouchDB

- Docker: The docker is based on a golang image, which is itself based on debian. Since their is no updated package of couchDB for Debian, it has been necessary to find a manual installation script and configure it to work directly in our context. Globally, using Docker has been a great addition as we were working on different environnements. It also permitted us to deploy the application on a server in a few minutes.

- Go: The language used on the server side has come with a few challenges. First of all, Go comes with a native handling of JSON but it requires to create structs for each exchange, which is far from practical. Because of that, it has been necessary to manually parse and create the json used for the REST API, and for storing into CouchDB. On the other side, Go is a great language, with incredible performances. It really looks and feels like improved C which is nice. However just like C, each advanced functionnality needs a dedicated library (server: gin-gonic, database: couchdb-go).

Globally the chosen stack was risky, especially in a context of rush, but it has been really enjoyable and instructive for us to work with these technologies.
