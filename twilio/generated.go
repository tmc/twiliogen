package twilio

import "context"

type Services struct {
	Account struct {
		Accounts AccountsService
	}
	Application struct {
		ApiKeys               ApiKeysService
		Applications          ApplicationsService
		AuthorizedConnectApps AuthorizedConnectAppsService
		ConnectApps           ConnectAppsService
		Keys                  KeysService
		NewKeys               NewKeysService
		NewSigningKeys        NewSigningKeysService
		SigningKeys           SigningKeysService
		Tokens                TokensService
		ValidationRequests    ValidationRequestsService
	}
	Chat struct {
		ChatChannels    ChatChannelsService
		ChatCredentials ChatCredentialsService
		ChatRoles       ChatRolesService
		ChatServices    ChatServicesService
		ChatUsers       ChatUsersService
	}
	Fax struct {
		Fax      FaxService
		FaxMedia FaxMediaService
	}
	Lookup struct {
		LookupPhoneNumbers LookupPhoneNumbersService
	}
	Notify struct {
		Notify            NotifyService
		NotifyBindings    NotifyBindingsService
		NotifyCredentials NotifyCredentialsService
		NotifySegments    NotifySegmentsService
		NotifyServices    NotifyServicesService
		NotifyUsers       NotifyUsersService
	}
	PhoneNumbers struct {
		Addresses                     AddressesService
		AvailablePhoneNumberCountries AvailablePhoneNumberCountriesService
		IncomingPhoneNumbers          IncomingPhoneNumbersService
		OutgoingCallerIds             OutgoingCallerIdsService
		ShortCodes                    ShortCodesService
	}
	Sip struct {
		CredentialLists      CredentialListsService
		Domains              DomainsService
		IpAccessControlLists IpAccessControlListsService
	}
	SipTrunking struct {
		TrunkCredentialLists      TrunkCredentialListsService
		TrunkIpAccessControlLists TrunkIpAccessControlListsService
		TrunkOriginationUrls      TrunkOriginationUrlsService
		TrunkPhoneNumbers         TrunkPhoneNumbersService
		Trunks                    TrunksService
	}
	Sms struct {
		AlphaSenders        AlphaSendersService
		MessagePhoneNumbers MessagePhoneNumbersService
		MessageServices     MessageServicesService
		MessageShortCodes   MessageShortCodesService
		Messages            MessagesService
	}
	Sync struct {
		SyncDocuments SyncDocumentsService
		SyncLists     SyncListsService
		SyncMaps      SyncMapsService
		SyncServices  SyncServicesService
	}
	Taskrouter struct {
		Activities          ActivitiesService
		Events              EventsService
		TaskChannels        TaskChannelsService
		TaskQueues          TaskQueuesService
		Tasks               TasksService
		Workers             WorkersService
		Workflows           WorkflowsService
		WorkspaceStatistics WorkspaceStatisticsService
		Workspaces          WorkspacesService
	}
	Usage struct {
		UsageRecords  UsageRecordsService
		UsageTriggers UsageTriggersService
	}
	Video struct {
		VideoRecordings VideoRecordingsService
		VideoRooms      VideoRoomsService
	}
	Voice struct {
		Calls          CallsService
		Conferences    ConferencesService
		Queues         QueuesService
		Recordings     RecordingsService
		Transcriptions TranscriptionsService
	}
	Wireless struct {
		WirelessCommands WirelessCommandsService
		WirelessSims     WirelessSimsService
		WirelessUsages   WirelessUsagesService
	}
}

// AccountsService provides access to Accounts
type AccountsService interface {
	// Create - Create an Account
	Create(context.Context, AccountsCreateRequest) (AccountsCreateResponse, error)
	// Fetch - View an Account
	Fetch(context.Context, AccountsFetchRequest) (AccountsFetchResponse, error)
	// Read - View Accounts List
	Read(context.Context, AccountsReadRequest) (AccountsReadResponse, error)
	// Update - Modify an Account
	Update(context.Context, AccountsUpdateRequest) (AccountsUpdateResponse, error)
}

// ApiKeysService provides access to Api Keys
type ApiKeysService interface {
	// Delete - Delete an Api Key
	Delete(context.Context, ApiKeysDeleteRequest) (ApiKeysDeleteResponse, error)
	// Fetch - View an Api Key
	Fetch(context.Context, ApiKeysFetchRequest) (ApiKeysFetchResponse, error)
	// Read - View Api Keys List
	Read(context.Context, ApiKeysReadRequest) (ApiKeysReadResponse, error)
	// Update - Modify an Api Key
	Update(context.Context, ApiKeysUpdateRequest) (ApiKeysUpdateResponse, error)
}

// ApplicationsService provides access to Applications
type ApplicationsService interface {
	// Create - Create an Application
	Create(context.Context, ApplicationsCreateRequest) (ApplicationsCreateResponse, error)
	// Delete - Delete an Application
	Delete(context.Context, ApplicationsDeleteRequest) (ApplicationsDeleteResponse, error)
	// Fetch - View an Application
	Fetch(context.Context, ApplicationsFetchRequest) (ApplicationsFetchResponse, error)
	// Read - View Applications List
	Read(context.Context, ApplicationsReadRequest) (ApplicationsReadResponse, error)
	// Update - Modify an Application
	Update(context.Context, ApplicationsUpdateRequest) (ApplicationsUpdateResponse, error)
}

// AuthorizedConnectAppsService provides access to Authorized Connect Apps
type AuthorizedConnectAppsService interface {
	// Fetch - View an Authorized Connect App
	Fetch(context.Context, AuthorizedConnectAppsFetchRequest) (AuthorizedConnectAppsFetchResponse, error)
	// Read - View Authorized Connect Apps List
	Read(context.Context, AuthorizedConnectAppsReadRequest) (AuthorizedConnectAppsReadResponse, error)
}

// ConnectAppsService provides access to Connect Apps
type ConnectAppsService interface {
	// Fetch - View a Connect App
	Fetch(context.Context, ConnectAppsFetchRequest) (ConnectAppsFetchResponse, error)
	// Read - View Connect Apps List
	Read(context.Context, ConnectAppsReadRequest) (ConnectAppsReadResponse, error)
	// Update - Modify a Connect App
	Update(context.Context, ConnectAppsUpdateRequest) (ConnectAppsUpdateResponse, error)
}

// KeysService provides access to Keys
type KeysService interface {
	// Delete - Delete a Key
	Delete(context.Context, KeysDeleteRequest) (KeysDeleteResponse, error)
	// Fetch - View a Key
	Fetch(context.Context, KeysFetchRequest) (KeysFetchResponse, error)
	// Read - View Keys List
	Read(context.Context, KeysReadRequest) (KeysReadResponse, error)
	// Update - Modify a Key
	Update(context.Context, KeysUpdateRequest) (KeysUpdateResponse, error)
}

// NewKeysService provides access to New Keys
type NewKeysService interface {
	// Create - Create a New Key
	Create(context.Context, NewKeysCreateRequest) (NewKeysCreateResponse, error)
}

// NewSigningKeysService provides access to New Signing Keys
type NewSigningKeysService interface {
	// Create - Create a New Signing Key
	Create(context.Context, NewSigningKeysCreateRequest) (NewSigningKeysCreateResponse, error)
}

// SigningKeysService provides access to Signing Keys
type SigningKeysService interface {
	// Delete - Delete a Signing Key
	Delete(context.Context, SigningKeysDeleteRequest) (SigningKeysDeleteResponse, error)
	// Fetch - View a Signing Key
	Fetch(context.Context, SigningKeysFetchRequest) (SigningKeysFetchResponse, error)
	// Read - View Signing Keys List
	Read(context.Context, SigningKeysReadRequest) (SigningKeysReadResponse, error)
	// Update - Modify a Signing Key
	Update(context.Context, SigningKeysUpdateRequest) (SigningKeysUpdateResponse, error)
}

// TokensService provides access to Tokens
type TokensService interface {
	// Create - Create a Token
	Create(context.Context, TokensCreateRequest) (TokensCreateResponse, error)
}

// ValidationRequestsService provides access to Validation Requests
type ValidationRequestsService interface {
	// Create - Create a Validation Request
	Create(context.Context, ValidationRequestsCreateRequest) (ValidationRequestsCreateResponse, error)
}

