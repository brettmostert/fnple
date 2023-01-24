# fnple - TODO

fnple - The simple fintech platform for people

### Todo
- [ ] bob - add project to include has DB to init DB, allow override of db name and replace config with config we actually want.
- [ ] boballow get projects (list) and get dbs (list)
- [ ] use get db from bob to get list of dbs to init in postgres
- [ ] place pop in another repo at some point
- [ ] bob run project in context
- [ ] build ci/cd
- [ ] deploy - k8s
- [ ] api-gateway
- [ ] add sharding for DB per user per org
- [ ] auth - between components
- [ ] auth - register/login etc
- [ ] habit - remove
- [ ] habit - modify
- [ ] auth - between components
- [ ] add a behaviour - make it a habit, later add other types like tasks
- [ ] move context in specific projects to pkg folder
- [ ] updated readme
- [ ]update builds to get vendor folder, pre commit hook for
- [ ] go mod tidy
- [ ]go mod vendor

### In Progress
- [ ] add golangci-lint to bootstrap
- [ ] fix all lint issues
- [ ] ensure gofumpt is run (replace gofmt)


### Done ✓
- [x]should the vendor folder be committed?
- [x] habit - add (db)
- [x] bob - add project for go cli cmd - should create config and folder structure
- [x] bob - remove project cli cmd
- [x] bob - init - create config file and folder structure - if projects defined create if does not exist
- [x] bob - run tern in context for db migration etc
