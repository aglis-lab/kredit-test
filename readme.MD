# Kredit Plus Test

## Running the project

- Maybe you need to restore the database first
- And then Simply just use docker compose already given by me. (Just that)

## Cons

- Not implemented code for repository (to much code to be implemented) (it's should be just copy paste the code from service and wrapper into repo function)
- No translation for error yet (to much code to be implmented) (it's simply create translation asset and call translation function when error occured, don't forget to include header 'Accept-Language'. All the headers already defined by the way)

## Pros

- Concept transaction is relatively complex, the transaction will be send into worker and worker will continue those transaction for other process. Like send payment into partner for realtime payment or doing other heavy thing
- Already implement for tracer and it's log
- For metric I use stdout for loging

## Tech and tools used

- opentelemetry for tracer, logs, and metric
- jaeger
- kafka
- zookeeper
- mysql
