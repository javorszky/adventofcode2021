package day20

import "fmt"

func task1(enhance, image string) int {
	img := NewImage(image, enhance)

	fmt.Printf("image:\n-----\n%s\n-----\n\n", image)
	fmt.Printf("enhance\n\n%s\n\n", enhance)

	fmt.Printf("starter image state:\n\n%s\n\n", img.String())

	for i := 0; i < 2; i++ {
		img = img.tick()
		fmt.Printf("\n\nimg after %d ticks\n\n%s\n\n", i+1, img.String())
	}

	return img.Lights()
}
