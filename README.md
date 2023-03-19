# maya

Maya is web systems control bot that can manage AWS EC2 instance lifecycle through Discord chat.

Presently, it can list, get status, start, and stop EC2 instances.

## Quick Reference

- `/describe-instances` - Get a list of all instances manageable by Maya.
- `/describe-instance-status [instance-id]` Get an instance's reachability and system health status.
- `/start-instance [instance-id]` - Start a specified instance.
- `/stop-instance [instance-id]` - Stop a specified instance.
