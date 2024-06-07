// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package calculatorplus

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = vmerrs.ErrOutOfGas
	_ = big.NewInt
	_ = common.Big0
	_ = require.New
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	expectedModuloPlusOutcome, _ = PackModuloPlusOutput(ModuloPlusOutput{big.NewInt(8), big.NewInt(0)})
	expectedPowOfThreeOutcome, _ = PackPowOfThreeOutput(PowOfThreeOutput{big.NewInt(4), big.NewInt(8), big.NewInt(16)})
	expectedSimplFracOutcome, _  = PackSimplFracOutput(SimplFracOutput{big.NewInt(1), big.NewInt(2)})
	tests                        = map[string]testutils.PrecompileTest{
		"insufficient gas for moduloPlus should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := ModuloPlusInput{big.NewInt(1), big.NewInt(1)}
				input, err := PackModuloPlus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: ModuloPlusGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"insufficient gas for powOfThree should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// set test input to a value here
				var testInput *big.Int
				testInput = big.NewInt(1)
				input, err := PackPowOfThree(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: PowOfThreeGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"insufficient gas for simplFrac should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SimplFracInput{big.NewInt(1), big.NewInt(1)}
				input, err := PackSimplFrac(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SimplFracGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"testing modulo plus": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				value1 := big.NewInt(16)
				value2 := big.NewInt(2)
				testInput := ModuloPlusInput{value1, value2}
				input, err := PackModuloPlus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: ModuloPlusGasCost,
			ReadOnly:    true,
			ExpectedRes: expectedModuloPlusOutcome,
		},
		"testing power of three": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				testInput := big.NewInt(2)
				input, err := PackPowOfThree(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: PowOfThreeGasCost,
			ReadOnly:    true,
			ExpectedRes: expectedPowOfThreeOutcome,
		},
		"testing simpl func": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				value1 := big.NewInt(8)
				value2 := big.NewInt(16)
				testInput := SimplFracInput{value1, value2}
				input, err := PackSimplFrac(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SimplFracGasCost,
			ReadOnly:    true,
			ExpectedRes: expectedSimplFracOutcome,
		},
	}
)

// TestCalculatorplusRun tests the Run function of the precompile contract.
func TestCalculatorplusRun(t *testing.T) {
	// Run tests.
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Run(t, Module, state.NewTestStateDB(t))
		})
	}
}

func BenchmarkCalculatorplus(b *testing.B) {
	// Benchmark tests.
	for name, test := range tests {
		b.Run(name, func(b *testing.B) {
			test.Bench(b, Module, state.NewTestStateDB(b))
		})
	}
}
