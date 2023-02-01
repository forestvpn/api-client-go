# Go API client for forestvpn_api

ForestVPN - Fast, secure, and modern VPN. It offers Distributed Computing, Crypto Mining, P2P, Ad Blocking, TOR over VPN,
30+ locations, and a free version with unlimited data.


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import forestvpn_api "github.com/forestvpn/api-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), forestvpn_api.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), forestvpn_api.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), forestvpn_api.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), forestvpn_api.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.forestvpn.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdsApi* | [**CreateAdUnitRequestLog**](docs/AdsApi.md#createadunitrequestlog) | **Post** /ads/request-logs/ | Create ad unit request log
*AdsApi* | [**ListAdPlacements**](docs/AdsApi.md#listadplacements) | **Get** /ads/placements/ | Get ad placement list
*AnalyticsApi* | [**GetDataUsageStats**](docs/AnalyticsApi.md#getdatausagestats) | **Get** /analytics/data-usage/ | Data Usage Stats
*AppApi* | [**GetCurrentUserDevice**](docs/AppApi.md#getcurrentuserdevice) | **Get** /app/devices/current/ | Get user device info
*AppApi* | [**UpdateCurrentUserDevice**](docs/AppApi.md#updatecurrentuserdevice) | **Patch** /app/devices/current/ | Update user device
*AppleApi* | [**VerifyAppStoreReceipt**](docs/AppleApi.md#verifyappstorereceipt) | **Post** /purchase/apple/verify/ | App store receipt verification
*AuthApi* | [**AuthorizeAccessTokenRequest**](docs/AuthApi.md#authorizeaccesstokenrequest) | **Post** /auth/access-token-requests/{requestID}/authorize/ | Authorize access token request
*AuthApi* | [**CreateAccessTokenRequest**](docs/AuthApi.md#createaccesstokenrequest) | **Post** /auth/access-token-requests/ | Create access token request
*AuthApi* | [**GetAccessTokenRequest**](docs/AuthApi.md#getaccesstokenrequest) | **Get** /auth/access-token-requests/{requestID}/ | Get access token request details
*AuthApi* | [**LoginToken**](docs/AuthApi.md#logintoken) | **Post** /auth/token/login/ | Login with JWT token
*AuthApi* | [**MigrateLegacyAuth**](docs/AuthApi.md#migratelegacyauth) | **Get** /legacy/auth/ | Legacy auth migration
*AuthApi* | [**ObtainToken**](docs/AuthApi.md#obtaintoken) | **Get** /auth/token/obtain/ | Obtain JWT token
*AuthApi* | [**RevokeAccessTokenRequest**](docs/AuthApi.md#revokeaccesstokenrequest) | **Post** /auth/access-token-requests/{requestID}/revoke/ | Revoke access token request
*AuthApi* | [**UpdateUserProfile**](docs/AuthApi.md#updateuserprofile) | **Patch** /auth/profile/ | Update profile
*AuthApi* | [**UserProfile**](docs/AuthApi.md#userprofile) | **Get** /auth/profile/ | Profile
*AuthApi* | [**WhoAmI**](docs/AuthApi.md#whoami) | **Get** /auth/whoami/ | Who am I
*BillingApi* | [**CancelSubscription**](docs/BillingApi.md#cancelsubscription) | **Delete** /billing/subscriptions/{subscriptionID}/ | Cancel subscription
*BillingApi* | [**CreatePaymentMethodStripeSetupIntent**](docs/BillingApi.md#createpaymentmethodstripesetupintent) | **Post** /billing/payment-methods/stripe/setup-intents/ | Create stripe&#39;s setup intent for add new payment method.
*BillingApi* | [**CreateSubscription**](docs/BillingApi.md#createsubscription) | **Post** /billing/subscriptions/ | Create subscription
*BillingApi* | [**DeletePaymentMethod**](docs/BillingApi.md#deletepaymentmethod) | **Delete** /billing/payment-methods/{paymentMethodID}/ | Delete payment method
*BillingApi* | [**GetBillingAccount**](docs/BillingApi.md#getbillingaccount) | **Get** /billing/account/ | Billing account info
*BillingApi* | [**GetBillingBundle**](docs/BillingApi.md#getbillingbundle) | **Get** /billing/bundles/{bundleID}/ | Bundle info
*BillingApi* | [**GetBillingPaymentOption**](docs/BillingApi.md#getbillingpaymentoption) | **Get** /billing/payment-options/{paymentOptionID}/ | Payment option info
*BillingApi* | [**GetBillingProduct**](docs/BillingApi.md#getbillingproduct) | **Get** /billing/products/{productID}/ | Product info
*BillingApi* | [**GetPaymentMethod**](docs/BillingApi.md#getpaymentmethod) | **Get** /billing/payment-methods/{paymentMethodID}/ | Payment method info
*BillingApi* | [**GetSubscription**](docs/BillingApi.md#getsubscription) | **Get** /billing/subscriptions/{subscriptionID}/ | Subscription info
*BillingApi* | [**GetSubscriptionItem**](docs/BillingApi.md#getsubscriptionitem) | **Get** /billing/subscription-items/{subscriptionItemID}/ | Subscription item info
*BillingApi* | [**ListBillingBundles**](docs/BillingApi.md#listbillingbundles) | **Get** /billing/bundles/ | Billing bundles list
*BillingApi* | [**ListBillingFeatures**](docs/BillingApi.md#listbillingfeatures) | **Get** /billing/features/ | Billing feature list
*BillingApi* | [**ListBillingPaymentOptions**](docs/BillingApi.md#listbillingpaymentoptions) | **Get** /billing/payment-options/ | Billing payment option list
*BillingApi* | [**ListBillingProducts**](docs/BillingApi.md#listbillingproducts) | **Get** /billing/products/ | Billing products list
*BillingApi* | [**ListPaymentMethods**](docs/BillingApi.md#listpaymentmethods) | **Get** /billing/payment-methods/ | Payment method list
*BillingApi* | [**ListSubscriptionItems**](docs/BillingApi.md#listsubscriptionitems) | **Get** /billing/subscription-items/ | Subscription items list
*BillingApi* | [**ListSubscriptions**](docs/BillingApi.md#listsubscriptions) | **Get** /billing/subscriptions/ | Billing subscriptions list
*BillingApi* | [**ReactivateSubscription**](docs/BillingApi.md#reactivatesubscription) | **Post** /billing/subscriptions/{subscriptionID}/reactivate/ | Reactivate subscription
*BillingApi* | [**UpdateBillingAccount**](docs/BillingApi.md#updatebillingaccount) | **Patch** /billing/account/ | Update billing account
*BillingApi* | [**UpdateSubscriptionItem**](docs/BillingApi.md#updatesubscriptionitem) | **Patch** /billing/subscription-items/{subscriptionItemID}/ | Update subscription item
*CheckoutApi* | [**ApplyCouponCheckoutSession**](docs/CheckoutApi.md#applycouponcheckoutsession) | **Post** /checkout/sessions/{sessionID}/apply_coupon/ | Apply coupon to session
*CheckoutApi* | [**CreateCheckoutSession**](docs/CheckoutApi.md#createcheckoutsession) | **Post** /checkout/sessions/ | Create checkout session
*CheckoutApi* | [**CreateWaitListRequest**](docs/CheckoutApi.md#createwaitlistrequest) | **Post** /checkout/wait-list/ | Create request to add country in wait list
*CheckoutApi* | [**ExpireCheckoutSession**](docs/CheckoutApi.md#expirecheckoutsession) | **Post** /checkout/sessions/{sessionID}/expire/ | Expire checkout session
*CheckoutApi* | [**GetCheckoutSession**](docs/CheckoutApi.md#getcheckoutsession) | **Get** /checkout/sessions/{sessionID}/ | Checkout session details
*CheckoutApi* | [**GetStripeCheckoutSession**](docs/CheckoutApi.md#getstripecheckoutsession) | **Get** /checkout/sessions/{sessionID}/stripe/checkout/session/ | Stripe checkout session details
*CheckoutApi* | [**GetStripePaymentIntent**](docs/CheckoutApi.md#getstripepaymentintent) | **Get** /checkout/sessions/{sessionID}/stripe/payment/intent/ | Stripe payment intent details
*CheckoutApi* | [**ProcessCloudPaymentsAuth**](docs/CheckoutApi.md#processcloudpaymentsauth) | **Post** /checkout/sessions/{sessionID}/cloud-payments/auth/ | Cloud payments auth
*CheckoutApi* | [**ProcessCloudPaymentsPost3ds**](docs/CheckoutApi.md#processcloudpaymentspost3ds) | **Post** /checkout/sessions/{sessionID}/cloud-payments/post3ds/ | Cloud payments post3ds
*DeviceApi* | [**CreateDevice**](docs/DeviceApi.md#createdevice) | **Post** /devices/ | Create new device
*DeviceApi* | [**CreateDevicePortForwarding**](docs/DeviceApi.md#createdeviceportforwarding) | **Post** /devices/{deviceID}/port-forwarding/ | Create new device port forwarding
*DeviceApi* | [**DeleteDevice**](docs/DeviceApi.md#deletedevice) | **Delete** /devices/{deviceID}/ | Delete Device
*DeviceApi* | [**DeleteDevicePortForwarding**](docs/DeviceApi.md#deletedeviceportforwarding) | **Delete** /devices/{deviceID}/port-forwarding/{portForwardingID}/ | Delete Device&#39;s Port Forwarding
*DeviceApi* | [**GetDevice**](docs/DeviceApi.md#getdevice) | **Get** /devices/{deviceID}/ | Device Info
*DeviceApi* | [**GetDeviceStats**](docs/DeviceApi.md#getdevicestats) | **Get** /devices/{deviceID}/stats/{statsID}/ | Device&#39;s stats detail
*DeviceApi* | [**GetDeviceWireGuard**](docs/DeviceApi.md#getdevicewireguard) | **Get** /devices/{deviceID}/wireguards/{wireGuardID}/ | Device&#39;s wireguard profile detail
*DeviceApi* | [**ListDeviceBindings**](docs/DeviceApi.md#listdevicebindings) | **Get** /devices/{deviceID}/bindings/ | Device bindings
*DeviceApi* | [**ListDeviceConnectionModes**](docs/DeviceApi.md#listdeviceconnectionmodes) | **Get** /devices/{deviceID}/connection-modes/ | Device connection modes
*DeviceApi* | [**ListDeviceDetailStats**](docs/DeviceApi.md#listdevicedetailstats) | **Get** /devices/{deviceID}/detail-stats/ | Device&#39;s detail stats list
*DeviceApi* | [**ListDevicePortForwardings**](docs/DeviceApi.md#listdeviceportforwardings) | **Get** /devices/{deviceID}/port-forwarding/ | Device Port Forwarding List
*DeviceApi* | [**ListDeviceStats**](docs/DeviceApi.md#listdevicestats) | **Get** /devices/{deviceID}/stats/ | Device&#39;s stats list
*DeviceApi* | [**ListDeviceWireGuardPeers**](docs/DeviceApi.md#listdevicewireguardpeers) | **Get** /devices/{deviceID}/wireguards/{wireGuardID}/peers/ | Device&#39;s wireguard peers
*DeviceApi* | [**ListDeviceWireGuards**](docs/DeviceApi.md#listdevicewireguards) | **Get** /devices/{deviceID}/wireguards/ | Device&#39;s wireguard profiles list
*DeviceApi* | [**ListDevices**](docs/DeviceApi.md#listdevices) | **Get** /devices/ | Device List
*DeviceApi* | [**UpdateDevice**](docs/DeviceApi.md#updatedevice) | **Patch** /devices/{deviceID}/ | Update device properties
*DeviceApi* | [**UpdateDevicePortForwarding**](docs/DeviceApi.md#updatedeviceportforwarding) | **Patch** /devices/{deviceID}/port-forwarding/{portForwardingID}/ | Update device&#39;s port forwarding
*FcmApi* | [**CreateFCMDevice**](docs/FcmApi.md#createfcmdevice) | **Post** /notification/fcm/ | Device registration for push notification through out Firebase Cloud Messaging
*FcmApi* | [**DeleteFCMDevice**](docs/FcmApi.md#deletefcmdevice) | **Delete** /notification/fcm/{registrationID}/ | Delete fcm device
*FcmApi* | [**GetFCMDevice**](docs/FcmApi.md#getfcmdevice) | **Get** /notification/fcm/{registrationID}/ | Device info
*FcmApi* | [**UpdateFCMDevice**](docs/FcmApi.md#updatefcmdevice) | **Patch** /notification/fcm/{registrationID}/ | Update device fcm properties
*FriendshipApi* | [**AcceptFriendshipInvitation**](docs/FriendshipApi.md#acceptfriendshipinvitation) | **Patch** /friendship/invitations/{code}/ | Accept friendship invitation
*FriendshipApi* | [**CreateFriendshipInvitation**](docs/FriendshipApi.md#createfriendshipinvitation) | **Post** /friendship/invitations/ | Create friendship invitation
*FriendshipApi* | [**DeleteFriend**](docs/FriendshipApi.md#deletefriend) | **Delete** /friendship/friends/{id}/ | Delete friend
*FriendshipApi* | [**GetFriend**](docs/FriendshipApi.md#getfriend) | **Get** /friendship/friends/{id}/ | Friend details
*FriendshipApi* | [**GetFriendshipInvitation**](docs/FriendshipApi.md#getfriendshipinvitation) | **Get** /friendship/invitations/{code}/ | Friendship invitation details
*FriendshipApi* | [**ListFriends**](docs/FriendshipApi.md#listfriends) | **Get** /friendship/friends/ | Get friends list
*FriendshipApi* | [**RejectFriendshipInvitation**](docs/FriendshipApi.md#rejectfriendshipinvitation) | **Delete** /friendship/invitations/{code}/ | Reject friendship invitation
*GeoApi* | [**ListCountries**](docs/GeoApi.md#listcountries) | **Get** /geo/countries/ | Countries list
*GeoApi* | [**ListCurrencies**](docs/GeoApi.md#listcurrencies) | **Get** /geo/currencies/ | Correncies list
*GeoApi* | [**ListLocations**](docs/GeoApi.md#listlocations) | **Get** /locations/ | Location list
*GoogleApi* | [**VerifyPlayStorePurchase**](docs/GoogleApi.md#verifyplaystorepurchase) | **Post** /purchase/google/verify/ | Play store purchase verification
*NotificationsApi* | [**GetNotificationsUnreadCount**](docs/NotificationsApi.md#getnotificationsunreadcount) | **Get** /notifications/unread_count/ | Get unread notifications count
*NotificationsApi* | [**ListNotifications**](docs/NotificationsApi.md#listnotifications) | **Get** /notifications/all_list/ | Get notifications list
*NotificationsApi* | [**UpdateNotificationMarkRead**](docs/NotificationsApi.md#updatenotificationmarkread) | **Get** /notifications/mark_as_read/{slug}/ | Mark notification as read by user
*NotificationsApi* | [**UpdateNotificationMarkReadAll**](docs/NotificationsApi.md#updatenotificationmarkreadall) | **Get** /notifications/mark_all_as_read/ | Mark all notifications as read by user
*SupportApi* | [**CreateSupportTicket**](docs/SupportApi.md#createsupportticket) | **Post** /support/tickets/ | Create support ticket
*SupportApi* | [**GetSupportTicketCategory**](docs/SupportApi.md#getsupportticketcategory) | **Get** /support/ticket-categories/ | Get ticket categories
*WireguardApi* | [**GetWireGuardPeerInfo**](docs/WireguardApi.md#getwireguardpeerinfo) | **Get** /wireguard/peers/{pubKey}/ | Wireguard peer info
*WireguardApi* | [**ListWireGuardPeers**](docs/WireguardApi.md#listwireguardpeers) | **Get** /wireguard/peers/ | Wireguard peers list


## Documentation For Models

 - [AccessTokenRequest](docs/AccessTokenRequest.md)
 - [AdPlacement](docs/AdPlacement.md)
 - [AdProvider](docs/AdProvider.md)
 - [AdReward](docs/AdReward.md)
 - [AdUnit](docs/AdUnit.md)
 - [AggregatedDataUsageStats](docs/AggregatedDataUsageStats.md)
 - [AppStoreReceiptVerificationRequest](docs/AppStoreReceiptVerificationRequest.md)
 - [BillingAccount](docs/BillingAccount.md)
 - [BillingFeature](docs/BillingFeature.md)
 - [Bundle](docs/Bundle.md)
 - [CheckoutSession](docs/CheckoutSession.md)
 - [CheckoutSessionProduct](docs/CheckoutSessionProduct.md)
 - [CloudPaymentsAuth](docs/CloudPaymentsAuth.md)
 - [CloudPaymentsPost3ds](docs/CloudPaymentsPost3ds.md)
 - [CloudPaymentsSecure3d](docs/CloudPaymentsSecure3d.md)
 - [ConnectionMode](docs/ConnectionMode.md)
 - [Constraint](docs/Constraint.md)
 - [Country](docs/Country.md)
 - [CouponCheckoutSession](docs/CouponCheckoutSession.md)
 - [CreateAdUnitRequestLogRequest](docs/CreateAdUnitRequestLogRequest.md)
 - [CreateCheckoutSessionProduct](docs/CreateCheckoutSessionProduct.md)
 - [CreateCheckoutSessionRequest](docs/CreateCheckoutSessionRequest.md)
 - [CreateCloudPaymentsAuth](docs/CreateCloudPaymentsAuth.md)
 - [CreateCloudPaymentsPost3ds](docs/CreateCloudPaymentsPost3ds.md)
 - [CreateCouponCheckoutSession](docs/CreateCouponCheckoutSession.md)
 - [CreateFCMDeviceRequest](docs/CreateFCMDeviceRequest.md)
 - [CreateOrUpdateDeviceRequest](docs/CreateOrUpdateDeviceRequest.md)
 - [CreateOrUpdateDeviceRequestInfo](docs/CreateOrUpdateDeviceRequestInfo.md)
 - [CreateOrUpdatePortForwardingRequest](docs/CreateOrUpdatePortForwardingRequest.md)
 - [CreateStripeSetupIntentRequest](docs/CreateStripeSetupIntentRequest.md)
 - [CreateSubscriptionRequest](docs/CreateSubscriptionRequest.md)
 - [CreateTokenLogin](docs/CreateTokenLogin.md)
 - [Currency](docs/Currency.md)
 - [Device](docs/Device.md)
 - [DeviceRecord](docs/DeviceRecord.md)
 - [DeviceStats](docs/DeviceStats.md)
 - [DeviceType](docs/DeviceType.md)
 - [Discount](docs/Discount.md)
 - [Environment](docs/Environment.md)
 - [Error](docs/Error.md)
 - [FCMDevice](docs/FCMDevice.md)
 - [Friendship](docs/Friendship.md)
 - [FriendshipInvitation](docs/FriendshipInvitation.md)
 - [LegacyAuthMigrationToken](docs/LegacyAuthMigrationToken.md)
 - [Location](docs/Location.md)
 - [NetworkService](docs/NetworkService.md)
 - [Notification](docs/Notification.md)
 - [NotificationAllList](docs/NotificationAllList.md)
 - [NotificationType](docs/NotificationType.md)
 - [NotificationUnreadCount](docs/NotificationUnreadCount.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentMethodCard](docs/PaymentMethodCard.md)
 - [PaymentMethodType](docs/PaymentMethodType.md)
 - [PaymentOption](docs/PaymentOption.md)
 - [PlayStorePurchaseVerificationRequest](docs/PlayStorePurchaseVerificationRequest.md)
 - [PortForwarding](docs/PortForwarding.md)
 - [Price](docs/Price.md)
 - [Product](docs/Product.md)
 - [ProductWithoutPrice](docs/ProductWithoutPrice.md)
 - [Recurring](docs/Recurring.md)
 - [Server](docs/Server.md)
 - [StripeCheckoutSession](docs/StripeCheckoutSession.md)
 - [StripePaymentIntent](docs/StripePaymentIntent.md)
 - [StripeSetupIntent](docs/StripeSetupIntent.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionItem](docs/SubscriptionItem.md)
 - [SubscriptionSource](docs/SubscriptionSource.md)
 - [SubscriptionStatus](docs/SubscriptionStatus.md)
 - [TicketCategory](docs/TicketCategory.md)
 - [TokenLogin](docs/TokenLogin.md)
 - [TokenObtain](docs/TokenObtain.md)
 - [UpdateFCMDeviceRequest](docs/UpdateFCMDeviceRequest.md)
 - [UpdateUserDeviceRequest](docs/UpdateUserDeviceRequest.md)
 - [User](docs/User.md)
 - [UserDevice](docs/UserDevice.md)
 - [WireGuard](docs/WireGuard.md)
 - [WireGuardPeer](docs/WireGuardPeer.md)
 - [WireGuardPeerDevice](docs/WireGuardPeerDevice.md)
 - [WireGuardPeerInfo](docs/WireGuardPeerInfo.md)
 - [WireGuardPeerUser](docs/WireGuardPeerUser.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


### wireguardAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

ForestVPN - support@forestvpn.com

## Credits

- ForestVPN.com [Free VPN](https://forestvpn.com) for all
- SpaceV.net [VPN for teams](https://spacev.net)
