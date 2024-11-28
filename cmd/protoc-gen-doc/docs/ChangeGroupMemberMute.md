# ChangeGroupMemberMute

### Feature Introduction

- mute or cancel mute a group member

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| ChangeGroupMemberMute | [ChangeGroupMemberMuteReq](#openim.sdk.group.ChangeGroupMemberMuteReq) | [ChangeGroupMemberMuteResp](#openim.sdk.group.ChangeGroupMemberMuteResp) |

### ChangeGroupMemberMuteReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to mute |
| userID | [string](#string) |  | user id you want to mute |
| mutedSeconds | [uint32](#uint32) |  | mute time (unit: seconds), 0 cancels the mute |


### ChangeGroupMemberMuteResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