// ChatChannelsService provides access to Chat Channels
type ChatChannelsService interface {
	// Create - Create a Chat Channel
	Create(context.Context, ChatChannelsCreateRequest) (ChatChannelsCreateResponse, error)
	// Delete - Delete a Chat Channel
	Delete(context.Context, ChatChannelsDeleteRequest) (ChatChannelsDeleteResponse, error)
	// Fetch - View a Chat Channel
	Fetch(context.Context, ChatChannelsFetchRequest) (ChatChannelsFetchResponse, error)
	// InvitesCreate - Create an Invite
	InvitesCreate(context.Context, ChatChannelsInvitesCreateRequest) (ChatChannelsInvitesCreateResponse, error)
	// InvitesDelete - Delete an Invite
	InvitesDelete(context.Context, ChatChannelsInvitesDeleteRequest) (ChatChannelsInvitesDeleteResponse, error)
	// InvitesFetch - View an Invite
	InvitesFetch(context.Context, ChatChannelsInvitesFetchRequest) (ChatChannelsInvitesFetchResponse, error)
	// InvitesRead - View Invites List
	InvitesRead(context.Context, ChatChannelsInvitesReadRequest) (ChatChannelsInvitesReadResponse, error)
	// MembersCreate - Create a Member
	MembersCreate(context.Context, ChatChannelsMembersCreateRequest) (ChatChannelsMembersCreateResponse, error)
	// MembersDelete - Delete a Member
	MembersDelete(context.Context, ChatChannelsMembersDeleteRequest) (ChatChannelsMembersDeleteResponse, error)
	// MembersFetch - View a Member
	MembersFetch(context.Context, ChatChannelsMembersFetchRequest) (ChatChannelsMembersFetchResponse, error)
	// MembersRead - View Members List
	MembersRead(context.Context, ChatChannelsMembersReadRequest) (ChatChannelsMembersReadResponse, error)
	// MembersUpdate - Modify a Member
	MembersUpdate(context.Context, ChatChannelsMembersUpdateRequest) (ChatChannelsMembersUpdateResponse, error)
	// MessagesCreate - Create a Message
	MessagesCreate(context.Context, ChatChannelsMessagesCreateRequest) (ChatChannelsMessagesCreateResponse, error)
	// MessagesDelete - Delete a Message
	MessagesDelete(context.Context, ChatChannelsMessagesDeleteRequest) (ChatChannelsMessagesDeleteResponse, error)
	// MessagesFetch - View a Message
	MessagesFetch(context.Context, ChatChannelsMessagesFetchRequest) (ChatChannelsMessagesFetchResponse, error)
	// MessagesRead - View Messages List
	MessagesRead(context.Context, ChatChannelsMessagesReadRequest) (ChatChannelsMessagesReadResponse, error)
	// MessagesUpdate - Modify a Message
	MessagesUpdate(context.Context, ChatChannelsMessagesUpdateRequest) (ChatChannelsMessagesUpdateResponse, error)
	// Read - View Chat Channels List
	Read(context.Context, ChatChannelsReadRequest) (ChatChannelsReadResponse, error)
	// Update - Modify a Chat Channel
	Update(context.Context, ChatChannelsUpdateRequest) (ChatChannelsUpdateResponse, error)
}

// ChatCredentialsService provides access to Chat Credentials
type ChatCredentialsService interface {
	// Create - Create a Chat Credential
	Create(context.Context, ChatCredentialsCreateRequest) (ChatCredentialsCreateResponse, error)
	// Delete - Delete a Chat Credential
	Delete(context.Context, ChatCredentialsDeleteRequest) (ChatCredentialsDeleteResponse, error)
	// Fetch - View a Chat Credential
	Fetch(context.Context, ChatCredentialsFetchRequest) (ChatCredentialsFetchResponse, error)
	// Read - View Chat Credentials List
	Read(context.Context, ChatCredentialsReadRequest) (ChatCredentialsReadResponse, error)
	// Update - Modify a Chat Credential
	Update(context.Context, ChatCredentialsUpdateRequest) (ChatCredentialsUpdateResponse, error)
}

// ChatRolesService provides access to Chat Roles
type ChatRolesService interface {
	// Create - Create a Chat Role
	Create(context.Context, ChatRolesCreateRequest) (ChatRolesCreateResponse, error)
	// Delete - Delete a Chat Role
	Delete(context.Context, ChatRolesDeleteRequest) (ChatRolesDeleteResponse, error)
	// Fetch - View a Chat Role
	Fetch(context.Context, ChatRolesFetchRequest) (ChatRolesFetchResponse, error)
	// Read - View Chat Roles List
	Read(context.Context, ChatRolesReadRequest) (ChatRolesReadResponse, error)
	// Update - Modify a Chat Role
	Update(context.Context, ChatRolesUpdateRequest) (ChatRolesUpdateResponse, error)
}

// ChatServicesService provides access to Chat Services
type ChatServicesService interface {
	// Create - Create a Chat Service
	Create(context.Context, ChatServicesCreateRequest) (ChatServicesCreateResponse, error)
	// Delete - Delete a Chat Service
	Delete(context.Context, ChatServicesDeleteRequest) (ChatServicesDeleteResponse, error)
	// Fetch - View a Chat Service
	Fetch(context.Context, ChatServicesFetchRequest) (ChatServicesFetchResponse, error)
	// Read - View Chat Services List
	Read(context.Context, ChatServicesReadRequest) (ChatServicesReadResponse, error)
	// Update - Modify a Chat Service
	Update(context.Context, ChatServicesUpdateRequest) (ChatServicesUpdateResponse, error)
}

// ChatUsersService provides access to Chat Users
type ChatUsersService interface {
	// Create - Create a Chat User
	Create(context.Context, ChatUsersCreateRequest) (ChatUsersCreateResponse, error)
	// Delete - Delete a Chat User
	Delete(context.Context, ChatUsersDeleteRequest) (ChatUsersDeleteResponse, error)
	// Fetch - View a Chat User
	Fetch(context.Context, ChatUsersFetchRequest) (ChatUsersFetchResponse, error)
	// Read - View Chat Users List
	Read(context.Context, ChatUsersReadRequest) (ChatUsersReadResponse, error)
	// Update - Modify a Chat User
	Update(context.Context, ChatUsersUpdateRequest) (ChatUsersUpdateResponse, error)
	// UserChannelsRead - View User Channels List
	UserChannelsRead(context.Context, ChatUsersUserChannelsReadRequest) (ChatUsersUserChannelsReadResponse, error)
}

// FaxService provides access to Programmable Fax
type FaxService interface {
	// Create - Create a Fax
	Create(context.Context, FaxCreateRequest) (FaxCreateResponse, error)
	// Delete - Delete a Fax
	Delete(context.Context, FaxDeleteRequest) (FaxDeleteResponse, error)
	// FaxMediaDelete - Delete a Fax Media
	FaxMediaDelete(context.Context, FaxFaxMediaDeleteRequest) (FaxFaxMediaDeleteResponse, error)
	// FaxMediaFetch - View a Fax Media
	FaxMediaFetch(context.Context, FaxFaxMediaFetchRequest) (FaxFaxMediaFetchResponse, error)
	// FaxMediaRead - View Fax Media List
	FaxMediaRead(context.Context, FaxFaxMediaReadRequest) (FaxFaxMediaReadResponse, error)
	// Fetch - View a Fax
	Fetch(context.Context, FaxFetchRequest) (FaxFetchResponse, error)
	// Read - View Fax List
	Read(context.Context, FaxReadRequest) (FaxReadResponse, error)
	// Update - Modify a Fax
	Update(context.Context, FaxUpdateRequest) (FaxUpdateResponse, error)
}

// FaxMediaService provides access to Fax Media
type FaxMediaService interface {
	// Delete - Delete a Fax Media
	Delete(context.Context, FaxMediaDeleteRequest) (FaxMediaDeleteResponse, error)
	// Fetch - View a Fax Media
	Fetch(context.Context, FaxMediaFetchRequest) (FaxMediaFetchResponse, error)
	// Read - View Fax Media List
	Read(context.Context, FaxMediaReadRequest) (FaxMediaReadResponse, error)
}

// LookupPhoneNumbersService provides access to Phone Number
type LookupPhoneNumbersService interface {
	// Fetch - View a Lookup Phone Number
	Fetch(context.Context, LookupPhoneNumbersFetchRequest) (LookupPhoneNumbersFetchResponse, error)
}

// NotifyService provides access to Notify
type NotifyService interface {
	// Create - Create a Notify
	Create(context.Context, NotifyCreateRequest) (NotifyCreateResponse, error)
}

// NotifyBindingsService provides access to Notify Bindings
type NotifyBindingsService interface {
	// Create - Create a Notify Binding
	Create(context.Context, NotifyBindingsCreateRequest) (NotifyBindingsCreateResponse, error)
	// Delete - Delete a Notify Binding
	Delete(context.Context, NotifyBindingsDeleteRequest) (NotifyBindingsDeleteResponse, error)
	// Fetch - View a Notify Binding
	Fetch(context.Context, NotifyBindingsFetchRequest) (NotifyBindingsFetchResponse, error)
	// Read - View Notify Bindings List
	Read(context.Context, NotifyBindingsReadRequest) (NotifyBindingsReadResponse, error)
}

// NotifyCredentialsService provides access to Notify Credentials
type NotifyCredentialsService interface {
	// Create - Create a Notify Credential
	Create(context.Context, NotifyCredentialsCreateRequest) (NotifyCredentialsCreateResponse, error)
	// Delete - Delete a Notify Credential
	Delete(context.Context, NotifyCredentialsDeleteRequest) (NotifyCredentialsDeleteResponse, error)
	// Fetch - View a Notify Credential
	Fetch(context.Context, NotifyCredentialsFetchRequest) (NotifyCredentialsFetchResponse, error)
	// Read - View Notify Credentials List
	Read(context.Context, NotifyCredentialsReadRequest) (NotifyCredentialsReadResponse, error)
	// Update - Modify a Notify Credential
	Update(context.Context, NotifyCredentialsUpdateRequest) (NotifyCredentialsUpdateResponse, error)
}

// NotifySegmentsService provides access to Notify Segments
type NotifySegmentsService interface {
	// Read - View Notify Segments List
	Read(context.Context, NotifySegmentsReadRequest) (NotifySegmentsReadResponse, error)
}

