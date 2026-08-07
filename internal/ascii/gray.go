package ascii

func colorToGray(rgb []byte) (gray []byte) {
	gray = make([]byte, 0, len(rgb)/3)

	for i := 0; i < len(rgb); i += 3 {
		r, g, b := float64(rgb[i]), float64(rgb[i+1]), float64(rgb[i+2])

		grayScale := r*0.2126 + g*0.7152 + b*0.0722

		gray = append(gray, byte(grayScale))
	}

	return gray
}
