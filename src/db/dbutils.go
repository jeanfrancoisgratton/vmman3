// vmman3 : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// source/db/db-utils.go
// 2022-08-25 13:32:28

package db

// This might get converted to generics, at some point
func interface2struct(hyps []DbHypervisors, sps []dbStoragePools, vms []dbVmStates, vmc []dbClusters) ([]interface{}, []interface{}, []interface{}, []interface{}) {
	dbH := make([]interface{}, len(hyps))
	for i, v := range hyps {
		dbH[i] = v
	}
	dbSP := make([]interface{}, len(sps))
	for i, v := range sps {
		dbSP[i] = v
	}
	dbVMs := make([]interface{}, len(vms))
	for i, v := range vms {
		dbVMs[i] = v
	}
	dbC := make([]interface{}, len(vmc))
	for i, v := range vmc {
		dbC[i] = v
	}

	return dbH, dbSP, dbVMs, dbC
}
