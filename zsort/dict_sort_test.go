package zsort

import (
	"testing"
)

func TestDictSort(t *testing.T) {
	var dict = []string{"Nonce", "SecretId", "Region", "Timestamp", "clusterIds.n", "clusterName", "limit", "orderField", "orderType", "status", "Action"}
	d := DictSort(dict)

	var result = []string{"Action", "Nonce", "Region", "SecretId", "Timestamp", "clusterIds.n", "clusterName", "limit", "orderField", "orderType", "status"}
	for i, _ := range d {
		if d[i] != result[i] {
			t.Error("Dict Sort Error!")
		}
	}
}
