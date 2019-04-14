package cpuid

// Flags contains detected cpu features and characteristics
type X86Flag uint64


const (
	CMOV        X86Flag = iota // i686 CMOV
	NX                         // NX (No-Execute) bit
	AMD3DNOW                   // AMD 3DNOW
	AMD3DNOWEXT                // AMD 3DNowExt
	MMX                        // standard MMX
	MMXEXT                     // SSE integer functions or AMD MMX ext
	SSE                        // SSE functions
	SSE2                       // P4 SSE functions
	SSE3                       // Prescott SSE3 functions
	SSSE3                      // Conroe SSSE3 functions
	SSE4                       // Penryn SSE4.1 functions
	SSE4A                      // AMD Barcelona microarchitecture SSE4a instructions
	SSE42                      // Nehalem SSE4.2 functions
	AVX                        // AVX functions
	AVX2                       // AVX2 functions
	FMA3                       // Intel FMA 3
	FMA4                       // Bulldozer FMA4 functions
	XOP                        // Bulldozer XOP functions
	F16C                       // Half-precision floating-point conversion
	BMI1                       // Bit Manipulation Instruction Set 1
	BMI2                       // Bit Manipulation Instruction Set 2
	TBM                        // AMD Trailing Bit Manipulation
	LZCNT                      // LZCNT instruction
	POPCNT                     // POPCNT instruction
	AESNI                      // Advanced Encryption Standard New Instructions
	CLMUL                      // Carry-less Multiplication
	HTT                        // Hyperthreading (enabled)
	HLE                        // Hardware Lock Elision
	RTM                        // Restricted Transactional Memory
	RDRAND                     // RDRAND instruction is available
	RDSEED                     // RDSEED instruction is available
	ADX                        // Intel ADX (Multi-Precision Add-Carry Instruction Extensions)
	SHA                        // Intel SHA Extensions
	AVX512F                    // AVX-512 Foundation
	AVX512DQ                   // AVX-512 Doubleword and Quadword Instructions
	AVX512IFMA                 // AVX-512 Integer Fused Multiply-Add Instructions
	AVX512PF                   // AVX-512 Prefetch Instructions
	AVX512ER                   // AVX-512 Exponential and Reciprocal Instructions
	AVX512CD                   // AVX-512 Conflict Detection Instructions
	AVX512BW                   // AVX-512 Byte and Word Instructions
	AVX512VL                   // AVX-512 Vector Length Extensions
	AVX512VBMI                 // AVX-512 Vector Bit Manipulation Instructions
	MPX                        // Intel MPX (Memory Protection Extensions)
	ERMS                       // Enhanced REP MOVSB/STOSB
	RDTSCP                     // RDTSCP Instruction
	CX16                       // CMPXCHG16B Instruction
	SGX                        // Software Guard Extensions
	IBPB                       // Indirect Branch Restricted Speculation (IBRS) and Indirect Branch Predictor Barrier (IBPB)
	STIBP                      // Single Thread Indirect Branch Predictors
)

//go:generate stringer -type=X86Flag

/*
// String returns a string representation of the detected
// CPU features.
func (f Flags) String() string {
	return strings.Join(f.Strings(), ",")
}
*/

// Strings returns and array of the detected features.
func (f Flags) Strings() []string {
	s := support()
	r := make([]string, 0, 20)
	for i := uint(0); i < 64; i++ {
		key := Flags(1 << i)
		val := flagNames[key]
		if s&key != 0 {
			r = append(r, val)
		}
	}
	return r
}