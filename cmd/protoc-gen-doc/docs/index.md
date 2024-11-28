# group.proto

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------- |
| QuitGroup | [QuitGroupReq](#openim.sdk.group.QuitGroupReq) | [QuitGroupResp](#openim.sdk.group.QuitGroupResp) | 1This is a leading comment for a message |
| DismissGroup | [DismissGroupReq](#openim.sdk.group.DismissGroupReq) | [DismissGroupResp](#openim.sdk.group.DismissGroupResp) | dismiss a group |
| ChangeGroupMute | [ChangeGroupMuteReq](#openim.sdk.group.ChangeGroupMuteReq) | [ChangeGroupMuteResp](#openim.sdk.group.ChangeGroupMuteResp) | mute or cancel mute a group |
| ChangeGroupMemberMute | [ChangeGroupMemberMuteReq](#openim.sdk.group.ChangeGroupMemberMuteReq) | [ChangeGroupMemberMuteResp](#openim.sdk.group.ChangeGroupMemberMuteResp) | mute or cancel mute a group member |
| TransferGroupOwner | [TransferGroupOwnerReq](#openim.sdk.group.TransferGroupOwnerReq) | [TransferGroupOwnerResp](#openim.sdk.group.TransferGroupOwnerResp) | transfer group owner |
| KickGroupMember | [KickGroupMemberReq](#openim.sdk.group.KickGroupMemberReq) | [KickGroupMemberResp](#openim.sdk.group.KickGroupMemberResp) | kick a group member |
| SetGroupInfo | [SetGroupInfoReq](#openim.sdk.group.SetGroupInfoReq) | [SetGroupInfoResp](#openim.sdk.group.SetGroupInfoResp) | set group information |
| SetGroupMemberInfo | [SetGroupMemberInfoReq](#openim.sdk.group.SetGroupMemberInfoReq) | [SetGroupMemberInfoResp](#openim.sdk.group.SetGroupMemberInfoResp) | set group member information |
| IsJoinGroup | [IsJoinGroupReq](#openim.sdk.group.IsJoinGroupReq) | [IsJoinGroupResp](#openim.sdk.group.IsJoinGroupResp) | is join group |
| GetUsersInGroup | [GetUsersInGroupReq](#openim.sdk.group.GetUsersInGroupReq) | [GetUsersInGroupResp](#openim.sdk.group.GetUsersInGroupResp) | get users in group |
| InviteUserToGroup | [InviteUserToGroupReq](#openim.sdk.group.InviteUserToGroupReq) | [InviteUserToGroupResp](#openim.sdk.group.InviteUserToGroupResp) | invite user to group |


<a name="ChangeGroupMemberMuteReq"></a>
### ChangeGroupMemberMuteReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to mute |
| userID | [string](#string) |  | user id you want to mute |
| mutedSeconds | [uint32](#uint32) |  | mute time (unit: seconds), 0 cancels the mute |


<a name="ChangeGroupMemberMuteResp"></a>
### ChangeGroupMemberMuteResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="ChangeGroupMuteReq"></a>
### ChangeGroupMuteReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to mute |
| mute | [bool](#bool) |  | mute or cancel mute |


<a name="ChangeGroupMuteResp"></a>
### ChangeGroupMuteResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="DismissGroupReq"></a>
### DismissGroupReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you dismissed |


<a name="DismissGroupResp"></a>
### DismissGroupResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="GetUsersInGroupReq"></a>
### GetUsersInGroupReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to get |
| userIDs | [string](#string) | repeated | user id you want to get |


<a name="GetUsersInGroupResp"></a>
### GetUsersInGroupResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userIDs | [string](#string) | repeated | user id in the group |


<a name="InviteUserToGroupReq"></a>
### InviteUserToGroupReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to invite |
| userIDs | [string](#string) | repeated | user id you want to invite |
| reason | [string](#string) |  | invite reason |


<a name="InviteUserToGroupResp"></a>
### InviteUserToGroupResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="IsJoinGroupReq"></a>
### IsJoinGroupReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to check |
| userID | [string](#string) |  | user id you want to check |


<a name="IsJoinGroupResp"></a>
### IsJoinGroupResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| joined | [bool](#bool) |  | whether to join the group |


<a name="KickGroupMemberReq"></a>
### KickGroupMemberReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to kick |
| kickedUserIDs | [string](#string) | repeated | user id you want to kick |
| reason | [string](#string) |  | kick reason |


<a name="KickGroupMemberResp"></a>
### KickGroupMemberResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="QuitGroupReq"></a>
### QuitGroupReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you quit |


<a name="QuitGroupResp"></a>
### QuitGroupResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="SetGroupInfoReq"></a>
### SetGroupInfoReq

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


<a name="SetGroupInfoResp"></a>
### SetGroupInfoResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="SetGroupMemberInfoReq"></a>
### SetGroupMemberInfoReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to set |
| userID | [string](#string) |  | user id you want to set |
| nickname | [string](#string) | optional | user nickname |
| faceURL | [string](#string) | optional | user face url |
| roleLevel | [int32](#int32) | optional | user role level |
| ex | [string](#string) | optional | extension field |


<a name="SetGroupMemberInfoResp"></a>
### SetGroupMemberInfoResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


<a name="TransferGroupOwnerReq"></a>
### TransferGroupOwnerReq

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to transfer |
| ownerUserID | [string](#string) |  | new owner user id |


<a name="TransferGroupOwnerResp"></a>
### TransferGroupOwnerResp

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


