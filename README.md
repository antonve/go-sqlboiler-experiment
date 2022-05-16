# go-sqlboiler-experiment

Personal opinion about `sqlboiler`

## Pros

* Everything is configurable, which allows for a lot of flexibility
* Documentation for most common use-cases seems well written

## Cons

* Everything being configurable is a double-edged sword. Not all options are well documented and it can be difficult to figure out how to configure something
* Only TOML is officially supported for configration, which I personally find quite confusing to use
* Doesn't seem to support PostGIS well and/or is severely lacking documentation on how to use spatial types
* Not a fan of generating the models based on the actual database (which can be dirty)
* Overall seems too complex compared to other options out there

## Generating models

```bash
# Boot up development environment
$ docker compose up

# Run tests so the database migrations have been applied
$ docker compose exec dev ./run-test.sh

# Generate new models based on the latest state of the database
$ docker compose exec bash -c 'go generate'
```
