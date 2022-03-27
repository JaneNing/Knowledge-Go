package fuzz

//go 1.18 Beta版本支持

//func FuzzTest(f *testing.F) {
//	f.Fuzz(func(t *testing.T, i int) {
//		_, err := test(i)
//		if err != nil {
//			fmt.Println(err)
//		}
//	})
//}
//
//func test(in int) (int, error) {
//	if in == 0 {
//		return 0, errors.New("test err")
//	} else {
//		return 0, nil
//	}
//}
