package mathutil

func MapInt(v int, inMin, inMax, outMin, outMax int) int {
	var unit = (outMax - outMin) / (inMax - inMin)
	return (v-inMin)*unit + outMin
}

func MapFloat64(v float64, inMin, inMax, outMin, outMax float64) float64 {
	var unit = (outMax - outMin) / (inMax - inMin)
	return (v-inMin)*unit + outMin
}

func MapFloat32(v float32, inMin, inMax, outMin, outMax float32) float32 {
	var unit = (outMax - outMin) / (inMax - inMin)
	return (v-inMin)*unit + outMin
}
