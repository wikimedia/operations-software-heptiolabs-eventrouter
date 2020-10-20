/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package rockset
import (
    "bytes"
    "encoding/json"
    "fmt"
    
)

type User struct {
	// has value `user` for a user object
	Type_ string `json:"type,omitempty"`
	// ISO-8601 date
	CreatedAt string `json:"created_at,omitempty"`
	// user email
	Email string `json:"email"`
	// user first name
	FirstName string `json:"first_name,omitempty"`
	// user last name
	LastName string `json:"last_name,omitempty"`
	// List of roles for a given user
	Roles []string `json:"roles,omitempty"`
	// state of user - NEW / ACTIVE
	State string `json:"state,omitempty"`
}
func (m User) PrintResponse() {
    r, err := json.Marshal(m)
    var out bytes.Buffer
    err = json.Indent(&out, []byte(string(r)), "", "    ")
    if err != nil {
        fmt.Println("error parsing string")
        return
    }

    fmt.Println(out.String())
}

