/*
Package yext provides bindings for Yext Location Cloud APIs.

For full documentation visit http://developer.yext.com/docs/api-reference/

Usage

Create an authenticated client (requires an API key)

	client := yext.NewClient(yext.NewDefaultConfig().WithApiKey("[API KEY]"))

List all locations

	locs, err := client.LocationService.ListAll()

Fetch a single location

	loc, _, err := client.LocationService.Get("JB-01")

Create a new location (see full documentation for required fields)

	loc := &yext.Location{
		Id: yext.SingleString("JB-02"),
		Name: yext.SingleString("Joe's Bake Shop"),
	}
	loc, err := client.LocationService.Create(loc)

Edit an existing location

	loc := &yext.Location{
		Id: yext.SingleString("JB-02"),
		Name: yext.SingleString("Joe's Pastry Emporium"),
	}
	loc, err := client.LocationService.Edit(loc)

Configuration

The behavior of the API client can be controlled with a Config instance.  The Config type exposes chainable utility methods to make construction simpler.

	// Config with sane settings (prod host, 3 retries)
	yext.NewDefaultConfig()

	// Set authentication parameters
	yext.NewDefaultConfig().WithApiKey("[API KEY]")

	// Set authentication parameters from environment: $YEXT_API_KEY (required) and $YEXT_API_ACCOUNTID (optional)
	yext.NewDefaultConfig().WithEnvCredentials()

	// Communicate with the Yext Sandbox
	yext.NewConfig().WithSandboxHost()

By default, clients will retry API requests up to 3 times in the case of non-4xx errors including HTTP transport, 5xx responses, etc.  This can be modified via Config:

	// No retries
	yext.NewDefaultConfig().WithRetries(0)

Models

In order to support partial object updates, many of the struct attributes are represented as pointers in order to differentiate between "not-present" and "zero-valued".  Helpers are provided to make it easier to work with the pointers:

	l := &yext.Location{
		Id: yext.SingleString("JB-01"),
		SuppressAddress: yext.Bool(true),
		DisplayLat: yext.Float(38.813),
		Keywords: yext.Strings([]string{"pastries", "bakery", "food"})
	}

In addition, accessors are provided to make extracting data from the model objects simpler:

	l.GetId() // => "JB-01"
	l.GetSuppressAddress() // => true
	l.GetDisplayLat() // => 38.813
	l.GetKeywords() // => ["pastries", "bakery", "food"]

Error handling

Errors returned from the API are surfaced as Errors objects.  The Errors object is comprised of a list of errors, each with a Message, Code, and Type.  A full list of expected errors can be found here:  http://developer.yext.com/support/error-messages/

	_, _, err := client.LocationService.Create(&yext.Location{})

	// yext.Errors([{"code": "2068", message": "location.phone: The field location.phone is required", type: "FATAL_ERROR"}])

Response Metadata

Most functions that interact with the API will return at least two parameters - a yext.Response object and an error.  Response contains a Meta substructure that in turn has a UUID and Errors attribute.  The UUID can be used to look up individual requests in the developer.yext.com portal, useful for debugging requests.

	locs, resp, err := client.LocationService.List(nil)
	resp.Meta.UUID // "a219023e-d090-4e4e-9732-b823e9b66c8a"

Services

Most of the functionality within the API is exposed via domain-specific services available under the `Client` object.  For example, if you are interacting with Locations, use the Client instance's LocationService.  If you need to interact with Users, you'd use the UserService.

 * CategoryService
 * CustomFieldService
 * ListService
 * FolderService
 * LocationService
 * UserService

Each service provides a set of common data-access functions that you can use to interact with objects under the service's domain.

	Get() // Fetch by known ID
	List() // Fetch all (or first page in paged endpoints)
	ListAll() // Available in paged endpoints
	Edit() // Update and return API version
	Create() // Create and return API version

Where appropriate, services will expose additional, domain-specific functionality.

*/
package yext
