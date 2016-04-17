/*
Package yext provides bindings for Yext Platform API.

For full Platform API documentation visit https://www.yext.com/support/platform-api/

Usage

Create an authenticated client

	// `user`, `pass`, and `customerid` should be populated with
	// account credentials for a user that has the "API Access" role.
	client := yext.NewClient(yext.NewDefaultConfig().WithCredentials("[user]", "[pass]", "[customerid]"))

List all locations

	locs, err := client.LocationService.List()

Fetch a single location

	loc, err := client.LocationService.Get("JB-01")

Create a new location (see full documentation for required fields)

	loc := &yext.Location{
		Id: yext.String("JB-02"),
		Name: yext.String("Joe's Bake Shop"),
	}
	loc, err := client.LocationService.Create(loc)

Edit an existing location

	loc := &yext.Location{
		Id: yext.String("JB-02"),
		Name: yext.String("Joe's Pastry Emporium"),
	}
	loc, err := client.LocationService.Edit(loc)

Configuration

The behavior of the API client can be controlled with a Config instance.  The Config type exposes chainable utility methods to make construction simpler.

	// Config with sane settings (prod host, 3 retries)
	yext.NewDefaultConfig()

	// Set authentication parameters
	yext.NewDefaultConfig().WithCredentials("[user]", "[password]", "[customerid]")

	// Set authentication parameters from environment
	yext.NewDefaultConfig().WithEnvCredentials()

	// Communicate with Yext Sandbox
	yext.NewConfig().WithSandboxHost()

Models

In order to support partial object updates, many of the struct attributes are represented as pointers in order to differentiate between "not-present" and "zero-valued".  Helpers are provided to make it easier to work with the pointers:

	l := &yext.Location{
		Id: yext.String("JB-01"),
		SuppressAddress: yext.Bool(true),
		DisplayLat: yext.Float(38.813),
	}

In addition, accessors are provided to make extracting data from the model objects simpler:

	l.GetId() // => "JB-01"
	l.GetSuppressAddress() // => true
	l.GetDisplayLat() // => 38.813

Error handling

Errors returned from the API are surfaced as ErrorResponse objects. The ErrorReponse contains a list of encountered errors, each with a Message and Code, as well as the raw http.Response object.  A full list a expected errors can be found here:  https://www.yext.com/support/platform-api/#API_Conventions_and_Authentication.htm

	_, err := client.LocationService.Create(&yext.Location{})

	// {"errors": [{"code": "123", message": "Id cannot be blank"}]}

One exception to the above is 404 "Not Found" errors which are returned as the named error `ResourceNotFound`.

All non-4xx errors (transport, API 5xx responses, etc.) are elligible for automatic retries.

Services

Most of the functionality within the API is exposed via domain-specific services available under the `Client` object.  For example, if you are interacting with Locations, use the Client instance's LocationService.  If you need to interact with Users, you'd use the UserService.

 * CustomFieldService
 * ECLService
 * FolderService
 * LicenseService
 * LocationService
 * UserService

Each service provides a set of common data-access functions that you can use to interact with objects under the service's domain.

	Get() // Fetch by known ID
	List() // Fetch all
	Edit() // Update and return API version
	Create() // Create and return API version

Where appropriate, services will expose additional, domain-specific functionality.

*/
package yext
