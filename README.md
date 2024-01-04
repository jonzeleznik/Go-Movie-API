# Project Title

## Table of Contents
+ [About](#about)
+ [Getting Started](#getting_started)
+ [Usage](#usage)
+ [Endpoints](../Endpoints)

## About <a name = "about"></a>
This is movie API but you own your data.

The focus is on the `/search/` endpoint. It checks if, the movie title you are searching for, is in your DB. If there is no movie with thaat title, then the API searches for that title in their API and then write to your own data base. 

## Getting Started <a name = "getting_started"></a>

### Prerequisites
- Go
- TMDB api key
- Cloned repo

### Dependencies
- fiber
- gin
- /joho/godotenv

### Installing

Install depoendecies

```
task install
```

Build app
```
task build
```

Run the app
```
task dev
```

## Endpoints <a name = "usage"></a>

- POST `/movies` - create a movie
- GET `/movies` - get all movies
- GET `/search/:title` - search movies
- POST `/watchlist` - add movie to watchlist
- GET `/watchlist` - get all movies from watchlist