// NotifyServicesService provides access to Notify Services
type NotifyServicesService interface {
	// Create - Create a Notify Service
	Create(context.Context, NotifyServicesCreateRequest) (NotifyServicesCreateResponse, error)
	// Delete - Delete a Notify Service
	Delete(context.Context, NotifyServicesDeleteRequest) (NotifyServicesDeleteResponse, error)
	// Fetch - View a Notify Service
	Fetch(context.Context, NotifyServicesFetchRequest) (NotifyServicesFetchResponse, error)
	// Read - View Notify Services List
	Read(context.Context, NotifyServicesReadRequest) (NotifyServicesReadResponse, error)
	// Update - Modify a Notify Service
	Update(context.Context, NotifyServicesUpdateRequest) (NotifyServicesUpdateResponse, error)
}

// NotifyUsersService provides access to Notify Users
type NotifyUsersService interface {
	// Create - Create a Notify User
	Create(context.Context, NotifyUsersCreateRequest) (NotifyUsersCreateResponse, error)
	// Delete - Delete a Notify User
	Delete(context.Context, NotifyUsersDeleteRequest) (NotifyUsersDeleteResponse, error)
	// Fetch - View a Notify User
	Fetch(context.Context, NotifyUsersFetchRequest) (NotifyUsersFetchResponse, error)
	// Read - View Notify Users List
	Read(context.Context, NotifyUsersReadRequest) (NotifyUsersReadResponse, error)
	// SegmentMembershipsCreate - Create a Segment Membership
	SegmentMembershipsCreate(context.Context, NotifyUsersSegmentMembershipsCreateRequest) (NotifyUsersSegmentMembershipsCreateResponse, error)
	// SegmentMembershipsDelete - Delete a Segment Membership
	SegmentMembershipsDelete(context.Context, NotifyUsersSegmentMembershipsDeleteRequest) (NotifyUsersSegmentMembershipsDeleteResponse, error)
	// SegmentMembershipsFetch - View a Segment Membership
	SegmentMembershipsFetch(context.Context, NotifyUsersSegmentMembershipsFetchRequest) (NotifyUsersSegmentMembershipsFetchResponse, error)
	// UserBindingsCreate - Create an User Binding
	UserBindingsCreate(context.Context, NotifyUsersUserBindingsCreateRequest) (NotifyUsersUserBindingsCreateResponse, error)
	// UserBindingsDelete - Delete an User Binding
	UserBindingsDelete(context.Context, NotifyUsersUserBindingsDeleteRequest) (NotifyUsersUserBindingsDeleteResponse, error)
	// UserBindingsFetch - View an User Binding
	UserBindingsFetch(context.Context, NotifyUsersUserBindingsFetchRequest) (NotifyUsersUserBindingsFetchResponse, error)
	// UserBindingsRead - View User Bindings List
	UserBindingsRead(context.Context, NotifyUsersUserBindingsReadRequest) (NotifyUsersUserBindingsReadResponse, error)
}

// AddressesService provides access to Addresses
type AddressesService interface {
	// Create - Create an Addres
	Create(context.Context, AddressesCreateRequest) (AddressesCreateResponse, error)
	// Delete - Delete an Addres
	Delete(context.Context, AddressesDeleteRequest) (AddressesDeleteResponse, error)
	// Fetch - View an Addres
	Fetch(context.Context, AddressesFetchRequest) (AddressesFetchResponse, error)
	// PhoneNumbersRead - View Phone Numbers List
	PhoneNumbersRead(context.Context, AddressesPhoneNumbersReadRequest) (AddressesPhoneNumbersReadResponse, error)
	// Read - View Address List
	Read(context.Context, AddressesReadRequest) (AddressesReadResponse, error)
	// Update - Modify an Addres
	Update(context.Context, AddressesUpdateRequest) (AddressesUpdateResponse, error)
}

// AvailablePhoneNumberCountriesService provides access to Available Phone Number Countries
type AvailablePhoneNumberCountriesService interface {
	// AvailablePhoneNumberLocalsRead - View Available Phone Number Locals List
	AvailablePhoneNumberLocalsRead(context.Context, AvailablePhoneNumberCountriesAvailablePhoneNumberLocalsReadRequest) (AvailablePhoneNumberCountriesAvailablePhoneNumberLocalsReadResponse, error)
	// AvailablePhoneNumberMobilesRead - View Available Phone Number Mobiles List
	AvailablePhoneNumberMobilesRead(context.Context, AvailablePhoneNumberCountriesAvailablePhoneNumberMobilesReadRequest) (AvailablePhoneNumberCountriesAvailablePhoneNumberMobilesReadResponse, error)
	// AvailablePhoneNumberTollFreesRead - View Available Phone Number Toll Frees List
	AvailablePhoneNumberTollFreesRead(context.Context, AvailablePhoneNumberCountriesAvailablePhoneNumberTollFreesReadRequest) (AvailablePhoneNumberCountriesAvailablePhoneNumberTollFreesReadResponse, error)
	// Fetch - View an Available Phone Number Country
	Fetch(context.Context, AvailablePhoneNumberCountriesFetchRequest) (AvailablePhoneNumberCountriesFetchResponse, error)
	// Read - View Available Phone Number Countries List
	Read(context.Context, AvailablePhoneNumberCountriesReadRequest) (AvailablePhoneNumberCountriesReadResponse, error)
}

// IncomingPhoneNumbersService provides access to Incoming Phone Numbers
type IncomingPhoneNumbersService interface {
	// Create - Create an Incoming Phone Number
	Create(context.Context, IncomingPhoneNumbersCreateRequest) (IncomingPhoneNumbersCreateResponse, error)
	// Delete - Delete an Incoming Phone Number
	Delete(context.Context, IncomingPhoneNumbersDeleteRequest) (IncomingPhoneNumbersDeleteResponse, error)
	// Fetch - View an Incoming Phone Number
	Fetch(context.Context, IncomingPhoneNumbersFetchRequest) (IncomingPhoneNumbersFetchResponse, error)
	// IncomingPhoneNumberLocalsCreate - Create an Incoming Phone Number Local
	IncomingPhoneNumberLocalsCreate(context.Context, IncomingPhoneNumbersIncomingPhoneNumberLocalsCreateRequest) (IncomingPhoneNumbersIncomingPhoneNumberLocalsCreateResponse, error)
	// IncomingPhoneNumberLocalsRead - View Incoming Phone Number Locals List
	IncomingPhoneNumberLocalsRead(context.Context, IncomingPhoneNumbersIncomingPhoneNumberLocalsReadRequest) (IncomingPhoneNumbersIncomingPhoneNumberLocalsReadResponse, error)
	// IncomingPhoneNumberMobilesCreate - Create an Incoming Phone Number Mobile
	IncomingPhoneNumberMobilesCreate(context.Context, IncomingPhoneNumbersIncomingPhoneNumberMobilesCreateRequest) (IncomingPhoneNumbersIncomingPhoneNumberMobilesCreateResponse, error)
	// IncomingPhoneNumberMobilesRead - View Incoming Phone Number Mobiles List
	IncomingPhoneNumberMobilesRead(context.Context, IncomingPhoneNumbersIncomingPhoneNumberMobilesReadRequest) (IncomingPhoneNumbersIncomingPhoneNumberMobilesReadResponse, error)
	// IncomingPhoneNumberTollFreesCreate - Create an Incoming Phone Number Toll Free
	IncomingPhoneNumberTollFreesCreate(context.Context, IncomingPhoneNumbersIncomingPhoneNumberTollFreesCreateRequest) (IncomingPhoneNumbersIncomingPhoneNumberTollFreesCreateResponse, error)
	// IncomingPhoneNumberTollFreesRead - View Incoming Phone Number Toll Frees List
	IncomingPhoneNumberTollFreesRead(context.Context, IncomingPhoneNumbersIncomingPhoneNumberTollFreesReadRequest) (IncomingPhoneNumbersIncomingPhoneNumberTollFreesReadResponse, error)
	// Read - View Incoming Phone Numbers List
	Read(context.Context, IncomingPhoneNumbersReadRequest) (IncomingPhoneNumbersReadResponse, error)
	// Update - Modify an Incoming Phone Number
	Update(context.Context, IncomingPhoneNumbersUpdateRequest) (IncomingPhoneNumbersUpdateResponse, error)
}

// OutgoingCallerIdsService provides access to Caller Ids
type OutgoingCallerIdsService interface {
	// Delete - Delete an Outgoing Caller Id
	Delete(context.Context, OutgoingCallerIdsDeleteRequest) (OutgoingCallerIdsDeleteResponse, error)
	// Fetch - View an Outgoing Caller Id
	Fetch(context.Context, OutgoingCallerIdsFetchRequest) (OutgoingCallerIdsFetchResponse, error)
	// Read - View Outgoing Caller Ids List
	Read(context.Context, OutgoingCallerIdsReadRequest) (OutgoingCallerIdsReadResponse, error)
	// Update - Modify an Outgoing Caller Id
	Update(context.Context, OutgoingCallerIdsUpdateRequest) (OutgoingCallerIdsUpdateResponse, error)
}

// ShortCodesService provides access to Short Codes
type ShortCodesService interface {
	// Fetch - View a Short Code
	Fetch(context.Context, ShortCodesFetchRequest) (ShortCodesFetchResponse, error)
	// Read - View Short Codes List
	Read(context.Context, ShortCodesReadRequest) (ShortCodesReadResponse, error)
	// Update - Modify a Short Code
	Update(context.Context, ShortCodesUpdateRequest) (ShortCodesUpdateResponse, error)
}

