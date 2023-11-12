# UN (Ugur's Notes)

UN is an online, cli oriented note/task manager. The main idea is to create a task manager that is
easy to use for a cli-nerd like me, and the data should be accessible from anywhere.

- It should be easy to use
- The data should be available online but for those that only use single computer, it should support local backend
- We don't need TUI, also TUI's are incompatible with how unix cli usually works.
- Fancy things are good!
- The user should be able to take advantage of their usual workflow, like neovim, grep etc.

The project status: *Just Starting*™

## UN-API

This is the REST API service the client connects to for the remote (cloud based) backend.
UN-API will use the GCP services for storing data. Since this is a hobby project, I am deliberately choosing the 
free-tier services from GCP.

Current plan is: Cloud Run + Firestore

## UN-CLI

UN-CLI is the application for the enduser. Depending on the configuration, it will use the cloud backend or local backend.

### Cloud Backend

Cloud backend is the backend handler which connects to the un-api for all operations. The module is `api-handler`.

### Local Backend

Local backend is the backend handler which uses an embeded database for all operations. The chosen solution is 
[Storm](https://github.com/asdine/storm) which uses the [BoltDB](https://github.com/etcd-io/bbolt) as database.

## UN-COMMON

This module keeps the common data structures
