We use kafka as a multiple-topic message queue with gRPC for communication.

This broker operates two topics:

1. persistent-log
2. postgres-submission

### persistent-log
Used for queueing user responses to the log micro-service which stores the server and user logs to backup log files.

### postgres-submission
Stores user responses in the database for chat feed later