// CredentialListsService provides access to SIP Credential List
type CredentialListsService interface {
	// Create - Create a Credential List
	Create(context.Context, CredentialListsCreateRequest) (CredentialListsCreateResponse, error)
	// Delete - Delete a Credential List
	Delete(context.Context, CredentialListsDeleteRequest) (CredentialListsDeleteResponse, error)
	// Fetch - View a Credential List
	Fetch(context.Context, CredentialListsFetchRequest) (CredentialListsFetchResponse, error)
	// Read - View Credential Lists List
	Read(context.Context, CredentialListsReadRequest) (CredentialListsReadResponse, error)
	// SipCredentialsCreate - Create a Sip Credential
	SipCredentialsCreate(context.Context, CredentialListsSipCredentialsCreateRequest) (CredentialListsSipCredentialsCreateResponse, error)
	// SipCredentialsDelete - Delete a Sip Credential
	SipCredentialsDelete(context.Context, CredentialListsSipCredentialsDeleteRequest) (CredentialListsSipCredentialsDeleteResponse, error)
	// SipCredentialsFetch - View a Sip Credential
	SipCredentialsFetch(context.Context, CredentialListsSipCredentialsFetchRequest) (CredentialListsSipCredentialsFetchResponse, error)
	// SipCredentialsRead - View Sip Credentials List
	SipCredentialsRead(context.Context, CredentialListsSipCredentialsReadRequest) (CredentialListsSipCredentialsReadResponse, error)
	// SipCredentialsUpdate - Modify a Sip Credential
	SipCredentialsUpdate(context.Context, CredentialListsSipCredentialsUpdateRequest) (CredentialListsSipCredentialsUpdateResponse, error)
	// Update - Modify a Credential List
	Update(context.Context, CredentialListsUpdateRequest) (CredentialListsUpdateResponse, error)
}

// DomainsService provides access to SIP Domains
type DomainsService interface {
	// Create - Create a Domain
	Create(context.Context, DomainsCreateRequest) (DomainsCreateResponse, error)
	// Delete - Delete a Domain
	Delete(context.Context, DomainsDeleteRequest) (DomainsDeleteResponse, error)
	// Fetch - View a Domain
	Fetch(context.Context, DomainsFetchRequest) (DomainsFetchResponse, error)
	// Read - View Domains List
	Read(context.Context, DomainsReadRequest) (DomainsReadResponse, error)
	// SipCredentialListMappingsCreate - Create a Sip Credential List Mapping
	SipCredentialListMappingsCreate(context.Context, DomainsSipCredentialListMappingsCreateRequest) (DomainsSipCredentialListMappingsCreateResponse, error)
	// SipCredentialListMappingsDelete - Delete a Sip Credential List Mapping
	SipCredentialListMappingsDelete(context.Context, DomainsSipCredentialListMappingsDeleteRequest) (DomainsSipCredentialListMappingsDeleteResponse, error)
	// SipCredentialListMappingsFetch - View a Sip Credential List Mapping
	SipCredentialListMappingsFetch(context.Context, DomainsSipCredentialListMappingsFetchRequest) (DomainsSipCredentialListMappingsFetchResponse, error)
	// SipCredentialListMappingsRead - View Sip Credential List Mappings List
	SipCredentialListMappingsRead(context.Context, DomainsSipCredentialListMappingsReadRequest) (DomainsSipCredentialListMappingsReadResponse, error)
	// SipIpAccessControlListMappingsCreate - Create a Sip Ip Access Control List Mapping
	SipIpAccessControlListMappingsCreate(context.Context, DomainsSipIpAccessControlListMappingsCreateRequest) (DomainsSipIpAccessControlListMappingsCreateResponse, error)
	// SipIpAccessControlListMappingsDelete - Delete a Sip Ip Access Control List Mapping
	SipIpAccessControlListMappingsDelete(context.Context, DomainsSipIpAccessControlListMappingsDeleteRequest) (DomainsSipIpAccessControlListMappingsDeleteResponse, error)
	// SipIpAccessControlListMappingsFetch - View a Sip Ip Access Control List Mapping
	SipIpAccessControlListMappingsFetch(context.Context, DomainsSipIpAccessControlListMappingsFetchRequest) (DomainsSipIpAccessControlListMappingsFetchResponse, error)
	// SipIpAccessControlListMappingsRead - View Sip Ip Access Control List Mappings List
	SipIpAccessControlListMappingsRead(context.Context, DomainsSipIpAccessControlListMappingsReadRequest) (DomainsSipIpAccessControlListMappingsReadResponse, error)
	// Update - Modify a Domain
	Update(context.Context, DomainsUpdateRequest) (DomainsUpdateResponse, error)
}

// IpAccessControlListsService provides access to IP Access Control Lists
type IpAccessControlListsService interface {
	// Create - Create an Ip Access Control List
	Create(context.Context, IpAccessControlListsCreateRequest) (IpAccessControlListsCreateResponse, error)
	// Delete - Delete an Ip Access Control List
	Delete(context.Context, IpAccessControlListsDeleteRequest) (IpAccessControlListsDeleteResponse, error)
	// Fetch - View an Ip Access Control List
	Fetch(context.Context, IpAccessControlListsFetchRequest) (IpAccessControlListsFetchResponse, error)
	// Read - View Ip Access Control Lists List
	Read(context.Context, IpAccessControlListsReadRequest) (IpAccessControlListsReadResponse, error)
	// SipIpAddressCreate - Create a Sip Ip Addre
	SipIpAddressCreate(context.Context, IpAccessControlListsSipIpAddressCreateRequest) (IpAccessControlListsSipIpAddressCreateResponse, error)
	// SipIpAddressDelete - Delete a Sip Ip Addre
	SipIpAddressDelete(context.Context, IpAccessControlListsSipIpAddressDeleteRequest) (IpAccessControlListsSipIpAddressDeleteResponse, error)
	// SipIpAddressFetch - View a Sip Ip Addre
	SipIpAddressFetch(context.Context, IpAccessControlListsSipIpAddressFetchRequest) (IpAccessControlListsSipIpAddressFetchResponse, error)
	// SipIpAddressRead - View Sip Ip Addres List
	SipIpAddressRead(context.Context, IpAccessControlListsSipIpAddressReadRequest) (IpAccessControlListsSipIpAddressReadResponse, error)
	// SipIpAddressUpdate - Modify a Sip Ip Addre
	SipIpAddressUpdate(context.Context, IpAccessControlListsSipIpAddressUpdateRequest) (IpAccessControlListsSipIpAddressUpdateResponse, error)
	// Update - Modify an Ip Access Control List
	Update(context.Context, IpAccessControlListsUpdateRequest) (IpAccessControlListsUpdateResponse, error)
}

// TrunkCredentialListsService provides access to Trunk Credential Lists
type TrunkCredentialListsService interface {
	// Create - Create a Trunk Credential List
	Create(context.Context, TrunkCredentialListsCreateRequest) (TrunkCredentialListsCreateResponse, error)
	// Delete - Delete a Trunk Credential List
	Delete(context.Context, TrunkCredentialListsDeleteRequest) (TrunkCredentialListsDeleteResponse, error)
	// Fetch - View a Trunk Credential List
	Fetch(context.Context, TrunkCredentialListsFetchRequest) (TrunkCredentialListsFetchResponse, error)
	// Read - View Trunk Credential Lists List
	Read(context.Context, TrunkCredentialListsReadRequest) (TrunkCredentialListsReadResponse, error)
}

// TrunkIpAccessControlListsService provides access to Trunk IP Access Control Lists
type TrunkIpAccessControlListsService interface {
	// Create - Create a Trunk Ip Access Control List
	Create(context.Context, TrunkIpAccessControlListsCreateRequest) (TrunkIpAccessControlListsCreateResponse, error)
	// Delete - Delete a Trunk Ip Access Control List
	Delete(context.Context, TrunkIpAccessControlListsDeleteRequest) (TrunkIpAccessControlListsDeleteResponse, error)
	// Fetch - View a Trunk Ip Access Control List
	Fetch(context.Context, TrunkIpAccessControlListsFetchRequest) (TrunkIpAccessControlListsFetchResponse, error)
	// Read - View Trunk Ip Access Control Lists List
	Read(context.Context, TrunkIpAccessControlListsReadRequest) (TrunkIpAccessControlListsReadResponse, error)
}

// TrunkOriginationUrlsService provides access to Trunk Origination Urls
type TrunkOriginationUrlsService interface {
	// Create - Create a Trunk Origination Url
	Create(context.Context, TrunkOriginationUrlsCreateRequest) (TrunkOriginationUrlsCreateResponse, error)
	// Delete - Delete a Trunk Origination Url
	Delete(context.Context, TrunkOriginationUrlsDeleteRequest) (TrunkOriginationUrlsDeleteResponse, error)
	// Fetch - View a Trunk Origination Url
	Fetch(context.Context, TrunkOriginationUrlsFetchRequest) (TrunkOriginationUrlsFetchResponse, error)
	// Read - View Trunk Origination Urls List
	Read(context.Context, TrunkOriginationUrlsReadRequest) (TrunkOriginationUrlsReadResponse, error)
	// Update - Modify a Trunk Origination Url
	Update(context.Context, TrunkOriginationUrlsUpdateRequest) (TrunkOriginationUrlsUpdateResponse, error)
}

