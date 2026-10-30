package operations_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/VIZ-Blockchain/viz-go-lib/encoding/transaction"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
)

// The expected hex below was produced by the VIZ testnet node itself:
// database_api.get_transaction_hex over a fixed transaction header. These are the
// consensus wire bytes, so a mismatch means the Go encoder disagrees with the chain.
// Stripped: ref_block_num=1, ref_block_prefix=2, expiration=2026-10-31T00:00:00,
// the operation count, and the trailing transaction extensions/signatures.
var opVectors = []struct {
	name string
	json string
	hex  string
}{
	{
		name: "transfer",
		json: `[["transfer",{"from":"alice","to":"bob","amount":"1.000 VIZ","memo":"hi"}]]`,
		hex:  "0205616c69636503626f62e8030000000000000356495a00000000026869",
	},
	{
		name: "validator_update",
		json: `[["validator_update",{"owner":"alice","url":"https://alice.example","block_signing_key":"VIZ6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"}]]`,
		hex:  "0605616c6963651568747470733a2f2f616c6963652e6578616d706c6502c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf",
	},
	{
		name: "account_validator_vote",
		json: `[["account_validator_vote",{"account":"alice","validator":"bob","approve":true}]]`,
		hex:  "0705616c69636503626f6201",
	},
	{
		name: "account_validator_proxy",
		json: `[["account_validator_proxy",{"account":"alice","proxy":"bob"}]]`,
		hex:  "0805616c69636503626f62",
	},
	{
		name: "use_invite_balance",
		json: `[["use_invite_balance",{"initiator":"alice","receiver":"bob","invite_secret":"secret"}]]`,
		hex:  "3a05616c69636503626f6206736563726574",
	},
	{
		name: "fixed_award",
		json: `[["fixed_award",{"initiator":"alice","receiver":"bob","reward_amount":"2.500 VIZ","max_energy":100,"custom_sequence":7,"memo":"thanks","beneficiaries":[{"account":"carol","weight":1000}]}]]`,
		hex:  "3c05616c69636503626f62c4090000000000000356495a0000000064000700000000000000067468616e6b7301056361726f6ce803",
	},
	{
		name: "target_account_sale",
		json: `[["target_account_sale",{"account":"alice","account_seller":"bob","target_buyer":"carol","account_offer_price":"5.000 VIZ","account_on_sale":true}]]`,
		hex:  "3d05616c69636503626f62056361726f6c88130000000000000356495a0000000001",
	},
	{
		name: "set_reward_sharing",
		json: `[["set_reward_sharing",{"owner":"alice","sharing_rate":2500}]]`,
		hex:  "4005616c696365c409",
	},
	{
		name: "pm_oracle_register",
		json: `[["pm_oracle_register",{"owner":"alice","insurance":"10.000 VIZ","fee_percent":100,"fixed_fee":"0.500 VIZ","rules_url":"https://alice.example/rules","auto_accept_creator":"bob","auto_accept_resolver":"carol","auto_accept":true}]]`,
		hex:  "4205616c69636510270000000000000356495a000000006400f4010000000000000356495a000000001b68747470733a2f2f616c6963652e6578616d706c652f72756c657303626f62056361726f6c0100",
	},
	{
		name: "pm_oracle_update",
		json: `[["pm_oracle_update",{"owner":"alice","insurance_delta":"1.000 VIZ","fee_percent":200,"fixed_fee":"0.250 VIZ","rules_url":"https://alice.example/rules2","auto_accept_creator":"bob","auto_accept_resolver":"carol","auto_accept":false}]]`,
		hex:  "4305616c69636501e8030000000000000356495a0000000001c80001fa000000000000000356495a00000000011c68747470733a2f2f616c6963652e6578616d706c652f72756c6573320103626f6201056361726f6c010000",
	},
	{
		name: "pm_create_market",
		json: `[["pm_create_market",{"creator":"alice","oracle":"bob","market_type":0,"outcomes":["yes","no"],"url":"https://alice.example/market","oracle_fee_percent":100,"oracle_fixed_fee":"0.100 VIZ","creator_fee_percent":50,"liquidity_fee_percent":25,"liquidity":"100.000 VIZ","lmsr_b":0,"betting_expiration":"2026-12-01T00:00:00","result_expiration":"2026-12-02T00:00:00","time_penalty_type":1,"time_penalty_value":3600,"penalty_curve_type":2,"allow_early_resolution":true,"allow_cancellation":false,"allow_batch":true,"allow_instant_bet":false,"endogeneity_tier":1,"dispute_mode":1,"dispute_resolver":"carol","dispute_penalty_percent":-500,"metadata":"{\"k\":1}"}]]`,
		hex:  "4405616c69636503626f62000203796573026e6f1c68747470733a2f2f616c6963652e6578616d706c652f6d61726b6574640064000000000000000356495a0000000032001900a0860100000000000356495a000000000000000000000000000e0e6b805f0f6b01100e000002010001000101056361726f6c0cfe077b226b223a317d00",
	},
	{
		name: "pm_oracle_accept_market",
		json: `[["pm_oracle_accept_market",{"oracle":"bob","market_id":42,"accept":true,"oracle_fee_percent":90,"oracle_fixed_fee":"0.090 VIZ"}]]`,
		hex:  "4503626f622a00000000000000015a005a000000000000000356495a0000000000",
	},
	{
		name: "pm_place_bet",
		json: `[["pm_place_bet",{"account":"alice","market_id":42,"side":1,"outcome_index":0,"amount":"1.000 VIZ","min_tokens":5,"mode":1}]]`,
		hex:  "4605616c6963652a00000000000000010000e8030000000000000356495a0000000005000000000000000100",
	},
	{
		name: "pm_commit_bet",
		json: `[["pm_commit_bet",{"account":"alice","market_id":42,"commitment":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","escrow_amount":"3.000 VIZ","no_reveal_fee_percent":200}]]`,
		hex:  "4705616c6963652a00000000000000e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855b80b0000000000000356495a00000000c80000",
	},
	{
		name: "pm_reveal_bet",
		json: `[["pm_reveal_bet",{"account":"alice","commit_id":55,"side":-1,"outcome_index":1,"amount":"3.000 VIZ","salt":"pepper","min_tokens":0}]]`,
		hex:  "4805616c6963653700000000000000ff0100b80b0000000000000356495a0000000006706570706572000000000000000000",
	},
	{
		name: "pm_cancel_bet",
		json: `[["pm_cancel_bet",{"account":"alice","bet_id":77,"min_return":9}]]`,
		hex:  "4905616c6963654d00000000000000090000000000000000",
	},
	{
		name: "pm_add_liquidity",
		json: `[["pm_add_liquidity",{"provider":"alice","market_id":42,"amount":"20.000 VIZ"}]]`,
		hex:  "4a05616c6963652a00000000000000204e0000000000000356495a0000000000",
	},
	{
		name: "pm_withdraw_liquidity",
		json: `[["pm_withdraw_liquidity",{"provider":"alice","liquidity_id":99,"amount":"5.000 VIZ"}]]`,
		hex:  "4b05616c696365630000000000000088130000000000000356495a0000000000",
	},
	{
		name: "pm_resolve_market",
		json: `[["pm_resolve_market",{"oracle":"bob","market_id":42,"winning_outcome":1,"decision_url":"https://bob.example/d","decision_reason":"final"}]]`,
		hex:  "4c03626f622a0000000000000001001568747470733a2f2f626f622e6578616d706c652f640566696e616c00",
	},
	{
		name: "pm_no_contest",
		json: `[["pm_no_contest",{"oracle":"bob","market_id":42,"reason":"void"}]]`,
		hex:  "4d03626f622a0000000000000004766f696400",
	},
	{
		name: "pm_dispute_create",
		json: `[["pm_dispute_create",{"disputer":"carol","market_id":42,"proposed_outcome":0,"reason":"wrong"}]]`,
		hex:  "4e056361726f6c2a0000000000000000000577726f6e6700",
	},
	{
		name: "pm_dispute_vote",
		json: `[["pm_dispute_vote",{"voter":"carol","market_id":42,"vote_outcome":0,"vote_percent":6000}]]`,
		hex:  "4f056361726f6c2a000000000000000000701700",
	},
	{
		name: "pm_dispute_resolve",
		json: `[["pm_dispute_resolve",{"resolver":"dave","market_id":42,"correct_outcome":0,"penalty_amount":"2.000 VIZ","ban_oracle":true,"ban_oracle_until":"2026-12-31T00:00:00","ban_creator":false,"ban_creator_until":"1970-01-01T00:00:00"}]]`,
		hex:  "5004646176652a000000000000000000d0070000000000000356495a0000000001009b356b000000000000",
	},
	{
		name: "pm_transfer_position",
		json: `[["pm_transfer_position",{"from":"alice","bet_id":77,"to":"bob","amount":12,"memo":"gift"}]]`,
		hex:  "5105616c6963654d0000000000000003626f620c00000000000000046769667400",
	},
	{
		name: "pm_lazy_deposit",
		json: `[["pm_lazy_deposit",{"account":"alice","amount":"4.000 VIZ"}]]`,
		hex:  "5205616c696365a00f0000000000000356495a0000000000",
	},
	{
		name: "pm_lazy_withdraw",
		json: `[["pm_lazy_withdraw",{"account":"alice","shares":13,"emergency":true}]]`,
		hex:  "5305616c6963650d000000000000000100",
	},
	{
		name: "pm_leverage_open",
		json: `[["pm_leverage_open",{"account":"alice","market_id":42,"outcome_index":1,"collateral":"10.000 VIZ","loan":"30.000 VIZ","min_tokens":3,"max_slippage_percent":150}]]`,
		hex:  "5b05616c6963652a00000000000000010010270000000000000356495a0000000030750000000000000356495a000000000300000000000000960000",
	},
	{
		name: "pm_leverage_close",
		json: `[["pm_leverage_close",{"account":"alice","position_id":88,"min_return":4}]]`,
		hex:  "5c05616c6963655800000000000000040000000000000000",
	},
	{
		name: "pm_leverage_convert",
		json: `[["pm_leverage_convert",{"account":"alice","position_id":88,"conversion_profit_cost":700}]]`,
		hex:  "5d05616c6963655800000000000000bc0200",
	},
	{
		name: "pm_dispute_oracle_respond",
		json: `[["pm_dispute_oracle_respond",{"oracle":"bob","market_id":42,"response":"evidence"}]]`,
		hex:  "6203626f622a000000000000000865766964656e636500",
	},
	{
		name: "pm_unban",
		json: `[["pm_unban",{"resolver":"dave","target":"bob","unban_oracle":true,"unban_creator":false}]]`,
		hex:  "63046461766503626f62010000",
	},
	{
		name: "set_agent_permission",
		json: `[["set_agent_permission",{"account":"alice","agent_name":"bot","agent_key":"VIZ6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV","operations":["transfer","pm_place_bet"],"expiration":"2026-12-31T00:00:00","addons":["vizhub"]}]]`,
		hex:  "6905616c69636503626f7402c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf020c706d5f706c6163655f626574087472616e73666572009b356b010676697a68756200",
	},
}

