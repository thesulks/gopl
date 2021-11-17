package main

import (
	"testing"
)

func BenchmarkMandelbrotImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImage()
	}
}

// func BenchmarkMandelbrotImageWaitGroup(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroup()
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(10)
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken100(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(100)
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken1000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(1000)
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken10000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(10000)
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken100000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(100000)
// 	}
// }

// func BenchmarkMandelbrotImageWaitGroupWithToken1000000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mandelbrotImageWaitGroupWithToken(1000000)
// 	}
// }

func BenchmarkMandelbrotImageLocality2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(2)
	}
}

func BenchmarkMandelbrotImageLocality4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(4)
	}
}

func BenchmarkMandelbrotImageLocality8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(8)
	}
}

func BenchmarkMandelbrotImageLocality16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(16)
	}
}

func BenchmarkMandelbrotImageLocality32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(32)
	}
}

func BenchmarkMandelbrotImageLocality64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(64)
	}
}

func BenchmarkMandelbrotImageLocality128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(128)
	}
}

func BenchmarkMandelbrotImageLocality256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(256)
	}
}

func BenchmarkMandelbrotImageLocality512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(512)
	}
}

func BenchmarkMandelbrotImageLocality1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrotImageLocality(1024)
	}
}