// TrunkPhoneNumbersService provides access to Trunk Phone Numbers
type TrunkPhoneNumbersService interface {
	// Create - Create a Trunk Phone Number
	Create(context.Context, TrunkPhoneNumbersCreateRequest) (TrunkPhoneNumbersCreateResponse, error)
	// Delete - Delete a Trunk Phone Number
	Delete(context.Context, TrunkPhoneNumbersDeleteRequest) (TrunkPhoneNumbersDeleteResponse, error)
	// Fetch - View a Trunk Phone Number
	Fetch(context.Context, TrunkPhoneNumbersFetchRequest) (TrunkPhoneNumbersFetchResponse, error)
	// Read - View Trunk Phone Numbers List
	Read(context.Context, TrunkPhoneNumbersReadRequest) (TrunkPhoneNumbersReadResponse, error)
}

// TrunksService provides access to Trunks
type TrunksService interface {
	// Create - Create a Trunk
	Create(context.Context, TrunksCreateRequest) (TrunksCreateResponse, error)
	// Delete - Delete a Trunk
	Delete(context.Context, TrunksDeleteRequest) (TrunksDeleteResponse, error)
	// Fetch - View a Trunk
	Fetch(context.Context, TrunksFetchRequest) (TrunksFetchResponse, error)
	// Read - View Trunks List
	Read(context.Context, TrunksReadRequest) (TrunksReadResponse, error)
	// Update - Modify a Trunk
	Update(context.Context, TrunksUpdateRequest) (TrunksUpdateResponse, error)
}

// AlphaSendersService provides access to Alpha Senders
type AlphaSendersService interface {
	// Create - Create an Alpha Sender
	Create(context.Context, AlphaSendersCreateRequest) (AlphaSendersCreateResponse, error)
	// Delete - Delete an Alpha Sender
	Delete(context.Context, AlphaSendersDeleteRequest) (AlphaSendersDeleteResponse, error)
	// Fetch - View an Alpha Sender
	Fetch(context.Context, AlphaSendersFetchRequest) (AlphaSendersFetchResponse, error)
	// Read - View Alpha Senders List
	Read(context.Context, AlphaSendersReadRequest) (AlphaSendersReadResponse, error)
}

// MessagePhoneNumbersService provides access to Phone Numbers
type MessagePhoneNumbersService interface {
	// Create - Create a Message Phone Number
	Create(context.Context, MessagePhoneNumbersCreateRequest) (MessagePhoneNumbersCreateResponse, error)
	// Delete - Delete a Message Phone Number
	Delete(context.Context, MessagePhoneNumbersDeleteRequest) (MessagePhoneNumbersDeleteResponse, error)
	// Fetch - View a Message Phone Number
	Fetch(context.Context, MessagePhoneNumbersFetchRequest) (MessagePhoneNumbersFetchResponse, error)
	// Read - View Message Phone Numbers List
	Read(context.Context, MessagePhoneNumbersReadRequest) (MessagePhoneNumbersReadResponse, error)
}

// MessageServicesService provides access to Services
type MessageServicesService interface {
	// Create - Create a Message Service
	Create(context.Context, MessageServicesCreateRequest) (MessageServicesCreateResponse, error)
	// Delete - Delete a Message Service
	Delete(context.Context, MessageServicesDeleteRequest) (MessageServicesDeleteResponse, error)
	// Fetch - View a Message Service
	Fetch(context.Context, MessageServicesFetchRequest) (MessageServicesFetchResponse, error)
	// Read - View Message Services List
	Read(context.Context, MessageServicesReadRequest) (MessageServicesReadResponse, error)
	// Update - Modify a Message Service
	Update(context.Context, MessageServicesUpdateRequest) (MessageServicesUpdateResponse, error)
}

// MessageShortCodesService provides access to Short Codes
type MessageShortCodesService interface {
	// Create - Create a Message Short Code
	Create(context.Context, MessageShortCodesCreateRequest) (MessageShortCodesCreateResponse, error)
	// Delete - Delete a Message Short Code
	Delete(context.Context, MessageShortCodesDeleteRequest) (MessageShortCodesDeleteResponse, error)
	// Fetch - View a Message Short Code
	Fetch(context.Context, MessageShortCodesFetchRequest) (MessageShortCodesFetchResponse, error)
	// Read - View Message Short Codes List
	Read(context.Context, MessageShortCodesReadRequest) (MessageShortCodesReadResponse, error)
}

// MessagesService provides access to Messages
type MessagesService interface {
	// Create - Create a Message
	Create(context.Context, MessagesCreateRequest) (MessagesCreateResponse, error)
	// Delete - Delete a Message
	Delete(context.Context, MessagesDeleteRequest) (MessagesDeleteResponse, error)
	// Fetch - View a Message
	Fetch(context.Context, MessagesFetchRequest) (MessagesFetchResponse, error)
	// MediaDelete - Delete a Media
	MediaDelete(context.Context, MessagesMediaDeleteRequest) (MessagesMediaDeleteResponse, error)
	// MediaFetch - View a Media
	MediaFetch(context.Context, MessagesMediaFetchRequest) (MessagesMediaFetchResponse, error)
	// MediaRead - View Media List
	MediaRead(context.Context, MessagesMediaReadRequest) (MessagesMediaReadResponse, error)
	// MessageFeedbacksCreate - Create a Message Feedback
	MessageFeedbacksCreate(context.Context, MessagesMessageFeedbacksCreateRequest) (MessagesMessageFeedbacksCreateResponse, error)
	// Read - View Messages List
	Read(context.Context, MessagesReadRequest) (MessagesReadResponse, error)
	// Update - Modify a Message
	Update(context.Context, MessagesUpdateRequest) (MessagesUpdateResponse, error)
}

// SyncDocumentsService provides access to Sync Documents
type SyncDocumentsService interface {
	// Create - Create a Sync Document
	Create(context.Context, SyncDocumentsCreateRequest) (SyncDocumentsCreateResponse, error)
	// Delete - Delete a Sync Document
	Delete(context.Context, SyncDocumentsDeleteRequest) (SyncDocumentsDeleteResponse, error)
	// DocumentPermissionsDelete - Delete a Document Permission
	DocumentPermissionsDelete(context.Context, SyncDocumentsDocumentPermissionsDeleteRequest) (SyncDocumentsDocumentPermissionsDeleteResponse, error)
	// DocumentPermissionsFetch - View a Document Permission
	DocumentPermissionsFetch(context.Context, SyncDocumentsDocumentPermissionsFetchRequest) (SyncDocumentsDocumentPermissionsFetchResponse, error)
	// DocumentPermissionsRead - View Document Permissions List
	DocumentPermissionsRead(context.Context, SyncDocumentsDocumentPermissionsReadRequest) (SyncDocumentsDocumentPermissionsReadResponse, error)
	// DocumentPermissionsUpdate - Modify a Document Permission
	DocumentPermissionsUpdate(context.Context, SyncDocumentsDocumentPermissionsUpdateRequest) (SyncDocumentsDocumentPermissionsUpdateResponse, error)
	// Fetch - View a Sync Document
	Fetch(context.Context, SyncDocumentsFetchRequest) (SyncDocumentsFetchResponse, error)
	// Read - View Sync Documents List
	Read(context.Context, SyncDocumentsReadRequest) (SyncDocumentsReadResponse, error)
	// Update - Modify a Sync Document
	Update(context.Context, SyncDocumentsUpdateRequest) (SyncDocumentsUpdateResponse, error)
}

// SyncListsService provides access to Sync Lists
type SyncListsService interface {
	// Create - Create a Sync List
	Create(context.Context, SyncListsCreateRequest) (SyncListsCreateResponse, error)
	// Delete - Delete a Sync List
	Delete(context.Context, SyncListsDeleteRequest) (SyncListsDeleteResponse, error)
	// Fetch - View a Sync List
	Fetch(context.Context, SyncListsFetchRequest) (SyncListsFetchResponse, error)
	// Read - View Sync Lists List
	Read(context.Context, SyncListsReadRequest) (SyncListsReadResponse, error)
	// SyncListItemsCreate - Create a Sync List Item
	SyncListItemsCreate(context.Context, SyncListsSyncListItemsCreateRequest) (SyncListsSyncListItemsCreateResponse, error)
	// SyncListItemsDelete - Delete a Sync List Item
	SyncListItemsDelete(context.Context, SyncListsSyncListItemsDeleteRequest) (SyncListsSyncListItemsDeleteResponse, error)
	// SyncListItemsFetch - View a Sync List Item
	SyncListItemsFetch(context.Context, SyncListsSyncListItemsFetchRequest) (SyncListsSyncListItemsFetchResponse, error)
	// SyncListItemsRead - View Sync List Items List
	SyncListItemsRead(context.Context, SyncListsSyncListItemsReadRequest) (SyncListsSyncListItemsReadResponse, error)
	// SyncListItemsUpdate - Modify a Sync List Item
	SyncListItemsUpdate(context.Context, SyncListsSyncListItemsUpdateRequest) (SyncListsSyncListItemsUpdateResponse, error)
	// SyncListPermissionsDelete - Delete a Sync List Permission
	SyncListPermissionsDelete(context.Context, SyncListsSyncListPermissionsDeleteRequest) (SyncListsSyncListPermissionsDeleteResponse, error)
	// SyncListPermissionsFetch - View a Sync List Permission
	SyncListPermissionsFetch(context.Context, SyncListsSyncListPermissionsFetchRequest) (SyncListsSyncListPermissionsFetchResponse, error)
	// SyncListPermissionsRead - View Sync List Permissions List
	SyncListPermissionsRead(context.Context, SyncListsSyncListPermissionsReadRequest) (SyncListsSyncListPermissionsReadResponse, error)
	// SyncListPermissionsUpdate - Modify a Sync List Permission
	SyncListPermissionsUpdate(context.Context, SyncListsSyncListPermissionsUpdateRequest) (SyncListsSyncListPermissionsUpdateResponse, error)
}