func TestOperationWireFormat(t *testing.T) {
	for _, tc := range opVectors {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var ops operations.Operations
			if err := json.Unmarshal([]byte(tc.json), &ops); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.json, err)
			}
			if len(ops) != 1 {
				t.Fatalf("expected 1 operation, got %d", len(ops))
			}

			var buf bytes.Buffer
			enc := transaction.NewEncoder(&buf)
			if err := enc.Encode(ops[0]); err != nil {
				t.Fatalf("encode: %v", err)
			}
			if got := hex.EncodeToString(buf.Bytes()); got != tc.hex {
				t.Fatalf("wire mismatch\n got %s\nwant %s", got, tc.hex)
			}
		})
	}
}

// TestOperationCodes pins the consensus op-id of the operations this library encodes.
// The id is the index in the node static_variant and is append-only.
func TestOperationCodes(t *testing.T) {
	cases := []struct {
		op   operations.OpType
		want uint16
	}{
		{operations.TypeTransfer, 2},
		{operations.TypeValidatorUpdate, 6},
		{operations.TypeAccountValidatorVote, 7},
		{operations.TypeAccountValidatorProxy, 8},
		{operations.TypeShutdownValidator, 30},
		{operations.TypeValidatorReward, 42},
		{operations.TypeAccountSale, 57},
		{operations.TypeUseInviteBalance, 58},
		{operations.TypeExpireEscrowRatification, 59},
		{operations.TypeFixedAward, 60},
		{operations.TypeTargetAccountSale, 61},
		{operations.TypeStakeholderReward, 65},
		{operations.TypePMOracleRegister, 66},
		{operations.TypePMCreateMarket, 68},
		{operations.TypePMOracleAcceptMarket, 69},
		{operations.TypePMPlaceBet, 70},
		{operations.TypePMLazyWithdraw, 83},
		{operations.TypePMLeverageOpen, 91},
		{operations.TypePMDisputeOracleRespond, 98},
		{operations.TypePMUnban, 99},
		{operations.TypeSetAgentPermission, 105},
		{operations.TypeValidatorEquivocation, 106},
	}
	for _, c := range cases {
		if got := c.op.Code(); got != c.want {
			t.Errorf("%s: code %d, want %d", c.op, got, c.want)
		}
	}
}
