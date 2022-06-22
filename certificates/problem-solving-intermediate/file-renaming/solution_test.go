package main

import (
	"strings"
	"testing"
)

// not optimized - timeout
func solutionRecursive(str string, substr string) int32 {
	var len1 = len(str)
	var len2 = len(substr)

	if len1 < len2 {
		return 0
	}

	if len2 == 1 {
		return int32(strings.Count(str, substr))
	}

	var count int32 = 0

	for i := 0; i < (len1-len2)+1; i++ {
		if str[i] == substr[0] {
			count += solutionRecursive(str[i+1:len1], substr[1:len2])
		}
	}

	return count
}

// optimized
func solutionOptimized(str string, substr string) int32 {
	var len1 = len(str)
	var len2 = len(substr)

	if len1 < len2 {
		return 0
	}

	var m = make([]int, len1+1, len1+1)
	for k := 0; k <= len1; k++ {
		m[k] = 1
	}

	for j := len2 - 1; j >= 0; j-- {
		var n = make([]int, len1+1, len1+1)
		for i := len1 - 1; i >= 0; i-- {
			if str[i] == substr[j] {
				n[i] = m[i+1] + n[i+1]
			} else {
				n[i] = n[i+1]
			}
		}
		// fmt.Println(n)
		m = n
	}

	return int32(m[0])
}

type CntCache struct {
	Size int32
	Data []int32
}

func (c *CntCache) GetCount(sidx int32) int32 {
	var i int32 = sidx + 1
	var cnt int32
	for ; i < c.Size; i++ {
		if c.Data[i] != 0 {
			cnt += c.Data[i]
		}
	}
	return cnt
}
func NewCntCache(size int32) CntCache {
	var c = CntCache{}
	c.Size = size
	var i int32 = 0
	for ; i < size; i++ {
		c.Data = append(c.Data, 0)
	}
	return c
}

type Cache struct {
	Size int32
	Data []CntCache
}

func NewCache(sszie, fsize int32) Cache {
	var cache = Cache{}
	cache.Size = sszie
	var i int32 = 0
	for ; i < sszie; i++ {
		cache.Data = append(cache.Data, NewCntCache(fsize))
	}
	return cache
}

// optimized
func solutionStruct(str string, substr string) int32 {
	size1 := int32(len(substr))
	size2 := int32(len(str))
	var cache = NewCache(size1, size2)
	var i, j int32

	for j = size1 - 1; j >= 0; j-- {
		for i = j; i < size2-size1+j+1; i++ {
			// fmt.Println(j, i)
			if substr[j] != str[i] {
				continue
			}
			if j == size1-1 {
				cache.Data[j].Data[i] = 1
			} else {
				cache.Data[j].Data[i] = cache.Data[j+1].GetCount(i)
			}
		}
		// fmt.Println(cache.Data[j].Data)
	}
	return cache.Data[0].GetCount(-1)
}

func TestFindCount(t *testing.T) {

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

	t.Run("Test FindCount", func(t *testing.T) {
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