// SyncMapsService provides access to Sync Maps
type SyncMapsService interface {
	// Create - Create a Sync Map
	Create(context.Context, SyncMapsCreateRequest) (SyncMapsCreateResponse, error)
	// Delete - Delete a Sync Map
	Delete(context.Context, SyncMapsDeleteRequest) (SyncMapsDeleteResponse, error)
	// Fetch - View a Sync Map
	Fetch(context.Context, SyncMapsFetchRequest) (SyncMapsFetchResponse, error)
	// Read - View Sync Maps List
	Read(context.Context, SyncMapsReadRequest) (SyncMapsReadResponse, error)
	// SyncMapItemsCreate - Create a Sync Map Item
	SyncMapItemsCreate(context.Context, SyncMapsSyncMapItemsCreateRequest) (SyncMapsSyncMapItemsCreateResponse, error)
	// SyncMapItemsDelete - Delete a Sync Map Item
	SyncMapItemsDelete(context.Context, SyncMapsSyncMapItemsDeleteRequest) (SyncMapsSyncMapItemsDeleteResponse, error)
	// SyncMapItemsFetch - View a Sync Map Item
	SyncMapItemsFetch(context.Context, SyncMapsSyncMapItemsFetchRequest) (SyncMapsSyncMapItemsFetchResponse, error)
	// SyncMapItemsRead - View Sync Map Items List
	SyncMapItemsRead(context.Context, SyncMapsSyncMapItemsReadRequest) (SyncMapsSyncMapItemsReadResponse, error)
	// SyncMapItemsUpdate - Modify a Sync Map Item
	SyncMapItemsUpdate(context.Context, SyncMapsSyncMapItemsUpdateRequest) (SyncMapsSyncMapItemsUpdateResponse, error)
	// SyncMapPermissionsDelete - Delete a Sync Map Permission
	SyncMapPermissionsDelete(context.Context, SyncMapsSyncMapPermissionsDeleteRequest) (SyncMapsSyncMapPermissionsDeleteResponse, error)
	// SyncMapPermissionsFetch - View a Sync Map Permission
	SyncMapPermissionsFetch(context.Context, SyncMapsSyncMapPermissionsFetchRequest) (SyncMapsSyncMapPermissionsFetchResponse, error)
	// SyncMapPermissionsRead - View Sync Map Permissions List
	SyncMapPermissionsRead(context.Context, SyncMapsSyncMapPermissionsReadRequest) (SyncMapsSyncMapPermissionsReadResponse, error)
	// SyncMapPermissionsUpdate - Modify a Sync Map Permission
	SyncMapPermissionsUpdate(context.Context, SyncMapsSyncMapPermissionsUpdateRequest) (SyncMapsSyncMapPermissionsUpdateResponse, error)
}

// SyncServicesService provides access to Sync Services
type SyncServicesService interface {
	// Create - Create Sync Service Instance
	Create(context.Context, SyncServicesCreateRequest) (SyncServicesCreateResponse, error)
	// Delete - Delete Sync Service Instance
	Delete(context.Context, SyncServicesDeleteRequest) (SyncServicesDeleteResponse, error)
	// Fetch - View Sync Service Instance
	Fetch(context.Context, SyncServicesFetchRequest) (SyncServicesFetchResponse, error)
	// Read - View Sync Services List
	Read(context.Context, SyncServicesReadRequest) (SyncServicesReadResponse, error)
	// Update - Update Service Instance
	Update(context.Context, SyncServicesUpdateRequest) (SyncServicesUpdateResponse, error)
}

// ActivitiesService provides access to Activities
type ActivitiesService interface {
	// Create - Create an Activity
	Create(context.Context, ActivitiesCreateRequest) (ActivitiesCreateResponse, error)
	// Delete - Delete an Activity
	Delete(context.Context, ActivitiesDeleteRequest) (ActivitiesDeleteResponse, error)
	// Fetch - View an Activity
	Fetch(context.Context, ActivitiesFetchRequest) (ActivitiesFetchResponse, error)
	// Read - View Activities List
	Read(context.Context, ActivitiesReadRequest) (ActivitiesReadResponse, error)
	// Update - Modify an Activity
	Update(context.Context, ActivitiesUpdateRequest) (ActivitiesUpdateResponse, error)
}

// EventsService provides access to Events
type EventsService interface {
	// Fetch - View an Event
	Fetch(context.Context, EventsFetchRequest) (EventsFetchResponse, error)
	// Read - View Events List
	Read(context.Context, EventsReadRequest) (EventsReadResponse, error)
}

// TaskChannelsService provides access to Task Channels
type TaskChannelsService interface {
	// Fetch - View a Task Channel
	Fetch(context.Context, TaskChannelsFetchRequest) (TaskChannelsFetchResponse, error)
	// Read - View Task Channels List
	Read(context.Context, TaskChannelsReadRequest) (TaskChannelsReadResponse, error)
}

// TaskQueuesService provides access to Task Queues
type TaskQueuesService interface {
	// Create - Create a Task Queue
	Create(context.Context, TaskQueuesCreateRequest) (TaskQueuesCreateResponse, error)
	// Delete - Delete a Task Queue
	Delete(context.Context, TaskQueuesDeleteRequest) (TaskQueuesDeleteResponse, error)
	// Fetch - View a Task Queue
	Fetch(context.Context, TaskQueuesFetchRequest) (TaskQueuesFetchResponse, error)
	// Read - View Task Queues List
	Read(context.Context, TaskQueuesReadRequest) (TaskQueuesReadResponse, error)
	// TaskQueueStatisticsFetch - View a Task Queue Statistic
	TaskQueueStatisticsFetch(context.Context, TaskQueuesTaskQueueStatisticsFetchRequest) (TaskQueuesTaskQueueStatisticsFetchResponse, error)
	// TaskQueuesStatisticsRead - View Task Queues Statistics List
	TaskQueuesStatisticsRead(context.Context, TaskQueuesTaskQueuesStatisticsReadRequest) (TaskQueuesTaskQueuesStatisticsReadResponse, error)
	// Update - Modify a Task Queue
	Update(context.Context, TaskQueuesUpdateRequest) (TaskQueuesUpdateResponse, error)
}

// TasksService provides access to Tasks
type TasksService interface {
	// Create - Create a Task
	Create(context.Context, TasksCreateRequest) (TasksCreateResponse, error)
	// Delete - Delete a Task
	Delete(context.Context, TasksDeleteRequest) (TasksDeleteResponse, error)
	// Fetch - View a Task
	Fetch(context.Context, TasksFetchRequest) (TasksFetchResponse, error)
	// Read - View Tasks List
	Read(context.Context, TasksReadRequest) (TasksReadResponse, error)
	// TaskReservationsFetch - View a Task Reservation
	TaskReservationsFetch(context.Context, TasksTaskReservationsFetchRequest) (TasksTaskReservationsFetchResponse, error)
	// TaskReservationsRead - View Task Reservations List
	TaskReservationsRead(context.Context, TasksTaskReservationsReadRequest) (TasksTaskReservationsReadResponse, error)
	// TaskReservationsUpdate - Modify a Task Reservation
	TaskReservationsUpdate(context.Context, TasksTaskReservationsUpdateRequest) (TasksTaskReservationsUpdateResponse, error)
	// Update - Modify a Task
	Update(context.Context, TasksUpdateRequest) (TasksUpdateResponse, error)
}

// WorkersService provides access to Workers
type WorkersService interface {
	// Create - Create a Worker
	Create(context.Context, WorkersCreateRequest) (WorkersCreateResponse, error)
	// Delete - Delete a Worker
	Delete(context.Context, WorkersDeleteRequest) (WorkersDeleteResponse, error)
	// Fetch - View a Worker
	Fetch(context.Context, WorkersFetchRequest) (WorkersFetchResponse, error)
	// Read - View Workers List
	Read(context.Context, WorkersReadRequest) (WorkersReadResponse, error)
	// Update - Modify a Worker
	Update(context.Context, WorkersUpdateRequest) (WorkersUpdateResponse, error)
	// WorkerChannelsFetch - View a Worker Channel
	WorkerChannelsFetch(context.Context, WorkersWorkerChannelsFetchRequest) (WorkersWorkerChannelsFetchResponse, error)
	// WorkerChannelsRead - View Worker Channels List
	WorkerChannelsRead(context.Context, WorkersWorkerChannelsReadRequest) (WorkersWorkerChannelsReadResponse, error)
	// WorkerChannelsUpdate - Modify a Worker Channel
	WorkerChannelsUpdate(context.Context, WorkersWorkerChannelsUpdateRequest) (WorkersWorkerChannelsUpdateResponse, error)
	// WorkerInstanceStatisticsFetch - View a Worker Instance Statistic
	WorkerInstanceStatisticsFetch(context.Context, WorkersWorkerInstanceStatisticsFetchRequest) (WorkersWorkerInstanceStatisticsFetchResponse, error)
	// WorkerReservationsFetch - View a Worker Reservation
	WorkerReservationsFetch(context.Context, WorkersWorkerReservationsFetchRequest) (WorkersWorkerReservationsFetchResponse, error)
	// WorkerReservationsRead - View Worker Reservations List
	WorkerReservationsRead(context.Context, WorkersWorkerReservationsReadRequest) (WorkersWorkerReservationsReadResponse, error)
	// WorkerReservationsUpdate - Modify a Worker Reservation
	WorkerReservationsUpdate(context.Context, WorkersWorkerReservationsUpdateRequest) (WorkersWorkerReservationsUpdateResponse, error)
	// WorkerStatisticsFetch - View a Worker Statistic
	WorkerStatisticsFetch(context.Context, WorkersWorkerStatisticsFetchRequest) (WorkersWorkerStatisticsFetchResponse, error)
}

