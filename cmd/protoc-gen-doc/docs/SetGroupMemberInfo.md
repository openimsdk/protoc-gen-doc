# SetGroupMemberInfo

### Feature Introduction

- set group member information

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| SetGroupMemberInfo | [SetGroupMemberInfoReq](#openim.sdk.group.SetGroupMemberInfoReq) | [SetGroupMemberInfoResp](#openim.sdk.group.SetGroupMemberInfoResp) |

### SetGroupMemberInfoReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to set |
| userID | [string](#string) |  | user id you want to set |
| nickname | [string](#string) | optional | user nickname |
| faceURL | [string](#string) | optional | user face url |
| roleLevel | [int32](#int32) | optional | user role level |
| ex | [string](#string) | optional | extension field |


### SetGroupMemberInfoResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


