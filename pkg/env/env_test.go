package env

import (
       "testing"
)

func Testinit(t *testing.T) {

     // test for anything which may slip through env error checking

     // make sure we have APP_HOME env set
     // make sure it is type map[string][string]
     // check we have required items (apparently mostly everything is required ...

     t.Error("testing init")     

}

