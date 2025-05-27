# MVP

This application will scrape pod logs and parse important information from them
Each log line will be written to durable storage and saved externally
A UI will let the user view logs and filter by namespace, deployment, pod name, log level, etc

# High Level Diagram



# Components

## Storage Engine

We could use a database to store each log line or a no-sql document db that stores entire log files at once
We should be able to query for time frames eventually, reading all individual files is not ideal

