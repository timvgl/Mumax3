package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var copypadmul_code cu.Function

type copypadmul_args struct {
	arg_dst     unsafe.Pointer
	arg_D0      int
	arg_D1      int
	arg_D2      int
	arg_src     unsafe.Pointer
	arg_S0      int
	arg_S1      int
	arg_S2      int
	arg_volmask unsafe.Pointer
	arg_Bsat    float32
	argptr      [10]unsafe.Pointer
}

// Wrapper for copypadmul CUDA kernel, asynchronous.
func k_copypadmul_async(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, volmask unsafe.Pointer, Bsat float32, cfg *config, str cu.Stream) {
	if copypadmul_code == 0 {
		copypadmul_code = fatbinLoad(copypadmul_map, "copypadmul")
	}

	var a copypadmul_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_D0 = D0
	a.argptr[1] = unsafe.Pointer(&a.arg_D0)
	a.arg_D1 = D1
	a.argptr[2] = unsafe.Pointer(&a.arg_D1)
	a.arg_D2 = D2
	a.argptr[3] = unsafe.Pointer(&a.arg_D2)
	a.arg_src = src
	a.argptr[4] = unsafe.Pointer(&a.arg_src)
	a.arg_S0 = S0
	a.argptr[5] = unsafe.Pointer(&a.arg_S0)
	a.arg_S1 = S1
	a.argptr[6] = unsafe.Pointer(&a.arg_S1)
	a.arg_S2 = S2
	a.argptr[7] = unsafe.Pointer(&a.arg_S2)
	a.arg_volmask = volmask
	a.argptr[8] = unsafe.Pointer(&a.arg_volmask)
	a.arg_Bsat = Bsat
	a.argptr[9] = unsafe.Pointer(&a.arg_Bsat)

	args := a.argptr[:]
	cu.LaunchKernel(copypadmul_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for copypadmul CUDA kernel, synchronized.
func k_copypadmul(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, volmask unsafe.Pointer, Bsat float32, cfg *config) {
	str := stream()
	k_copypadmul_async(dst, D0, D1, D2, src, S0, S1, S2, volmask, Bsat, cfg, str)
	syncAndRecycle(str)
}

var copypadmul_map = map[int]string{0: "",
	20: copypadmul_ptx_20,
	30: copypadmul_ptx_30,
	35: copypadmul_ptx_35}

const (
	copypadmul_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u32 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u64 copypadmul_param_8,
	.param .f32 copypadmul_param_9
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<35>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd6, [copypadmul_param_0];
	ld.param.u32 	%r18, [copypadmul_param_1];
	ld.param.u32 	%r19, [copypadmul_param_2];
	ld.param.u32 	%r20, [copypadmul_param_3];
	ld.param.u64 	%rd7, [copypadmul_param_4];
	ld.param.u32 	%r21, [copypadmul_param_5];
	ld.param.u32 	%r22, [copypadmul_param_6];
	ld.param.u32 	%r23, [copypadmul_param_7];
	ld.param.u64 	%rd5, [copypadmul_param_8];
	ld.param.f32 	%f3, [copypadmul_param_9];
	cvta.to.global.u64 	%rd1, %rd6;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 2 15 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r24, %r1, %r2, %r3;
	.loc 2 16 1
	mov.u32 	%r4, %ntid.x;
	mov.u32 	%r5, %ctaid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r25, %r4, %r5, %r6;
	.loc 2 18 1
	setp.ge.s32 	%p1, %r25, %r23;
	setp.ge.s32 	%p2, %r24, %r22;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r24, %r19;
	or.pred  	%p5, %p3, %p4;
	setp.ge.s32 	%p6, %r25, %r20;
	or.pred  	%p7, %p5, %p6;
	@%p7 bra 	BB0_6;

	.loc 3 210 5
	min.s32 	%r7, %r21, %r18;
	.loc 2 24 1
	setp.lt.s32 	%p8, %r7, 1;
	@%p8 bra 	BB0_6;

	mad.lo.s32 	%r33, %r23, %r24, %r25;
	mul.lo.s32 	%r9, %r23, %r22;
	mad.lo.s32 	%r32, %r20, %r24, %r25;
	mul.lo.s32 	%r11, %r20, %r19;
	mov.u32 	%r34, 0;

BB0_3:
	.loc 2 27 1
	cvt.s64.s32 	%rd4, %r33;
	setp.eq.s64 	%p9, %rd5, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 2 26 1
	@%p9 bra 	BB0_5;

	shl.b64 	%rd8, %rd4, 2;
	add.s64 	%rd9, %rd3, %rd8;
	ld.global.f32 	%f8, [%rd9];

BB0_5:
	.loc 2 27 1
	mul.wide.s32 	%rd10, %r33, 4;
	add.s64 	%rd11, %rd2, %rd10;
	ld.global.f32 	%f5, [%rd11];
	mul.f32 	%f6, %f8, %f3;
	mul.f32 	%f7, %f6, %f5;
	mul.wide.s32 	%rd12, %r32, 4;
	add.s64 	%rd13, %rd1, %rd12;
	st.global.f32 	[%rd13], %f7;
	.loc 2 24 1
	add.s32 	%r33, %r33, %r9;
	add.s32 	%r32, %r32, %r11;
	.loc 2 24 52
	add.s32 	%r34, %r34, 1;
	.loc 2 24 1
	setp.lt.s32 	%p10, %r34, %r7;
	@%p10 bra 	BB0_3;

BB0_6:
	.loc 2 29 2
	ret;
}


`
	copypadmul_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u32 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u64 copypadmul_param_8,
	.param .f32 copypadmul_param_9
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<35>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd6, [copypadmul_param_0];
	ld.param.u32 	%r18, [copypadmul_param_1];
	ld.param.u32 	%r19, [copypadmul_param_2];
	ld.param.u32 	%r20, [copypadmul_param_3];
	ld.param.u64 	%rd7, [copypadmul_param_4];
	ld.param.u32 	%r21, [copypadmul_param_5];
	ld.param.u32 	%r22, [copypadmul_param_6];
	ld.param.u32 	%r23, [copypadmul_param_7];
	ld.param.u64 	%rd5, [copypadmul_param_8];
	ld.param.f32 	%f3, [copypadmul_param_9];
	cvta.to.global.u64 	%rd1, %rd6;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 2 15 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r24, %r1, %r2, %r3;
	.loc 2 16 1
	mov.u32 	%r4, %ntid.x;
	mov.u32 	%r5, %ctaid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r25, %r4, %r5, %r6;
	.loc 2 18 1
	setp.ge.s32 	%p1, %r25, %r23;
	setp.ge.s32 	%p2, %r24, %r22;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r24, %r19;
	or.pred  	%p5, %p3, %p4;
	setp.ge.s32 	%p6, %r25, %r20;
	or.pred  	%p7, %p5, %p6;
	@%p7 bra 	BB0_6;

	.loc 3 210 5
	min.s32 	%r7, %r21, %r18;
	.loc 2 24 1
	setp.lt.s32 	%p8, %r7, 1;
	@%p8 bra 	BB0_6;

	mad.lo.s32 	%r33, %r23, %r24, %r25;
	mul.lo.s32 	%r9, %r23, %r22;
	mad.lo.s32 	%r32, %r20, %r24, %r25;
	mul.lo.s32 	%r11, %r20, %r19;
	mov.u32 	%r34, 0;

BB0_3:
	.loc 2 27 1
	cvt.s64.s32 	%rd4, %r33;
	setp.eq.s64 	%p9, %rd5, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 2 26 1
	@%p9 bra 	BB0_5;

	shl.b64 	%rd8, %rd4, 2;
	add.s64 	%rd9, %rd3, %rd8;
	ld.global.f32 	%f8, [%rd9];

BB0_5:
	.loc 2 27 1
	mul.wide.s32 	%rd10, %r33, 4;
	add.s64 	%rd11, %rd2, %rd10;
	ld.global.f32 	%f5, [%rd11];
	mul.f32 	%f6, %f8, %f3;
	mul.f32 	%f7, %f6, %f5;
	mul.wide.s32 	%rd12, %r32, 4;
	add.s64 	%rd13, %rd1, %rd12;
	st.global.f32 	[%rd13], %f7;
	.loc 2 24 1
	add.s32 	%r33, %r33, %r9;
	add.s32 	%r32, %r32, %r11;
	.loc 2 24 52
	add.s32 	%r34, %r34, 1;
	.loc 2 24 1
	setp.lt.s32 	%p10, %r34, %r7;
	@%p10 bra 	BB0_3;

BB0_6:
	.loc 2 29 2
	ret;
}


`
	copypadmul_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry copypadmul(
	.param .u64 copypadmul_param_0,
	.param .u32 copypadmul_param_1,
	.param .u32 copypadmul_param_2,
	.param .u32 copypadmul_param_3,
	.param .u64 copypadmul_param_4,
	.param .u32 copypadmul_param_5,
	.param .u32 copypadmul_param_6,
	.param .u32 copypadmul_param_7,
	.param .u64 copypadmul_param_8,
	.param .f32 copypadmul_param_9
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<33>;
	.reg .f32 	%f<9>;
	.reg .s64 	%rd<14>;


	ld.param.u64 	%rd6, [copypadmul_param_0];
	ld.param.u32 	%r18, [copypadmul_param_1];
	ld.param.u32 	%r19, [copypadmul_param_2];
	ld.param.u32 	%r20, [copypadmul_param_3];
	ld.param.u64 	%rd7, [copypadmul_param_4];
	ld.param.u32 	%r21, [copypadmul_param_5];
	ld.param.u32 	%r22, [copypadmul_param_6];
	ld.param.u32 	%r23, [copypadmul_param_7];
	ld.param.u64 	%rd5, [copypadmul_param_8];
	ld.param.f32 	%f3, [copypadmul_param_9];
	cvta.to.global.u64 	%rd1, %rd6;
	cvta.to.global.u64 	%rd2, %rd7;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 3 15 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r24, %r1, %r2, %r3;
	.loc 3 16 1
	mov.u32 	%r4, %ntid.x;
	mov.u32 	%r5, %ctaid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r25, %r4, %r5, %r6;
	.loc 3 18 1
	setp.ge.s32 	%p1, %r25, %r23;
	setp.ge.s32 	%p2, %r24, %r22;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r24, %r19;
	or.pred  	%p5, %p3, %p4;
	setp.ge.s32 	%p6, %r25, %r20;
	or.pred  	%p7, %p5, %p6;
	@%p7 bra 	BB2_6;

	.loc 4 210 5
	min.s32 	%r7, %r21, %r18;
	.loc 3 24 1
	setp.lt.s32 	%p8, %r7, 1;
	@%p8 bra 	BB2_6;

	mad.lo.s32 	%r31, %r23, %r24, %r25;
	mul.lo.s32 	%r9, %r23, %r22;
	mad.lo.s32 	%r30, %r20, %r24, %r25;
	mul.lo.s32 	%r11, %r20, %r19;
	mov.u32 	%r32, 0;

BB2_3:
	.loc 3 27 1
	cvt.s64.s32 	%rd4, %r31;
	setp.eq.s64 	%p9, %rd5, 0;
	mov.f32 	%f8, 0f3F800000;
	.loc 3 26 1
	@%p9 bra 	BB2_5;

	shl.b64 	%rd8, %rd4, 2;
	add.s64 	%rd9, %rd3, %rd8;
	ld.global.nc.f32 	%f8, [%rd9];

BB2_5:
	.loc 3 27 1
	mul.wide.s32 	%rd10, %r31, 4;
	add.s64 	%rd11, %rd2, %rd10;
	ld.global.nc.f32 	%f5, [%rd11];
	mul.f32 	%f6, %f8, %f3;
	mul.f32 	%f7, %f6, %f5;
	mul.wide.s32 	%rd12, %r30, 4;
	add.s64 	%rd13, %rd1, %rd12;
	st.global.f32 	[%rd13], %f7;
	.loc 3 24 1
	add.s32 	%r31, %r31, %r9;
	add.s32 	%r30, %r30, %r11;
	.loc 3 24 52
	add.s32 	%r32, %r32, 1;
	.loc 3 24 1
	setp.lt.s32 	%p10, %r32, %r7;
	@%p10 bra 	BB2_3;

BB2_6:
	.loc 3 29 2
	ret;
}


`
)