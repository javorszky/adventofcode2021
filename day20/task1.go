package day20

func task1(enhance, image string) int {
	img := NewImage(image, enhance)

	for i := 0; i < 2; i++ {
		img = img.tick()
	}

	return img.Lights()
}
