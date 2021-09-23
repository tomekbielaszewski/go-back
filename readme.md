# Goback

I wanted something to help me organize my AWS Glacier backups.

### Setup

- aws credentials
- backup directory (bucket)

### Commands

- list
- upload
- zip & upload
- download

### TODO:

- [ ] Hierarchical tree-like printing of the bucket contents
- [ ] Cache the archives list in local file
- [ ] Force the list update with a flag switch
- [ ] Update the list in cache after some time (same as force update, but without a switch
  specified)
- [ ] Upload specified file
- [ ] Upload whole directory recursively
- [ ] Notify about offline state
- [ ] Upload whole directory by zipping it up first
- [ ] When directory contains big number of files - ask for zipping it first
- [ ] Display Archive size in human friendly way
- [ ] Display cost estimations
- [ ] Display download cost estimations
- [ ] Display files list and information in form of a table
- [ ] JSON display mode for integrations
- [ ] Request glacier archive download