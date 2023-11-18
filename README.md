# UN (Ugur's Notes)

UN is an online, cli oriented note/task manager. The main idea is to create a task manager that is
easy to use for a cli-nerd like me, and the data should be accessible from anywhere.

- It should be easy to use
- User should be able to keep the data online or offline.
- Be fancy but don't interrupt the usual flow of cli.

The project status: *Just Starting*™

## UN-API

This is the REST API service for the cloud backend. Using the un-api, the user can keep the data in cloud easily.
UN-API will use the GCP services for storing data. Since this is a hobby project, I am deliberately choosing the 
free-tier services from GCP.

Current plan is: Cloud Run + Firestore

## UN-CLI

UN-CLI is the cli application for the end-user. Depending on the configuration, it will use the cloud backend or embedded backend.

### Cloud Backend

Cloud backend is the backend handler which connects to the un-api for all data related operations. The module is `api-handler`.

### Embedded Backend

Embedded backend is the backend handler which uses an embedded database for all data related operations. The module is `embedded-handler`.
The chosen data storage solution is [BoltDB](https://github.com/etcd-io/bbolt) as database.

## UN-COMMON

This module keeps the common data structures for the un-cli and un-api.
