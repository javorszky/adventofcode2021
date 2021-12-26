package day20

func task2(enhance, image string) int {
	img := NewImage(image, enhance)

	for i := 0; i < 50; i++ {
		img = img.tick()
	}

	return img.Lights()
}
