# rwth-ical-filter

A tiny HTTP proxy that fetches an iCal feed from RWTHonline and filters out unwanted events by their LV number.

## Build & Run

```bash
go build -o rwth-ical-filter ./cmd/rwth-ical-filter
./rwth-ical-filter
```

## Usage
```
GET http://localhost:8080/?pStud=<ID>&pToken=<TOKEN>&lv=<CODE>&lv=<CODE>…
```