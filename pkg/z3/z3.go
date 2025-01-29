package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"
import (
	"strconv"
	"unsafe"
)

// CONTEXT

type Context struct {
	val C.Z3_context
}

func MakeContext() Context {
	config := C.Z3_mk_config()
	ctx := C.Z3_mk_context(config)
	C.Z3_del_config(config)

	return Context{ctx}
}

func (ctx Context) Delete() {
	C.Z3_del_context(ctx.val)
}

// OPTIMIZE

type Optimize struct {
	val C.Z3_optimize
	ctx C.Z3_context
}

type Bool struct {
	val C.Z3_ast
	ctx C.Z3_context
}

func MakeOptimize(ctx Context) Optimize {
	opt := C.Z3_mk_optimize(ctx.val)
	C.Z3_optimize_inc_ref(ctx.val, opt)

	return Optimize{opt, ctx.val}
}

func (opt Optimize) Delete() {
	C.Z3_optimize_dec_ref(opt.ctx, opt.val)
}

func (opt Optimize) Assert(bool Bool) {
	C.Z3_optimize_assert(opt.ctx, opt.val, bool.val)
}

func (opt Optimize) Minimize(t BV) {
	C.Z3_optimize_minimize(opt.ctx, opt.val, t.val)
}

func (opt Optimize) Check() (sat bool, is_undef bool) {
	res := C.Z3_optimize_check(opt.ctx, opt.val, 0, nil)
	return res == 1, res == 0
}

func (opt Optimize) Eval(t BV) (value BV, success bool) {
	model := C.Z3_optimize_get_model(opt.ctx, opt.val)
	C.Z3_model_inc_ref(opt.ctx, model)
	defer C.Z3_model_dec_ref(opt.ctx, model)

	var cval C.Z3_ast
	success = C.Z3_model_eval(opt.ctx, model, t.val, true, &cval) == true

	return BV{cval, t.size, opt.ctx}, success
}

// BIT VECTOR FACTORY

type BVFactory struct {
	sort C.Z3_sort
	size uint
	ctx  C.Z3_context
}

type BV struct {
	val  C.Z3_ast
	size uint
	ctx  C.Z3_context
}

func MakeBVFactory(ctx Context, size uint) BVFactory {
	sort := C.Z3_mk_bv_sort(ctx.val, C.uint(size))
	return BVFactory{sort, size, ctx.val}
}

func (bvFac BVFactory) MakeConst(symbol string) BV {
	ctx := bvFac.ctx
	symbolStr := C.CString(symbol)
	defer C.free(unsafe.Pointer(symbolStr))

	return BV{
		C.Z3_mk_const(ctx, C.Z3_mk_string_symbol(ctx, symbolStr), bvFac.sort),
		bvFac.size,
		ctx,
	}
}

func (bvFac BVFactory) MakeInt(num int) BV {
	ctx := bvFac.ctx
	return BV{
		C.Z3_mk_int(ctx, C.int(num), bvFac.sort),
		bvFac.size,
		ctx,
	}
}

// BIT VECTOR

func (bv BV) ShiftRight(other BV) BV {
	return BV{C.Z3_mk_bvlshr(bv.ctx, bv.val, other.val), bv.size, bv.ctx}
}

func (bv BV) Xor(other BV) BV {
	return BV{C.Z3_mk_bvxor(bv.ctx, bv.val, other.val), bv.size, bv.ctx}
}

func (bv BV) And(other BV) BV {
	return BV{C.Z3_mk_bvand(bv.ctx, bv.val, other.val), bv.size, bv.ctx}
}

func (bv BV) Eq(other BV) Bool {
	return Bool{C.Z3_mk_eq(bv.ctx, bv.val, other.val), bv.ctx}
}

func (bv BV) Ne(other BV) Bool {
	return Bool{C.Z3_mk_not(bv.ctx, C.Z3_mk_eq(bv.ctx, bv.val, other.val)), bv.ctx}
}

func (bv BV) ToIntValue() int {
	strVal := C.GoString(C.Z3_ast_to_string(bv.ctx, bv.val))
	val, err := strconv.ParseInt(strVal[2:], 16, int(bv.size))

	if err != nil {
		panic(err)
	}

	return int(val)
}
