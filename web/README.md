# Chat UI

This is a VueJS application that embedded into the Golang HTTP Server as static html pages.

## Setup

You can also run this SPA web app separately using NodeJS, by installing the dependencies first, with following comment:

```bash
yarn install
```

## Dev

Start the development server in your local machine.

```bash
yarn dev
```

## Build

You may want to build the static html that will be serve in HTTP server, you can do this:

```bash
yarn build
```

Then the build files will be live under `./dist` directory.
