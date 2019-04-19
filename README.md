For now this is just notes.

I created this using operator-sdk to get the pkg and cmd stuff in place.

Using this for a starting point: https://developers.google.com/calendar/quickstart/go

# Development Plan

Design idea is to build in layers:

1. executable service (systemctl)
2. run as a container
3. run as a k8s deployment (kubectl apply)
4. run as an k8s operator (kubectl apply)
5. run via OLM (openshift CatalogSource and Subscription)

# make

## default

target: build

## format

Reformats code.  Makes changes.

## build

PHONY, reference to the binary file.

## clean

Clears all built artifacts.

## run

Assumes you have a `credentials.json` file.  Will prompt you to authenticate via browser and writes `token.json`.  These two files are in `.gitignore` so it's safe to put them in the root of this repo.

Will run the `build` target first, but won't clean.

# Packages

## cmd/service

Wrapper making it executable as a service.

## pkg/utility

Logic for everything.

## pkg/controller

The k8s controllers.

## pkg/apis

The API stuff, CRDs probably.