// WorkflowsService provides access to Workflows
type WorkflowsService interface {
	// Create - Create a Workflow
	Create(context.Context, WorkflowsCreateRequest) (WorkflowsCreateResponse, error)
	// Delete - Delete a Workflow
	Delete(context.Context, WorkflowsDeleteRequest) (WorkflowsDeleteResponse, error)
	// Fetch - View a Workflow
	Fetch(context.Context, WorkflowsFetchRequest) (WorkflowsFetchResponse, error)
	// Read - View Workflows List
	Read(context.Context, WorkflowsReadRequest) (WorkflowsReadResponse, error)
	// Update - Modify a Workflow
	Update(context.Context, WorkflowsUpdateRequest) (WorkflowsUpdateResponse, error)
	// WorkflowStatisticsFetch - View a Workflow Statistic
	WorkflowStatisticsFetch(context.Context, WorkflowsWorkflowStatisticsFetchRequest) (WorkflowsWorkflowStatisticsFetchResponse, error)
}

// WorkspaceStatisticsService provides access to Workspace Statistics
type WorkspaceStatisticsService interface {
	// Fetch - View a Workspace Statistic
	Fetch(context.Context, WorkspaceStatisticsFetchRequest) (WorkspaceStatisticsFetchResponse, error)
}

// WorkspacesService provides access to Workspaces
type WorkspacesService interface {
	// Create - Create a Workspace
	Create(context.Context, WorkspacesCreateRequest) (WorkspacesCreateResponse, error)
	// Delete - Delete a Workspace
	Delete(context.Context, WorkspacesDeleteRequest) (WorkspacesDeleteResponse, error)
	// Fetch - View a Workspace
	Fetch(context.Context, WorkspacesFetchRequest) (WorkspacesFetchResponse, error)
	// Read - View Workspaces List
	Read(context.Context, WorkspacesReadRequest) (WorkspacesReadResponse, error)
	// Update - Modify a Workspace
	Update(context.Context, WorkspacesUpdateRequest) (WorkspacesUpdateResponse, error)
}

// UsageRecordsService provides access to Usage Records
type UsageRecordsService interface {
	// Read - View Usage Records List
	Read(context.Context, UsageRecordsReadRequest) (UsageRecordsReadResponse, error)
	// UsageRecordAllTimesRead - View Usage Record All Times List
	UsageRecordAllTimesRead(context.Context, UsageRecordsUsageRecordAllTimesReadRequest) (UsageRecordsUsageRecordAllTimesReadResponse, error)
	// UsageRecordDailiesRead - View Usage Record Dailies List
	UsageRecordDailiesRead(context.Context, UsageRecordsUsageRecordDailiesReadRequest) (UsageRecordsUsageRecordDailiesReadResponse, error)
	// UsageRecordLastMonthsRead - View Usage Record Last Months List
	UsageRecordLastMonthsRead(context.Context, UsageRecordsUsageRecordLastMonthsReadRequest) (UsageRecordsUsageRecordLastMonthsReadResponse, error)
	// UsageRecordMonthliesRead - View Usage Record Monthlies List
	UsageRecordMonthliesRead(context.Context, UsageRecordsUsageRecordMonthliesReadRequest) (UsageRecordsUsageRecordMonthliesReadResponse, error)
	// UsageRecordThisMonthsRead - View Usage Record This Months List
	UsageRecordThisMonthsRead(context.Context, UsageRecordsUsageRecordThisMonthsReadRequest) (UsageRecordsUsageRecordThisMonthsReadResponse, error)
	// UsageRecordTodaysRead - View Usage Record Todays List
	UsageRecordTodaysRead(context.Context, UsageRecordsUsageRecordTodaysReadRequest) (UsageRecordsUsageRecordTodaysReadResponse, error)
	// UsageRecordYearliesRead - View Usage Record Yearlies List
	UsageRecordYearliesRead(context.Context, UsageRecordsUsageRecordYearliesReadRequest) (UsageRecordsUsageRecordYearliesReadResponse, error)
	// UsageRecordYesterdaysRead - View Usage Record Yesterdays List
	UsageRecordYesterdaysRead(context.Context, UsageRecordsUsageRecordYesterdaysReadRequest) (UsageRecordsUsageRecordYesterdaysReadResponse, error)
}

// UsageTriggersService provides access to Usage Triggers
type UsageTriggersService interface {
	// Create - Create an Usage Trigger
	Create(context.Context, UsageTriggersCreateRequest) (UsageTriggersCreateResponse, error)
	// Delete - Delete an Usage Trigger
	Delete(context.Context, UsageTriggersDeleteRequest) (UsageTriggersDeleteResponse, error)
	// Fetch - View an Usage Trigger
	Fetch(context.Context, UsageTriggersFetchRequest) (UsageTriggersFetchResponse, error)
	// Read - View Usage Triggers List
	Read(context.Context, UsageTriggersReadRequest) (UsageTriggersReadResponse, error)
	// Update - Modify an Usage Trigger
	Update(context.Context, UsageTriggersUpdateRequest) (UsageTriggersUpdateResponse, error)
}

// VideoRecordingsService provides access to Video Recordings
type VideoRecordingsService interface {
	// Delete - Delete a Video Recording
	Delete(context.Context, VideoRecordingsDeleteRequest) (VideoRecordingsDeleteResponse, error)
	// Fetch - View a Video Recording
	Fetch(context.Context, VideoRecordingsFetchRequest) (VideoRecordingsFetchResponse, error)
	// Read - View Video Recordings List
	Read(context.Context, VideoRecordingsReadRequest) (VideoRecordingsReadResponse, error)
	// RecordingMediasFetch - View a Recording Media
	RecordingMediasFetch(context.Context, VideoRecordingsRecordingMediasFetchRequest) (VideoRecordingsRecordingMediasFetchResponse, error)
}

// VideoRoomsService provides access to Video Rooms
type VideoRoomsService interface {
	// Create - Create a Video Room
	Create(context.Context, VideoRoomsCreateRequest) (VideoRoomsCreateResponse, error)
	// Fetch - View a Video Room
	Fetch(context.Context, VideoRoomsFetchRequest) (VideoRoomsFetchResponse, error)
	// Read - View Video Rooms List
	Read(context.Context, VideoRoomsReadRequest) (VideoRoomsReadResponse, error)
	// RoomRecordingMediasFetch - View a Room Recording Media
	RoomRecordingMediasFetch(context.Context, VideoRoomsRoomRecordingMediasFetchRequest) (VideoRoomsRoomRecordingMediasFetchResponse, error)
	// RoomRecordingsFetch - View a Room Recording
	RoomRecordingsFetch(context.Context, VideoRoomsRoomRecordingsFetchRequest) (VideoRoomsRoomRecordingsFetchResponse, error)
	// RoomRecordingsRead - View Room Recordings List
	RoomRecordingsRead(context.Context, VideoRoomsRoomRecordingsReadRequest) (VideoRoomsRoomRecordingsReadResponse, error)
	// Update - Modify a Video Room
	Update(context.Context, VideoRoomsUpdateRequest) (VideoRoomsUpdateResponse, error)
}

