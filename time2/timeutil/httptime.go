package timeutil

import (
	"net/http"
	"strconv"
	"time"
)

func TimeHandler(w http.ResponseWriter, _ *http.Request) {
	t1 := float64(time.Now().UnixNano()) / 1e6 // ms
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//{"t1":1668210427454.1458,"t2":1668210427454.1487}
	var buffer []byte = make([]byte, 0, 50)
	buffer = append(buffer, "{\"t1\":"...)
	buffer = strconv.AppendFloat(buffer, t1, 'f', 4, 64)
	buffer = append(buffer, ",\"t2\":"...)
	t2 := float64(time.Now().UnixNano()) / 1e6 // ms
	buffer = strconv.AppendFloat(buffer, t2, 'f', 4, 64)
	buffer = append(buffer, "}"...)
	w.Write(buffer)
}
