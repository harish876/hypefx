# HypeFx - Framework to create SPA's using Go, HTMX and Templ, the Lebron James of all stacks. 
CLI Tool providing a project scaffolding and a compact component library for apps written using Go, HTMX and Templ. 

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Usage](#usage)
- [Dependencies Used](#dependencies-used)
- [Todos and Future Plans](#todos-and-future-plans)

## Overview
This project is aimed at proividing a project scaffolding structure and a lightweight component library using Go, HTMX and Templ to create web apps.

## Features
 1. CLI Tool out of the box to create a projects scaffolding.
 2. Dev Features like Hot Module Reloading, Routing, Template Rendering , Template Generation all covered by the CLI.
 3. Comes bundled with a tailwind config file to add supporting styles to the application.

# Usage
1. If you dont have the hypefx tool installed, then run the following command `go install github.com/harish876/hypefx@latest`.
2. Run the `go mod init 'project_name'` command to initialize a new golang project.
3. Now run `hypefx generate 'project_name'` generated from the first command to define your project workspace or module.
4. Run `go mod tidy` to install all accompanying third party dependencies.
5. Now run `npm install` or `npm i` command to install dev dependencies for your tailwind and other HMR utilities.
6. Run `npm run start` command to start the dev server for your project and get started with your first Golang + HTMX + Templ project.
7. To add the component library and make it your own use the following command `hypefx add grid`.
8. Now a component can depend on mutiple other components, hence onus on you to currently add all dependent components. I will make some sort of recursive dependency resolver soon.

## Dependecies Used
1. Uses Echo for HTTP Routing, Route grouping , Template Rendering etc.
2. HTMX as a data fetching library.
3. a-h Templ for templating HTML.
4. Cobra for the CLI tool.
5. Tailwind for styling.

## Todos and Future Plans 
1. Focus on the Component Library currently. Make each component more composable.
2. Look more into ELM architecture and see if it can replace the current MVC model.
3. Add bubble tea to the terminal and make it more interactive.


