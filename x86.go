package cpuid

// CPU contains information about the CPU as detected on startup,
// or when Detect last was called.
//
// Use this as the primary entry point to you data,
// this way queries are
var X86 X86Info

// X86Info contains information about the detected system CPU.
type X86Info struct {
	CPUInfo
	Family         int    // CPU family number
	Model          int    // CPU model number
	SGX       SGXSupport
	maxFunc   uint32
	maxExFunc uint32
}

