# gogogadget

☂️ [**WIP**] Self-hosted interactive JSON API builder written in Go.

⏱ **Create and deploy an API 30 seconds: [gogogadget.live](https://gogogadget.live)**

## Table of Contents

- [gogogadget](#gogogadget)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [How to Use](#how-to-use)
  - [Tasks](#tasks)
    - [Documentation](#documentation)
    - [Server-Side (Go)](#server-side-go)
    - [Client-Side (HTML/CSS/Vanilla JS)](#client-side-htmlcssvanilla-js)
      - [Ops](#ops)

## Features

## How to Use

1. Paste a JSON data structure into the editor and assign the JSON an endpoint name that best represents the content.
2. Submit the form.
3. Click the link that displays upon submission.
   - This link is a permanent URL that returns the submitted JSON!
4. Profit! Call your API from anywhere!

## Tasks

### Documentation

- [ ] create architecture diagram for `README`
- [ ] add badges to `README`
- [ ] `dep status -dot` dependency diagram

### Server-Side (Go)

- [] db schema
- [ ] endpoints model
- [ ] stub routes
- [ ] cURL command generator func
- [ ] validate `JSON` server-side on `/create`
- [ ] visitor tracking on `/api/:slug`
- [ ] (user option) API key [generated](https://echo.labstack.com/middleware/key-auth) on `/create`
- [ ] (user option) API key enforced on `/api/:slug`
- [ ] write tests

### Client-Side (HTML/CSS/Vanilla JS)

- [x] find template
- [ ] add `highlight.js`, `axios.js`
- [ ] workflows:
  - [ ] login / signup
  - [ ] form (+ client-side JSON validation)
  - [ ] gallery
  - [ ] profile

#### Ops

- [ ] heroku setup with heroku-postgres
- [ ] integrate circleci
