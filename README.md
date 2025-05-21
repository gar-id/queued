# QUEUED
## Roadmap
- [x] Goroutine for program process
- [x] Grouping program and process
- [x] Using numProcs for multiple process
- [x] Dynamic log location using text template
- [x] Using mutex when updating caches data 
- [x] Using mutex when updating caches channel 
- [ ] Rebrand from QueueD to RunAll and RunCtl
- [ ] Handling FATAL status for process
- [ ] Change global variable to go-cache
- [ ] API for control QueueD
	- [ ] API whitelist
	- [ ] API ACL
	- [x] API for program action with REST API
	- [x] API for program tail with SSE
	- [x] API for update program from config without restart QueueD Server
	- [x] API for add and remove program(s)
- [ ] Web GUI
- [x] Execute with specific user
- [x] Graceful exit
- [x] Custom command
- [x] Custom stdout and stderr location
- [x] Slow start program
- [x] Auto start and auto restart program
- [x] Read config from env first, then config file
- [x] Load only yaml and yml files
- [ ] Startsecs for restart program
- [ ] QueueD Control CLI
	- [x] Ability to stop, start and restart program
	- [x] Ability to tail logs
	- [ ] Ability to add and remove program without restart QueueD Server
- [ ] Alert channel

## Known Issue
- [ ] Race condition
- [ ] Slow tailing logs