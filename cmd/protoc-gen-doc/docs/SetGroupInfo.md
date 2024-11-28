# SetGroupInfo

### Feature Introduction

* set group information


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------- |
| SetGroupInfo | [SetGroupInfoReq](#openim.sdk.group.SetGroupInfoReq) | [SetGroupInfoResp](#openim.sdk.group.SetGroupInfoResp) | set group information |
 

#### SetGroupInfoReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to set |
| groupName | [string](#string) | optional | group name |
| notification | [string](#string) | optional | group notification |
| introduction | [string](#string) | optional | group introduction |
| faceURL | [string](#string) | optional | group face url |
| ex | [string](#string) | optional | extension field |
| needVerification | [int32](#int32) | optional | whether to verify the group |
| lookMemberInfo | [int32](#int32) | optional | whether to view the group member information |
| applyMemberFriend | [int32](#int32) | optional | whether to allow group members to add friends |
 

#### SetGroupInfoResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


