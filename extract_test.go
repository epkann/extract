package extract

import (
    "testing"
)

// TestHighestLevelMap tests input whose highest level is a map.
func TestHighestLevelMap(t *testing.T) {
    const jsonBlob = `
{
    "type": "error",
    "status": 409,
    "code": "conflict",
    "context_info": {
        "errors": [
            {
                "reason": "invalid_parameter",
                "name": "group_tag_name",
                "message": "Invalid value 'All Box '. A resource with value 'All Box ' already exists"
            }
        ]
    },
    "help_url": "http://developers.box.com/docs/#errors",
    "message": "Resource with the same name already exists",
    "request_id": "2132632057555f584de87b7"
}`

    want := "Resource with the same name already exists"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestUnknownError tests a JSON blob with no error message.
func TestUnknownError(t *testing.T) {
    const jsonBlob = `
{
    "type": "error",
    "status": 409,
    "code": "conflict",
    "context_info": {
        "errors": [
            {
                "reason": "invalid_parameter",
                "name": "group_tag_name"
            }
        ]
    },
    "help_url": "http://developers.box.com/docs/#errors",
    "request_id": "2132632057555f584de87b7"
}`

    want := "unknown error"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestEmptyString tests an empty JSON string.
func TestEmptyString(t *testing.T) {
    const jsonBlob = ""

    want := "unexpected end of JSON input"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestInvalidJSON
func TestInvalidJSON(t *testing.T) {
    const jsonBlob = "blahblahblahblahblah"

    want := "invalid character 'b' looking for beginning of value"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestNonStringMessage tests a message key whose value is not a string.
func TestNonStringMessage(t *testing.T) {
const jsonBlob = `
{
    "type": "error",
    "status": 409,
    "code": "conflict",
    "context_info": {
        "errors": [
            {
                "reason": "invalid_parameter",
                "name": "group_tag_name",
        "message": 18
            }
        ]
    },
    "help_url": "http://developers.box.com/docs/#errors",
    "request_id": "2132632057555f584de87b7",
    "message": true
}`

    want := "unknown error"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestHighestLevelSlice tests input whose highest level is a slice.
func TestHighestLevelSlice(t *testing.T) {
const jsonBlob = `
[   18, false, 3.14, "Smaug",
            {
                "reason": "invalid_parameter",
                "name": "group_tag_name",
        "errors": [
            {
                "blah": 2,
                "arghk": true,
                "dijifjg": null,
                "message": "Everything is terrible."
            }
        ]
            }
]`

    want := "Everything is terrible."

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}



// TestMessageBomb tests whether extractErrorFromJSON() works when there are messages all 
// over the place.
func TestMessageBomb(t *testing.T) {
const jsonBlob = `
[   18, false, "Smaug",
            {
        "message": 3,
                "reason": "invalid_parameter",
                "name": "group_tag_name",
        "message": [
            {
                "message": 2.718,
                "message": true,
                "dijifjg": null,
                "message": [2, 3, 4],
                "message": "Voldemort",
                "zzzz": [
                    {
                        "message": "Sauron",
                        "gihif": "abababab",
                        "hha": [11, 12, 0],
                        "ie": [
                            {
                                "message": "Mordred",
                                "message": "ice giant",
                                "afkg": 20                          
                            }
                        ]
                    }                   
                ]
            }
        ]
            }
]`

    want := "Voldemort"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


// TestDeepEmbedded
func TestDeepEmbedded(t *testing.T) {
    const jsonBlob = `
{
"this":[    18, false, "Smaug",
            {
        "fijg": 3,
                "reason": "invalid_parameter",
                "name": "group_tag_name",
        "capybara": [
            {
                "fig": 2.718,
                "sloth": true,
                "burrow": null,
                "squirrel": [2, 3, 4],
                "velociraptor": "fossil",
                "vole": [
                    {
                        "hello": "Sauron",
                        "mouse": "abababab",
                        "avocado": [11, 12, 0],
                        "ie": [
                            {
                                "lol": "Mordred",    
                                "message": "ice giant",
                                "afkg": 20                          
                            }
                        ]
                    }                   
                ]
            }
        ]
            }
]
}`

    want := "ice giant"

    got := ExtractErrorFromJSON(jsonBlob)
    if got.Error() != want {
        t.Errorf("Want: %q \nGot: %q", want, got)
    }
}