// CallsService provides access to Calls
type CallsService interface {
	// CallFeedbackSummariesCreate - Create a Call Feedback Summary
	CallFeedbackSummariesCreate(context.Context, CallsCallFeedbackSummariesCreateRequest) (CallsCallFeedbackSummariesCreateResponse, error)
	// CallFeedbackSummariesDelete - Delete a Call Feedback Summary
	CallFeedbackSummariesDelete(context.Context, CallsCallFeedbackSummariesDeleteRequest) (CallsCallFeedbackSummariesDeleteResponse, error)
	// CallFeedbackSummariesFetch - View a Call Feedback Summary
	CallFeedbackSummariesFetch(context.Context, CallsCallFeedbackSummariesFetchRequest) (CallsCallFeedbackSummariesFetchResponse, error)
	// CallFeedbacksCreate - Create a Call Feedback
	CallFeedbacksCreate(context.Context, CallsCallFeedbacksCreateRequest) (CallsCallFeedbacksCreateResponse, error)
	// CallFeedbacksFetch - View a Call Feedback
	CallFeedbacksFetch(context.Context, CallsCallFeedbacksFetchRequest) (CallsCallFeedbacksFetchResponse, error)
	// CallFeedbacksUpdate - Modify a Call Feedback
	CallFeedbacksUpdate(context.Context, CallsCallFeedbacksUpdateRequest) (CallsCallFeedbacksUpdateResponse, error)
	// CallNotificationsDelete - Delete a Call Notification
	CallNotificationsDelete(context.Context, CallsCallNotificationsDeleteRequest) (CallsCallNotificationsDeleteResponse, error)
	// CallNotificationsFetch - View a Call Notification
	CallNotificationsFetch(context.Context, CallsCallNotificationsFetchRequest) (CallsCallNotificationsFetchResponse, error)
	// CallNotificationsRead - View Call Notifications List
	CallNotificationsRead(context.Context, CallsCallNotificationsReadRequest) (CallsCallNotificationsReadResponse, error)
	// CallRecordingsDelete - Delete a Call Recording
	CallRecordingsDelete(context.Context, CallsCallRecordingsDeleteRequest) (CallsCallRecordingsDeleteResponse, error)
	// CallRecordingsFetch - View a Call Recording
	CallRecordingsFetch(context.Context, CallsCallRecordingsFetchRequest) (CallsCallRecordingsFetchResponse, error)
	// CallRecordingsRead - View Call Recordings List
	CallRecordingsRead(context.Context, CallsCallRecordingsReadRequest) (CallsCallRecordingsReadResponse, error)
	// Create - Make a Call
	Create(context.Context, CallsCreateRequest) (CallsCreateResponse, error)
	// Delete - Delete a Call
	Delete(context.Context, CallsDeleteRequest) (CallsDeleteResponse, error)
	// Fetch - View a Call
	Fetch(context.Context, CallsFetchRequest) (CallsFetchResponse, error)
	// Read - View Calls List
	Read(context.Context, CallsReadRequest) (CallsReadResponse, error)
	// Update - Modify a Live Call
	Update(context.Context, CallsUpdateRequest) (CallsUpdateResponse, error)
}

// ConferencesService provides access to Conferences
type ConferencesService interface {
	// Fetch - View a Conference
	Fetch(context.Context, ConferencesFetchRequest) (ConferencesFetchResponse, error)
	// ParticipantsCreate - Create a Participant
	ParticipantsCreate(context.Context, ConferencesParticipantsCreateRequest) (ConferencesParticipantsCreateResponse, error)
	// ParticipantsDelete - Delete a Participant
	ParticipantsDelete(context.Context, ConferencesParticipantsDeleteRequest) (ConferencesParticipantsDeleteResponse, error)
	// ParticipantsFetch - View a Participant
	ParticipantsFetch(context.Context, ConferencesParticipantsFetchRequest) (ConferencesParticipantsFetchResponse, error)
	// ParticipantsRead - View Participants List
	ParticipantsRead(context.Context, ConferencesParticipantsReadRequest) (ConferencesParticipantsReadResponse, error)
	// ParticipantsUpdate - Modify a Participant
	ParticipantsUpdate(context.Context, ConferencesParticipantsUpdateRequest) (ConferencesParticipantsUpdateResponse, error)
	// Read - View Conferences List
	Read(context.Context, ConferencesReadRequest) (ConferencesReadResponse, error)
	// Update - Modify a Conference
	Update(context.Context, ConferencesUpdateRequest) (ConferencesUpdateResponse, error)
}

// QueuesService provides access to Queues
type QueuesService interface {
	// Create - Create a Queue
	Create(context.Context, QueuesCreateRequest) (QueuesCreateResponse, error)
	// Delete - Delete a Queue
	Delete(context.Context, QueuesDeleteRequest) (QueuesDeleteResponse, error)
	// Fetch - View a Queue
	Fetch(context.Context, QueuesFetchRequest) (QueuesFetchResponse, error)
	// MembersFetch - View a Member
	MembersFetch(context.Context, QueuesMembersFetchRequest) (QueuesMembersFetchResponse, error)
	// MembersRead - View Members List
	MembersRead(context.Context, QueuesMembersReadRequest) (QueuesMembersReadResponse, error)
	// MembersUpdate - Modify a Member
	MembersUpdate(context.Context, QueuesMembersUpdateRequest) (QueuesMembersUpdateResponse, error)
	// Read - View Queues List
	Read(context.Context, QueuesReadRequest) (QueuesReadResponse, error)
	// Update - Modify a Queue
	Update(context.Context, QueuesUpdateRequest) (QueuesUpdateResponse, error)
}

// RecordingsService provides access to Recordings
type RecordingsService interface {
	// Delete - Delete a Recording
	Delete(context.Context, RecordingsDeleteRequest) (RecordingsDeleteResponse, error)
	// Fetch - View a Recording
	Fetch(context.Context, RecordingsFetchRequest) (RecordingsFetchResponse, error)
	// Read - View Recordings List
	Read(context.Context, RecordingsReadRequest) (RecordingsReadResponse, error)
	// RecordingAddOnResultPayloadDatasFetch - View a Recording Add On Result Payload Data
	RecordingAddOnResultPayloadDatasFetch(context.Context, RecordingsRecordingAddOnResultPayloadDatasFetchRequest) (RecordingsRecordingAddOnResultPayloadDatasFetchResponse, error)
	// RecordingAddOnResultPayloadsDelete - Delete a Recording Add On Result Payload
	RecordingAddOnResultPayloadsDelete(context.Context, RecordingsRecordingAddOnResultPayloadsDeleteRequest) (RecordingsRecordingAddOnResultPayloadsDeleteResponse, error)
	// RecordingAddOnResultPayloadsFetch - View a Recording Add On Result Payload
	RecordingAddOnResultPayloadsFetch(context.Context, RecordingsRecordingAddOnResultPayloadsFetchRequest) (RecordingsRecordingAddOnResultPayloadsFetchResponse, error)
	// RecordingAddOnResultPayloadsRead - View Recording Add On Result Payloads List
	RecordingAddOnResultPayloadsRead(context.Context, RecordingsRecordingAddOnResultPayloadsReadRequest) (RecordingsRecordingAddOnResultPayloadsReadResponse, error)
	// RecordingAddOnResultsDelete - Delete a Recording Add On Result
	RecordingAddOnResultsDelete(context.Context, RecordingsRecordingAddOnResultsDeleteRequest) (RecordingsRecordingAddOnResultsDeleteResponse, error)
	// RecordingAddOnResultsFetch - View a Recording Add On Result
	RecordingAddOnResultsFetch(context.Context, RecordingsRecordingAddOnResultsFetchRequest) (RecordingsRecordingAddOnResultsFetchResponse, error)
	// RecordingAddOnResultsRead - View Recording Add On Results List
	RecordingAddOnResultsRead(context.Context, RecordingsRecordingAddOnResultsReadRequest) (RecordingsRecordingAddOnResultsReadResponse, error)
	// RecordingTranscriptionsDelete - Delete a Recording Transcription
	RecordingTranscriptionsDelete(context.Context, RecordingsRecordingTranscriptionsDeleteRequest) (RecordingsRecordingTranscriptionsDeleteResponse, error)
	// RecordingTranscriptionsFetch - View a Recording Transcription
	RecordingTranscriptionsFetch(context.Context, RecordingsRecordingTranscriptionsFetchRequest) (RecordingsRecordingTranscriptionsFetchResponse, error)
	// RecordingTranscriptionsRead - View Recording Transcriptions List
	RecordingTranscriptionsRead(context.Context, RecordingsRecordingTranscriptionsReadRequest) (RecordingsRecordingTranscriptionsReadResponse, error)
}

// TranscriptionsService provides access to Transcriptions
type TranscriptionsService interface {
	// Delete - Delete a Transcription
	Delete(context.Context, TranscriptionsDeleteRequest) (TranscriptionsDeleteResponse, error)
	// Fetch - View a Transcription
	Fetch(context.Context, TranscriptionsFetchRequest) (TranscriptionsFetchResponse, error)
	// Read - View Transcriptions List
	Read(context.Context, TranscriptionsReadRequest) (TranscriptionsReadResponse, error)
}

// WirelessCommandsService provides access to Wireless Commands
type WirelessCommandsService interface {
	// Create - Create a Wireless Command
	Create(context.Context, WirelessCommandsCreateRequest) (WirelessCommandsCreateResponse, error)
	// Fetch - View a Wireless Command
	Fetch(context.Context, WirelessCommandsFetchRequest) (WirelessCommandsFetchResponse, error)
	// Read - View Wireless Commands List
	Read(context.Context, WirelessCommandsReadRequest) (WirelessCommandsReadResponse, error)
}

// WirelessSimsService provides access to Wireless Sims
type WirelessSimsService interface {
	// Fetch - View a Wireless Sim
	Fetch(context.Context, WirelessSimsFetchRequest) (WirelessSimsFetchResponse, error)
	// Read - View Wireless Sims List
	Read(context.Context, WirelessSimsReadRequest) (WirelessSimsReadResponse, error)
	// Update - Modify a Wireless Sim
	Update(context.Context, WirelessSimsUpdateRequest) (WirelessSimsUpdateResponse, error)
	// UsageRecordsRead - View Usage Records List
	UsageRecordsRead(context.Context, WirelessSimsUsageRecordsReadRequest) (WirelessSimsUsageRecordsReadResponse, error)
}

// WirelessUsagesService provides access to Wireless Usages
type WirelessUsagesService interface {
	// Read - View Wireless Usages List
	Read(context.Context, WirelessUsagesReadRequest) (WirelessUsagesReadResponse, error)
}
