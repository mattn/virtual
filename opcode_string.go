// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT

package virtual

import "fmt"

const _Opcode_name = "NopAPAddF64AddI32AddPtrAddPtrsAddSPAnd32And8Argument32Argument64Argument8ArgumentsArgumentsFPBPBoolI32BoolI64BoolI8CallCallFPConvF32F64ConvF64F32ConvF64I32ConvF64I8ConvI32F32ConvI32F64ConvI32I64ConvI32I8ConvI64I32ConvI8I32CopyDSDSI32DSI64DivF64DivI32DivU64Dup32Dup64EqI32EqI64ExtFloat32Float64FuncGeqI32GeqU64GtI32GtI64GtU64IndexIndexI32Int32Int64JmpJnzJzLabelLeqI32Load32Load64Load8LshI32LtI32LtU64MulF64MulI32NeqI32NeqI64NotOr32PanicPostIncI32PostIncI8PostIncPtrPreIncI32PreIncI8PreIncPtrPtrDiffRemU64ReturnRshI32RshI8Store32Store64Store8StoreBits8SubF64SubI32TextVariable32Variable64Variable8Xor32Zero32Zero64abortacosasinatanceilcoscoshexitexpfabsfclosefgetcfgetsfloorfopenfprintffreadfwriteloglog10memcmpmemcpymemsetpowprintfroundsinsinhsprintfsqrtstrcatstrchrstrcmpstrcpystrlenstrncmpstrncpystrrchrtantanhtolower"

var _Opcode_index = [...]uint16{0, 3, 5, 11, 17, 23, 30, 35, 40, 44, 54, 64, 73, 82, 93, 95, 102, 109, 115, 119, 125, 135, 145, 155, 164, 174, 184, 194, 203, 213, 222, 226, 228, 233, 238, 244, 250, 256, 261, 266, 271, 276, 279, 286, 293, 297, 303, 309, 314, 319, 324, 329, 337, 342, 347, 350, 353, 355, 360, 366, 372, 378, 383, 389, 394, 399, 405, 411, 417, 423, 426, 430, 435, 445, 454, 464, 473, 481, 490, 497, 503, 509, 515, 520, 527, 534, 540, 550, 556, 562, 566, 576, 586, 595, 600, 606, 612, 617, 621, 625, 629, 633, 636, 640, 644, 647, 651, 657, 662, 667, 672, 677, 684, 689, 695, 698, 703, 709, 715, 721, 724, 730, 735, 738, 742, 749, 753, 759, 765, 771, 777, 783, 790, 797, 804, 807, 811, 818}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
