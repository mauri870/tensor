package execution

import "testing"

var benchSizes = []int{64, 4096, 1 << 20}

func BenchmarkAddVSF32(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]float32, n)
		for i := range a {
			a[i] = float32(i) * 0.01
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				AddVSF32(a, 1.5)
			}
		})
	}
}

func BenchmarkMulVSF32(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]float32, n)
		for i := range a {
			a[i] = float32(i) * 0.01
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				MulVSF32(a, 1.5)
			}
		})
	}
}

func BenchmarkAddVSF64(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]float64, n)
		for i := range a {
			a[i] = float64(i) * 0.01
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				AddVSF64(a, 1.5)
			}
		})
	}
}

func BenchmarkMulVSF64(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]float64, n)
		for i := range a {
			a[i] = float64(i) * 0.01
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				MulVSF64(a, 1.5)
			}
		})
	}
}

func BenchmarkVecAddI32(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]int32, n)
		bv := make([]int32, n)
		for i := range a {
			a[i] = int32(i)
			bv[i] = int32(i) * 2
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				VecAddI32(a, bv)
			}
		})
	}
}

func BenchmarkVecMulI32(b *testing.B) {
	for _, n := range benchSizes {
		a := make([]int32, n)
		bv := make([]int32, n)
		for i := range a {
			a[i] = int32(i)
			bv[i] = int32(i) * 2
		}
		b.Run(itoa(n), func(b *testing.B) {
			for b.Loop() {
				VecMulI32(a, bv)
			}
		})
	}
}

func itoa(n int) string {
	switch n {
	case 64:
		return "64"
	case 4096:
		return "4096"
	case 1 << 20:
		return "1048576"
	default:
		return "unknown"
	}
}
