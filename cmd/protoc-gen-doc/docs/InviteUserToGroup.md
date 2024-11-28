# InviteUserToGroup

### Feature Introduction

* invite user to group


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------- |
| InviteUserToGroup | [InviteUserToGroupReq](#openim.sdk.group.InviteUserToGroupReq) | [InviteUserToGroupResp](#openim.sdk.group.InviteUserToGroupResp) | invite user to group |
 

#### InviteUserToGroupReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to invite |
| userIDs | [string](#string) | repeated | user id you want to invite |
| reason | [string](#string) |  | invite reason |
 

#### InviteUserToGroupResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


