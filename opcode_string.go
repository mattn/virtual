// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT.

package virtual

import "fmt"

const _Opcode_name = "NopAPAddF32AddF64AddC64AddC128AddI32AddI64AddPtrAddPtrsAddSPAnd16And32And64And8ArgumentArgument16Argument32Argument64Argument8ArgumentsArgumentsFPBPBitfieldI8BitfieldI16BitfieldI32BitfieldI64BitfieldU8BitfieldU16BitfieldU32BitfieldU64BoolC128BoolF32BoolF64BoolI16BoolI32BoolI64BoolI8CallCallFPConvC64C128ConvF32C128ConvF32C64ConvF32F64ConvF32I32ConvF32I64ConvF32U32ConvF64C128ConvF64F32ConvF64I32ConvF64I64ConvF64I8ConvF64U16ConvF64U32ConvF64U64ConvI16I32ConvI16I64ConvI16U32ConvI32C128ConvI32C64ConvI32F32ConvI32F64ConvI32I16ConvI32I64ConvI32I8ConvI64F64ConvI64I16ConvI64I32ConvI64I8ConvI64U16ConvI8I16ConvI8I32ConvI8I64ConvI8U32ConvU16I32ConvU16I64ConvU16U32ConvU32F32ConvU32F64ConvU32I16ConvU32I64ConvU32U8ConvU8I16ConvU8I32ConvU8U32ConvU8U64CopyCpl32Cpl64Cpl8DSDSC128DSI16DSI32DSI64DSI8DSNDivC128DivC64DivF32DivF64DivI32DivI64DivU32DivU64Dup32Dup64Dup8EqF32EqF64EqI32EqI64EqI8ExtFFIReturnFPField16Field64Field8FuncGeqF32GeqF64GeqI32GeqI64GeqI8GeqU32GeqU64GtF32GtF64GtI32GtI64GtU32GtU64IndexIndexI16IndexI32IndexI64IndexI8IndexU32IndexU64IndexU8JmpJmpPJnzJzLabelLeqF32LeqF64LeqI32LeqI64LeqI8LeqU32LeqU64LoadLoad16Load32Load64Load8LshI16LshI32LshI64LshI8LtF32LtF64LtI32LtI64LtU32LtU64MulC128MulC64MulF32MulF64MulI32MulI64NegF32NegF64NegI16NegI32NegI64NegI8NegIndexI32NegIndexI64NegIndexU32NegIndexU64NeqC128NeqC64NeqF32NeqF64NeqI32NeqI64NeqI8NotOr32Or64PanicPostIncF64PostIncI16PostIncI32PostIncI64PostIncI8PostIncPtrPostIncU32BitsPostIncU64BitsPreIncI16PreIncI32PreIncI64PreIncI8PreIncPtrPreIncU32BitsPreIncU64BitsPtrDiffPush16Push32Push64Push8PushC128RemI32RemI64RemU32RemU64ReturnRshI16RshI32RshI64RshI8RshU16RshU32RshU64RshU8StoreStore16Store32Store64StoreC128Store8StoreBits16StoreBits32StoreBits64StoreBits8StrNCopySubF32SubF64SubI32SubI64SubPtrsTextVariableVariable16Variable32Variable64Variable8Xor32Xor64Zero8Zero16Zero32Zero64abortabsacosallocaasinatanbswap64callocceilcimagfclose_clrsbclrsblclrsbllclzclzlclzllcopysigncoscoshcrealfctzctzlctzllexitexpfabsfcloseffsffslffsllfgetcfgetsfloorfopenfprintfframeAddressfreadfreefwriteisinfisinffisinflisprintloglog10mallocmemcmpmemcpymemmovemempcpymemsetopenparityparitylparityllpopcountpopcountlpopcountllpowprintfqsortreadreturnAddressroundsign_bitsign_bitfsinsinhsprintfsqrtstrcatstrchrstrcmpstrcpystrlenstrncmpstrncpystrrchrtantanhtolowervfprintfvprintfwrite"

var _Opcode_index = [...]uint16{0, 3, 5, 11, 17, 23, 30, 36, 42, 48, 55, 60, 65, 70, 75, 79, 87, 97, 107, 117, 126, 135, 146, 148, 158, 169, 180, 191, 201, 212, 223, 234, 242, 249, 256, 263, 270, 277, 283, 287, 293, 304, 315, 325, 335, 345, 355, 365, 376, 386, 396, 406, 415, 425, 435, 445, 455, 465, 475, 486, 496, 506, 516, 526, 536, 545, 555, 565, 575, 584, 594, 603, 612, 621, 630, 640, 650, 660, 670, 680, 690, 700, 709, 718, 727, 736, 745, 749, 754, 759, 763, 765, 771, 776, 781, 786, 790, 793, 800, 806, 812, 818, 824, 830, 836, 842, 847, 852, 856, 861, 866, 871, 876, 880, 883, 892, 894, 901, 908, 914, 918, 924, 930, 936, 942, 947, 953, 959, 964, 969, 974, 979, 984, 989, 994, 1002, 1010, 1018, 1025, 1033, 1041, 1048, 1051, 1055, 1058, 1060, 1065, 1071, 1077, 1083, 1089, 1094, 1100, 1106, 1110, 1116, 1122, 1128, 1133, 1139, 1145, 1151, 1156, 1161, 1166, 1171, 1176, 1181, 1186, 1193, 1199, 1205, 1211, 1217, 1223, 1229, 1235, 1241, 1247, 1253, 1258, 1269, 1280, 1291, 1302, 1309, 1315, 1321, 1327, 1333, 1339, 1344, 1347, 1351, 1355, 1360, 1370, 1380, 1390, 1400, 1409, 1419, 1433, 1447, 1456, 1465, 1474, 1482, 1491, 1504, 1517, 1524, 1530, 1536, 1542, 1547, 1555, 1561, 1567, 1573, 1579, 1585, 1591, 1597, 1603, 1608, 1614, 1620, 1626, 1631, 1636, 1643, 1650, 1657, 1666, 1672, 1683, 1694, 1705, 1715, 1723, 1729, 1735, 1741, 1747, 1754, 1758, 1766, 1776, 1786, 1796, 1805, 1810, 1815, 1820, 1826, 1832, 1838, 1843, 1846, 1850, 1856, 1860, 1864, 1871, 1877, 1881, 1887, 1893, 1898, 1904, 1911, 1914, 1918, 1923, 1931, 1934, 1938, 1944, 1947, 1951, 1956, 1960, 1963, 1967, 1973, 1976, 1980, 1985, 1990, 1995, 2000, 2005, 2012, 2024, 2029, 2033, 2039, 2044, 2050, 2056, 2063, 2066, 2071, 2077, 2083, 2089, 2096, 2103, 2109, 2113, 2119, 2126, 2134, 2142, 2151, 2161, 2164, 2170, 2175, 2179, 2192, 2197, 2205, 2214, 2217, 2221, 2228, 2232, 2238, 2244, 2250, 2256, 2262, 2269, 2276, 2283, 2286, 2290, 2297, 2305, 2312, 2317}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
