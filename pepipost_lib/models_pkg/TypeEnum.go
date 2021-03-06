/*
 * pepipost_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for TypeEnum enum
 */
type TypeEnum int

/**
 * Value collection for TypeEnum enum
 */
const (
    Type_AMP TypeEnum = 1 + iota
    Type_HTML
)

func (r TypeEnum) MarshalJSON() ([]byte, error) { 
    s := TypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *TypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  TypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts TypeEnum to its string representation
 */
func TypeEnumToValue(typeEnum TypeEnum) string {
    switch typeEnum {
        case Type_AMP:
    		return "amp"		
        case Type_HTML:
    		return "html"		
        default:
        	return "amp"
    }
}

/**
 * Converts TypeEnum Array to its string Array representation
*/
func TypeEnumArrayToValue(typeEnum []TypeEnum) []string {
    convArray := make([]string,len( typeEnum))
    for i:=0; i<len(typeEnum);i++ {
        convArray[i] = TypeEnumToValue(typeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func TypeEnumFromValue(value string) TypeEnum {
    switch value {
        case "AMP":
            return Type_AMP
        case "HTML":
            return Type_HTML
        default:
            return Type_AMP
    }
}
