package main

import (
	"testing"
)

func solutionOptimized(str string, substr string) int32 {
	var n, m = len(substr), len(str)

	if m < n {
		return 0
	}

	var dp = make([]int, m+1, m+1)
	for k := 0; k <= m; k++ {
		dp[k] = 1
	}

	for j := n - 1; j >= 0; j-- {
		var dpp = make([]int, m+1, m+1)
		for i := m - 1; i >= 0; i-- {
			if str[i] == substr[j] {
				dpp[i] = dp[i+1] + dpp[i+1]
			} else {
				dpp[i] = dpp[i+1]
			}
		}
		dp = dpp
	}

	return int32(dp[0])
}

func TestSolution(t *testing.T) {

	cases := []struct {
		Args           []string
		ExpectedOutput int32
	}{
		{[]string{"ababa", "aba"}, 4},
		{[]string{"cccc", "ccc"}, 4},
		{[]string{"abcabcabc", "abcd"}, 0},
		{[]string{"abababababababa", "aba"}, 84},
		{[]string{"abababababababacaxbababababbababababababbababbabababbababbabababbababba", "aba"}, 6455},
		{[]string{"abababababababacaxbababababbababababababbababbabababbababbabababbababbababababbabababbabbababbababbababbabbababbababbababbababbabababbabbabababbababbabababbabbababbabbab", "aba"}, 82732},
		{[]string{"abcababacbabacbabacaxbababababbabcababababcabbababbabababbacbabbacbababbacbabbababababcbabababbabcbabdabbababbababbabbabacbbababbcababbababbabababbcabdbabababbababbabacbabbabbababbabbabd", "abcd"}, 24956},
		{[]string{"abcababacbabacbabacaxbababababbabcababababcabbababbabababbacbabbacbababbacbabbababababcbabababbabcbabdabbababbababbabbabacbbababbcababbababbabababbcabdbabababbababbabacbabbabbababbabbabdabcdefghjijlkabcdaaaaaaaaaaaaaaaaaaaaaaaaacdbbbbbbbbbbbbbbbbsssswwwaaaaccccddddbbbbbwaaaaddd", "abcd"}, 402410},
		{[]string{"abcababacbabacamnamdfasldkfalsdfasjfqpglncvplzxv;kxzc.,dfnwpqofj;lzm.ajfoehgkjbvlcjv.ana,msncm,xjcpksx.zmxznlkjd.ancz,xnldjs;lkrlghnkbepbg[wjf'dsnf,axcn,sdaffl;ew,fm';qgpn.xznc,zncmx;lwmefleqwjgand.a,nf,masdnfkfj;mcx,knfs.ams.mcn xz,ncz'ldfma,c ,mcnzj;df,asn.,cx.xz,ma.kf;mfa, vdfbdlfgjerpgjs;ljw;ljm,zmncl,zxncslkjf';ewfm,nfdskfjpegh[hbbc,mxznclkjdalsjkddnnmxcnxcdmfnlkfja,mncxmcnxjlsdkjfadnfsmnxmcnxlkjcksdnamncxmnxdbabacaxbababababbabcababababcabbababbabababbacbabbacbababbacbabbababababcbabababbabcbabdabbababbababbabbabacbbababbcababbababbabababbcabdbabababbababbalmnopqrstuvwqabcdebacbabbabbababbabbabdabcdefghjijlkabcdaaaaaaaaaaaaaaaaaaaaaaaaacdsssssoqpqppabcdeeqpsrtsaabbbbbbbbbbbbbbbbsssswwwaaaaccccddddbbbbbwaaaaddd", "almebelsopq"}, 41492276},
		{[]string{"abcababacbabacamnamdfasldkfalsdfasjfqpglncvplzxv;kxzc.,dfnwpqofj;lzm.ajfoehgkjbvlcjv.ana,msncm,xjcpksx.zmxznlkjd.ancz,xnldjs;lkrlghnkbepbg[wjf'dsnf,axcn,sdaffl;ew,fm';qgpn.xznc,zncmx;lwmefleqwjgand.a,nf,masdnfkfj;mcx,knfs.ams.mcn xz,ncz'ldfma,c ,mcnzj;df,asn.,cx.xz,ma.kf;mfa, vdfbdlfgjerpgjs;ljw;ljm,zmncl,zxncslkjf';ewfm,nfdskfjpegh[hbbc,mxznclkjdalsjkddnnmxcnxcdmfnlkfja,mncxmcnxjlsdkjfadnfsmnxmcnxlkjcksdnamncxmnxdbabacaxbababababbabcababababcabbababbabababbacbabbacbababbacbabbababababcbabababbabcbabdabbababbababbabbabacbbababbcababbababbabababbcabdbabababbababbalmnopqrstuvwqabcdebacbabbabbababbabbabdabcdefghjijlkabcdaaaaaaaaaaaaaaaaasdfasdflknohadjf;asnfbnvxzk.c,.zxfjpowjgkm.anlaskfbg,mnv,.xzcmz,ndjbfhgoquptoquithofhaslkfnanv'aj;jls'anv,mcxbjkfdhioegujqpnxcnv.,xcmzxaaaaaaaaacdsssssoqpqppabcdeeqpsrtsaabbbbbbbbbbbbbbbbsssswwwaaaaccccddddbbbbbwaaaaddd", "almebelsowolxnczowaxpq"}, 0},
	}

	t.Run("Test Solution", func(t *testing.T) {
		for _, tc := range cases {
			var actualOutput int32 = solutionOptimized(tc.Args[0], tc.Args[1])
			if tc.ExpectedOutput != actualOutput {
				t.Errorf("Wrong output for args: %v, expected: %v, got: %v",
					tc.Args, tc.ExpectedOutput, actualOutput)
			}
		}
		// if tc.ExpectedExit != actualExit {
		// 	T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
		// 		tc.Args, tc.ExpectedExit, actualExit)
		// }
	})

}
